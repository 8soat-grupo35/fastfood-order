docker build -t fastfood-order-app:latest ../.
kubectl apply -f ./postgres-dbinit-configmap.yaml
kubectl apply -f ./postgres-pv.yaml
kubectl apply -f ./postgres-pvc.yaml
kubectl apply -f ./postgres-deploy.yaml
kubectl apply -f ./postgres-service.yaml
kubectl apply -f ./fastfood-deployment.yaml
kubectl apply -f ./fastfood-service.yaml
kubectl apply -f ./fastfood-hpa.yaml