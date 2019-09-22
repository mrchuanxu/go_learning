## 谈谈kubernetes的本质 从容器到容器云
Docker 项目，一个容器，实际上是由Linux Namespace、Linux Cgroups和rootfs三种技术构建出来的进程隔离环境。Linux容器，其实可以被一分为二地看待
* 一组联合挂载在 /var/lib/docker/aufs/mnt 上的rootfs，这一部分我们称为 容器镜像(Container Image)，是容器的静态视图
* 一个由Namespace+Cgroups构成的隔离环境，这一部分我们称为 容器运行时(Container Runtime)，是容器的动态视图

在整个开发-测试-发布的流程中，真正承载着容器信息进行传递，是容器镜像，而不是容器运行。<br>
通过容器镜像，就可以和潜在用户(开发者)直接关联起来。<br>
从一个开发者和单一的容器镜像，到无数开发者和庞大的容器集群，容器技术实现了从容器到容器云的飞跃，标志着它真正得到了市场和生态的认可。<br>
#### 容器编排技术
容器从一个开发者手里的小工具，一跃成为云计算领域的绝对主角；而能够定义容器组织和管理规范的容器编排技术，当仁不让地坐上了容器技术领域的头把交椅。<br>
### Kubernetes 项目想要解决什么问题？ 编排？调度？容器云？还是集群管理？
有了应用的容器镜像，请帮我在一个给定的集群上把这个应用运行起来。希望k8s能够提供路由网管、水平扩展、监控、备灾恢复等一系列运维能力。<br>
Kubernetes全局架构。<br>
![kubernetes架构](./img/k8s_struct.png)

k8s项目架构由Master和Node两种节点组成，而这两种角色分别对应着控制节点和计算节点。<br>

* master由负责api服务的kube-apiserver、负责调度的kube-scheduler、以及负责容器编排的kube-controller manager。整个集群的持久化数据，则由kube-apiserver处理后保存在Etcd中。<br>

核心是kubelet组件。主要负责同容器运行时的(比如Docker项目)打交道。而这个交互所依赖的是CRI(Container RUntime Interface)的远程调用接口，这个接口定义了容器运行时的各项核心操作，比如: 启动一个容器需要的所有参数。<br>

### Pod
Pod 容器共享同一个Network Namespace、同一组数据卷，从而达到高效率交换信息的目的。<br>

![Net](./img/k8s_pod.png)
