package models

type Namespace struct {
	Name        string   `json: "name"`
	Pods        []string `json: "pods"`
	Services    []string `json: "services"`
	Deployments []string `json: "deployments"`
	Ingresses   []string `json: "ingresses"`
}
