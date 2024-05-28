## Go语言高并发与微服务实战
近年来云原生技术发展迅猛，帮助开发者在云上快速和频繁地构建、发布和部署应用，以提高开发效率和快速定位故障。

微服务作为开展云原生技术落地的核心，它将复杂的单体应用按照业务划分并进行有效地拆分，每个微服务都可以进行独立部署和开发，大大提升了应用开发效率。Go语言作为新生代的编译型编程语言，具备语法简单、高并发性能良好和编译速度快等特点，是微服务架构落地实践的绝妙利器。

本书围绕云原生、微服务和Go语言，介绍如何通过Go语言进行微服务架构开发与实践。旨在帮助读者深入理解微服务架构并使用Go语言快速加入到微服务开发中。

### 环境搭建

#### golang 安装
本书基于的 Go 版本为 1.12.4，使用 go mod 管理依赖。具体使用可参见本书基础部分。

#### setup

```sh
go mod tidy
```
### Q&A
如有问题，可在 [链接](https://github.com/longjoy/micro-go-book) 提出 issue。

### 个人修改

    Go版本基于 1.18.1

### 本书的微服务架构

    基于consul的服务注册与发现
    consul 是一个服务网关的作用
    微服务网关的作用就是保护、增强和控制外部请求对于API服务的访问(RESTful是一种API规范)

    本书如果不以微服务的结构来做的话，可以从单例服务的角度，从go-kit出发，学习服务的分层处理
    
### go-kit

    go-kit是一个分布式的开发工具集，属于微服务体系
    从 github.com/go-kit/kit/transport/http/server.go 入手
        type Server struct {
            e            endpoint.Endpoint
            dec          DecodeRequestFunc
            enc          EncodeResponseFunc
            before       []RequestFunc
            after        []ServerResponseFunc
            errorEncoder ErrorEncoder
            finalizer    []ServerFinalizerFunc
            errorHandler transport.ErrorHandler
        }
    
        和函数
        func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request)
        
        分析go-kit的层级间的调用

    go-kit与MVC一样，也是三层架构，分别是：
    transport：处于微服务的最上层，
            定义路由，定义解码和编码器代码，链接endpoint层
            规定了从接收请求到响应的全流程
    endpoint：从解码器接收输入 链接servive层，调用方法，将结果输出到编码器
            每个makeXXXEndpoint方法返回一个endpoint，接受请求，调用service处理请求，返回结果
    service：业务逻辑
            方法入参为解码器输出，出参为编码器输入
    model: 数据模型
        定义了请求和响应的格式