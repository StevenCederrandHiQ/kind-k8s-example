# Kind K8s Example Project
*Note: this project for educational purposes*

Hi, I know that nobody asked, but let me suggest how you might be able to learn to use kubernetes without having to go broke
using a cloud provider.

## About
Learning how to use Kubernetes shouldn't be hard or cost you an arm and a leg is my philosophy. So I've put together this 
small example project that uses KiND([Kubernetes in Docker](https://kind.sigs.k8s.io/)), in order to setup a kubernetes instance 
on your machine.

### Why KiND over ex minikube
KiND is simple, lightweight and easy to use. Making it perfect to learn work with it. You can easily scale up to have mutliple kind
clusters and so on. Using local docker images in KiND is a piece of cake compared to minikube(wtf..). It is also rather easy to configure
ingresses with KiND(imo).

## Setup

### Mandatory
A nice bonus with this project is that it should provide you with a good enough baseline to get started with the tools that are mandatory.

This project assumes that you have the following installed:
* [helm](https://helm.sh/)
* [Docker](https://docs.docker.com/get-docker/) **I recommend docker desktop**
* [go](https://go.dev/doc/install)
* [KiND](https://kind.sigs.k8s.io/)
* [kubectl](https://kubernetes.io/docs/reference/kubectl/)

#### Good To Have
* [k9s](https://k9scli.io/) or [lens](https://k8slens.dev/)

These(☝️) two tools makes it easier to visualise your kubernetes resources. But if you prefer doing things raw then `kubectl` is good enough.

## How To Run
Simplicity is key. So to deploy the entire project(including the example service), then just run:
`$ make maiden_deploy`

Now you can open up the postman collection *Http User Service Collection.json* in postman and send requests to the deployed service.
*Note: the user id is provided in the response header `id`*

After you are down then you can use to `$ make teardown` in order to clean up the reosurces.

## Deep Dive
What does this project do, and where can you take this?

### KiND
The KiND container is configured using the `kind.config.yaml`.
This is not a mandatory file when using KiND in general, but for this project it is. The `kind.config.yaml` configures the cluster control 
plane with an extra port mapping, routing all requests to `127.0.0.1:30080` to the container port `80`. The port `80` is then used by
the nginx proxy to route trafik to the correct service(by the use of the [custom ingress](./http-user-service/templates/http-user-service-ingress.yaml)).
It also sets up 3 worker nodes.
#### Excercise:
You can add to the `extraPortMappings` list:
``` yaml
  - containerPort: 30061
    hostPort: 30060
    protocol: TCP
```
Thus routing trafik from `localhost:30060` to the container port `30061` that you can then use to configure a custom service to use,
in order to directly send requests to it.

### Helm
Helm is used here in order to deploy the custom service `http-user-service`. This is done by templating all of the kubernetes resources
that it requires.
You can read up more on the role that helm plays on their website.
#### Example:
A good excercise would be write a very simple docker service, and update the helm templates to deploy your service to KiND.

##### Note:
You have to upload the docker image to KiND internally so that it can actually fetch the image and deploy it. The exception
being if you are pushing the image to ex: dockerhub but I doubt that you are.

You can load the images with the following command:
```sh
$ kind load docker-image <insert image name>:<insert image tag>
```

A variation to this would be to write your code, containerise it and package it in an entirely different folder and then 
fetching it here as a separate helm chart. 
