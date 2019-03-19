# dashboard

https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml

docker pull registry.cn-hangzhou.aliyuncs.com/tdycloud/kubernetes-dashboard-amd64:v1.10.1

docker tag registry.cn-hangzhou.aliyuncs.com/tdycloud/kubernetes-dashboard-amd64:v1.10.1 k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.1

docker rmi registry.cn-hangzhou.aliyuncs.com/tdycloud/kubernetes-dashboard-amd64:v1.10.1

## kubectl create -f kubernetes-dashboard.yaml

kubectl edit svc kubernetes-dashboard-*** -n kube-system
Type:NodePort

kubectl get svc -n kube-system
获得暴露的port

打开代理，否则只能从本地访问
kubectl proxy --address=local_ip --disable-filter=true &

http://10.0.0.10:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/login


创建ServiceAccount
kubectl create serviceaccount dashboard-admin -n default


绑定相关role
kubectl create clusterrolebinding dashboard-admin --clusterrole=cluster-admin --serviceaccount=default:dashboard-admin


获取ServiceAccount使用的Secret

kubectl describe sa dashboard-admin -n default
