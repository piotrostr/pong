# pong server

Generally, there is a number of ways of getting traffic into the kubernetes
cluster, mostly depending on where it is to be deployed.

1. locally using the `manifest-gke.yaml` and minikube or k3s

   ```sh
   kubectl apply -f manifest-gke.yaml
   ```

   If there is a service, it can be forwarded straight to, as done in the
   snippet above.

   Otherwise, the deployment has to be exposed to become a service.

   ```sh
   kubectl expose deployment [deployment]
   ```

   It can then be forwarded a port from the node:

   ```sh
   kubectl port-forward service/[deployment] [port-host]:[port-container]
   ```

   This could potentially be useful for debugging a single deployment.

   There is also an option to use the `kubectl proxy` and interact directly
   with the local kubernetes api, but it is kind of hassleful compared to the
   other options above.

2. using docker desktop kubernetes provider (single node)

   This is a nice one, as since docker-desktop uses vpnkit to expose any load
   balancers and forward traffic into the cluster.

   ```sh
   kubectl apply -f manifest-docker.yaml
   ```

   Making it ready to be `curl`'ed.

3. On GKE (Google Kubernetes Engine)

   After including the ingress resource in the manifest and changing the
   `LoadBalancer` service to `NodePort`, the manifest can be used to provision
   a cluster on GCP cloud.

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

   The load balancer will be provisioned from GCP and will forward any traffic
   into the cluster.

## TODO

Look into [nginx](https://kubernetes.github.io/ingress-nginx/).
