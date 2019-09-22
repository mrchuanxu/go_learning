## 深入理解容器镜像
Namespace和Cgroups，让人明白容器的本质是一种特殊的进程。容器进程看到的文件系统又是什么样子呢？
#### Mount Namespace，容器里的应用进程，理应看到一份完全独立的文件系统
![Docker全景图](./img/k8s_docjer2.png)