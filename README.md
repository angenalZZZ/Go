# Go
Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。

 > [应用&库&工具](https://github.com/avelino/awesome-go)、[中文文档](https://studygolang.com/pkgdoc)

 * 通常用于服务器编程，网络编程，分布式系统，内存数据库，云平台等。
 * 1.部署简单 2.静态编译语言(又像动态解释语言) 3.自动回收机制GC 4.语言层面支持高并发 5.丰富的第三方库,并且开源.
~~~shell
# 1.部署简单：编译成机器码，复制后，别人就能直接用[Go环境也不用装] > 可通过<linux> ldd 查看可执行文件hello依赖的库文件
$ ldd hello # 不依赖库，不像其它语言C|C++|Java|.Net|Swift..依赖系统环境库才能运行
~~~

#### ① [搭建开发环境](https://juejin.im/book/5b0778756fb9a07aa632301e/section/5b0d466bf265da08ee7edd20)
    安装版本> go version
    环境配置> go env

> Windows - src: %GOPATH%\src - 配置 set: cd %USERPROFILE% (C:\Users\Administrator)

    https://studygolang.com/dl/golang/go1.12.windows-amd64.msi
    GOROOT=D:\Program\Go\
    GOPATH=C:\Users\Administrator\go
    PATH=D:\Program\Go\bin;%GOPATH%\bin;%PATH%

> Linux - src: $GOPATH/src - 配置 export: cd $HOME (/root 或 /home)
    
    wget https://studygolang.com/dl/golang/go1.12.linux-amd64.tar.gz
    GO_INSTALL_DIR=/usr/local # 默认安装目录: 可更改 (选项 tar -C)
    tar -xvzf go1.12.linux-amd64.tar.gz -C $GO_INSTALL_DIR
    GOROOT=/usr/local/go
    GOPATH=/home/go
    PATH=/usr/local/go/bin:$GOPATH/bin:$PATH

> 安装依赖包 [集成工具](https://godoc.org/golang.org/x/tools)
~~~bash
# 安装-全局依赖-镜像包 (解决网络问题)
git clone https://github.com/golang/arch.git %GOPATH%/src/golang.org/x/arch     # 数据结构
git clone https://github.com/golang/build.git %GOPATH%/src/golang.org/x/build   # 构建、发布
git clone https://github.com/golang/crypto.git %GOPATH%/src/golang.org/x/crypto # 加密、安全
git clone https://github.com/golang/debug.git %GOPATH%/src/golang.org/x/debug   # 调试、跟踪
git clone https://github.com/golang/image.git %GOPATH%/src/golang.org/x/image   # 图片库
git clone https://github.com/golang/lint.git %GOPATH%/src/golang.org/x/lint     # 语法检查
git clone https://github.com/golang/mobile.git %GOPATH%/src/golang.org/x/mobile # 移动端
git clone https://github.com/golang/net.git %GOPATH%/src/golang.org/x/net       # 网络库
git clone https://github.com/golang/oauth2.git %GOPATH%/src/golang.org/x/oauth2 # OAuth 2.0 认证授权
git clone https://github.com/golang/perf.git %GOPATH%/src/golang.org/x/perf     # 性能测量、存储和分析
git clone https://github.com/golang/sync.git %GOPATH%/src/golang.org/x/sync     # 并发访问-同步锁
git clone https://github.com/golang/sys.git %GOPATH%/src/golang.org/x/sys       # 系统底层
git clone https://github.com/golang/text.git %GOPATH%/src/golang.org/x/text     # 文本处理
git clone https://github.com/golang/time.git %GOPATH%/src/golang.org/x/time     # 时间处理
git clone https://github.com/golang/tools.git %GOPATH%/src/golang.org/x/tools   # 工具包
git clone https://github.com/golang/tour.git %GOPATH%/src/golang.org/x/tour     # 其他

# 开发工具-VSCode语言支持
go get -u -v github.com/nsf/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/zmb3/gogetdoc
go get -u -v github.com/golang/lint/golint
go get -u -v github.com/lukehoban/go-outline
go get -u -v github.com/lukehoban/go-find-references
go get -u -v github.com/sqs/goreturns
go get -u -v github.com/tpng/gopkgs
go get -u -v github.com/golang/tools/cmd/goimports
go get -u -v github.com/golang/tools/cmd/gorename
go get -u -v github.com/golang/tools/cmd/guru
go get -u -v github.com/newhook/go-symbols
go get -u -v github.com/fatih/gomodifytags
go get -u -v github.com/cweill/gotests/...

# 管理项目依赖包
go get -u github.com/kardianos/govendor
  > govendor init             # 项目依赖vendor目录
  > govendor add +e           # 添加本地$GOPATH包[go get]
  > govendor fetch            # 获取远程vendor.json包[govendor get]
go get -u github.com/golang/dep/cmd/dep
  > dep init                  # 初始化项目
  > dep ensure -add [package] # 添加一个包
  > dep ensure                # 安装依赖包(速度慢)
go get -u github.com/sparrc/gdm
go get -u github.com/Masterminds/glide # <Mac or Linux> curl https://glide.sh/get | sh
  > glide --version ; glide help # https://glide.sh
  > glide create                 # Start a new workspace
  > glide get github.com/foo/bar#^1.2.3 # Get a package, add to glide.yaml; https://glide.sh/docs/glide.yaml
  > glide install -v          # Install packages and dependencies
  > glide update              # Update the latest dependency tree
  > glide list                # See installed packages
  > glide tree                # See imported packages
  > go build

# 学习playground*
go get github.com/golang/playground
go get github.com/golang/example/hello
go get github.com/shen100/golang123         # shen100
go get github.com/golang/leveldb            # 内存数据库
git clone https://github.com/adonovan/gopl.io.git %GOPATH%/src/github.com/adonovan/gopl.io # Example programs

~~~

> Docker 编译器 Golang + custom build tools

~~~shell
docker pull jetbrainsinfra/golang:1.11.5  # build tools: Glide, gdm, go-test-teamcity
docker run --name golang -itd -p 8080:8080 -v %GOPATH%\src\apiserver:/go/src/apiserver 
  jetbrainsinfra/golang:1.11.5 bash
~~~

#### ② [功能、框架、基础库、应用、工具](https://github.com/avelino/awesome-go)

 * [QT跨平台应用框架](https://github.com/therecipe/qt)
 * [其他应用](https://github.com/avelino/awesome-go)

#### ③ [构建企业级的 RESTful API 服务](https://juejin.im/book/5b0778756fb9a07aa632301e)
~~~
# 开发
cd %GOPATH%/src                                                                 # 项目框架 Gin 
git clone https://github.com/lexkong/apiserver_demos apiserver                  # 项目源码-复制^demo至-工作目录
git clone https://github.com/lexkong/vendor                                     # 项目依赖-govendor
go get github.com/StackExchange/wmi                                             # 项目依赖-缺失的包
# 编译 
cd %GOPATH%/src/apiserver && gofmt -w . && go tool vet . && go build -v .
# 运行
%GOPATH%/src/apiserver/apiserver.exe
~~~

#### ④ [中文标准库文档](https://studygolang.com/pkgdoc)

#### ⑤ 阅读相关文章

 * 高性能
    * [高并发架构解决方案](https://studygolang.com/articles/15479)


----

