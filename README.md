# pong server

```sh
minikube start
kubectl apply -f ./manifest.yaml
kubectl port-forward services/pong-lb 8000:80
```

Generally, there is a number of ways of getting traffic into the kubernetes cluster.

1. the simple way (using port forwarding)

   If there is a service, it can be forwarded straight to, as done in the
   snippet above.

   Otherwise, the deployment has to be exposed to become a service.

   ```sh
   kubectl expose deployment [deployment]
   ```

   it can then be forwarded a port from the node:

   ```sh
   kubectl port-forward service/[deployment] [port-host]:[port-container]
   ```

   this could potentially be useful for debugging a single deployment

2. the right way (using ingress)

   add the ingress resource to the manifest, here is the basic one it could
   also be an ingress from `helm`, like the
   [nginx](https://kubernetes.github.io/ingress-nginx/) one.

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
           - path: /hello/
             pathType: Prefix
             backend:
               service:
                 name: pong-api
                 port:
                   number: 80
   ```

   however in practice, it seems like there is a load balancer automatically
   assigned for any cluster from the cloud provider
