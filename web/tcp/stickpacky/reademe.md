# tcp 粘包和解决方案

1. 定长协议
2. 分隔符协议
    客户端连续发送：
        Hello World\nHow are you\nI am fine\n
        服务端收到粘包数据：
        Hello World\nHow are you\nI am fine\n
        服务端根据 \n 拆分为三条消息。
3. 自定义协议头，包含数据长度

4. 自定义协议设置协议头和尾
