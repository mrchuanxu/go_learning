## 隔离与限制🚫
#### 隔离
Namespace技术实际上修改了应用进程看待整个计算机的视图，也就是视线被操作系统做了限制🚫，只能看到某些指定的内容。但对于宿主机来说，这些被隔离的进程跟其他进程并没有太大的区别。<br>
![容器示意图](./img/k8s_docker.jpg)

docker其实就是用户运行在容器里的应用进程，跟宿主机上的其他进程一样，都由宿主机操作系统统一管理，只不过这些被隔离的进程拥有额外设置过的Namespace参数。Docker项目在这里扮演的角色，更多的是旁路式的辅助和管理工作。<br>
敏捷和高性能是容器相较于虚拟机最大的优势，也是它能够在PaaS这种更细粒度的资源管理平台上的大行其道的重要原因。<br>
#### Linux Namespace隔离🉐️不够彻底
1. 容器只是运行在宿主机上的一种特殊的进程，那么多个容器之间实用的就还是同一个宿主机的操作系统内核。<br>
内核是固定的，所以容器的运行系统也是固定的，所以wins不能执行linux容器，linux不能执行wins容器。<br>
2. linux内核种，很多资源都不能被Namespace化，例如时间。<br>
如果你的容器中的程序实用settimeofday(2)系统调用修改时间，整个宿主机的时间都会被随之修改，显然不符合用户预期。<br>
#### 限制🚫 Cgroups
Linux Cgroups的全称是Linux Control group。主要作用，就是限制一个进程组能够实用的资源上限，包括CPU、内存、磁盘、网络带宽等等。<br>
此外，Cgroups给用户暴露出来的操作借口是文件系统，即它以文件和目录的方式组织在操作系统的/sys/fs/cgroup路径下。
```sh
mount -t cgroup 
cpuset on /sys/fs/cgroup/cpuset type cgroup (rw,nosuid,nodev,noexec,relatime,cpuset)
cpu on /sys/fs/cgroup/cpu type cgroup (rw,nosuid,nodev,noexec,relatime,cpu)
cpuacct on /sys/fs/cgroup/cpuacct type cgroup (rw,nosuid,nodev,noexec,relatime,cpuacct)
blkio on /sys/fs/cgroup/blkio type cgroup (rw,nosuid,nodev,noexec,relatime,blkio)
memory on /sys/fs/cgroup/memory type cgroup (rw,nosuid,nodev,noexec,relatime,memory)
...
```
可以看到，/sys/fs/cgroup下面有很多诸如cpuset、cpu、memory这样的子目录，也叫子系统。比如
```sh
ls /sys/fs/cgroup/cpu
cgroup.clone_children 
cpu.cfs_period_us 
cpu.rt_period_us  
cpu.shares notify_on_release
cgroup.procs     
cpu.cfs_quota_us  
cpu.rt_runtime_us 
cpu.stat  tasks
```
cfs_period和cfs_quota这组合，可以用来限制进程在长度为cfs_period的一段时间内，只能被分配到总量为cfs_quota的CPU时间。<br>
```sh
root@ubuntu:/sys/fs/cgroup/cpu$ mkdir container
root@ubuntu:/sys/fs/cgroup/cpu$ ls container/
cgroup.clone_children cpu.cfs_period_us cpu.rt_period_us  cpu.shares notify_on_release
cgroup.procs      cpu.cfs_quota_us  cpu.rt_runtime_us cpu.stat  tasks
```
可以发现，操作系统会在新创建的container目录下，自动生成该子系统对应的资源限制文件。<br>
* Cgroups blkio为块设备设定IO限制，一般用于磁盘等设备。
* cpuset，为进程分配单独的CPU核和对应的内存节点
* memory 为进程设定内存实用的限制

Linux Cgroups的设计还是比较易用，子系统目录加上一组资源限制文件的组合。至于控制组下面的资源文件填写啥，得靠用户
```sh
docker run -it --cpu-period=100000 --cpu-quota=20000 ubuntu /bin/bash
```
启动容器后查看
```sh
cat /sys/fs/cgroup/cpu/docker/5d5c9f67d/cpu.cfs_period_us 
100000
cat /sys/fs/cgroup/cpu/docker/5d5c9f67d/cpu.cfs_quota_us 
20000
```
修复容器中的top指令以及/proc文件系统中的信息 lxcfs<br>
top是从/prof/stats目录下获取数据，

