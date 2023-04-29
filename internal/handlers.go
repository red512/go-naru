package internal

import (
	"net/http"

	"github.com/red512/go-naru/pkg"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func GetDataFromMongoDBHandler(c echo.Context) error {
	client := pkg.Client()
	namespaceData, err := pkg.GetK8sDataFromMongoDB(client)
	if err != nil {
		logrus.Error(err)
		return c.String(http.StatusInternalServerError, "Error retrieving data from MongoDB")
	}
	return c.JSON(http.StatusOK, namespaceData)
}
