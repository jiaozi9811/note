# RBAC

[TOC]

## RBAC(role-based access control)
基于角色的访问控制(role-based access control)使用”rbac.authorization.k8s.io” API Group实现授权决策，允许管理员通过Kubernetes API动态配置策略

## Role与ClusterRole
RBAC API中,一个角色包含了一套表示一组权限的规则
- 角色可以由namespace内的Role对象定义
- 整个kubernetes集群范围的有效角色通过ClusterRole对象实现


- Role对象只能用于授予对某个单一namespace中的资源的访问权限
- ClusterRole对象可以授予与Role相同的权限,但由于它们属于集群范围对象，也可以使用它们授予对以下资源的访问权限
- - 集群范围的资源(如节点node)
- - 非资源类型Endpoint(如/healthz)
- - 跨所有namespace的命名空间范围的资源(如pod,需要运行命令kubectl get pods --all-namespaces来查询集群中所有的pod)

## RoleBinding与ClusterRoleBinding
角色绑定将一个角色中定义的权限授予一个或一组用户
**角色绑定包含了一组相关主体(即subject,包括user,group,or serviceAccount)以及对被授予角色的引用**




## service account
service account是为了方便pod里面的进程调用kubernetes API或其他外部服务而设计的

创建service account
```
kubectl create serviceaccount jenkins
kubectl get serviceaccounts jenkins -o yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  creationTimestamp: 2017-05-27T14:32:25Z
  name: jenkins
  namespace: default
  resourceVersion: "45559"
  selfLink: /api/v1/namespaces/default/serviceaccounts/jenkins
  uid: 4d66eb4c-42e9-11e7-9860-ee7d8982865f
secrets:
- name:jenkins-token-l9v7v
```
