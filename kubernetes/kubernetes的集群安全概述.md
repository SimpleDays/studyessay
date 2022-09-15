---
title: kubernetes的集群安全概述
date: 2022-09-15 17:53:00
tags: kubernetes
author: 小笼包
categories: Kubernetes
---

> 作者: 小笼包  
> 2022-09-15 大雨  

## k8s的集群安全机制概述

- 访问k8s集群时候，需要经过三个步骤完成，“认证”， “鉴权（授权）”， “准入控制”。

- 进行访问的时候，过程中都需要经过apiserver，apiserver做统一的协调，类似门卫。访问过程中需要证书、token、用户名+密码，如果访问pod，需要serviceAccount。

### 认证（传输安全）

- 传输安全： 对外不暴露8080端口，只能内部访问，对外使用端口6443

- 认证：
  
    客户端身份认证方式：

    一、https证书，基于ca证书。

    二、http token认证，通过token识别用户。

    三、http基本认证，用户名+密码认证。

### 鉴权（授权）

- 基于RBAC进行鉴权操作。

- 基于角色访问控制。

### 准入控制

- 就是准入控制器的列表， 如果列表有请求的内容，通过，没有就拒绝。

## 什么是RBAC

RBAC就是基于角色的控制。

**角色**： Role ClusterRole
**主体**： user、 group、 serviceaccount

角色与主体进行绑定，称为角色绑定。

角色的规划： 资源 （pod， node......），可以拥有（get，create）等操作。

![图一](./diagram/k8s-rbac.dio)

### 角色

- Role： 特定的命名空间访问权限

- ClusterRole： 所有的命名空间访问的权限

### 角色绑定

- roleBinding： 把角色绑定到主体上

- ClusterRoleBinding： 把集群的角色绑定到主体

### 主体  

- user： 用户

- group： 用户组

- serviceAccount： 服务账户

### 实际操作授权下

#### 1、创建一个待授权的k8s命名空间

```shell
[root@master01 dongliang]# kubectl create ns roledemo
namespace/roledemo created
[root@master01 dongliang]# 
```

确认查看下创建命名空间是否存在

```shell
[root@master01 dongliang]# kubectl get ns
NAME                   STATUS   AGE
default                Active   2d1h
kube-flannel           Active   2d1h
kube-node-lease        Active   2d1h
kube-public            Active   2d1h
kube-system            Active   2d1h
kubernetes-dashboard   Active   2d
roledemo               Active   28s
[root@master01 dongliang]# 
```

#### 2、在新的命名空间下创建一个pod

```shell
[root@master01 dongliang]# kubectl run nginx --image=nginx -n roledemo
pod/nginx created
[root@master01 dongliang]# 
```

查看确认下pod是否创建成功

```shell
[root@master01 dongliang]# kubectl get pods -n roledemo -o wide
NAME    READY   STATUS    RESTARTS   AGE   IP             NODE     NOMINATED NODE   READINESS GATES
nginx   1/1     Running   0          81s   10.244.1.107   node01   <none>           <none>
[root@master01 dongliang]# 
```

#### 3、创建一个可以访问roledemo的角色

我这里配置了一个role-rbac.yaml的文件来配置这个角色。

```yaml
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  namespace: roledemo
  name: jacky
rules:
- apiGroups: [""] # "" indicates the core API group
  resources: ["pods"]
  verbs: ["get", "watch", "list"]
```

根据这个yaml文件，可以看到角色“jacky”只提供命名空间为“roledemo”访问pod的权限，并且之后获取和监听、列表等操作信息。

然后我通过命令来创建这个角色

```shell
[root@master01 dongliang]# kubectl apply -f role-rbac.yaml 
role.rbac.authorization.k8s.io/jacky created
[root@master01 dongliang]#
```

查看下这个角色是否已经创建成功

```shell
[root@master01 dongliang]# kubectl get role -n roledemo
NAME    CREATED AT
jacky   2022-09-15T10:11:00Z
[root@master01 dongliang]# 
```

#### 4、创建角色的绑定

我这里配置了一个rolebinding-rbac.yaml的文件来配置这个角色绑定。

```yaml
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: read-pods
  namespace: roledemo
subjects:
- kind: User
  name: jacky # Name is case sensitive
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: pod-reader
  apiGroup: rbac.authorization.k8s.io
```

根据上面yaml文件可以看到角色绑定了一个类型为User的用户名叫“jacky”的角色，并为此角色绑定起名为read-pods。

通过命令创建角色绑定

```shell
[root@master01 jacky]# kubectl apply -f rolebinding-rbac.yaml 
rolebinding.rbac.authorization.k8s.io/read-pods created
[root@master01 jacky]# 
```

查看确认角色绑定已经成功

```shell
[root@master01 jacky]# kubectl get rolebinding -n roledemo
NAME        ROLE              AGE
read-pods   Role/pod-reader   2m19s
[root@master01 jacky]#
```

#### 5、使用证书识别身份

通过shell脚本文件去生成角色的token。

```shell
# 创建ca-config.json配置文件
cat > ca-config.json <<EOF
{
    "signing": {
        "default": {
            "expiry": "43800h"
        },
        "profiles": {
            "server": {
                "expiry": "43800h",
                "usages": [
                    "signing",
                    "key encipherment",
                    "server auth"
                ]
            },
            "client": {
                "expiry": "43800h",
                "usages": [
                    "signing",
                    "key encipherment",
                    "client auth"
                ]
            },
            "peer": {
                "expiry": "43800h",
                "usages": [
                    "signing",
                    "key encipherment",
                    "server auth",
                    "client auth"
                ]
            },
            "kubernetes": {
                "expiry": "43800h",
                "usages": [
                    "signing",
                    "key encipherment",
                    "server auth",
                    "client auth"
                ]
            },
            "etcd": {
                "expiry": "43800h",
                "usages": [
                    "signing",
                    "key encipherment",
                    "server auth",
                    "client auth"
                ]
            }
        }
    }
}
EOF

cat > jacky-csr.json <<EOF
{
    "CN": "jacky",
    "hosts":[],
    "key": {
       "algo": "rsa",
       "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "L": "BeiJing",
            "ST": "BeiJing"
        }
    ]
}
EOF

cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes jacky-csr.json | cfssljson -bare jacky

kubectl config set-cluster kubernetes \
  --certificate-authority=ca.pem \
  --embed-certs=true \
  --server=https://192.168.20.130:6443 \
  --kubeconfig=jacky-kubeconfig

kubectl config set-credentials jacky \
  --client-key=jacky-key.pem \
  --client-certificate=jacky.pem \
  --embed-certs=true \
  --kubeconfig=jacky-kubeconfig

kubectl config set-context default \
  --cluster=kubernetes \
  --user=jacky \
  --kubeconfig=jacky.kubeconfig

kubectl config use-context default --kubeconfig=jacky-kubeconfig
```

目前对k8s的ca申请权限还有很多疑问，暂时写到这里
