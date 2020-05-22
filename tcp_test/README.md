#### socket简易聊天室编程练习
1. c/s模式,使用tcp进行通信
2. 服务端启动
go run service/service.go
3. 客服端启动
go run Client/client.go

4. 命令行消息发送
// 加入操作
{"mtype":1,"s_type":2,"m_user":1,"m_data":"","receiver":0}
{"mtype":1,"s_type":2,"m_user":2,"m_data":"","receiver":0}
{"mtype":1,"s_type":2,"m_user":3,"m_data":"","receiver":0}

// 单播操作
{"mtype":2,"s_type":1,"m_user":1,"m_data":"xiaobaijun","receiver":2}

// 多播操作
{"mtype":2,"s_type":2,"m_user":2,"m_data":"xiaozhong","receiver":0}

### 未完成
1. 离开聊天室,服务端断开 客户端检测
2. 使用数据DB或者Cache存储数据
3. 多服务器转发