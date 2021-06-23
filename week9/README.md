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
  #### fix length :  
    即每次发送固定缓冲区大小.客户端和服务器约定每次发送请求的大小.例如客户端发送1024个字节，服务器接受1024个字节。  
    这样可以解决粘包的问题，但是如果发送的数据小于1024个字节，就会导致数据内存冗余和浪费；且如果发送请求大于1024字节，会出现半包的问题。  
  #### delimiter based : 
    基于定界符来判断是不是一个请求（例如结尾'\n'). 客户端发送过来的数据，每次以\n结束，服务器每接受到一个 \n 就以此作为一个请求。  
  #### length field based frame decoder :  
    在TCP协议头里面写入每次发送请求的长度。 客户端在协议头里面带入数据长度，服务器在接收到请求后，根据协议头里面的数据长度来决定接受多少数据。
    但这种方式的成本比delimiter based 方式高，因为协议头里面的要存储的字节比加\n更大。  
# 实现一个从 socket connection 中解码出 goim 协议的解码器。  
