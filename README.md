# pong server

Let's assume there is an API ready to be deployed and one decides to
use the industry leading orchestrator - Kubernetes.

Originating from Google, kubernetes in greek means helmsman or pilot and it
serves exactly that purpose - commanding the fleet of, in this case,
containers.

The core concepts are easy to grasp, but since running the "real" or more
formally, the production cluster, requires a range of components, most of the
time one gets started with the development environment alternatives, like
`minikube`, `microk8s` or the built-in (enableable) Docker Desktop Kubernetes engine.

Even though it is very logical at first and it's easy to get started, it might
be a bit difficult to enable access to the cluster without port forwarding.
Personally, I am not a big fan of port forwarding and thus even developing
local applications, I stand by developing them in a way that makes them ready
to be deployed into production without any breaking changes.

Generally, there is a number of ways of getting traffic into the kubernetes
cluster, mostly depending on where it is to be deployed.

1. locally using minikube or k3s

   ```sh
   kubectl create deployment [deployment] --image=piotrostr/pong
   kubectl expose deployment [deployment]
   ```

   It can then be forwarded a port from the node:

   ```sh
   kubectl port-forward service/[deployment] [port-host]:[port-container]
   ```

   This could potentially be useful for debugging a single deployment.

   There is also an option to use the `kubectl proxy` and interact directly
   with the local kubernetes api, but it is kind of hassleful compared to the
   other options below.

2. Using Docker Desktop Kubernetes provider

   This is a nice one, as since docker-desktop uses vpnkit to expose any load
   balancers and forward traffic into the cluster.

   ```sh
   kubectl apply -f manifest-docker.yaml
   ```

   Making it ready to be `curl`'ed.

   I would say this is the go-to for debugging simple applications.

3. On Google Kubernetes Engine (EKS) [`manifest-gke.yaml`]

   Note: requires the `gcloud` to be configured with the right project and GKE
   enabled.

   After including the ingress resource in the the manifest can be used to
   provision a cluster on GCP cloud quite seamlessly.

   ```yaml
   apiVersion: networking.k8s.io/v1
   kind: Ingress
   metadata:
     name: minimal-ingress
   annotations:
     ingress.kubernetes.io/ssl-redirect: 'false'
   spec:
   rules:
     - http:
         paths:
           - path: /
             pathType: Exact
             backend:
               service:
                 name: pong-api
                 port:
                   number: 80
   ```

   Configure the `kubectl` to use the gcloud context:

   ```sh
   gcloud container clusters create-auto [cluster-name]
   gcloud container clusters get-credentials [cluster-name]
   ```

   After applying the yaml, the load balancer will be provisioned from GCP and will
   forward any traffic into the cluster.

4. Using the [nginx](https://kubernetes.github.io/ingress-nginx/) ingress
   [`manifest-nginx.yaml`]

Install it with

```sh
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace
```

By including the same ingress resource as in the _3._ and adding
`ingressClassName: nginx` under `spec` (in order to define which controller to
use) it allows external traffic into the cluster and deployments without GKE or
EKS (Elastic Kubernetes Service from AWS). This manifest can deployed on a
single node cluster on an virtual machine quite easily, enabling one to benefit
from the Kubernetes awesome features like auto-scaling and auto-healing while
not being forced to use the AWS/GCP load balancing services and cluster costs,
which can pile up for small applications.

Every api call is distributed between the five pods (change `replicas` in the)

Still, the preffered way should be to use GKE or EKS, since those offer
autoscaling and automatical provisioning of nodes.
