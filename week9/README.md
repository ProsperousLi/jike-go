# 总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用  

-什么是粘包和半包？  
  粘包问题是指当发送两条消息时，比如发送了 ABC 和 DEF，但另一端接收到的却是 ABCD，像这种一次性读取了两条数据的情况就叫做粘包（正常情况应该是一条一条读取的）。  
  ![image](https://github.com/ProsperousLi/jike-go/blob/main/week9/image.png)  
  半包问题是指，当发送的消息是 ABC 时，另一端却接收到的是 AB 和 C 两条信息，像这种情况就叫做半包  
  ![image](https://github.com/ProsperousLi/jike-go/blob/main/week9/半包.png)  
-为什么会有粘包和半包？  
  这是因为 TCP 是面向连接的传输协议，TCP 传输的数据是以流的形式，而流数据是没有明确的开始结尾边界，所以 TCP 也没办法判断哪一段流属于一个消息。  

粘包的主要原因：  
  发送方每次写入数据 < 套接字（Socket）缓冲区大小；  
  接收方读取套接字（Socket）缓冲区数据不够及时。   
半包的主要原因：  
  发送方每次写入数据 > 套接字（Socket）缓冲区大小；  
  发送的数据大于协议的 MTU (Maximum Transmission Unit，最大传输单元)，因此必须拆包。  
##  粘包的解包方式总结  
  fix length :  
  delimiter based :  
  length field based frame decoder :  
# 实现一个从 socket connection 中解码出 goim 协议的解码器。  
