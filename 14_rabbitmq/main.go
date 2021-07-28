package _4_rabbitmq

//docker run

//docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management

// -d 后台运行容器；
//--name 指定容器名；
//-p 指定服务运行的端口（5672：应用访问端口；15672：控制台Web端口号）；
//-v 映射目录或文件；
//--hostname 主机名（RabbitMQ的一个重要注意事项是它根据所谓的 “节点名称” 存储数据，默认为主机名）；
//-e 指定环境变量；（RABBITMQDEFAULTVHOST：默认虚拟机名；RABBITMQDEFAULTUSER：默认的用户名；RABBITMQDEFAULTPASS：默认用户名的密码）

//阿里云开放 15672端口
//IP:15672访问
