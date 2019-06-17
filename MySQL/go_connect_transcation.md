## 一个事务的故事
![hello](./img/xiaoming.png)

小明要给小王转100块钱，而银行卡只有100块钱，银行就要做增删查改，一个事务完成。一个事物要么是完成，要么是失败。在MySQL中，事物支持是在引擎层实现的。 **MySQL是一个支持多引擎的系统，但并不是所有的引擎都支持事务。** 比如MySQL原生的MyISAM引擎就不支持事务，这也是MyISAM被InnoDB取代的重要原因之一。<br>
#### 隔离性与隔离级别
事务ACID肯定要想到(Atomicity, Consistency, Isolation, Durability), 其中I，就是隔离性。<br>
当数据库上有多个事务同时执行的时候，就可能出现脏读(dirty read)、不可重复读(non-repeatable read)、幻读(phantom read)问题，为了解决这些问题，就有了隔离级别的概念。<br>
##### 隔离得越严实，效率就越低！
SQL标准的事务隔离级别包括: A.读未提交(read uncommited)、B.读提交(read committed)、C.可重复读(Repeatable read)和D.串行化(serializable)。<br>
* A. 一个事务还没提交，它做的变更就能被别的事务看到
* B. 一个事务提交之后，它做的变更才会被其他事务看到
* C. 一个事务执行过程中看到的数据，总是跟这个事务在启动时看到的数据是一致的。当然在可重复读隔离的级别下，未提交的变更对其他事务也是不可见的。
* D. 对于同一行记录，写会加写锁，读会加读锁。当出现读写锁冲突时，后访问的事务必须等前一个事务执行完成，才能继续执行

实际上， **数据库会创建一个视图，访问的时候以视图的逻辑结果为准。** 在C隔离级别下，这个视图是在事务启动时创建的，整个事务存在期间都用这个视图。 在B 隔离级别下，这个视图是在每个SQL语句开始执行时创建的。也就是说，每个SQL语句都需要执行一次创建视图，这样每个视图就会因为某个事务执行阶段的数据变换而不同。A隔离级别下，直接返回记录上的最新值，没有视图概念。 D隔离级别下，直接用加锁方式避免并行访问。<br>

##### transaction-isolation的值设置成READ-COMMITTED
```mysql
mysql> show variables like 'transaction_isolation';

+-----------------------+----------------+

| Variable_name | Value |

+-----------------------+----------------+

| transaction_isolation | READ-COMMITTED |

+-----------------------+----------------+

```
#### 事务隔离的实现
在MySQL中，实际上每一条记录在更新的时候都会同时记录一条回滚操作。记录上的最新值，通过回滚操作，都可以得到前一个状态的值。<br>
不同时刻启动的事务会有不同的read-view。同一条记录在系统中可以村早多个版本。这就是数据库的多版本并发控制(MVCC)。<br>
⚠️ 删除这里说的不是很清楚❓<br>
长事务意味着系统里面会存在很老的事务视图。由于事务随时可能访问数据库里面的任何数据，所以这个事务提交之前，数据库里面它可能用到的回滚记录都必须保留，这就会导致大量占用存储空间。<br>
长事务占用资源，可能拖垮整个库<br>
#### 事务的启动方式
1. 显式启动begin或start transaction...commit...rollback
2. set autocmmit=0这个命令会将这个线程自动提交关掉。意味着如果你只执行一个select语句，这个事务就启动了，而且并不会自动提交。这个事务持续存在直到你主动执行commit或rollback语句，或者断开链接

set autocommit=0的命令，导致如果是长链接，就会导致以为的长事务，所以建议设为1。<br>
可以在information_schema库innodb_trx这个表中查询长事务。
```mysql
select * from information_schema.innodb_trx where TIME_TO_SEC(timediff(now(),trx_started))>60
```


