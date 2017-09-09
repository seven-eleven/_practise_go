# mass_connect  

## 功能
- 实现CLIENT与SERVER的建链。CLIENT节点数和SERVER节点可配置。
- 实现CLIENT与SERVER之间的KEY-VALUE更新。KEY采用客户端连接的IP:PORT，VALUE为客户端访问计数。
- 配置功能。可配置SERVER节点IP列表、端口列表、单HOST客户端节点数、日志开关、CLIENT/SERVER可执行程序名（方便跨平台）、KEY-VALUE更新操作时间间隔等
- 日志功能。打开日志开关，可跟踪程序执行过程、定位异常。默认关闭。
- 启动程序一键运行系统。startup.go编译得到的可执行程序即为启动程序。
- 提供操作维护客户端。omclient.go编译得到的可执行程序即为操作维护客户端程序。支持获取帮助、打印SERVER数据、关闭服务、查询配置信息等。

## 设计思路
采用可配置的分布式架构，方便扩展系统规模  
每个CLIENT与所有SERVER节点建立TCP连接，并循环发起数据更新操作  
每个SERVER接受到所有CLIENT的数据更新操作，更新本地并回复确认消息  
对于N * M的架构（N个CLIENT，M个SERVER），存在N * M个链路，但是对于每个CLIENT存在M+1个线程，每个SERVER存在N+1个线程  

内核数据存储采用map[string]string结构，key为每个客户端的"IP:PORT"，value为访问次数  
为保证线程安全，使用了读写锁  

为了方便系统监控、问题定位，提供日志、操作维护界面  

## 运行
- 获取代码
- 构建程序
以windows为例：  
go build -v -o server.exe server.go  
go build -v -o client.exe client.go  
go build -v -o omclient.exe omclient.go  
go build -v -o startup.exe startup.go  
- 填写配置文件  
"server_ip_list": "127.0.0.1",  ## SERVER IP地址列表，每个HOST填写一个。样例只使用了一个HOST  
"server_port_list": "38677;38678;38679",  ## SERVER使用的端口列表  
"server_app_name": "server.exe",  ## SERVER可执行程序名  
"client_num_per_host": 100,  ## 单HOST启动的CLIENT数  
"client_app_name": "client.exe",  ## CLIENT可执行程序名  
"log_switch": false,  ## 日志开关  
"update_intervals": 10  ## 更新key-value操作时间间隔，以ms为单位  
- 启动  
1）执行startup.exe （或对应的其它平台构建得到的可执行程序）。执行完后系统就已经启动，可从操作系统界面看到多个SERVER和CLIENT进程。  
2）执行omclient.exe (或对应的其它平台构建得到的可执行程序）。获取系统内部信息。  
