# jike-go

"我可以写代码一整天"  
![image](https://github.com/ProsperousLi/jike-go/blob/main/docs/u%3D787858893%2C1297713883%26fm%3D26%26gp%3D0.png)

go进阶训练营1期-作业-极客时间

## 第一周总结

1. SOA 微服务 ：小即是美；API先交付，再实现功能；一个服务只做一个功能（单一职责） 
2. 接口设计：版本兼容；数据接受要各种情况考虑；数据返回要尽可能的简洁精炼 
3. grpc的基本使用 

## 第二周作业

问： 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？ 

答： 需要由wrap包装抛给上层。因为当前调用者执行玩sql的语句后，语句报错，外部调用者可能想知道是哪个操作导致了sql.ErrNoRows,此时返回方可以使用wrap包装一层返回出去。同时在真正处理错误的地方，可以打印出调用堆栈信息，方便更快地定位问题。  
作业地址： https://github.com/ProsperousLi/jike-go/tree/main/week2  

## 第三周总结

1.go协程的正确打开方式：知道什么时候结束；知道怎么去结束它。  
2.生命周期管理  
3.内存模型  
4.COW ：copy on write  
5.store buffer  
6.automic.value 无锁访问共享内存  
7.源码errorgroup sync  

## 第三章作业

问：如何做一个应用的生命周期的管理？（errorgroup、wiatgroup；参考 https://github.com/go-kratos/kratos ）

答：TODO  
作业地址：https://github.com/ProsperousLi/jike-go/tree/main/week3  

问：基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出  

答： https://github.com/ProsperousLi/jike-go/tree/main/week3  

## 第四周总结  
1. 工程目录框架：  
2. grpc 大一统。 grp的TCP单线多路复用；头压缩，protobuf数据格式；
3. api 设计的向下兼容性；错误码的定义（大错误码+小错误信息）；参考 Google API Design Guide (谷歌API设计指南)；方法结构体字段部分更新 FieldMask
   protobuf+options： 使用protobuf接受配置文件的内容，然后调用转换方法返回真正需要使用配置的方法的结构体。
4. 《谷歌测试之道》 附下载链接 https://github.com/ProsperousLi/jike-go/blob/main/Google%E8%BD%AF%E4%BB%B6%E6%B5%8B%E8%AF%95%E4%B9%8B%E9%81%93.docx


## Week04 作业题目：

1. 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

思路：TODO 使用 gin框架搭建简单的web项目  

## 第五周总结  
1. 隔离  
   物理隔离；动静隔离。  
2. 超时：rpc到一定时间没返回，直接返回调用方超时  
3. 限流：  
   分布式限流：根据服务接口级别的优先级和当前资源的占用率来选择性的丢弃一些优先级较低的请求。 qps的分配优先满足于quota比较少的接口调用。  
   熔断：client端侧直接拒绝请求，概率放一些请求去rpc尝试，多个请求 成功则关闭熔断。 熔断的触发条件：qps达到一定程度，同时错误概率达到一定程度。  
        Gutter限流：以redis为例，某个redis触发熔断时，将多余的请求转发到其他小（小集群）redis，让小redis来处理。（双熔断，小redis触发熔断，流量会写回原来的redis）  
        移动端限流：例如双11，用户积极尝试不可达服务。限制请求的频次

## 第六周总结

评论系统设计
1. kafaka 削峰  
2. kafaka 某个partition因为热点事件，导致partition成为热点，如何解决？  
   进程自适应的去发现热点：使用滑动窗口（环形数组）去统计窗口内QPS超过一定次数的事件，将其缓存到我们的local cache，下次直接访问local cache
3. 归并回源 （1） 单飞 singleflight：多个人查询cache miss 导致缓存穿透的时候，只让其中一个人去mysql查询数据（waitgroup.Add(1)）,其他人等待（waitgroup.wait()）,待查询完成后，其他人共享
           （2） 二级节点：二级节点查询一级节点的mysql，用户挂载在二级节点查询数据。 1w用户 10个二级节点 1个一级节点，对于mysql来说只被查询了10次，每个二级节点被查询1000次，这样查询都被归并在二级节点，配合单飞解决缓存穿透的问题。


## 第八周总结  
分布式事务&分布式缓存  
1.memcache 基本使用和内存实现以及redis的应用  
2.多级缓存：底部的缓存过期时间一定要大于上层；缓存的清除一定是先删除底层再删除上层缓存  
3.分布式事务：内部系统之间调用，使用消息服务。不同系统之间靠不停的回调（支付宝与游戏道具）
4.如何保证分布式事务消息的可靠存储（消息丢了，钱就丢了，那就gg了）  
### Transactional outbox：
   支付宝本地创建一个sql事务，扣钱，同时更新消息表。插入一条加钱消息。
### Polling publisher
   定时轮询服务，定时轮询消息表，努力送达加钱消息到余额宝。
   缺点：轮询频次高，sql压力大；轮询频次低，延迟高。
### Transaction log tailing  
   canal服务订阅msg表的binlog，支付宝插入一条扣钱消息后，canal拿到扣钱消息发送到Kafka，然后余额宝去消费Kafka的加钱消息  ，去加钱。
### 幂等性  
   消息重复投递/消费，使用全局唯一id + 去重表（消息应用状态表），余额宝加钱后，插入数据到消息应用状态表（或者直接更新状态）、  
###  2PC （两阶段提交消息队列）  
![image](https://github.com/ProsperousLi/jike-go/blob/main/docs/%E5%9B%BE%E7%89%871.png)

生产者失败的情况以及处理方式：  
   1. 发送prepare消息失败，直接返回错误  
   2. 数据库事务提交失败，等待消息队列回调询问，prepare消息是否丢掉  
   3. 数据库提交成功，确认成功消息发送失败。等待消息队列回调询问，prepare消息是否确认成功。  
消费者：
   人工介入  
### TCC （Try , Confirm, Cancel）预处理、确认、撤销

## 第八周作业

1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。  
![image](https://github.com/ProsperousLi/jike-go/blob/main/week8/redis-10.png)
![image](https://github.com/ProsperousLi/jike-go/blob/main/week8/redis-100.png)
![image](https://github.com/ProsperousLi/jike-go/blob/main/week8/redis-1k.png)
![image](https://github.com/ProsperousLi/jike-go/blob/main/week8/redis-20.png)
![image](https://github.com/ProsperousLi/jike-go/blob/main/week8/redis-200.png)
![image](https://github.com/ProsperousLi/jike-go/blob/main/week8/redis-2k.png)
![image](https://github.com/ProsperousLi/jike-go/blob/main/week8/redis-50.png)

2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。  

# reference  
引用的阅读和总结  

[答疑文档以及reference](https://shimo.im/docs/JcrTccXkKjJvJdjJ/read)  

1. [Google API设计指南](https://www.bookstack.cn/read/API-design-guide/API-design-guide-README.md)
2. [为什么像王者荣耀这样的游戏Server不愿意使用微服务？](https://blog.csdn.net/github_shequ/article/details/109302632)
3. [美团发号器：Leaf：美团分布式ID生成服务开源](https://tech.meituan.com/2019/03/07/open-source-project-leaf.html)
4. 框架的使用：社区要活跃；一定要规范。  
5. [go mod 之痛](https://xargin.com/go-mod-hurt-gophers/)
6. [DDD 实践手册(番外篇: 事件风暴-概念)](https://zhuanlan.zhihu.com/p/110979132)
7. [Golang中的nil，没有人比我更懂nil！](https://zhuanlan.zhihu.com/p/151140497)  
