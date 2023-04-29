# go-naru

Go-Naru - alpha version
Go-Naru is a Go-based small REST API that leverages Echo framework to retrieve Kubernetes custom data including namespaces, pods, services, deployments, and ingresses from the cluster, and provides convenient CLI.

Installation
Before running Go-Naru, make sure that you have the following prerequisites installed:

```
kubectl
kubectx
kubens
```

Then, clone the repository and run go run main.go to start the server.

Usage
Go-Naru provides several command-line interface (CLI) tools that can be run by entering the following commands:

- This command will run the server.

```
naru server
```

- This command will send data to MongoDB.

```
naru sendDataToMongo
```

Directories:

```
.
├── build
├── cmd
├── internal
├── models
├── pkg
├── utils
└── web
```

Please note that this is a demo/alpha version and more features will be added in the future.

Go-Naru - alpha version
Go-Naru is a Go-based small REST API that leverages Echo framework to retrieve Kubernetes custom data including namespaces, pods, services, deployments, and ingresses from the cluster, and provides convenient CLI.

Installation
Before running Go-Naru, make sure that you have the following prerequisites installed:

```
kubectl
kubectx
kubens
```

Then, clone the repository and run go run main.go to start the server.

Usage
Go-Naru provides several command-line interface (CLI) tools that can be run by entering the following commands:

- This command will run the server.

```
naru server
```

- This command will send data to MongoDB.

```
naru sendDataToMongo
```

Directories:

```
.
├── build
├── cmd
├── internal
├── models
├── pkg
├── utils
└── web
```

Please note that this is a demo/alpha version and more features will be added in the future.
