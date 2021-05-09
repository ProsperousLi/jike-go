# jike-go
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
4. 《谷歌测试之道》


## Week04 作业题目：

1. 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

思路：TODO 使用 gin框架搭建简单的web项目
