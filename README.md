# pong server

```bash
minikube start
kubectl apply -f ./manifest.yaml
kubectl port-forward services/pong-lb 8000:80
```
