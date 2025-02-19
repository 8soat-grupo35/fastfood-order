docker build -t fastfood-order-app:latest ../.

kubectl apply -f ./fastfood-order-secrets.yaml
kubectl apply -f ./fastfood-order-deployment.yaml
kubectl apply -f ./fastfood-order-service.yaml
kubectl apply -f ./fastfood-order-hpa.yaml