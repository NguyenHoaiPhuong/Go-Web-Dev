# Kubernetes example 01

## Commands

```
gcloud auth configure-docker

kubectl create secret docker-registry regcred --docker-server=https://asia.gcr.io/v2/ --docker-username=akagi2106 --docker-password=Nothingisimpossible --docker-email=hoaiphuong.nguyen.vn@gmail.com

kubectl get secret regcred --output=yaml
```

```
kubectl apply -f deployment.yml
kubectl apply -f service.yml
kubectl get pods
kubectl describe pod kubernetes-example01
kubectl get services
minikube service kubernetes-example01-service

kubectl delete deployment kubernetes-example01
kubectl delete service kubernetes-example01-service
```

```
kubectl create secret generic regcred \
    --from-file=.dockerconfigjson=~/.docker/config.json> \
    --type=kubernetes.io/dockerconfigjson

eval $(minikube docker-env)
```


## References:
https://www.digitalocean.com/community/tutorials/how-to-deploy-resilient-go-app-digitalocean-kubernetes