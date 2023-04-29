package pkg

import (
	"context"
	"strings"

	"github.com/red512/go-naru/models"
	u "github.com/red512/go-naru/utils"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetContexts() (string, error) {
	contexts, err := u.CmdExecutor("kubectx")
	if err != nil {
		logrus.Error()
	}
	return contexts, nil
}

func GetNamespaces() ([]string, error) {
	namespaces, err := u.CmdExecutor("kubens")
	if err != nil {
		logrus.Error()
	}

	return strings.Split(namespaces, "\n"), nil
}

func GetPods(namespace string) ([]string, error) {

	json, err := u.CmdExecutor("kubectl", "get", "po", "-n", namespace, "-ojson")
	pods := gjson.Get(json, "items.#.metadata.name")

	if err != nil {
		logrus.Error()
	}
	return strings.Split(pods.String(), ","), nil
}

func GetServices(namespace string) ([]string, error) {
	json, err := u.CmdExecutor("kubectl", "get", "svc", "-n", namespace, "-ojson")
	services := gjson.Get(json, "items.#.metadata.name")

	if err != nil {
		logrus.Error()
	}
	return strings.Split(services.String(), ","), nil
}

func GetDeployments(namespace string) ([]string, error) {
	json, err := u.CmdExecutor("kubectl", "get", "deployments", "-n", namespace, "-ojson")
	deployments := gjson.Get(json, "items.#.metadata.name")

	if err != nil {
		logrus.Error()
	}
	return strings.Split(deployments.String(), ","), nil
}

func GetNamespacesData(namespaces []string) ([]models.Namespace, error) {
	namespaceDataSlice := make([]models.Namespace, 0)

	for _, n := range namespaces {
		pods, _ := GetPods(n)
		services, _ := GetServices(n)
		ingresses, _ := GetServices(n)
		deployments, _ := GetDeployments(n)
		namespaceDataSlice = append(namespaceDataSlice, models.Namespace{
			Name:        n,
			Pods:        pods,
			Services:    services,
			Ingresses:   ingresses,
			Deployments: deployments,
		})
	}
	return namespaceDataSlice, nil
}

func SendK8sDataToMongoDB(data []models.Namespace, client *mongo.Client) {
	collection := client.Database("narudb").Collection("k8sinfo")
	for _, n := range data {
		filter := bson.M{"name": n.Name}
		update := bson.M{"$set": bson.M{
			"pods":        n.Pods,
			"services":    n.Services,
			"deployments": n.Deployments,
			"ingresses":   n.Ingresses,
		}}
		_, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
		if err != nil {
			logrus.Errorf("Failed to update MongoDB: %v", err)
		}
	}
}

func GetK8sDataFromMongoDB(client *mongo.Client) ([]models.Namespace, error) {
	collection := client.Database("narudb").Collection("k8sinfo")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var namespaces []models.Namespace
	for cursor.Next(context.Background()) {
		var namespace models.Namespace
		if err := cursor.Decode(&namespace); err != nil {
			return nil, err
		}
		namespaces = append(namespaces, namespace)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return namespaces, nil
}
