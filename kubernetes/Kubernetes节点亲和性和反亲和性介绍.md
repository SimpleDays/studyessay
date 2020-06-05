---
title: Kubernetes节点亲和性和反亲和性介绍
date: 2020-06-05 15:40:00
tags: kubernetes
author: 小笼包
categories: Kubernetes
---

## Kubernetes节点亲和性和反亲和性介绍

> 作者: 小笼包  
> 2020-06-05 大雨  

## nodeAffinity

## 官方文档

[将 Pod 分配给节点](https://kubernetes.io/zh/docs/concepts/configuration/assign-pod-node/)  

## 亲和与反亲和

亲和/反亲和功能极大地扩展了你可以表达约束的类型。关键的增强点是:

- 语言更具表现力（不仅仅是“完全匹配的 AND”）
- 你可以发现规则是“软”/“偏好”，而不是硬性要求，因此，如果调度器无法满足该要求，仍然调度该 pod
- 你可以使用节点上（或其他拓扑域中）的 pod 的标签来约束，而不是使用节点本身的标签，来允许哪些 pod 可以或者不可以被放置在一起。  

亲和功能包含两种类型的亲和，即“节点亲和”和“pod 间亲和/反亲和”。节点亲和就像现有的 nodeSelector（但具有上面列出的前两个好处），然而 pod 间亲和/反亲和约束 pod 标签而不是节点标签（在上面列出的第三项中描述，除了具有上面列出的第一和第二属性）。

**requiredDuringSchedulingIgnoredDuringExecution**  

- 硬性条件必需满足

**preferredDuringSchedulingIgnoredDuringExecution**  

- 软性条件，属于偏好部署(优先匹配原则)

<!-- more -->

## nodeSelector

nodeSelector 是节点选择约束的最简单推荐形式。nodeSelector 是 PodSpec 的一个字段。它指定键值对的映射。为了使 pod 可以在节点上运行，节点必须具有每个指定的键值对作为标签（它也可以具有其他标签）。最常用的是一对键值对。  

 **示例 :**  

**1、添加标签到节点**   

 执行 kubectl get nodes 命令获取集群的节点名称。选择一个你要增加标签的节点，然后执行 **kubectl label nodes [node-name] [label-key]=[label-value]**  命令将标签添加到你所选择的节点上。例如，如果你的节点名称为 **kubernetes-foo-node-1.c.a-robinson.internal**并且想要的标签是 **disktype=ssd**，则可以执行 **kubectl label nodes kubernetes-foo-node-1.c.a-robinson.internal disktype=ssd** 命令。  

 你可以通过重新运行 **kubectl get nodes --show-labels** 并且查看节点当前具有了一个标签来验证它是否有效。你也可以使用 **kubectl describe node "nodename"** 命令查看指定节点的标签完整列表。  

**2、添加nodeSelector字段到pod配置中**    

``` yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    env: test
spec:
  containers:
  - name: nginx
    image: nginx
    imagePullPolicy: IfNotPresent
  nodeSelector:
    disktype: ssd
```

当通过yaml文件进行部署时，pod 将会调度到将标签添加到的节点上。你可以通过运行 kubectl get pods -o wide 并查看分配给 pod 的 “NODE” 来验证其是否有效。  

**3、节点亲和的pod的实例**  

``` yaml
apiVersion: v1
kind: Pod
metadata:
  name: with-node-affinity
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: kubernetes.io/e2e-az-name
            operator: In
            values:
            - e2e-az1
            - e2e-az2
      preferredDuringSchedulingIgnoredDuringExecution:
      - weight: 1
        preference:
          matchExpressions:
          - key: another-node-label-key
            operator: In
            values:
            - another-node-label-value
  containers:
  - name: with-node-affinity
    image: k8s.gcr.io/pause:2.0
```

此节点亲和规则表示，pod 只能放置在具有标签键为 kubernetes.io/e2e-az-name 且 标签值为 e2e-az1 或 e2e-az2 的节点上。另外，在满足这些标准的节点中，具有标签键为 another-node-label-key 且标签值为 another-node-label-value 的节点应该优先使用。  

你可以在上面的例子中看到 In 操作符的使用。新的节点亲和语法支持下面的操作符： In，NotIn，Exists，DoesNotExist，Gt，Lt。你可以使用 NotIn 和 DoesNotExist 来实现节点反亲和行为，或者使用节点污点将 pod 从特定节点中驱逐。  

- In：label 的值在某个列表中
- NotIn：label 的值不在某个列表中
- Gt：label 的值大于某个值
- Lt：label 的值小于某个值
- Exists：某个 label 存在
- DoesNotExist：某个 label 不存在

如果你同时指定了 nodeSelector 和 nodeAffinity，两者必须都要满足，才能将 pod 调度到候选节点上。  

如果你指定了多个与 nodeAffinity 类型关联的 nodeSelectorTerms，则如果其中一个 nodeSelectorTerms 满足的话，pod将可以调度到节点上。  

如果你指定了多个与 nodeSelectorTerms 关联的 matchExpressions，则只有当所有 matchExpressions 满足的话，pod 才会可以调度到节点上。  

如果你修改或删除了 pod 所调度到的节点的标签，pod 不会被删除。换句话说，亲和选择只在 pod 调度期间有效。  

**preferredDuringSchedulingIgnoredDuringExecution** 中的 **weight** 字段值的范围是 1-100。对于每个符合所有调度要求（资源请求，RequiredDuringScheduling 亲和表达式等）的节点，调度器将遍历该字段的元素来计算总和，并且如果节点匹配对应的MatchExpressions，则添加“权重”到总和。然后将这个评分与该节点的其他优先级函数的评分进行组合。总分最高的节点是最优选的。(我理解就是，调度器把软亲和的MatchExpressions匹配的节点把当前**weight**字段值进行累加计算，最后通过各个匹配规则计算出最终结果，累加值最高的节点会是最优先选择的亲和性节点，比如有 “A”、“B”、“C”三个节点匹配规则结果分别是 “A”-3，"B"-2，”C“-0，最终会选择节点”A“进行部署）。  

## pod间亲和与反亲和

pod 间亲和与反亲和使你可以基于已经在节点上运行的 pod 的标签来约束 pod 可以调度到的节点，而不是基于节点上的标签。  

> **提示：实现内部亲和与反亲和需要数量可观的计算步骤，会明显降低pod调度的速度，集群规模越大、节点越多降速越明显，呈指数级增长，需要在使用此特性时考虑对调度速度的影响。**  

**内部pod亲和用podAffinity字段表示，内部pod反亲和用podAntiAffinity字段表示**  

requiredDuringSchedulingIgnoredDuringExecution 和 preferredDuringSchedulingIgnoredDuringExecution 两种配置。  

Pod 亲和与反亲和的合法操作符有 In，NotIn，Exists，DoesNotExist。  

topologyKey 是节点标签的键以便系统用来表示这样的拓扑域。  

原则上，topologyKey 可以是任何合法的标签键。然而，出于性能和安全原因，topologyKey 受到一些限制：  

- 对于亲和与 requiredDuringSchedulingIgnoredDuringExecution 要求的 pod 反亲和，**topologyKey 不允许为空**。

- 对于 requiredDuringSchedulingIgnoredDuringExecution 要求的 pod 反亲和，准入控制器 LimitPodHardAntiAffinityTopology 被引入来限制 topologyKey 不为 kubernetes.io/hostname。如果你想使它可用于自定义拓扑结构，你必须修改准入控制器或者禁用它。

- 对于 preferredDuringSchedulingIgnoredDuringExecution 要求的 pod 反亲和，空的 topologyKey 被解释为“所有拓扑结构”（这里的“所有拓扑结构”限制为 kubernetes.io/hostname，failure-domain.beta.kubernetes.io/zone 和 failure-domain.beta.kubernetes.io/region 的组合）。

- 除上述情况外，topologyKey 可以是任何合法的标签键。

除了 labelSelector 和 topologyKey，你也可以指定表示命名空间的 namespaces 队列，labelSelector 也应该匹配它（这个与 labelSelector 和 topologyKey 的定义位于相同的级别）。如果忽略或者为空，则默认为 pod 亲和/反亲和的定义所在的命名空间。  

所有与 requiredDuringSchedulingIgnoredDuringExecution 亲和与反亲和关联的 matchExpressions 必须满足，才能将 pod 调度到节点上。  

**示例1**

假设集群有五个工作节点，部署一个web应用，假设其用redis作内存缓存，共需要三个副本，通过反亲和将三个redis副本分别部署在三个不同的节点上，提高可用性，Deployment配置如下：  

``` yaml
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: redis-cache
    spec:
      selector:
        matchLabels:
          app: store
      replicas: 3
      template:
        metadata:
          labels:
            app: store
        spec:
          affinity:
            podAntiAffinity:
              requiredDuringSchedulingIgnoredDuringExecution:
              - labelSelector:
                  matchExpressions:
                  - key: app
                    operator: In
                    values:
                    - store
                topologyKey: "kubernetes.io/hostname"
          containers:
          - name: redis-server
            image: redis:3.2-alpine
```

反亲和会阻止同相的redis副本部署在同一节点上。

**示例2**

现在部署三个nginx web前端，要求三个副本分别部署在不同的节点上，通过与上例相似的反亲和实现。同时需要将三个web前端部署在其上已经部署redis的节点上，降低通信成本，通过亲和实现，配置如下：  

``` yaml
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: web-server
    spec:
      selector:
        matchLabels:
          app: web-store
      replicas: 3
      template:
        metadata:
          labels:
            app: web-store
        spec:
          affinity:
            podAntiAffinity:
              requiredDuringSchedulingIgnoredDuringExecution:
              - labelSelector:
                  matchExpressions:
                  - key: app
                    operator: In
                    values:
                    - web-store
                topologyKey: "kubernetes.io/hostname"
            podAffinity:
              requiredDuringSchedulingIgnoredDuringExecution:
              - labelSelector:
                  matchExpressions:
                  - key: app
                    operator: In
                    values:
                    - store
                topologyKey: "kubernetes.io/hostname"
          containers:
          - name: web-app
            image: nginx:1.12-alpine
```

## nodeName

nodeName 是节点选择约束的最简单方法，但是由于其自身限制，通常不使用它。nodeName 是 PodSpec 的一个字段。如果它不为空，调度器将忽略 pod，并且运行在它指定节点上的 kubelet 进程尝试运行该 pod。因此，如果 nodeName 在 PodSpec 中指定了，则它优先于上面的节点选择方法。  

使用 nodeName 来选择节点的一些限制：  

- 如果指定的节点不存在,pod将会调度失败并且其原因: 节点找不到或者 节点不存在。  

- 如果指定的节点没有资源来容纳 pod，pod 将会调度失败并且其原因将显示为，比如 OutOfmemory 或 OutOfcpu。  

- 云环境中的节点名称并非总是可预测或稳定的。  

## nodeAffinity、podAffinity、podAnitAffinity策略比较

| 调度策略 | 匹配标签 | 操作符 | 拓扑域支持 | 调度目标 |
| -------- | -------- | -------- | -------- | -------- |
| nodeAffinity     | 主机     | In, NotIn, Exists, DoesNotExist, Gt, Lt     |  否   | 指定主机     |
| podAffinity     | POD     | In, NotIn, Exists, DoesNotExist    |  是   | POD与指定POD同一拓扑域     |
| podAnitAffinity     | POD     | In, NotIn, Exists, DoesNotExist    |  是   | POD与指定POD不在同一拓扑域     |

## 污点（Taints）与容忍（tolerations）

对于nodeAffinity无论是硬策略还是软策略方式，都是调度 POD 到预期节点上，而Taints恰好与之相反，如果一个节点标记为 Taints ，除非 POD 也被标识为可以容忍污点节点，否则该 Taints 节点不会被调度pod。  

比如用户希望把 Master 节点保留给 Kubernetes 系统组件使用，或者把一组具有特殊资源预留给某些 POD，则污点就很有用了，POD 不会再被调度到 taint 标记过的节点。taint 标记节点举例如下：  

``` yaml
$ kubectl taint nodes 192.168.1.40 key=value:NoSchedule
node "192.168.1.40" tainted
```

如果仍然希望某个 POD 调度到 taint 节点上，则必须在 Spec 中做出Toleration定义，才能调度到该节点，举例如下：  

``` yaml
tolerations:
- key: "key"
operator: "Equal"
value: "value"
effect: "NoSchedule"
```

effect 共有三个可选项，可按实际需求进行设置：  

- NoSchedule：POD 不会被调度到标记为 taints 节点。

- PreferNoSchedule：NoSchedule 的软策略版本。

- NoExecute：该选项意味着一旦 Taint 生效，如该节点内正在运行的 POD 没有对应 Tolerate 设置，会直接被逐出。  

[参考文章](https://www.cnblogs.com/sunsky303/p/11130629.html)  
