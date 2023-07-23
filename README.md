## Go REST API with GitHub Actions and ArgoCD

### Introduction

This is a project in which we made a Go REST API with help of [Gin Framework](https://gin-gonic.com/docs/quickstart/).
Additionally, this project contains a CI/CD (Continuous Integration and Continuous Deployment) implementation. It offers a thorough illustration of how to create, test, and deploy a Go application using GitHub Actions and Argo CD.

### Prerequisites

- [Go(v1.16 +)](https://go.dev/doc/install)
- [Docker](https://www.docker.com/)
- [Kubernetes](https://kubernetes.io/docs/setup/)
- [GitHub Actions](https://docs.github.com/en/actions)
- [ArgoCD](https://argo-cd.readthedocs.io/en/stable/)

### Getting Started

Follow the steps below, to get started with the project:

1. **Clone the repository**

```
git clone https://github.com/ShivangShandilya/golang-api.git
```

2. **Install the dependencies**

```
go mod download
```

3. **Running the Application**

```
go run .
```

4. **You can now make GET,POST requests via Thunder Client extension on VS Code or using Postman**

<p align = "center">
  <img src = "https://github.com/ShivangShandilya/golang-api/assets/101946115/84ef207c-b32c-4ac3-99fd-5c1189f7a541" height = 400 width = 800 />
</p>

5. **Build the Dockerfile**

```
docker build . -t golang-api
```

**Note**: If you're using Docker Desktop, then you can assign any port number of your choice where you want your app to run.

<p align = "center">
<img src = "https://github.com/ShivangShandilya/golang-api/assets/101946115/63b23add-44d7-422b-93a5-ed9df91f57ad" height = 400 width = 800 />
</p>

### Overview of CI/CD Workflow

Refer to [ci.yaml](https://github.com/ShivangShandilya/golang-api/blob/main/.github/workflows/ci.yaml) and [publish-dockerhub.yaml](https://github.com/ShivangShandilya/golang-api/blob/main/.github/workflows/publish-dockerhub.yaml) files for configuration.

In order for the abpve GitHub Actions to function you must store these inside GitHub secrets of your repository:
- `DOCKERHUB_USERNAME` -  DockerHub Username
- `DOCKERHUB_PASSWORD` - DockerHub Password or Access Token


**The ci.yaml action consists of 4 steps**
- Checkout Code - This step uses the `actions/checkout@v3` action to clone the repository containing the code that needs to be built and tested.

- Setup Go - This step uses the `actions/setup-go@v4` action to set up the Go programming language environment on the runner.
  
- Build - This step builds the Go code present in the repository using the go build command. The -v flag enables verbose output, providing additional information during the build process. The . argument indicates that the current directory should be built, assuming it contains the main package of the project.
  
- Test - This step runs the tests for the Go code using the go test command. The -v flag again enables verbose output, giving detailed information about the executed tests. The ./tests argument specifies the path to the test directory or package that needs to be tested. It might contain test files with names like *_test.go.

**The publish-dockerhub.yaml action consists of 3 steps**
- DockerHub Login - As the name suggests this logs in to the provided DockerHub Account given in the secrets via `DOCKERHUB_USERNAME` and `DOCKERHUB_PASSWORD`.

-  Build the Docker image - This step builds our Docker Image.

- Docker Push -  Docker Push basically just pushes our image made in previous step over to DockerHub.

### ArgoCD - Continuos Deployment

A short description for ArgoCD would be - ArgoCD is a Kubernetes-native tool for automating application deployments. It ensures that the deployed applications match the desired state defined in a Git repository, simplifying the management of Kubernetes deployments.

6. **Start a K8s Cluster locally**

```
minikube start --memory=4098 --driver=docker
```

7. **If you don't have Operator Lifecycle Manager (OLM) installed, install from here:**

```
curl -sL https://github.com/operator-framework/operator-lifecycle-manager/releases/download/v0.25.0/install.sh | bash -s v0.25.0
```

8. **Then follow the steps below**

```
vim argocd-basic.yml
```

Insert this inside vim editor:

```
apiVersion: argoproj.io/v1alpha1
kind: ArgoCD
metadata:
  name: example-argocd
  labels:
    example: basic
spec: {}
```

Now apply this file and check if your pods have started running or not:

```
kubectl get pods -n operators
```

9. **Change ClusterIP to NodePort for argocd-service and get the external link for the same**

10. **Now configure your ArgoCD with your project repository and VOILA!!!**

<p align = "center">
  <img src = "https://github.com/ShivangShandilya/golang-api/assets/101946115/b2513cc9-dc95-42ef-9a43-6a85dc8a3f58" height = 400 width = 800 />
</p>

- To learn more about Argo CD and configure it in your cluster, refer the [documentation](https://argo-cd.readthedocs.io/en/stable/).

### CONTRIBUTE

This project is entirely open-source, if you spot any errors in documnetation or codebase or just wanna suggest an improvement in project, feel free to raise an issue.

### LICENSE

This project is licensed under [MIT License](https://github.com/ShivangShandilya/golang-api/blob/main/LICENSE).
