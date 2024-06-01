# GoWebPractice

## 📑 简介

**go-web-practice 是一个学习项目。通过使用 Golang 的原生库实现一个简单的 Web 应用。**

本项目是跟随 B 站 UP 主 [软件工艺师](https://space.bilibili.com/361469957) 的视频 《[Go Web 编程快速入门【Golang/Go语言】(完结)](https://www.bilibili.com/video/BV1Xv411k7Xn)》来进行学习与编写的。

在视频的基础上，做出了一些修改：

* 修整了项目文件结构。
* JSON 文件完成配置。
* 自定义 logger。
* 支持 Docker 以及 Docker-Compose 部署运行。

本项目已经过测试与部署，Linux 主机部署指南在下文中。

* 项目 Gitee 链接: https://gitee.com/Fxtack/go-web-practice (国内快一些)
* 项目 Github 链接: https://github.com/Fxtack/go-web-practice

项目结构:
```text
go-web-practice
.
├── LICENSE
├── README.md
├── build
│   └── package
│       ├── dockerfile.build_run
│       └── dockerfile.run
├── cmd
│   └── main.go
├── configs
│   └── config.json
├── deployments
│   └── docker-compose.yaml
├── go.mod
├── internal
│   ├── config
│   │   ├── config.go
│   │   └── config_test.go
│   ├── controller
│   │   ├── controller.go
│   │   ├── index.go
│   │   ├── look.go
│   │   └── welcome.go
│   ├── log
│   │   └── log.go
│   ├── middleware
│   │   ├── cross.go
│   │   ├── log.go
│   │   └── timeout.go
│   └── templates
│       └── template.go
└── web
    ├── img
    │   ├── favicon.ico
    │   ├── golang-down.png
    │   ├── golang-right.png
    │   └── golang.png
    ├── plugins
    │   ├── bootstrap
    │   │   ├── css
    │   │   │   └── bootstrap.min.css
    │   │   └── js
    │   │       └── bootstrap.min.js
    │   └── jquery
    │       └── js
    │           └── jquery.min.js
    └── templates
        ├── index.tmpl
        ├── look.tmpl
        ├── test.tmpl
        └── welcome.tmpl
```

> 注意
> 
> 以下几点是项目目前的问题 ~~或者说是拿来做大学作业的时候要注意的东西~~。
> 
> * 项目使用 Goland 开发的，其他 IDE 的迁移情况未知。
> * 只有 main.go 里面写了一点点注释，其他文件的注释以后更新。
> * 当前项目中的路由还有界面非常少，需要可以自己添加。
> * 没有数据库。
> * 使用 docker 以及 docker-compose 运行前需要自行对项目进行编译或将可执行文件放到项目目录下。

## 🔧 部署指南

### 编译

首先你得确保你安装好了 Golang 的环境。如果你修改了项目，那么你就需要重新编译，得到可执行文件，才能部署运行。**

Golang 的交叉编译提供了很大的便利，这使得我们可以在 Windows 系统下编译出可以在 Linux 环境下运行的可执行文件。

你可以使用以下 .bat 脚本在 windows 完成 windows 以及 linux 环境下的可执行文件编译。

```bat
# 请以项目目录为工作目录执行该脚本

# windows 可执行文件编译
go build -o go-web-practice.exe cmd\main.go

# linxu 可执行文件编译
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -o go-web-practice .\cmd\main.go

pause
```

如果你是 MacOS 或 Linux 开发，可以使用 `make` 构建。

```shell
make
```

### 配置

本项目使用 JSON 文件来进行配置。

查看 `config/conifg.json` 配置语义如下表，根据需求修改。

|                配置字段 | 语义                               |
|--------------------:|:---------------------------------|
|              static | 静态文件的路径                          |
|            template | go 模板的路径                         |
|             address | 服务启动的地址 (默认为空，可以实现公网部署)          |
|                port | 服务启动的端口                          |
| handleTimeoutSecond | request 处理超时时间（该配置在项目中未使用，作保留）   |
|               trace | trace logger 输出日志的路径（项目中未使用，作保留） |
|                info | info logger 输出日志的路径              |
|             warning | warning logger 输出日志的路径           |
|               error | error logger 输出日志的路径             |

### 运行

运行必须保证以下目录结构（以 Linux 运行为例）：

```bash
.
├── configs
│   └── config.json
├── go-web-practice
└── web
    └── ...
```

直接运行可执行文件即可看到以下提示表示运行成功：

![run](https://user-images.githubusercontent.com/59989422/173054601-10ec271d-5341-49d3-9b28-30e5d18642ba.png)

可以用 `curl localhost:8000/welcome` 测试一下，将会返回 welcome 界面的 html。

### 访问

本机访问（未修改端口配置）：http://localhost:8000。

**如果你是部署在服务器上，想要通过公网访问，记得设置防火墙开放相应的端口。如果是云服务器还要打开安全组。**

## 🐋 使用 Docker 编译/运行

master 分支已更新 dockerfile，并使用 docker 对 go-web-practice 的编译、运行进行了测试。

以下是 go-web-practice 的 Docker 编译/部署运行指南。

### 安装 Docker

Docker 是什么、怎么装、怎么用，这里不做详细的讲解介绍。

详细可参见 Docker 官方文档：[https://docs.docker.com/](https://docs.docker.com/)

### 方法一：使用 Docker 进行编译与运行

在项目目录下按照以下顺序执行命令，即可在镜像构建时完成编译，并在容器启动时运行服务。

#### 制作镜像

```bash
$ docker build -t go-web-practice -f ./build/package/dockerfile.build_run .
```

#### 运行容器

```bash
docker run -d -p 8000:8000 go-web-practice
```

### 方法二：本地编译后使用 Docker 运行

在本地编译生成可执行文件之后再构建运行镜像，这能保证镜像体积较小的情况下也能正常运行。

上文中的使用 Docker 编译并运行构建的镜像大小接近 1GB，而本节中构建的可运行镜像大概只有 100MB 左右。

> 注意：上文中构建过名为 go-web-practice 镜像后，再执行本节的指令前要先删除老镜像。

#### 准备镜像

先保证在项目目录下已有编译好的可执行文件。

在项目的 `/build/package` 目录下有准备好的 dockerfile 用于生成镜像文件，使用以下指令构建 docker 镜像。

```bash
$ docker build -t go-web-practice -f ./build/package/dockerfile.run .
```

#### 运行容器

后台运行容器输入以下指令：

```bash
$ docker run -d -p 8000:8000 go-web-practice
```

### 使用 `docker-compose` 部署运行

在构建了镜像后，也通过使用 docker-compose 指令直接部署运行：

```bash
$ docker-compose -f deployments/docker-compose.yaml up
```

> 注意，如果修改了端口配置则需修改 docker 运行指令以及 docker-compose.yaml 配置文件中的端口映射字段。

-------

若有问题可以邮箱联系我：fxtack@qq.com
