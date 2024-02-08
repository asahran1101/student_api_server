cd $(dirname $0)/..

minikube start

kubectl delete namespace student
kubectl delete pv postgres-pv
helm install student-server ./student-api-server 

echo "The server is running now."