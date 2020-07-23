### 计算机系统漫游
计算机系统是由硬件与系统软件组成，共同个工作来运行应用程序。<br>
[-] 信息就是位+上下文。1byte = 8bit。每个字节表示程序中的某些文本文件。<br>
[-] 程序被其他程序翻译成不同的格式 可执行目标文件。<br>
`gcc -o hello hello.c` 预处理器，编译器，汇编器和链接器 一起构成编译系统。<br>
hello.c(源程序文本)--预处理器cpp-->hello.i(修改了的源程序文本)--编译器ccl-->hello.s(汇编程序文本)--汇编器as-->hello.o+printf.o(可重定位目标程序 二进制)--链接器(ld)-->hello(可执行目标程序 二进制)<br>

##### 了解编译系统如何工作大有益处
* 优化程序性能
* 理解链接时出现的错误
* 避免安全漏洞

##### 处理器读并解析储存在内存中的指令
shell是一个命令行解析器。

##### 系统的硬件组成
* 总线

    总线贯穿整个系统的一组电子管道，携带信息字节并负责在各个部件间传递。通常设计成传送定长的字节块(字)。
* IO设备

    磁盘-控制器-适配器 通过总线相连。
* 主存

    临时存储设备 处理器执行程序时，用来存放程序和程序处理的数据。DRAM，线性字节数组，每个字节都有其唯一的地址(数组索引)。
* 处理器

    中央处理单元 解释或执行 存储在主存中指令的引擎。处理器的核心是一个大小为一个字(8bytes 或 4bytes)的存储设备(或寄存器)，称为程序计数器(PC)。在任何时刻，PC都指向主存中某条机器语言指令。<br>
    从通电到停电，处理器一直在不断地执行程序计数器指向的指令，再更新程序计数器，使其指向下一条指令。<br>
    处理器是由指令集架构模型决定的。在这个模型中，指令按照严格的顺序执行，而执行一条指令包含执行一系列的步骤。<br>
    处理器从程序计数器指向的内存处，读取指令，解释指令中的位，执行该指令指示的简单操作，然后更新PC，使其指向下一条指令，而这条指令并不一定和在内存中刚刚执行的指令相邻。<br>
    寄存器文件是一个小的存储设备，由一些单个字长的寄存器组成，每个寄存器都有唯一的命名。<br>
    ALU 计算新的数据和地址值。CPU在指令下可能会执行的操作，加载-存储-操作-跳转。<br>
    


