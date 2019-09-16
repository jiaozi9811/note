# secret 

[TOC]

## secret
secret解决了密码，token，秘钥等敏感数据的配置问题，而不需要把这些敏感数据暴露到镜像或pod spec中。secret可以以volume或环境变量的方式使用

## 类型
service account   用来访问Kubernetes API.由kubernetes自动创建，并且会自动挂载到pod的/run/secrets/kubernetes.io/serviceaccount目录中
Opaque            base64编码格式的secret，用于存储密码，秘钥等
kubernetes.io/dockerconfigjson  用于存储私有docker registry的认证信息


## 使用
创建好secret后，有两种方式使用
- 以volume方式
- 以环境变量方式
    
```
将secret挂载到volume中
apiVersion: v1
kind: Pod
metadata:
  labels:
    name: db
  name: db
spec:
  volumes:
  - name: secrets
    secret:
      secretName: mysecret
  containers:
  - image: gre.io/my_project_id/pg:v1
    name: db
    volumeMounts:
    - name: secrets
      mountPath: "/etc/secrets"
      readOnly: true
    ports:
    - name: cp
      containerPort: 5432
      hostPort: 5432
```      

```
将secret导出到环境变量中
apiVersion:extensions/v1beta1
kind: Deployment
metadata:
  name: wordpress-deployment
spec:
  replicas: 2
  strategy:
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: wordpress
        visualize: "true"
    spec:
      containers:
      - name: "wordpress:
        image: "wordpress"
        ports:
        - containerPort: 80
        env:
        - name: WORDPRESS_DB_USER
          valueFrom:
            secretKeyRef:
              name: mysecret
              key: username
        - name: WORDPRESS_DB_PASSWORD
          valueFrom:
            secretKeyRef:
              name: mysecret
              key: password
```
