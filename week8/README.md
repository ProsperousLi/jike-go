
## 第八周作业

1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。  
![image](https://github.com/ProsperousLi/jike-go/blob/main/week8/redis-10.png)
![image](https://github.com/ProsperousLi/jike-go/blob/main/week8/redis-20.png)
![image](https://github.com/ProsperousLi/jike-go/blob/main/week8/redis-50.png)
![image](https://github.com/ProsperousLi/jike-go/blob/main/week8/redis-100.png)
![image](https://github.com/ProsperousLi/jike-go/blob/main/week8/redis-200.png)
![image](https://github.com/ProsperousLi/jike-go/blob/main/week8/redis-1k.png)
![image](https://github.com/ProsperousLi/jike-go/blob/main/week8/redis-2k.png)
![image](https://github.com/ProsperousLi/jike-go/blob/main/week8/redis-5k.png)

2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。  
