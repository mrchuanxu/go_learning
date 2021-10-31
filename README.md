# go_learning
集合go的编程技巧 分布式 数据库MySQL Linux操作系统讲解 设计模式 架构思想总结 算法思想总结 还是持续搞起来，让这个项目更加丰满

## go 深入理解系列(多看点源码，看看人家的设计)

* [深入理解文本类型](./go_coding/go_text.md)
* [深入理解defer](./go_coding/go_defer.md)
* [实验楼游戏-2048](./go_project/g2048.go)
* [mysql增删查改实现](./go_project_mysql/mysql_test.go)
* [go深入理解interface{}](./go_coding/interface_test/go_interface.md)
* [go深入理解context]
* [go深入理解channel，锁]
* [go 认识IO](./go_lib/goio_test.go)
* [go相对路径示例]
* [goruntime详解]
* [go并发编程，一个简易聊天室](./go_bible/ch8/ch8chat/chat.go)
* [net/http源码研究以及深入实践](./gosource/net_http)

## go leetcode

* [two Sum，map的判断不是==，而是value,ok:=map[]...](./go_leetcode/two_sum.go)
* [Three Sum，三个指针 发现快排还是慢](./go_leetcode/three_Sum.go)
* [Three Sum Closet，发现太慢了！应该有更快的方法，但是空间用的贼少，典型时间换空间](./go_leetcode/threeSumCloset.go)
* [倒数第n个，k-n+1](./go_leetcode/theNththeend.go)
* [longestPalindrome最长回文，其实都是从中间开始向左右遍历，一个一个往后遍历](./go_leetcode/longestProline.go)
* [我地D普通人净可以谂到普通算法(x2-x1)*min(y2,y1)，指针最快](./go_leetcode/container_water.go)
* [MedianSortArray,其实要求O(log(m+n))  归并的思路可以尝试用一下 其实我觉得就是归并排序的变形](./go_leetcode/middienSortArray.go)
* [nQUeens我不知道为什么这个编译器是不是有问题，整个答案来说，就是求出唯一的，但是就是没想到怎么解决](./go_leetcode/nQueens.go)
* [rotateNums是array的一个经典题目，冒泡解法最慢，但是空间上是O(1)复杂度，O(k*n的复杂度)](./go_leetcode/roateArray.go)
* [来自momoso的笔试题，真的让我费劲脑子，写了一天，终于想清楚关系，并用两个递归实现了，瞬间露出老母亲的笑脸](./go_leetcode/flattentoNested.go)
* [需要了解清楚runtime的用法，调度Gosched,Goexit方法，交替打印数字与字母](./go_leetcode/chanSolve_test.go)
* [我真的服了，写了一天都没有推倒出来->k(2*numRows-2)](./go_leetcode/zigzag.go)

## sql 性能优化 leetcode

* [leftjoin 尽量选择小表来join，这样性能比较优](./MySQL/left_join.sql)

## go 数据结构与算法之美

* [链表实现CRUD，链表反转，环的检测，链表合并递归完美](./go_alorigthm/listgo/golist_test.go)
* [栈的实现，用切片实现，自动扩容](./go_alorigthm/stackgo/stackgo_test.go)
* [go实现冒泡，插入，选择排序](./go_alorigthm/sortgo/sortgo_test.go)
* [go实现快排，归并排序 求第K大元素 O(n)的时间复杂度](./go_alorigthm/quickMergeSort/quickMergeSort_test.go)
* [桶排序，计数排序，基数排序O(n)](./go_alorigthm/bucketSort/bucketgo_test.go)
* [二分查找以及其变体，十个二分九个错](./go_alorigthm/binarySearch/binarySearch_test.go)
* [跳表实现我就不实现了，我们需要知道这个性能可以媲美红黑树](https://github.com/wangzheng0822/algo/blob/master/go/17_skiplist/skiplist.go)
* [LRU缓存算法，实现了一个散列表和双向链表](./go_alorigthm/hashTable/hashTable_test.go)
* [二叉树中序遍历，前序遍历循环递归实现](./go_alorigthm/binaryTree/binaryTree_test.go)
* [普通二叉查找树，容易退化，衍生红黑树，红黑树平衡，但是实现难度大，但是性能求稳定，不求实现但求原理的理解以及解决的问题](./go_alorigthm/binaryTreeSearch/binaryTreeSearch_test.go)
* [堆排序，一个完全二叉树，基于数组，持续堆化，建堆排序，在数组的操作，一般搞下标就好了](./go_alorigthm/heapSort/heapSort_test.go)
* [图的广度与深度遍历](./go_alorigthm/graphgo/graph_test.go)
* [多模式串匹配实现敏感词过滤算法，trie树经典，比较有意思，等工作了再搞](./go_alorigthm/morePattern/morePattern_test.go)
* [贪心算法 我有贪心策略，此策略必有优解](./go_alorigthm/greedyAlgo/greedy_test.go)
* [分而治之，了解一下MapReduce的分治思想，遇到大问题，解决成小问题](./)
* [回溯算法(蝴蝶效应)8皇后与背包问题，要么选要么不选](./go_alorigthm/backDynamic/eightQueens_test.go)

#### 基础篇
* [循环队列实现](./go_date_program/20200514/queue_test.go)
## go k8s实战

* [容器，只不过是一种特殊的进程](./kubernetes/processasBegin.md)
* [容器，Cgroups限制，Namespace隔离](./kubernetes/depatchLimit.md)

## go MySQL系列

* [mysql是怎么执行语句?](./MySQL/go_connect_transcation.md)
* [mysql表级锁，行锁以及全局锁](./MySQL/lock.md)
* [mysql事务隔离？](./MySQL/transaction.md)
* [唯一索引与普通索引的选择,change buffer!=redo log](./MySQL/chooseIndex.md)
* [不好意思，一不小心刷完了45讲，没有写md，只能做一个脑图，来弥补一下缺失](./MySQL/knowledge.md)

## React 系列

## 深入浅出redis

## go—interview

* [面试题集锦1](./go_interview/bac.go)

### 最近在读编程珠玑

* 最近在读编程珠玑，打算每天做一些题目放上来，持续更新
* [第二章—二分-标识](./go_programpearls/ch1/ex1_test.go)

### go-高级编程
[string\slice\array](./go_highprogram/ch1/str_test.gp)

### go-设计模式
[多用组合少用继承](./go_design_pattern/lkp_test.go)
[DDD充血模型开发一个虚拟钱包](./go_design_pattern/wallet_test.go)
[单例模式](./go_design_pattern/singleholder.go)
[工厂模式](./go_design_pattern/factory.go)

### 深入理解计算机系统
[计算机系统漫游笔记](./CSAPP/ch1.md)
[程序结构和执行-信息的表示和处理](./CSAPP/ch2.md)

### 为什么系列
[为什么会64位系统调用不需要陷入内核态？](./Linux/system_call.md)
