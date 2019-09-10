## 容器 其实就是一个特殊的进程
* 容器技术的兴起源于PaaS技术的普及
* Docker项目通过容器镜像，解决了应用大包的这个根本性难题

容器本身没有价值，有价值的是 “容器编排”<br>

容器其实是一种沙盒技术。沙盒就是能够像集装箱一样，将应用装起来的技术。<br>

### 边界实现手段
进程，程序运行起来后的计算机执行环境总和。<br>
容器技术的核心功能就是通过约束和修改进程的动态表现，从而为其创造出一个"边界"。<br>
Cgroups技术用于制造约束的主要手段，而Namespace技术则是用来修改进程视图的主要方法。<br>
先来一个命令 
```sh
docker run -it busybox /bin/sh
/ #
```
docker run，创建一个容器，-it告诉Docker项目在启动容器后，需要分配一个文本输入输出的环境，也就是TTY。跟容器的标准输入相关联，这样就可以和这个Docker容器进行交互了。而/bin/sh就是我们要在Docker 容器里面运行的程序。busybox就是这个容器的名字。大概意思就是请帮我启动一个容器，在容器里面执行/bin/sh，并且给我分配一个命令行终端跟这个容器交互。<br>
从这里开始，ubuntu就成了一个宿主机，运行着一个/bin/sh的容器，抱在这个宿主机里面。<br>
```sh
/ # ps
PID  USER   TIME COMMAND
  1 root   0:00 /bin/sh
  10 root   0:00 ps
```
会发现，Docker执行/bin/sh，就是一号进程。ps命令发现自己完全隔离在一个跟宿主机完全不同的世界当中。<br>
#### 障眼法clone()
为什么呢？我们知道Linux创建进程的系统调用是clone()，这个系统调用会创建一个新的进程，进程号就是pid。其实fork也可以，但是clone()系统调用可以在参数里面指定CLONE_NEWPID参数，这样新建的进程就会看到一个全新的进程空间，PID就是1。但是这只是一种“障眼法”。
```sh
int pid = clone(main_function, stack_size, CLONE_NEWPID|SIGCHLD, NULL); 
```
当然也可以执行多次clone，创建多个Namespace。处理这个，还有Mount，UTS，IPC，Network和User这些Namespace，用于对各种不同的进程上下文进行障眼法操作。Mount就是隔离挂载点信息。Network就是隔离网络设备和配置。<br>
因此，Docker容器其实就是创建容器进程时，指定了这个进程所需要启用的一组Namespace参数。这样容器就只能看到当前Namespace所限定的资源，文件，设备，状态或者配置。对于宿主机就看不到了。<br>