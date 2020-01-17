# Go
Go是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。 罗伯特·格瑞史莫，罗勃·派克及肯·汤普逊于2007年9月开始设计Go，稍后Ian Lance Taylor、Russ Cox加入项目。Go是基于Inferno操作系统所开发的。

`[✨程序设计]`    [帮助文档](https://github.com/angenalZZZ/doc)
`[✨内存管理]`    [脸书的jemalloc](https://github.com/jemalloc/jemalloc)、[谷歌的tcmalloc](https://goog-perftools.sourceforge.net)

 > [官方中文文档](https://studygolang.com/pkgdoc)、[官方推荐项目](https://github.com/golang/go/wiki/Projects)、[Go资料收集](https://github.com/ty4z2008/Qix/blob/master/golang.md)、[*Go语言圣经*](https://docs.hacknode.org/gopl-zh)、[*✨Go高级编程*](https://chai2010.cn/advanced-go-programming-book)
 
 > [*搭建开发环境*](#-搭建开发环境) ；[*管理.构建*](#管理构建) + [*测试*](#测试) + [*性能优化*](#性能优化) <br> [语法速查表](#语法速查表) ；[*✨推荐功能.框架.基础库.应用.工具*](#-功能框架基础库应用工具) ；[*开源web框架*](#-开源的-web-框架) ；[*查询第三方库*](https://libs.garden/go) <br> [*云平台.公众平台.在线支付*](#云平台公众平台在线支付) ；[*google开源*](#Google开源) ；[*GUI-HTML/JS/CSS*](#gui---htmljscss) - [*WebAssembly*](#webassembly) - [*WebRTC*](#webrtc) <br> [awesome-go大全](https://github.com/avelino/awesome-go) ；[*github开源排名*](https://github.com/topics/go) 

 * 常用于服务器编程，网络编程，分布式系统，内存数据库，云平台... [freecodecamp.org](https://guide.freecodecamp.org/go)
 * 集成工具 [JetBrains/GoLand](https://www.7down.com/search.php?word=JetBrains+GoLand&s=3944206720423274504&nsid=0)（[^搭建开发环境$](#-搭建开发环境)）、[liteide](http://liteide.org/cn/)
 * 语言优势：开发效率、抽象能力、内存管理、稳定性和功能之间的矛盾、应对需求变更、人才和资源等。

 > `开发者`
    [Gopher-China技术交流大会](https://gopherchina.org)、[搜索50万Go语言项目](https://gowalker.org)、[API+SDK'排名'服务平台](https://sdk.cn)

~~~shell
# 1.部署简单：编译成机器码(像C一样,不被反编译)复制给别人后，就能直接运行(参考跨平台编译)
#   通过<linux>命令 ldd 查看可执行文件依赖的环境(库文件)
$   ldd hello # Go不像其它语言C|C++|Java|.Net|...依赖系统环境库才能运行(已编译成机器码)
# 2.静态编译语言(又像动态解释语言)，您不用再去关心变量是存在堆上还是栈的内存问题(编译器与运行时会帮您做到)
# 3.自动回收机制GC(除了CGO中`C语言那部分`管理的之外的内存；另外，Go指针不能被`值类型`长期保持--与其它语言不同)
# 4.语言层面支持高并发(goroutine是go适合高并发场景的重要原因)高性能goroutine池 go get github.com/panjf2000/ants/v2
# 5.丰富的第三方库,并且开源... github.com/avelino/awesome-go
~~~

 > 语法`关键字`

    break      default       func     interface   select
    case       defer         go       map         struct
    chan       else          goto     package     switch
    const      fallthrough   if       range       type
    continue   for           import   return      var

 > 内建的`常量`、`类型`、`函数`

    常量: true false iota nil
          const ptrSize = 4 << (^uintptr(0) >> 63) // unsafe.Sizeof(uintptr(0)) but an ideal const
    
    类型: bool int int8 int16 int32 int64  uint uint8 uint16 uint32 uint64
          float32 float64  complex64 complex128
          array chan func interface map ptr slice string struct
          uintptr  unsafe.Pointer                  // unsafe.Pointer is a safe version of uintptr used
          byte rune error  invalid    reflect.Type,Value,StringHeader,SliceHeader,SelectCase...底层结构
    
    函数: make len cap append delete new copy close    complex real imag    panic recover

 > 通道`chan`

 ![](http://tensor-programming.com/wp-content/uploads/2016/11/go-channel.jpg)

    读写: ch := make(chan<- int) #只读; ch := make(<-chan int) #只写; make(chan<- chan int) #只读chanOfchan;
    同步: ch := make(chan struct{}) // unbuffered channel, goroutine blocks for IO #空结构,无内存开销,更高效;
    异步: ch := make(chan int, 100) // buffered channel with capacity 100 (缓冲) 可避免阻塞,推荐select用法;
    管道: ch1, ch2 := make(chan int), make(chan int) ; ch1 <- 1 ; ch2 <- 2 * <-ch1; result := <-ch2 ;
    选择: select: 常规模式(for轮循次数=chan实际数量); 反射模式(reflect.Select([]reflect.SelectCase)..);
    时间: ch := time.After(300 * time.Second) #过期chan; ch := time.Tick(1 * time.Second) #轮循chan;
    更多: github.com/eapache/channels #Distribute分发1In*Out,Multiplex多路复用*In1Out,Pipe管道1In1Out,Batching*批量..

 > 指针`pointer`

 ![](http://tensor-programming.com/wp-content/uploads/2016/11/Pointer.png)
 
 > 协程(`轻量级线程`/`非抢占式线程`/`并发体`)`goroutine` + 上下文(同步通信)`context`<br>
	　　Go语言是基于`CSP消息并发模型`的集大成者，与Erlang不同的是Go语言的`Goroutine`之间是共享内存的。`Goroutine`和系统线程不是等价的。尽管两者的区别实际上只是一个量的区别，但正是这个量变引发了Go语言并发编程质的飞跃。<br>
	　　首先，每个系统级线程都会有一个固定大小的栈（一般默认可能是2MB），这个栈主要用来保存函数递归调用时参数和局部变量。固定了栈的大小导致了两个问题：一是对于很多只需要很小的栈空间的线程来说是一个巨大的浪费，二是对于少数需要巨大栈空间的线程来说又面临栈溢出的风险。针对这两个问题的解决方案是：要么降低固定的栈大小，提升空间的利用率；要么增大栈的大小以允许更深的函数递归调用，但这两者是没法同时兼得的。相反，一个`Goroutine`会以一个很小的栈启动（可能是2KB或4KB），当遇到深度递归导致当前栈空间不足时，`Goroutine`会根据需要动态地伸缩栈的大小（主流实现中栈的最大值可达到1GB）。因为启动的代价很小，所以我们可以轻易地启动成千上万个`Goroutine`。<br>
	　　Go运行时还包含了自己的调度器，这个调度器使用了一些技术手段，可以在n个操作系统线程上多工调度m个`Goroutine`。调度器的工作和内核的调度是相似的，但是这个调度器只关注单独的Go程序中的`Goroutine`。`Goroutine`采用的是`半抢占式`的协作调度，只有在当前`Goroutine`发生阻塞时(如`chan`线程同步通信)才会导致调度（或者在特定的代码行产生调度`runtime.Gosched()`切换点）；同时发生在用户态(运行时)，调度器会根据具体函数只保存必要的寄存器，切换的代价要比系统线程低得多(系统的`抢占式`线程调度机制,导致`运行时`产生的线程是用户被动使用,不能手动调度的)。运行时有一个变量`runtime.GOMAXPROCS(一般设置为CPU核心数)`用于控制当前运行正常非阻塞`Goroutine`的系统线程数目，`runtime.NumGoroutine()`返回运行时总的`goroutine`使用数量。

 ![](http://tensor-programming.com/wp-content/uploads/2016/11/gopher_pipe.png)

 > 包、模块(命名空间)`package`

    << 依赖`import` + 接口`interface` + 类型`type` + 函数`func` + 常量`Constants` + 变量`Variables` >>
    
----

#### ① [搭建开发环境](https://juejin.im/book/5b0778756fb9a07aa632301e/section/5b0d466bf265da08ee7edd20)
    环境配置> go env
    安装版本> go version
    帮助文档> godoc -http=:6060 -index  ↑↑查看本地文档; 在线文档→→ golang.org/doc
             :go^1.13需安装godoc: set GO111MODULE=on ; go get golang.org/x/tools/cmd/godoc
    开发工具> goland激活→→ idea.lanyus.com

> Windows - src: %GOPATH%\src - 配置 set: cd %USERPROFILE% (C:\Users\Administrator)

    https://studygolang.com/dl/golang/go1.13.5.windows-amd64.msi
    set GOPATH=C:\Users\Administrator\go
    set GOROOT=D:\Program\Go
    set GOTOOLS=%GOROOT%/pkg/tool  (可选项: GOOS=windows, GOARCH=amd64, CGO_ENABLED=0)
    set GO111MODULE=on             (可选项: 建议 GO111MODULE=auto )
    set GOPROXY=https://goproxy.io (可选项: 建议 网络代理)
    set PATH=%GOROOT%\bin;%GOPATH%\bin;%PATH%
    # go build 环境：CGO_ENABLED=1;GO_ENV=development # CGO_ENABLED=0禁用后兼容性更好;GO_ENV(dev>test>prod);[-ldflags "-H windowsgui"]可以让exe运行时不弹出cmd窗口
    set CGO_ENABLED=0 set GOOS=linux set GOARCH=amd64 go build -ldflags "-s -w" -o api_linux_amd64 ./api
    # go build 参数：-i -ldflags "-s -w" # -ldflags 自定义编译标记:"-s -w"去掉编译符号+调试信息(杜绝gdb调试)+缩小file
    # gofmt -w .    # 格式化代码         #build参数 -ldflags "-s -w -X importpath.varname=value" #-X编译时修改代码
    # GoLand环境设置：GOROOT, GOPATH ( √ Use GOPATH √ Index entire GOPATH?  √ Enable Go Modules[vgo go版本^1.11])

> Linux - src: $GOPATH/src - 配置 export: cd $HOME (/root 或 /home)

    wget https://studygolang.com/dl/golang/go1.13.5.linux-amd64.tar.gz
    GO_INSTALL_DIR=/usr/local # 默认安装目录: 可更改解压到的目录 (选项 tar -C)
    tar -zxf go1.13.5.linux-amd64.tar.gz -C $GO_INSTALL_DIR
    export GOPATH=~/go
    export GOROOT=/usr/local/go
    export GOTOOLS=$GOROOT/pkg/tool   (可选项: GOOS=linux, GOARCH=amd64, CGO_ENABLED=0)
    export GO111MODULE=on             (可选项: 建议 GO111MODULE=auto )
    export GOPROXY=https://goproxy.io (可选项: 建议 网络代理)
    export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
    sudo vi /etc/profile   # 添加以上export变量到profile文件结尾,然后启用配置文件 source /etc/profile
    # <跨平台编译> 查看支持的操作系统和对应平台: https://github.com/fatedier/frp/blob/master/README_zh.md
    go tool dist list
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o api_linux_amd64 ./api
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./api_windows_amd64.exe ./api
    # <按条件编译> 1.通过代码注释的形式(在包声明之前&空行隔开); 2.通过文件名后缀(比如:*_linux_amd64.go)
    go build -tags [linux|darwin|386|amd64]
    // +build darwin linux freebsd windows android js
    // +build 386 amd64 arm arm64 ppc64 wasm
    [空行]

> 编译器命令
~~~bash
go command [arguments]    // go 命令 [参数]
go build                  // 编译包和依赖包
go clean                  // 移除对象和缓存文件
go doc                    // 显示包的文档
go env                    // 打印go的环境变量信息
go bug                    // 报告bug
go fix                    // 更新包使用新的api
go fmt                    // 格式规范化代码
go generate               // 通过处理资源生成go文件
go get                    // 下载并安装包及其依赖
go install                // 编译和安装包及其依赖
go list                   // 列出所有包
go run                    // 编译和运行go程序
go test                   // 测试
go tool                   // 运行给定的go工具
go version                // 显示go当前版本
go vet                    // 发现代码中可能的错误
~~~

> 安装依赖包
~~~bash
# 代理设置 (解决网络问题) HTTP_PROXY, HTTPS_PROXY, NO_PROXY - defines HTTP proxy environment variables
> set http_proxy=http://127.0.0.1:5005     (临时有效) [设置环境变量linux $ export -> vim /etc/profile]
> set https_proxy=http://127.0.0.1:5005    (临时有效)
> set ftp_proxy=http://127.0.0.1:5005      (临时有效)
# 代理推荐
> $env:GOPROXY=https://goproxy.io         ## Windows PowerShell
$ export GOPROXY=https://goproxy.io       ## Linux Profile ~ GO111MODULE=on
$ echo "GOPROXY=https://goproxy.io" >> ~/.profile && source ~/.profile
# 1. https://goproxy.io                   ## 首选
# 2. https://mirrors.aliyun.com/goproxy/
# 3. https://athens.azurefd.net
# 4. https://gocenter.io
# 5. https://goproxy.cn
# go get github.com/gpmgo/gopm ##内地代理 https://gopm.io > gopm -h ; gopm bin [-d %GOPATH%/bin][包名]
# 内网代理推荐 Athens: https://docs.gomods.io/zh/
$ export ATHENS_STORAGE=~/athens-storage ##Docker 参考 https://docs.gomods.io/walkthrough/
$ mkdir -p $ATHENS_STORAGE
$ docker run -d -v $ATHENS_STORAGE:/var/lib/athens \
   -e ATHENS_DISK_STORAGE_ROOT=/var/lib/athens \
   -e ATHENS_STORAGE_TYPE=disk \
   --name athens-proxy \
   --restart always \
   -p 3000:3000 \
   gomods/athens:latest
$ export GO111MODULE=on                  ##Linux Bash use after run Docker
$ export GOPROXY=http://127.0.0.1:3000
- environment:                           ##Docker file
    GO111MODULE: on
    GOPROXY: http://127.0.0.1:3000

# 下载模块
go get -d         # 下载模块源码,不安装
go get -u         # 更新模块源码
go get -v         # 打印日志
go get -insecure  # 解决安全下载问题,允许用http(非https)

# 安装-全局依赖-镜像包 (解决网络问题)
git clone --depth=1 https://github.com/golang/arch.git %GOPATH%/src/golang.org/x/arch     # 数据结构
git clone --depth=1 https://github.com/golang/build.git %GOPATH%/src/golang.org/x/build   # 构建、发布
git clone --depth=1 https://github.com/golang/crypto.git %GOPATH%/src/golang.org/x/crypto # 加密、安全
git clone --depth=1 https://github.com/golang/debug.git %GOPATH%/src/golang.org/x/debug   # 调试、跟踪
git clone --depth=1 https://github.com/golang/exp.git %GOPATH%/src/golang.org/x/exp       # 实验和弃用的包
git clone --depth=1 https://github.com/golang/image.git %GOPATH%/src/golang.org/x/image   # 图片库
git clone --depth=1 https://github.com/golang/lint.git %GOPATH%/src/golang.org/x/lint     # 语法检查
git clone --depth=1 https://github.com/golang/mobile.git %GOPATH%/src/golang.org/x/mobile # 移动端
git clone --depth=1 https://github.com/golang/net.git %GOPATH%/src/golang.org/x/net       # 网络库
git clone --depth=1 https://github.com/golang/oauth2.git %GOPATH%/src/golang.org/x/oauth2 # OAuth 2.0 认证授权
git clone --depth=1 https://github.com/golang/perf.git %GOPATH%/src/golang.org/x/perf     # 性能测量、存储和分析
git clone --depth=1 https://github.com/golang/sync.git %GOPATH%/src/golang.org/x/sync     # 并发访问-同步锁
git clone --depth=1 https://github.com/golang/sys.git %GOPATH%/src/golang.org/x/sys       # 系统底层
git clone --depth=1 https://github.com/golang/text.git %GOPATH%/src/golang.org/x/text     # 文本处理
git clone --depth=1 https://github.com/golang/time.git %GOPATH%/src/golang.org/x/time     # 时间处理
git clone --depth=1 https://github.com/golang/tour.git %GOPATH%/src/golang.org/x/tour     # 开发文档
git clone --depth=1 https://github.com/googleapis/google-cloud-go.git %GOPATH%/src/cloud.google.com/go # 谷歌云

# 开发工具 VSCode✨  github.com/Microsoft/vscode-go
# 分享Go工具链✨  pan.baidu.com/s/13tfSyd2OeSXU4lNaUfMHpA 提取码: 41jq
git clone --depth=1 https://github.com/golang/tools.git %GOPATH%/src/golang.org/x/tools
go get github.com/ramya-rao-a/go-outline
go get github.com/acroca/go-symbols
go get github.com/mdempsky/gocode
go get github.com/rogpeppe/godef
go get golang.org/x/tools/cmd/godoc
go get github.com/zmb3/gogetdoc
go get golang.org/x/lint/golint
go get github.com/fatih/gomodifytags
go get golang.org/x/tools/cmd/gorename
go get sourcegraph.com/sqs/goreturns
go get golang.org/x/tools/cmd/goimports
go get github.com/cweill/gotests/...
go get golang.org/x/tools/cmd/guru
go get github.com/josharian/impl
go get github.com/haya14busa/goplay/cmd/goplay
go get github.com/uudashr/gopkgs/cmd/gopkgs
go get github.com/davidrjenni/reftools/cmd/fillstruct
go get github.com/alecthomas/gometalinter  && gometalinter --install
go get github.com/go-delve/delve/cmd/dlv  #debug: github.com/go-delve/delve/blob/master/Documentation/installation/README.md
~~~

#### 管理|构建
~~~bash
# ------------------------------------------------------------------------------------
#  谷歌开源的构建和测试工具，类似于Make、Maven、Gradle支持跨平台|语言|代码库|工具链 ✨ docs.bazel.build
#   /构建规则: Bazel rules for building protocol buffers +/- gRPC ✨ github.com/stackb/rules_proto
# ------------------------------------------------------------------------------------

# 构建
  > go fmt ./... && gofmt -s -w . && go vet ./... && go get ./... && go test ./... && \
  > golint ./... && gocyclo -avg -over 15 . && errcheck ./...
  > GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w"
  > gowatch                       # 热编译工具，提升开发效率 go get github.com/silenceper/gowatch

# 管理模块依赖( go版本^1.11.* 推荐) & 设置GoLand环境 √ Enable Go Modules(vgo)
# 集成 vgo 项目模块管理工具 (可用环境变量 GO111MODULE 开启或关闭模块支持:off,on,auto) #默认auto未开启
git clone --depth=1 https://github.com/golang/vgo.git %GOPATH%/src/golang.org/x/vgo ; go install #安装vgo
  #! github.com/golang/go/wiki/Modules research.swich.com/vgo blog.jetbrains.com/go
  > go help mod <command>       # 帮助 | 功能概述 go help modules
  > set GO111MODULE=on          # 开始前(临时开启) | <linux> $ export GO111MODULE=on && env
  > mkdir.\example.com\app      # 新建项目 | <linux> $ mkdir -p example.com/app
  > cd example.com/app          # 进入项目目录，此目录不再需要 in %GOPATH%
  #----------------------------------------------------------------------
  > go mod init [$MODULE_NAME]  # 1.默认生成 go.mod 文件，$MODULE_NAME默认为github.com/$GITHUB_USER_NAME/$PROJECT_NAME
  > go mod init example.com/app # 1.指定生成 go.mod 文件，依赖golang.org/...可能要翻墙，go.mod中用replace替换成github镜像
  > go get github.com/gin-gonic/gin # 安装项目依赖... 生成 go.sum 文件，锁定依赖的版本。
  > code .                      # 2.开始编码... 在 go module 下，不需要vendor目录(go~1.10.*)进行精确的版本管理
  #----------------------------------------------------------------------
  > go mod tidy || go get ./... # 2.下载依赖%GOPATH%/pkg/mod/... 文件夹(tidy保持依赖项目同步,舍弃无用的依赖)
  > go build                    # 3.构建使用%GOPATH%/pkg/mod/... 文件夹
  > go clean -r -cache .        # 4.清除构建&缓存文件
  #----------------------------------------------------------------------
  > go list -m all              # 2.查看当前版本
  > go list -m -u all           # 2.查看当前的依赖和模块版本更新 -json 支持json输出
  > depth [package-name]        # 查看某一个库的依赖 go get github.com/KyleBanks/depth/cmd/depth
  > go-callvis                  # 代码调用链图工具 go get github.com/TrueFurby/go-callvis
  > go mod graph                # 4.输出依赖关系,打印模块依赖图
  > go mod verify               # 5.验证依赖是否正确
  > go get -u || -u=patch       # 5.升级到最新依赖版本 || 升级到最新的修订版本 (也可指定版本)
  > go mod edit -fmt            # 5.格式化 go.mod 文件
  > go mod edit -require=path@ver # 2.添加或修改依赖版本
  > go mod download             # 2.下载依赖到%GOPATH%/pkg/mod/cache'共享缓存'
  #----------------------------------------------------------------------     (处理网络问题)
  > go mod edit -replace=google.golang.org/grpc=github.com/grpc/grpc-go@latest # 2.编辑镜像
  > go mod tidy
  > go mod vendor               # 3.拷贝依赖到./vendor/... 文件夹
  > go build -mod=vendor        # 3.构建时使用./vendor/... 文件夹
  > go build -mod=readonly      # 3.防止隐式修改go.mod
  > rm go.sum && go mod vendor # 删掉 go.sum 并重建, 解决 checksum mismatch?
  #----------------------------------------------------------------------
  > go mod init github.com/golang/app # 6.从旧项目迁移 GO111MODULE (读取vendor/vendor.json,gopkg.toml到go.mod)
  > go mod download             # 6.下载依赖到%GOPATH%/pkg/mod/... 缓存文件夹
  #----------------------------------------------------------------------
  > go mod download | go build              # $GOPATH/pkg/mod [缓存]
  > go mod vendor   | go build -mod=vendor  # ./vendor [方便复制打包]

# 管理模块依赖( go版本~1.10.* 推荐)
go get -u github.com/golang/dep/cmd/dep # 推荐使用 *12k
  > dep init                  # 初始化项目
  > dep ensure -add [package] # 添加一个包
  > dep ensure                # 安装依赖包(速度慢)
go get -u github.com/Masterminds/glide  # 推荐使用 *7k <Mac or Linux> curl https://glide.sh/get | sh
  > glide --version ; glide help # https://glide.sh
  > glide create                 # Start a new workspace
  > glide get github.com/foo/bar#^1.2.3 # Get a package, add to glide.yaml; https://glide.sh/docs/glide.yaml
  > glide install -v          # Install packages and dependencies
  > glide update              # Update the latest dependency tree
  > glide list                # See installed packages
  > glide tree                # See imported packages
  > go build
go get -u github.com/kardianos/govendor # 推荐使用 *4k
  > govendor init             # 项目依赖vendor目录
  > govendor add +e           # 添加本地$GOPATH包(未加入vendor目录时)[go get]
  > govendor update|remove    # 从$GOPATH更新包|移除包依赖vendor目录
  > govendor fetch|sync       # 获取远程vendor.json包[govendor get]
~~~

#### 测试
~~~bash
# -------------------------------------------------------------------------------
# 测试命令 >>
# -------------------------------------------------------------------------------
  > go help test                                   # 测试帮助文档
  > go test ./...                                  # 测试遍历递归目录...所有的*_test.go 并且代码中包名一致
  > go test -v -count=1 [package-name]             # 测试指定的包(默认目录path=. 默认次数count=1 -v打印详情)
  > go test -run=^$  -parallel=20 ./path           # 单元测试(t *testing.T) -run=查找TestXxx -parallel=并行数
  > go test -test.list=^Benchmark ./path           # 只打印匹配的测试函数
  # t.Log   t.Logf   # 正常信息 -> 类似HTTP状态码^200
  # t.Error t.Errorf # 测试失败信息，测试程序`报告`的错误信息 -> 类似HTTP状态码^400
  # t.Fatal t.Fatalf # 致命错误信息，测试程序`退出`的异常信息 -> 类似HTTP状态码^500
  # t.Fail     # 当前测试函数被标记为失败
  # t.Failed   # 查看当前测试函数失败标记
  # t.FailNow  # 标记失败，并终止当前测试函数的执行，需要注意的是，我们只能在运行测试函数的
               # Goroutine 中调用 t.FailNow 方法，而不能在我们在测试代码创建出的 Goroutine 中调用它
  # t.Skip     # 调用 t.Skip 方法相当于先后对 t.Log 和 t.SkipNow 方法进行调用，而调用t.Skipf方法则相当于先后对
               # t.Logf 和 t.SkipNow 方法进行调用。方法 t.Skipped 的结果值会告知我们当前的测试是否已被忽略
  # t.Parallel # 标记为可并行测试 (当test参数 -parallel 时)
  > go test -timeout=10s github.com/mpvl/errdare   # 远程测试超时10秒
  > go test -cover ./path                          # 检测代码覆盖率(testing使用到的代码行比例)
  > go test -bench=.* -cpu=2 -benchmem -benchtime=1s #`压测`基准测试(b *testing.B)在函数循环体指定b.N
# -------------------------------------------------------------------------------
# 测试工具 >>
# -------------------------------------------------------------------------------
  > go test -bench=. -memprofile=mem.prof ./path  # 生成mem性能测试两个文件path.test.exe,mem.prof
  > go test -bench=. -cpuprofile=cpu.prof ./path  # 生成cpu性能测试两个文件path.test.exe,cpu.prof
  > go tool pprof path.test.exe cpu.prof    # 分析函数调用(pprof)指令+> help,top,png生成图片;提前安装Graphviz
   $ go tool pprof path.test cpu.prof > web # 分析函数调用(svg)图+> yum install graphviz.x86_64  www.graphviz.org
   $ apt search graphviz ; sudo apt-get install graphviz/eoan ; sudo apt-get install graphviz-doc/eoan 
   $ go tool pprof -raw -seconds 30 http://localhost/debug/pprof/profile # 查看CPU性能火焰图 go-torch -h #out.svg
  > go test -coverprofile=c.out             # 生成代码覆盖率分析文件(标记出未测试到的代码行与代码比例)
  > go tool cover -func=c.out               # 分析代码覆盖率;检查哪些`函数`没测试或者没测试完全
  > go tool cover -html=c.out               # 分析代码覆盖率;查看网页格式html文件
# -------------------------------------------------------------------------------
# 代码质量 >>
# -------------------------------------------------------------------------------
  > go help vet                                    # 执行代码静态检查(语法检查)
  > go tool vet help                               # 查看工具vet支持哪些检查?
  > go list ./...|grep -v vendor|xargs go vet -v   # 检查时,排除目录vendor?
  > go tool vet -shadow main.go                    # 检查变量覆盖? 请提前安装 'shadow' analyzer tool
  > go get github.com/securego/gosec/cmd/gosec/... # 代码质量与安全分析工具> gosec -fmt=json -out=1.json ./... 
  > go errcheck|golint|unused|varcheck             # 其它的代码检测工具 go-linters
   # 推荐[1.结合github平台进行自动化的审查 https://golangci.com 2.本地src审查工具golangci-lint & gocritic]
  > revive -h                                    #1.1代码质量检测工具 go get github.com/mgechev/revive
  > golangci-lint run | golangci-lint run ./... #2.1代码运行与审查工具 github.com/golangci/golangci-lint
  > go get -v github.com/go-lintpack/lintpack/... && go get -v github.com/go-critic/go-critic/... #2.2审查工具
     && lintpack build -o gocritic -linter.version='v0.3.4' -linter.name='gocritic' github.com/go-critic/go-critic/checkers
  > gocritic check-project %gopath%/src/github.com/graphql-go/graphql/  #扫描GraphQL代码质量 gocritic check -help
  > go get github.com/fortytw2/leaktest         # 检测goroutine内存泄漏问题:leaktest.Check()CheckTimeout()CheckContext()
  
  # 测试HTTP负载，内置HTTP服务与请求速率，包含命令行实用工具和库 > go get github.com/tsenart/vegeta
  > vegeta [global flags] <command> [command flags]
  
  # 捕获HTTP请求,跟踪流量  github.com/buger/goreplay/wiki
  > gor --input-raw :80 --output-http="http://localhost:81" # 跟踪HTTP流量(:80), HTTP服务查阅结果(HTTP:81)
  > gor --input-raw :80 --output-stdout # 跟踪HTTP流量(:80)[--output-http-track-response],文件web结果 gor file-server :81
  > gor --input-raw :80 --output-file=requests.gor && gor --input-file requests.gor --output-http="http://localhost:8080"

  # 集成go-test,全自动web-UI,回归测试套件,测试复盖率,代码生成器,桌面通知`goconvey`
  > go get github.com/smartystreets/goconvey   # 优雅的单元测试 *5k (推荐) | Convey("test1",t,func(){So(v1,ShouldEqual,v2)})
  > go get github.com/stretchr/testify         # 通用的接口测试 *10k (强力推荐) | assert,http,mock,require,suite
  > go get github.com/appleboy/gofight/...     # API测试框架 beego,Gin.依赖上面的框架 github.com/stretchr/testify
  > go get github.com/astaxie/bat              # 接口调试增强curl *2k | testing, debugging, interacting servers
  > go get github.com/asciimoo/wuzz            # 用于http请求 | 交互式命令行工具 | 增强curl
  # Web基准测试命令 github.com/wg/wrk *20k      # +辅助生成图表 sudo apt-get -y install gnuplot --fix-missing
  $ wrk -t144 -c600 -d30s --latency <url>      # -t线程数 -c连接数 -d压测时间s --latency响应+n%延迟统计ms --timeout超时
  $ wrk2 -t144 -c412 -d30s -R14400 --latency <url> # -R每秒请求的速率[次/秒] --latency[-L]响应延迟统计 --timeout[-T]超时
  $ wrk2 -t144 -c412 -d30s -R14400 --u_latency <url> # --u_latency[-U]打印未校正的延迟统计;生成报告"未校正延迟直方图"
  # Web性能测试命令 github.com/codesenberg/bombardier *1.5k
  $ bombardier -n 10000 -c 600 -d 10s -m GET -t 3s --fasthttp -l <url> # -n请求数 -c连接数 -d压测时间s -l即--latencies
  > go get github.com/tsliwowicz/go-wrk        # Web性能测试工具 *0.4k > go-wrk -help
  $ go-wrk -M GET -c 600 -d 10 -no-ka <url>    #  -c并发连接数 -d压测时间10s -T超时3s -no-ka=no-keep-alive
  > go get github.com/goadapp/goad             # Web性能测试工具 *1.5k > goad -h
  > go get github.com/uber/go-torch            # Web性能测试与CPU火焰图生成工具 *3.5k > go-torch -h
  > go get github.com/smallnest/go-web-framework-benchmark # Web性能测试工具 > gowebbenchmark -help

# 测试代码书写`Testing Coding` (go语言推荐`表格数据驱动`代码写法;比传统写法:可读性更强+可维护性更好)
  > go get github.com/k0kubun/pp     # 彩色漂亮的打印输出
  > go get github.com/davecgh/go-spew/spew # 为数据结构实现了一个深度漂亮的打印输出，以帮助调试。
  > go get github.com/google/go-cmp  # 一个强大和安全的`Equal`替代方案(reflect.DeepEqual仅用于比较两个值在语义上是否相等)
  > go get github.com/go-stack/stack # 包堆栈实现了用于捕获、操作和格式化调用堆栈的实用程序。它提供了比包运行时更简单的API。

# 自动化工具`CI`
  # 构建+发布到Github | goreleaser.com | github.com/goreleaser/goreleaser
  $ wget https://github.com/goreleaser/goreleaser/releases/download/v0.111.0/goreleaser_Linux_x86_64.tar.gz
  $ tar zxf goreleaser_Linux_x86_64.tar.gz && sudo cp goreleaser /usr/local/bin/
  $ rm -f goreleaser && rm -f *.md
  $ goreleaser help release
~~~

#### 性能优化
 * 减少算法的时间复杂度
 * 根据业务逻辑，设计优化的数据结构 (一般需投入的精力:数据结构>程序算法)
 * 尽量减少磁盘IO次数 (现而今需关注)
 * 尽量复用资源 (分布式应用需关注)
 * 同步锁sync.Map,RWMutex,Mutex (锁的粒度尽量小;尽量使用无锁的方式)
 * 内存分配 (数据结构初始化时，尽量指定合适的容量 make 避免多次内存分配)
 * 固定的 go routine 数量 + 固定的 channel 数量, 提升单机性能
 * 原子操作配合互斥锁可以实现非常高效的单件模式(参考标准库sync.Once)。互斥锁的代价比普通整数的原子读写高很多，在性能敏感的地方可以增加一个数字型的标志位，通过原子检测标志位状态降低互斥锁的使用次数来提高性能。
 * [High performance go workshop](https://talks.godoc.org/github.com/davecheney/high-performance-go-workshop/high-performance-go-workshop.slide)
 * [An Introduction to go tool trace](https://about.sourcegraph.com/go/an-introduction-to-go-tool-trace-rhys-hiltner/)
 * [Writing and Optimizing Go code](https://github.com/dgryski/go-perfbook/blob/master/performance.md)
 * [Go tooling essentials](https://rakyll.org/go-tool-flags/) 、 [go-torch](https://www.cnblogs.com/li-peng/p/9391543.html)
 * [Profiling、Tracing、Debugging、Runtime statistics and events](https://cyningsun.github.io/07-21-2019/go-diagnostics-cn.html)
~~~
# ------------------------------------------------------------------------------------
# 程序代码排查`go tool pprof`的优化：
# ------------------------------------------------------------------------------------
#1. import _ "runtime/pprof" # 添加性能分析采集(可用于控制台程序、测试程序等)
go test -bench=.* -benchtime 10s -cpuprofile=cpu.prof -memprofile=mem.prof #测试与性能分析*testing.B
go tool pprof [binary] [profile] # 调用分析工具pprof(调用上面生成的分析结果文件;再调用svg可生成直观图)
go tool pprof -alloc_objects -inuse_objects [binary] [profile] # 生成对象数量、引用对象数量等分析结果
#go tool pprof -http=:8080 [binary] [profile] # GC对象扫描,函数占据大量CPU(如runtime.scanobject等问题)
# top ; top20 ; top -cum # 按累积取样计数来排序(top命令-默认只包含本地取样计数最大的前十个函数)
#2. import _ "net/http/pprof" # 添加HTTP性能分析采集(也是基于runtime/pprof的封装;用于暴露HTTP端口进行调试)
# 通过访问/debug/pprof查看cpu和内存状况 (通常:我们用wrk来访问，让服务处于高速运行状态，取样的结果会更准确)
# go tool pprof http://127.0.0.1:8080/debug/pprof/profile # 分析CPU采样信息(默认频率100Hz,即每10毫秒取样一次)
# git clone https://github.com/brendangregg/FlameGraph.git 后运行FlameGraph下的(拷贝flamegraph.pl到/usr/local/bin)
# go-torch -u http://127.0.0.1:8080 --seconds 60 -f <cpu.svg> # 火焰图分析CPU: 生成cpu.svg文件
# ... http://127.0.0.1:8080/debug/pprof/profile,heap,goroutine,mutex,block,threadcreate # 查看性能采集数据与分析结果
#2.1 import "expvar"; var v=expvar.NewInt("visits"); v.Publish(name string, v expvar.Var) # 全局注册表func
# 通过访问/debug/vars查看expvar包中注册的所有公共变量(两个指标:os.Args,runtime.Memstats)与自定义变量#globalVars
go get -u github.com/google/pprof # 更新pprof用于性能分析和采集数据可视化(分析cpu时间片,内存,死锁,火焰图等)
#  pprof -http=:8088 cpu.prof     # 后启动可视化界面 http://127.0.0.1:8088/ui 查看火焰图
go get github.com/uber/go-torch # Web性能测试与CPU火焰图生成工具 > go-torch demo.exe <cpu.prof|mem.prof>
go get github.com/prashantv/go_profiling_talk #剖析如何用pprof和go-torch识别性能瓶颈?视频youtu.be/N3PWzBeLX2M
# ------------------------------------------------------------------------------------
# 内存管理`GC`的优化：
# ------------------------------------------------------------------------------------
 # 内存分配性能测试
testing.AllocsPerRun()                          # 检测平均占用内存(分配对象数量)
 # 对象数量过多时(引用传递过多时)，导致`GC`三色算法耗费较多CPU（可利用耗费少量的内存，优化耗费的CPU）
map[string]NewStruct -> map[[32]byte]NewStruct  # key使用值类型,避免对map遍历
map[int]*NewStruct   -> map[int]NewStruct       # val使用值类型,避免对map遍历
someSlice []float64  -> someSlice [32]float64   # 可利用值类型Array代替对象类型Slice
someSlice []int      -> map[int]bool
# ------------------------------------------------------------------------------------
# 扩充`IO`容量（横向|纵向）>>分布式应用：
# ------------------------------------------------------------------------------------
 # 分片Sharding > 如何集群? 把数据划分成若干部分,1个部分映射1个Shard(内存中分配完成);把Shard分配到服务器节点上;
   节点node+副本replica
 # 策略 > 如何分片? <空间索引>把数据按空间范围划分成若干个最小单元Cell;按规则算法把部分单元Cells放入1个Shard分片;
   Cell队列中的数据可查找所在Shard/Cell;数据清理Clean
~~~

----

#### ② [功能、框架、基础库、应用、工具](https://github.com/avelino/awesome-go)

~~~
go get -d github.com/golang/example/hello  # hello
go get -d github.com/golang/playground     # playground  #本地教程# > tour  #在线教程# tour.go-zh.org
go get -d github.com/shen100/golang123     # 适合初学者
go get -d github.com/go-training/training  # 适合初学者培训
go get -d github.com/insionng/zenpress     # 适合学习 cms system
go get -u github.com/ponzu-cms/ponzu/...   # 用户友好、可扩展的CMS和管理后台(SSL+Push+BoltDB)
go get -d github.com/polaris1119/The-Golang-Standard-Library-by-Example # 标准库例子
go get -d github.com/muesli/cache2go       # 缓存库，代码量少，适合学习，锁、goroutines等
go get -d github.com/phachon/gis           # 图片上传，下载，存储，裁剪等
go get -d github.com/phachon/mm-wiki       # 轻量级的企业知识分享、文档管理、团队协同
go get -d github.com/getlantern/lantern    # 网络底层的东西，适合深入学习  *42k
go get -d github.com/Unknwon/the-way-to-go_ZH_CN # 中文入门教程 *2.7k  关注: Gogs, INI file, 音视频学习教程
go get -d github.com/Yesterday17/bili-archive-frontend # 前端实现*bili-bili
go get -d github.com/detectiveHLH/go-backend-starter   # 后端实现*gin, gorm
go get -d github.com/etcd-io/etcd/etcdserver           # 深度学习*grpc
git clone --depth=1 https://github.com/adonovan/gopl.io.git %GOPATH%/src/github.com/adonovan/gopl.io
-------------------------------------------------------------------------------------------------

go get github.com/teris-io/shortid         # super short, fully unique 9~10 chars(推荐) *0.5k -URL friendly
go get github.com/rs/xid                   # uuid shortuuid Snowflake MongoID xid(推荐) *1.5k -xid 20 chars
go get github.com/bwmarrin/snowflake       # 分布式id生成器:Twitter-snowflake算法:1毫秒2^12=4096条:1秒409万
go get github.com/sony/sonyflake           # 分布式id生成器:Twitter-snowflake扩展(推荐)  www.sony.net
go get github.com/google/uuid              # 基于RFC4122和DCE1.1身份验证和安全服务，生成uuid、检查uuid等
go get github.com/satori/go.uuid           # uuid generator, 支持5种版本(基于RFC4122)
go get github.com/kjk/betterguid           # guid generator, 20 chars
go get github.com/juju/utils               # Utility functions: arch,cache,cert,debug,deque,exec,file,hash,kv,os,parallel,proxy,ssh,tar,zip...
go get github.com/henrylee2cn/goutil       # Common and useful utils
go get github.com/shirou/gopsutil          # Utils(CPU, Memory, Disks, etc)
go get github.com/appleboy/com             # Random、Array、File、Convert
go get github.com/huandu/xstrings          # String functions to their friends in other languages
go get github.com/bradfitz/iter            # Range [0,n) | for i:=range iter.N(1e9) `内存分配`testing.AllocsPerRun()
go get gopkg.in/pipe.v2                    # io.Pipeline | github.com/go-pipe/pipe
go get golang.org/x/crypto/acme/autocert   # tls usage | github.com/go-ego/autotls
go get golang.org/x/sync/errgroup          # sync.group扩展 g:=new(errgroup.Group);g.Go(func(){});e:=g.Wait() godoc.org/golang.org/x/sync/errgroup
go get github.com/rafaeldias/async         # 超级好用+异步高并发处理(推荐)
go get github.com/Jeffail/tunny            # 工作线程池+Api并行处理请求限制goroutines(推荐) *1.5k
go get gopkg.in/go-playground/pool.v3      # 工作线程池+高效对象池(推荐) github.com/go-playground/pool
go get gopkg.in/go-playground/validator.v9 # 数据结构扩展验证功能(强力推荐) github.com/go-playground/validator
go get github.com/asaskevich/govalidator   # 字符串、数字、切片和自定义结构的验证器(推荐) *4k
go get github.com/bytedance/go-tagexpr     # 数据校验，参数Binding:{ B string `tagexpr:"len($)>1 && regexp('^\\w*$')"` }
go get github.com/xeipuuv/gojsonschema     # 数据校验，json schema 自定义错误校验       *1k
go get github.com/chrislusf/glow/...       # 大数据计算+分布式集群，像Hadoop-MapReduce,Spark,Flink,Storm  *2.5k
go get github.com/chrislusf/gleam/...      # 快速高并发可扩展分布式计算(推荐)MapReduce,dag,pipe,k8s,Read>HDFS&Kafka
go get github.com/reactivex/rxgo           # 响应式编程库rxgo
go get github.com/google/go-intervals/...  # 时间范围内执行操作
go get github.com/Knetic/govaluate         # 表达式引擎:Eval表达式:Functions:Accessors
go get github.com/cheekybits/genny         # 泛型语言支持 golang.org/doc/faq#generics
go get github.com/google/btree             # 数据结构 B-Trees
go get github.com/google/trillian          # 数据结构 Merkle tree, Verifiable Data Structures *2k
go get github.com/emirpasic/gods           # 数据结构(强力推荐)*7.2k Containers,Sets,Lists,Stacks,Maps,Trees,Comps,Iters…
go get github.com/TheAlgorithms/Go         # 各种算法的实现 github.com/TheAlgorithms/Python   *31k
go get github.com/cenkalti/backoff         # 指数退避算法backoff.v4,用于降低程序执行速率,如:最大尝试重连数等
go get gonum.org/v1/gonum/...              # 各种算数运行(强力推荐)*3.2k矩阵,线性代数统计,概率分析和抽样,分区&集成&优化,网络分析等
go get github.com/skelterjohn/go.matrix    # 线性代数统计库(推荐)
go get github.com/OneOfOne/xxhash          # 超快的非对称加密哈希算法(推荐)> xxhgo ; xxhsum -h ;C语言github.com/Cyan4973/xxHash
go get github.com/spaolacci/murmur3        # 超快的哈希分布均匀的算法(推荐)> murmur32 123456 ; murmur64 123456
go get github.com/bkaradzic/go-lz4         # 无损压缩算法LZ4> lz4go ; lz4 -h ;C语言 github.com/Cyan4973/lz4
go get github.com/mholt/archiver/cmd/arc   # 压缩/解压文件(zip,tar,rar)> arc archive|unarchive|extract|ls|compress|decompress
go get github.com/hpcloud/tail/...         # 从不断更新的文件读取.惠普.开源(推荐) log rotation tool: www.hpe.com
go get github.com/DataDog/zstd             # 实时数据压缩方法(强力推荐) DataDog: Facebook/Zstd: Fast-Stream-API
# 编码/解码:性能比拼: https://github.com/alecthomas/go_serialization_benchmarks
go get github.com/vipally/binary           # binary编码/解码 data和[]byte的互转(encoding/gob,encoding/binary)
go get github.com/linkedin/goavro          # Avro编码/解码 avro.apache.org
go get github.com/tinylib/msgp             # MessagePack编码/解码(推荐) 考虑结合缓存库使用
go get github.com/vmihailenco/msgpack      # MessagePack编码/解码(像JSON但更快更小) msgpack.org
go get github.com/niubaoshu/gotiny         # 效率非常的高，是golang自带序列化库gob的3倍以上(减少使用reflect库)
go get github.com/google/jsonapi           # 转换对象，HTTP请求的输入输出                       *1k
go get github.com/google/go-querystring/query # 转换对象，URL参数                              *1k
go get github.com/json-iterator/go         # json编码/解码的性能优化，替换原生(encoding/json)   *5k
go get github.com/tidwall/gjson            # json路径+过滤+to[array,map..] gjson.Valid(json)&&gjson.Get(json,"name.last").Exists()
go get github.com/PuerkitoBio/goquery      # 解析HTML像jQuery那样操作DOM                       *7k
go get github.com/rs/zerolog/log           # 日志记录-性能最高           *2k
go get github.com/uber-go/zap              # 日志记录-Uber开源-扩展插件 *8.5k
go get github.com/sirupsen/logrus          # 日志跟踪-功能最多-扩展插件 *13.4k
go get github.com/pkg/errors               # 错误处理库pkg/errors        *5k
go get github.com/juju/errors              # 用简单的方法描述错误而不丢失原始错误信息(强力推荐)  *1.1k

go get github.com/abice/go-enum            # 代码生成枚举类型的功能
go get github.com/clipperhouse/gen         # 代码生成类似泛型的功能 *1k  clipperhouse.com/gen/overview
go get github.com/rjeczalik/interfaces/cmd/interfacer # 生成接口代码使用
go get github.com/ahmetb/go-linq           # 推荐使用.NET_LINQ功能*1.8k From(slice).Where(predicate).Select(selector).Union(data)
go get github.com/ungerik/pkgreflect       # 生成包反射时使用pkgreflect.go
go get github.com/alecthomas/participle    # 超简单的Lexer解析器Parser(推荐使用,Lexer性能高于`反射`) *1.5k
go get github.com/blynn/nex                # 好用的Lexer解析器工具，生成go代码&YACC/Bison&正则表达式: nex -r -s lc.nex
go get github.com/antlr/antlr4/runtime/Go/antlr # 语言识别工具，强大的Parser生成器，读取、处理、执行或翻译文本或二进制文件 | www.antlr.org
go get github.com/go-ego/gpy               # 汉语拼音转换工具
go get github.com/levigross/grequests      # HTTP client Requests(推荐)
go get gopkg.in/h2non/gentleman.v2         # HTTP client library
go get github.com/sethgrid/pester          # HTTP client calls with retries, backoff, and concurrency.
go get github.com/haxpax/gosms             # 发短信 SMS gateway *1.2k
go get -d github.com/upspin/upspin         # 构建安全统一和全局命名、共享文件和数据的框架：一个排序的全局名称系统 *5k | git clone https://upspin.googlesource.com/upspin %GOPATH%\src\upspin.io

# https://github.com/etcd-io               # 分布式可靠键值存储，适用于分布式系统中最关键的数据；提供分享配置和服务发现
# client: http://play.etcd.io              # 数据中心 etcd | 下载 github.com/etcd-io/etcd/releases
go get github.com/hashicorp/serf/cmd/serf  # 数据中心 serf | 基于Gossip的Membership,P2P对等网络\去中心 | www.serf.io
go get github.com/spf13/viper && go get github.com/spf13/pflag # 配置(JSON,TOML,YAML,HCL)热加载;远程配置;缓存;加密
go get github.com/xordataexchange/crypt/bin/crypt 加密存储 secret keyring: gpg(gpg4win)用于安全传输(类似rsa)
go get github.com/minio/minio-go           # 云存储|分布式存储SDK|网盘|OSS | www.min.io  docs.min.io/cn
go get -d github.com/minio/mc              # 云存储|配置客户端, 指南 | docs.min.io/cn/minio-client-quickstart-guide.html
go get -d github.com/minio/minio           # 云存储|配置服务端, 运行: hidec /w minio.exe server d:\docker\app\minio\data
go get github.com/perkeep/perkeep/cmd/...  # Camlistore 个人存储系统：一种存储、同步、共享、建模和备份内容的方式
go get -d github.com/rclone/rclone         # 云存储的Sync: 用于各种文件存储服务的同步   *15k
go get -d github.com/s3git/s3git           # 云存储的Git: 用于数据的分布式版本控制系统  *1k
go get github.com/allegro/bigcache         # 缓存库[GB级大数据高效缓存+超快的GC](推荐) *3k
go get github.com/eko/gocache              # 缓存管理(推荐)Memory[Bigcache,Ristretto]+Memcache+Redis+[Chained,Metric]..
go get github.com/coocood/freecache        # cache and high concurrent performance
go get github.com/patrickmn/go-cache       # in-memory key:value store/cache (similar to Memcached)适用于单台应用程序
go get github.com/VictoriaMetrics/fastcache
go get github.com/chrislusf/seaweedfs/weed # 一个用于小文件的简单且高度可扩展的分布式文件系统，可集成其他云服务，如AWS..
go get github.com/bigfile/bigfile/artisan  # 提供http-api,rpc,ftp客户端文件管理(推荐) 中文文档 learnku.com/docs/bigfile/1.0
go get github.com/fsnotify/fsnotify        # 文件系统监控 # go get golang.org/x/sys/...
go get github.com/rjeczalik/notify         # 文件系统事件通知库
# 数据狗 - 云监控 (Modern monitoring & analytics)  https://www.datadoghq.com
go get github.com/cloudflare/cfssl/cmd/... # SSL证书 usage play.etcd.io/install#TLS  *4k
go get github.com/tidwall/evio             # 超快的事件/网络IO{http,redis..}-server   *4k
go get github.com/muesli/beehive           # 灵活的事件/代理/自动化系统                *3k
go get github.com/asaskevich/EventBus      # 异步的事件总线Subscribe/Publish/Wait/Callback *1k
go get github.com/nuclio/nuclio-sdk-go     # 高性能事件微服务和数据处理平台(结合MQ,Kafka,DB) *3k docker run -p 8070:8070 -v /var/run/docker.sock:/var/run/docker.sock -v /tmp:/tmp quay.io/nuclio/dashboard:stable-amd64

go get github.com/go-redis/cache
go get github.com/go-redis/redis           # 内存数据库,类型安全的Redis-client *6k (推荐使用,性能高于redigo)
go get github.com/gomodule/redigo/redis    # 内存数据库,集成原生的Redis-cli *6k
go get github.com/sent-hil/bitesized        # Redis位图计数> 统计分析、实时计算
go get github.com/yannh/redis-dump-go       # Redis导出导入> redis-dump-go -h ; redis-cli --pipe < backup.resp;redis-dump
go get github.com/syndtr/goleveldb/leveldb # 内存数据库,谷歌leveldb推荐
go get github.com/seefan/gossdb/example    # 内存数据库,替代Redis的ssdb | ssdb.io/zh_cn

go get github.com/dgraph-io/badger/...     # 高性能 key/value 数据库,支持事务,(强力推荐)LSM+tree,ACID,Stream,KV+version,SSDs
go get github.com/dgraph-io/dgraph/dgraph  # 高性能,具有可扩展、分布式、低延迟和高吞吐量功能的分布式位图索引数据库 *10k
go get github.com/boltdb/bolt/...          # 高性能 key/value 数据库,支持事务,B+tree,ACID,分桶 *10k | 性能低于badger
go get github.com/tidwall/buntdb           # 内存数据库,BuntDB is a low-level, in-memory, key/value store, persists to disk
go get github.com/tidwall/buntdb-benchmark # 性能测试 > buntdb-benchmark -n 10000 -q # 单机时超越Redis，有索引和geospatial功能
go get github.com/allegro/bigcache         # 高可用千兆级数据的高效 key/value 缓存   *2k
go get github.com/cockroachdb/cockroach    # 云数据存储系统，支持地理位置、事务等 *20k | www.cockroachlabs.com/docs/stable
go get -d github.com/tidwall/tile38        # 具有空间索引和实时地理位置数据库,如PostGIS *7k docker run -p 9851:9851 tile38/tile38
go get -d github.com/pingcap/tidb          # TiDB 支持包括传统 RDBMS 和 NoSQL 的特性 *18k | pingcap.com/docs-cn
go get github.com/influxdata/influxdb1-client/v2 # 分布式、事件、实时的可扩展时序数据库InfluxDB *19k | github.com/influxdata/influxdb
go get github.com/influxdata/influxdb-client-go # 时序数据库InfluxDB2.x客户端 | v2.docs.influxdata.com/v2.0/get-started
go get github.com/pilosa/pilosa            # Pilosa分布式位图索引+实时计算+大数据+列式存储 *16k | kuanshijiao.com/2017/06/12/pilosa1
go get github.com/pilosa/go-pilosa         # Pilosa分布式位图索引-客户端 | www.pilosa.com/docs/latest/installation/#docker
go get github.com/pilosa/pdk               # Pilosa开发套件+用例示例
go get github.com/melihmucuk/geocache      # 适用于地理位置处理, 基于应用程序的内存缓存 *1k
go get github.com/bluele/gcache            # 支持LFU、LRU 和 ARC 的缓存数据库 *1k
go get github.com/bradfitz/gomemcache/memcache # memcache 客户端库
go get github.com/couchbase/go-couchbase   # Couchbase 客户端

go get github.com/astaxie/beego/orm        # 数据库orm    *20k support mysql,postgres,sqlite3...
go get github.com/jinzhu/gorm              # 数据库gorm   *12k | gorm.io/docs
git clone --depth=1 https://github.com/rana/ora.git %GOPATH%/src/gopkg.in/rana/ora.v4 && go get gopkg.in/rana/ora.v4
go get github.com/mattn/go-oci8            # Oracle env: instantclient & MinGW-w64-gcc & pkgconfig/oci8.pc
go get github.com/go-sql-driver/mysql      # Mysql client and driver     *8k   github.com/siddontang/go-mysql
go get github.com/lib/pq                   # Postgres client and driver  *5k   github.com/prest/prest
go get github.com/jackc/pgx                # Postgres client and toolkit *2k
go get github.com/go-pg/pg/v9              # Postgres client and ORM     *3k
go get github.com/sosedoff/pgweb           # Postgres client and WebUI   *6k
go get github.com/denisenkom/go-mssqldb    # MsSql client and driver     *1k
go get gopkg.in/mgo.v2                     # MongoDB 驱动:集群,并发,一致性,Auth,GridFS *2k github.com/go-mgo/mgo labix.org/mgo
go get github.com/globalsign/mgo           # MongoDB^4 client and driver *2k
go get github.com/mattn/go-sqlite3         # SQLite client and driver    *3k
go get github.com/jmoiron/sqlx             # 数据库sql library  *6k  (extensions go's standard database/sql)
  go get github.com/heetch/sqalx             # sqlx & sqalx 支持嵌套的事务
  go get github.com/twiglab/sqlt             # sqlx & sqlt 模板拼接sql和java的数据库访问工具MyBatis的sql配置
  go get github.com/albert-widi/sqlt         # sqlx & sqlt 支持数据库主从数据源，读写分离
go get github.com/go-xorm/xorm             # 数据库xorm   *5k  support mysql,postgres,tidb,sqlite3,mssql,oracle
  go get github.com/go-xorm/builder          # ^xorm SQL Builder 增强-拼接sql
  go get github.com/xormplus/xorm            # ^xorm增强版*$ 支持sql模板,动态sql,嵌套事务,类ibatis配置等
                                             # ^xorm增强版*文档 https://www.kancloud.cn/xormplus/xorm/167077
go get github.com/didi/gendry              # 滴滴开源 SQL Builder 增强-拼接sql、连接池管理、结构映射.
go get github.com/golang-migrate/migrate   # 数据库 schema 迁移工具 *3k
go get github.com/rubenv/sql-migrate/...   # 数据库 schema 迁移工具，允许使用 go-bindata 将迁移嵌入到应用程序中 *1k
git clone --depth=1 https://github.com/go-gormigrate/gormigrate.git %GOPATH%/src/gopkg.in/gormigrate.v1 && go get gopkg.in/gormigrate.v1 
go get github.com/gchaincl/dotsql          # 帮助你将 sql 文件保存至某个地方并轻松使用sql
go get github.com/xo/xo                    # 命令行工具 xo --help  [DbFirst]生成 models/*.xo.go # gorm migrate
   > cp %GOPATH%/src/github.com/xo/xo/templates/* ./templates
   > xo mysql://root:123456@127.0.0.1:3306/AppAuth?parseTime=true -o ./models [--template-path templates]
   > xo mssql://sa:123456@localhost:1433/AppAuth?parseTime=true -o ./models [--template-path templates]
go get github.com/go-xorm/cmd/xorm         # 命令行工具 xorm help  [DbFirst]生成 models/*.go
   > cp %GOPATH%/src/github.com/go-xorm/cmd/xorm/templates/goxorm/* ./templates
   > xorm reverse mysql root:123456@tcp(127.0.0.1:3306)/AppAuth?charset=utf8 ./templates ./models [^表名前缀]
   > xorm reverse mssql "server=localhost;user id=sa;password=HGJ766GR767FKJU0;database=AppAuth" %GOPATH%/src/github.com/go-xorm/cmd/xorm/templates/goxorm ./models [^表名前缀]
go get github.com/variadico/scaneo         # 命令行工具 scaneo -h  [DbFirst]生成 models/*.go

go get github.com/cayleygraph/go-client    # 图数据库 Client API  *13k
go get github.com/cayleygraph/cayley       # 图数据库(推荐) Driven & RESTful API & LevelDB Stores
go get github.com/siesta/neo4j             # Neo4j 客户端  github.com/jmcvetta/neoism
go get github.com/go-ego/riot              # Riot 搜索引擎|分布式索引|中文分词|中文转拼音 *5k
go get github.com/blevesearch/bleve        # Bleve 现代文本搜索引擎  *6k
go get github.com/olivere/elastic          # Elasticsearch 6.0客户端 *4k
go get github.com/Qihoo360/poseidon        # 360开源|百亿级日志分布式搜索引擎&Hadoop *1.5k
go get github.com/DarthSim/imgproxy        # Fast image server: docker pull darthsim/imgproxy
go get willnorris.com/go/imageproxy/...    # Caching image proxy server & docker & nginx

# Web开发推荐如下：Router|Api框架 + MVC框架
go get github.com/julienschmidt/httprouter # 高性能Router框架(强力推荐) *10k (很多Web框架都是基于它进行二次开发)
go get github.com/gin-gonic/gin            # 后端WebSvr框架 *33k: Gin(推荐)  Star数最高的Web框架
go get github.com/astaxie/beego            # 后端WebSvr框架 *22k: API、Web、MVC  高度解耦的框架  beego.me/docs
# 基础模块：cache,config,context,httplibs,logs,orm,session,toolbox,plugins.. 管理工具bee   github.com/beego/bee
go get github.com/kataras/iris             # 最快WebSvr框架 *15k 中文文档 github.com/kataras/iris/blob/master/README_ZH.md
# 入门程序：[iris+xorm]github.com/yz124/superstar [gorm+jwt]github.com/snowlyg/IrisApiProject [postgres+angular]github.com/iris-contrib/parrot
go get github.com/go-martini/martini       # 强大中间件和模块化设计的web框架 *11k martini.codegangsta.io
go get gopkg.in/macaron.v1                 # 高生产力的和模块化设计的web框架(推荐)martini高级扩展+依赖注入 go-macaron.com/docs
go get goa.design/goa/v3/cmd/goa           # 高生产力的和集成开发的web框架+微服务工具链 *3.6k (推荐)
go get github.com/gorilla/{mux,sessions,schema,csrf,handlers,websocket} # 后端Web框架与工具链mux *10k
go get github.com/gohugoio/hugo            # 超快的静态网站生成工具(强力推荐) *37k   gohugo.io
go get github.com/mholt/caddy/caddy        # 全栈Web服务平台 *21k  配置apache+nginx   caddyserver.com
go get github.com/revel/cmd/revel          # 高生产率的全栈web框架 *11k > revel new -a my-app -r
go get github.com/graphql-go/graphql       # Facebook开源API查询语言 *5k  GraphQL中文网™ graphql.org.cn
go get github.com/graph-gophers/graphql-go # GraphQL api server      *3k
go get golang.org/x/oauth2                 # OAuth 2.0 认证授权       *2k   github.com/golang/oauth2
go get github.com/casbin/casbin            # 授权访问-认证服务(强力推荐)*5k  (ACL, RBAC, ABAC) casbin.org
go get github.com/volatiletech/authboss    # 授权访问-认证服务(强力推荐)*2k  CSRF,Throttle,Auth(Password|OAuth2|2fa[totp.sms]),Regist,Lock,Expire等
go get github.com/bitly/oauth2_proxy       # 反向代理-认证服务(推荐) *5k (OAuth2.0, OpenID Connect; Google,Github...
go get github.com/ory/fosite/...           # 访问控制-认证服务易扩展 *1k (OAuth2.0, OpenID Connect...官网 www.ory.sh
go get github.com/qor/auth                 # 模块化身份验证系统, 易于集成和二次开发(推荐) *1k
go get go.uber.org/ratelimit               # 速率限制 github.com/uber-go/ratelimit
go get github.com/juju/ratelimit           # 速率限制-由高效的令牌桶实现(推荐)*1k 调用Bucket方法及限流Read\Write
go get golang.org/x/time                   # 速率限制-调用Limiter接口方法 import golang.org/x/time/rate
go get github.com/sony/gobreaker           # 熔断功能-断路器模式(推荐)breaker.CircuitBreaker  www.sony.net
go get github.com/afex/hystrix-go          # 熔断功能-频率限制qps
go get github.com/jaegertracing/jaeger-client-go # 分布式链路追踪系统 *9.6k CNCF(推荐) github.com/jaegertracing/jaeger
go get github.com/fvbock/endless           # 站点零停机\重启
go get github.com/codegangsta/gin          # 站点热启动 > gin -h
go get github.com/ochinchina/supervisord   # 开机启动服务 > supervisord -d -c website.conf
go get github.com/sourcegraph/checkup/cmd/checkup # 分布式站点健康检查工具 > checkup --help
go get github.com/dgrijalva/jwt-go/cmd/jwt # JSON Web Tokens (JWT)   *6k
go get github.com/appleboy/gin-jwt         # JWT Middleware for Gin  *1k
go get github.com/thoas/stats              # Http Router Filter[计时] *1k
go get github.com/gorilla/sessions         # session & cookie authentication            *1.5k
go get github.com/kgretzky/evilginx2       # session & cookie, 2-factor authentication  *2.5k
go get github.com/dpapathanasiou/go-recaptcha # Google验证码|申请(推荐) www.google.com/recaptcha/admin/create
go get github.com/dchest/captcha           # 验证码|图片|声音(推荐)
go get github.com/mojocn/base64Captcha     # 验证码|展示 | captcha.mojotv.cn
go get github.com/emersion/go-imap/...     # 邮箱服务 IMAP library for clients and servers
go get github.com/sdwolfe32/trumail/...    # 邮箱验证 clients
go get github.com/matcornic/hermes/v2      # HTML e-mails, like: npm i mailgen  github.com/eladnava/mailgen
go get github.com/fagongzi/gateway         # 基于HTTP协议的restful的API网关, 可以作为统一的API接入层
go get github.com/wanghongfei/gogate       # 高性能Spring Cloud网关, 路由配置热更新、负载均衡、灰度、服务粒度的流量控制、服务粒度的流量统计
go get github.com/go-swagger/go-swagger/cmd/swagger # 后端API文档生成器 > swagger generate spec --scan-models -o docs/spec.json

# 微服务(分布式RPC框架)rpcx，支持Zookepper、etcd、consul服务发现&路由 *3k books.studygolang.com/go-rpc-programming-guide
go get -u -v -tags "reuseport quic kcp zookeeper etcd consul ping rudp utp" github.com/smallnest/rpcx/...
# 谷歌开源gRPC  grpc.io/docs/quickstart/go & 'HTTP/2'传输更快 http2.golang.org
 # 1.安装: protoc、genproto; <protoc>插件: protoc-gen-go、protoc-gen-gogo、protoc-gen-gofast;prototool(增强插件)
 > github.com/google/protobuf/releases     # 先下载protobuf-command > protoc.exe, protoc
 > git clone --depth=1 https://github.com/grpc/grpc-go.git %GOPATH%/src/google.golang.org/grpc
 > git clone --depth=1 https://github.com/google/go-genproto %GOPATH%/src/google.golang.org/genproto
 > go get github.com/golang/{text,net}                                       # 安装protoc的依赖 ↓
 > go get github.com/golang/protobuf/{proto,protoc-gen-go}                   # 安装插件protoc-gen-go ↓
   $ protoc --go_out=plugins=grpc:. *.proto                                  # 使用segmentfault.com/a/1190000009277748
 > go get github.com/gogo/protobuf/{proto,protoc-gen-gogo,protoc-gen-gofast} # 推荐gofast性能高于protoc-gen-go ↓
   $ protoc --gogo_out=plugins=grpc:. *.proto  ||  protoc --gofast_out=plugins=grpc:. *.proto  (使用gofast插件)
   # ⚡ gRPC-Gateway (gRPC to JSON proxy: 接口用例) + swagger + validators ↓
   > git clone --depth=1 https://github.com/gogo/grpc-example.git && set GO111MODULE=on && go build -mod=vendor && grpc-example.exe
  $ prototool help                         # 增强版protoc优步推荐 github.com/uber/prototool <ubuntu>
 > go get github.com/fullstorydev/grpcurl  # 查询工具grpcurl<服务列表&调用方法+反射服务> google.golang.org/grpc/reflection
 # 2.使用: gRPC-Examples > cd %GOPATH%/src/google.golang.org/grpc/examples/helloworld
 > protoc -I ./helloworld --go_out=plugins=grpc:./helloworld ./helloworld/helloworld.proto #2.1生成代码*.pb.go
 > go run ./greeter_server/main.go ; go run ./greeter_client/main.go                       #2.2启动服务&客户端
go get github.com/grpc-ecosystem/grpc-gateway/... # 谷歌开源网关(gRPC to JSON proxy: 读取protobuf,生成反向代理)
go get github.com/grpc-ecosystem/go-grpc-middleware #auth,logrus,prometheus⚡,opentracing,validator,recovery,ratelimit;retry
go get github.com/TykTechnologies/tyk  # Tyk开源|服务网关API:auth,grantKeyAccess&keyExpiry,ratelimit,analytics,quotas,webhooks,IP/Blacklist/Whitelist,restart,versioning
go get github.com/istio/istio              # 谷歌开源|微服务集群管理k8s  *17k istio.io | www.grpc.io
go get github.com/go-kit/kit/cmd/kitgen    # 阿里推荐|微服务构建框架gRPC *13k gokit.io (推荐)
go get github.com/apache/thrift/lib/go/thrift/... #滴滴推荐|Thrift协议的高性能RPC框架 *7k (推荐) thrift.apache.org > thrift -help
go get github.com/bilibili/kratos/tool/kratos # bilibili开源微服务框架|包含大量微服务框架工具 *6k
go get github.com/bilibili/sniper             # bilibili开源轻量级业务框架,mvc+rpc业务工具库(推荐) *1k
go get github.com/TarsCloud/TarsGo/tars    # 腾讯开源|基于Tars协议的高性能RPC框架 *2k 网关+容器化+服务治理(推荐)
go get github.com/micro/go-micro           # 开源Micro分布式RPC微服务 *7k (推荐)
go get github.com/jhump/protoreflect       # protobuf文件动态解析接口，可以实现反射相关的能力

go get github.com/gocolly/colly/...        # 高性能Web采集利器 *7k
go get github.com/henrylee2cn/pholcus      # 重量级爬虫Pholcus(幽灵蛛) *5k
go get github.com/MontFerret/ferret        # 声明式Web爬虫系统 *4k
go get github.com/tealeg/xlsx              # 读写Excel文件     *4k
go get github.com/360EntSecGroup-Skylar/excelize # 读写Excel文件(推荐) *5k
go get github.com/davyxu/tabtoy            # 高性能便捷电子表格导出器   *1k
go get github.com/claudiodangelis/qr-filetransfer # 二维码识别|qr转换  *3k
go get github.com/skip2/go-qrcode/...      # 二维码生成器 > qrcode     *1k
go get github.com/go-echarts/go-echarts/... # 数据可视化图表库:25+图表:400+地图 go-echarts.github.io/go-echarts
go get github.com/jung-kurt/gofpdf         # 创建PDF文件  *2.8k | 支持text,drawing,images
go get github.com/unidoc/unipdf/...        # 创建和处理PDF文件 *1k  unidoc.io
# Gotenberg is a Docker-powered stateless API for converting HTML, Markdown and Office documents to PDF.
# https://thecodingmachine.github.io/gotenberg/#url.basic.c_url
go get github.com/rakyll/statik            # 将静态资源文件嵌入到Go二进制文件中，提供http服务> statik -src=/path/to
go get github.com/yudai/gotty              # 终端扩展为Web网站服务 *12.3k
go get github.com/libp2p/go-libp2p         # 网络库模块p2p-serves
go get github.com/libp2p/go-libp2p-examples# 网络库模块p2p-examples

go get github.com/gorilla/websocket        # WebSocket Serve(推荐1) *10.k 一个快速，测试良好，广泛使用的WebSocket
go get github.com/joewalnes/websocketd     # Websocket Serve(推荐2) *14.k 将STDIN/STDOUT程序转换为WebSocket服务器
go get github.com/gotify/server            # WebSocket Serve(推荐3) *4.3k 提供Web管理及客户端App推送功能 gotify.net
go get github.com/googollee/go-socket.io   # WebSocket Serve(推荐4) *3.1k 提供完整的WebSocket接口处理 socket.io/docs
go get github.com/gobwas/ws                # WebSocket Serve(推荐5) *2.7k 支持百万级连接数 github.com/socketio/socket.io
# 聊天室 git clone --depth=1 https://github.com/GoBelieveIO/im_service.git && cd im_service && dep ensure && mkdir bin && make install
# 高并发 go get github.com/xiaojiaqi/10billionhongbaos  # 抢购系统：单机支持QPS达6万，可以满足100亿红包的压力测试
# https://github.com/oikomi/FishChatServer2 消息服务与聊天功能，支持容器部署 (Kubernetes + Docker)

go get github.com/enriquebris/goconcurrentqueue # 高并发-队列-线程安全(推荐)
go get github.com/beeker1121/goque         # 高性能-堆栈-队列-数据存储(推荐) & LevelDB Stores
go get github.com/eapache/channels         # 通道：`Distribute分发`1In*Out,`Multiplex多路复用`*In1Out,`Pipe管道`1In1Out,`BatchingChannel批量通道`...
go get github.com/robfig/cron              # 任务计划 a cron library *4k
go get github.com/gocraft/work             # do work of redis-queue *1k | github.com/gocraft/work#run-the-web-ui
go get github.com/lisijie/webcron          # 定时任务Web管理器 (基于beego框架) *1k
go get github.com/shunfei/cronsun          # 分布式容错任务管理系统 *1.5k
go get github.com/gocelery/gocelery        # 分布式任务队列管理系统 *1k client/server | www.celeryproject.org
go get github.com/RichardKnop/machinery/v1 # 分布式消息队列+异步任务(强烈推荐) *3.5k
go get github.com/benmanns/goworker        # 10万级并行的后台任务系统-基于Redis的workers(推荐) *2.3k
go get github.com/streadway/amqp           # RabbitMQ tutorials *3k | www.rabbitmq.com | github.com/rabbitmq/rabbitmq-tutorials
go get github.com/blackbeans/kiteq         # KiteQ 是一个基于 go + protobuff + zookeeper 实现的多种持久化方案的mq框架

go get github.com/nsqio/nsq                # 实时分布式消息平台nsq(推荐) *15k | nsqlookupd & nsqd & nsqadmin https://nsq.io
go get github.com/youzan/nsq               # 有赞科技的nsq重塑 | www.bilibili.com/video/av29142217
go get github.com/nats-io/nats-server      # 消息系统nats服务端 *7k | http://nats.io
go get github.com/nats-io/nats.go/         # 消息系统nats客户端 | docs.nats.io/developing-with-nats/tutorials/
go get github.com/appleboy/gorush          # 消息推送gorush服务(推荐)  *4k : api, notification queue, multiple workers
go get github.com/Shopify/sarama           # 消息系统Kafka客户端(推荐) *5k : github.com/bsm/sarama-cluster [集群客户端]
go get github.com/travisjeffery/jocko      # 消息系统Kafka原生实现Serve*3k : producing/consuming[生产/消费] cluster[zk集群]
go get github.com/mattermost/mattermost-server #通讯 *15k 为您带来跨PC和移动设备的消息+文件分享，提供归档+搜索功能+前端React

# 物联网IoT、物理计算Drones、机器人Robotics
go get -d -u gobot.io/x/gobot/...          # 物联网IoT开源框架 *5k | github.com/hybridgroup/gobot
# github.com/LiteOS/LiteOS                 # 华为LiteOS是华为面向物联网领域开发的一个基于实时内核的轻量级操作系统
go get -d github.com/emqx/emqx             # 百万级分布式开源物联网-消息服务平台 *4k | www.emqtt.com
go get -u -v github.com/davyxu/cellnet     # 游戏服务器RPC *2.5k | ARM设备<设备间网络通讯> | 证券软件<内部RPC>
go get -u -v github.com/liangdas/mqant     # 游戏服务器RPC *1.5k
git clone --depth=1 https://github.com/EasyDarwin/EasyDarwin.git %GOPATH%/src/github.com/EasyDarwin/EasyDarwin # RTSP流媒体服务
go get github.com/iikira/BaiduPCS-Go       # 百度网盘命令行客户端
go get github.com/inconshreveable/go-update # 自动更新应用程序
go get -d https://github.com/restic/restic  # 数据备份工具 | restic.readthedocs.io
cd %GOPATH%/src/github.com/restic/restic && go run -mod=vendor build.go --goos windows --goarch amd64
# ------------------------------------------------------------------------------------
# 部署-维护
# ------------------------------------------------------------------------------------
go get github.com/containous/yaegi/cmd/yaegi # 一个优雅的 Go 解释器
go get github.com/martinlindhe/gohash/...  # 编码解码> coder --help ; 推荐> hasher --help [-A加密算法][-E编码解码]
 > echo2 123456|hasher md5 -n --no-colors  # echo2解决各系统的换行问题(个人编译,用于替换echo)
 > echo2 hello |coder -e base64+hex > echo2 614756736247383d|coder -d hex+base64 > cat file.b64|coder -d base64
 > echo2 123456|hasher md5 [-e base64] [-n --no-colors] > hasher -i file.txt sha1 --bsd #加密文件&输出BSD格式化结果
go get github.com/hidevopsio/crypto        # 加密解密> crypto rsa -h [rsa -e -s hello][rsa -d -s ***][-k:rsa.key]
 $ wget https://github.com/smallstep/cli/releases/download/v0.11.0/step-cli_0.11.0_amd64.deb
 $ wget https://github.com/smallstep/certificates/releases/download/v0.11.0/step-certificates_0.11.0_amd64.deb
 $ sudo dpkg -i step-cli_0.11.0_amd64      # 自动化证书管理cli: X.509,TLS;OAuth OIDC|OTP;JSONWebEncrypt;JWT...
 $ sudo dpkg -i step-certificates_0.11.0_amd64.deb && step version && step-ca version
 $ step ca init #1.初始化ca
 $ step-ca $(step path)/config/ca.json #2.设置ca密码
 $ cat > srv.go #3.开发server
 $ step ca certificate localhost srv.crt srv.key #4.从StepCA获取服务器标识>step certificate inspect --bundle srv.crt
 $ go run srv.go & #5.运行server
 $ step ca root root.crt #6.从StepCA获取根证书
 $ curl --cacert root.crt https://localhost:8443/hi #7.验证CA,使用HTTP.TLS向服务器发出经过身份验证的加密curl请求
 $ step certificate inspect https://www.baidu.com # 查看网站证书Certificate
go get github.com/smallstep/autocert       # 自动化证书管理 for Docker kubernetes ^1.9
go get github.com/vbauerster/mpb/...       # 在终端为 Go 命令行应用程序显示进度条
go get github.com/shazow/ssh-chat          # 自定义 SSH server 用于替代 shell  *3.6k
go get github.com/elves/elvish             # <shell for unix>可编程：数组、字典、传递对象的增强型管道、闭包、模块机制、类型检查
go get github.com/mattn/sudo               # sudo for windows > sudo cmd /c dir ; sudo notepad c:\windows\system32\drivers\etc\hosts
go get github.com/google/gousb             # 用于访问USB设备的低级别接口
go get github.com/google/gops              # 用于列出并诊断Go应用程序进程
go get github.com/google/pprof             # 用于可视化和分析性能和数据的工具
go get github.com/google/mtail             # 用于从应用程序日志中提取白盒监视数据，以便收集到时间序列数据库中
go get github.com/google/godepq            # 用于查询程序依赖 > godepq -from github.com/google/pprof
go get github.com/google/ko/cmd/ko         # 用于构建和部署应用程序到Kubernetes的工具
go get github.com/go-task/task             # 一个任务运行/构建工具，旨在比 GNU Make 更简单易用 *2k
go get github.com/drakkan/sftpgo           # 全功能和高度可配置SFTP服务器
go get github.com/FiloSottile/mkcert       # 证书管理工具 *18k  > mkcert > https://localhost
# [申请Let's Encrypt永久免费SSL证书]         www.jianshu.com/p/3ae2f024c291
go get github.com/go-acme/lego/cmd/lego    # Let's Encrypt client and ACME library, DNS providers manager.
go get github.com/google/git-appraise/git-appraise # 用于Git版本管理的分布式代码审核
go get github.com/google/easypki/cmd/easypki # CA证书申请工具 | API: go get gopkg.in/google/easypki.v1
go get go.universe.tf/tcpproxy/cmd/tlsrouter # TLS代理根据握手的SNI（服务器名称指示）将连接路由到后端。它不携带加密密钥，无法解码其代理的流量
go get github.com/prometheus/prometheus/cmd/... # 服务监控系统和时间序列数据库 *23k | prometheus.io/community
go get github.com/grafana/grafana          # 漂亮的监测系统|指标分析|InfluxDB时序DB|Prometheus等(强力推荐) *30k
go get github.com/rsc/goversion            # 扫描目录中Go可执行文件的版本信息 > goversion /usr/bin
go get github.com/yinqiwen/gscan           # 扫描可用HTTPsIP、修复Hosts、可用GoogleIP; 可用于代理工具GSnova,GoAgent
go get github.com/BurntSushi/wingo/wingo-cmd # 一个功能齐全的窗口管理器 > wingo-cmd
# 源代码版本管理 git version manage
go get -d github.com/gogs/gogs  # 一款极易搭建的自助Git服务  *30k
go get -d github.com/github/hub # 轻松使用Github的命令行工具 *17k
go get gitea.com/lunny/gps                 # 地图坐标系转换
/** WGS84坐标系：即地球坐标系，国际上通用的坐标系。设备一般包含GPS芯片或者北斗芯片获取的经纬度为WGS84地理坐标系,
 * 谷歌地图采用的是WGS84地理坐标系（中国范围除外）;
 *  GCJ02坐标系：即火星坐标系，是由中国国家测绘局制订的地理信息系统的坐标系统。由WGS84坐标系经加密后的坐标系。
 * 谷歌中国地图和搜搜中国地图采用的是GCJ02地理坐标系; BD09坐标系：即百度坐标系，GCJ02坐标系经加密后的坐标系;
 * 搜狗坐标系、图吧坐标系等，估计也是在GCJ02基础上加密而成的。
 */

# 小米公司的互联网企业级监控系统 | book.open-falcon.org
# 各大 Go 模板引擎的对比及压力测试 | github.com/SlinSo/goTemplateBenchmark

~~~

----

#### 云平台|公众平台|在线支付
~~~
# 云计算、云数据库、点播、直播...
# ------------------------------------------------------------------------------------
# 亚马逊 AWS | www.amazonaws.cn/tools

# 谷歌云 Google Cloud Platform | cloud.google.com/go | github.com/GoogleCloudPlatform
go get -u github.com/google/go-cloud      # 云计算
go get -u cloud.google.com/go/storage     # 在 Cloud Storage 中存储和归档数据
go get -u cloud.google.com/go/bigquery    # 使用 Google BigQuery 执行数据分析
go get -u cloud.google.com/go/pubsub      # 使用 Pub/Sub 设置完全托管的事件驱动型消息传递系统
go get -u cloud.google.com/go/translate   # 使用 Translation API 翻译不同语言的文本
go get -u cloud.google.com/go/vision/apiv1# 使用 Vision API 分析图片

# 阿里云 | api.aliyun.com | developer.aliyun.com/sdk
# 云服务器 ECS、对象存储 OSS、阿里云关系型数据库、云数据库MongoDB版、CDN、VPC、
# 视频点播、音视频通信、媒体转码、负载均衡、云监控、容器服务、邮件推送、弹性伸缩、移动推送、日志服务、交易与账单管理
# 阿里云 SDK
go get -u github.com/aliyun/alibaba-cloud-sdk-go
go get -u github.com/aliyun/alibaba-cloud-sdk-go/sdk  # cloud sdk
# 存储服务
go get github.com/aliyun/aliyun-oss-go-sdk
# 日志服务
go get github.com/aliyun/aliyun-log-go-sdk
# 函数计算
go get github.com/aliyun/fc-go-sdk
# 消息服务
go get github.com/aliyun/aliyun-mns-go-sdk
# 表格存储tablestore
go get github.com/aliyun/aliyun-tablestore-go-sdk
# 消息队列 MQ
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/ons
# 消息队列 Kafka
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/alikafka
# API 网关
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/cloudapi
# 表格存储ots
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/ots
# 链路追踪
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/xtrace
# Web应用防火墙
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/waf_openapi
# 专有网络VPC
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/vpc
# 视频点播
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/vod
# 视频直播
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/live
# 媒体处理
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/mts
# 音视频通信
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/rtc
# 资源编排
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/ros
# 智能媒体管理
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/imm
# 商标服务
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/trademark
# 云通信网络加速
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/snsuapi
# 智能接入网关
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/smartag
# sls服务
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/sls
# 负载均衡
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/slb
# 敏感数据保护
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/sddp
# 实时计算（流计算）
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/cusanalytic_sc_online
# 云安全中心（态势感知）
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/sas
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/sas-api
# 风险识别
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/saf
# Serverless 应用引擎
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/sae
# 访问控制
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/ram
# 云解析 PrivateZone
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/pvtz
# 云数据库 POLARDB
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/polardb
# 云数据库 MongoDB 版
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/dds
# 云数据库 Redis 版
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/r_kvstore
# HybridDB for MySQL
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/petadata
# 分析型数据库PostgreSQL版
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/gpdb
# 分布式关系型数据库服务
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/drds
# 大数据E-MapReduce
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/emr
# 数据管理
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/dms-enterprise
# 数据库备份
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/dbs
# 开放搜索
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/opensearch
# 图像识别
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/imagesearch
# Data Lake Analytics
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/openanalytics
# 数据库审计
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/yundun_dbaudit
# 数据库和应用迁移
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/tesladam
# 操作审计
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/actiontrail
# 归档存储
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/foas
# 文件存储NAS
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/nas
# 密钥管理服务
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/kms
# 智能视觉
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/ivision
# 语音服务
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/dyvmsapi
# 人脸识别
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/linkface
# 物联网平台
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/iot
# 加密服务
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/hsm
# 函数工作流
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/fnf
# 弹性伸缩
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/ess
# 边缘节点服务 ENS
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/ens
# 弹性Web托管
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/elasticsearch
# 弹性高性能计算 E-HPC
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/ehpc
# 企业级分布式应用服务
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/edas
# 云服务器 ECS
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/ecs
# 弹性容器实例 ECI
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/eci
# 短信服务
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/dysmsapi
# 号码认证服务
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/dypnsapi
# 号码隐私保护
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/dyplsapi
# 数据传输
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/dts
# 域名
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/domain
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/domain-intl
# SSL证书（CA证书服务、数据安全）
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/cas
# 云解析DNS
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/alidns
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/httpdns
# 安全加速 SCDN
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/scdn
# 全站加速
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/dcdn
# CDN
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/cdn
# 企业工商注册服务
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/companyreg
# 云监控
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/cms
# 实人认证
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/cloudauth
# 凭证管理
go get github.com/aliyun/credentials-go
# 区块链服务
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/baas
# 应用实时监控服务
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/arms
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/arms4finance
# 应用高可用服务
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/ahas_openapi
# 容器镜像服务
go get github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/cr


# 腾讯云 | console.cloud.tencent.com/api/explorer
go get -u github.com/tencentcloud/tencentcloud-sdk-go
# 云服务器 黑石物理服务器 云硬盘 容器服务 容器实例服务 弹性伸缩 无服务器云函数 批量计算
# 负载均衡 私有网络 专线接入 云数据库 MySQL 云数据库 Redis 云数据库 MongoDB 数据传输服务 DTS 云数据库 MariaDB
# 分布式数据库 DCDB	云数据库 SQL'Server 云数据库 PostgreSQL 内容分发网络 主机安全 Web漏洞扫描 应用安全 云点播
# 云直播 智能语音服务 机器翻译 智能钛机器学习 催收机器人 智聆口语评测 腾讯优评 Elasticsearch'Service
# 物联网通信 TBaaS 云监控 迁移服务平台 电子合同服务 计费相关 渠道合作伙伴 人脸核身'云智慧眼
# 威胁情报云查 样本智能分析平台 数学作业批改 人脸融合 人脸识别 数字版权管理


# 七牛云 | developer.qiniu.com/sdk
go get -u github.com/qiniu/api.v7 # 对象存储

# CDN.内容分发网络
# 又拍云 | www.upyun.com

# 实时通信云
# 野狗 | www.wilddog.com 实时通信引擎（Sync），它帮助企业快速实现三大功能：基础实时通信、实时数据分发和实时数据持久化。

# 应用数据、开放数据API
# 聚合数据 | www.juhe.cn

# 微信公众平台SDK
go get -u gopkg.in/chanxuehong/wechat.v2/... # 微信公众平台、企业号、微信支付 github.com/chanxuehong/wechat 
# https://github.com/sidbusy/weixinmp
# https://github.com/arstd/weixin
# https://github.com/wizjin/weixin
# https://github.com/hoperong/RabbitGo
# https://github.com/Cheney-Su/go_weixin
# https://github.com/i11cn/go_weixin
# 微信支付SDK
# https://github.com/imzjy/wxpay
# 微信支付／支付宝支付
# https://github.com/philchia/gopay
# 微信公众平台/微信企业号/微信商户平台/微信支付
# https://github.com/philsong/wechat2
~~~

#### Google开源

*Go team:*

- https://github.com/google/go-github
- https://github.com/golang/protobuf
- https://github.com/golang/oauth2
- https://github.com/golang/glog
- https://github.com/golang/geo
- https://github.com/golang/groupcache
- https://github.com/golang/snappy
- https://github.com/golang/freetype
- https://github.com/google/gxui/...

*Google team:*

- https://github.com/googleapis/googleapis
- https://github.com/google/btree
- https://github.com/google/go-cloud
- https://github.com/google/gops
- https://github.com/google/gvisor
- https://github.com/google/google-api-go-client
- https://github.com/grpc/grpc-go

#### *WebAssembly*

- https://tip.golang.org/pkg/syscall/js
- https://github.com/golang/go/tree/master/misc/wasm
- https://github.com/chai2010/awesome-wasm-zh
- https://github.com/mbasso/awesome-wasm
- https://gopry.rice.sh/

#### *WebRTC*

- https://github.com/pion/webrtc  | Pion WebRTC v2 | https://github.com/pion/webrtc/tree/v2.1.4
- https://github.com/pion/turn  | An extendable TURN server
- https://github.com/pion/example-webrtc-applications  | Examples
- https://w3c.github.io/webrtc-pc  | Pion WebRTC API for JavaScript
- https://github.com/pion/webrtc/blob/master/examples/README.md#webassembly | WebAssembly
~~~bash
GO111MODULE=on go get github.com/pion/webrtc/v2
cd $GOPATH/src/github.com/pion/webrtc/examples
go run examples.go --address=":8236"       # WebRTC example
GOOS=js GOARCH=wasm go build -o demo.wasm  # WebAssembly demo
~~~

#### *GUI - HTML/JS/CSS*

 * [Electron](https://github.com/asticode/go-astilectron)
    * Install astilectron-bundler & Play Demos
~~~bash
# Download : astilectron & electron***
 - https://github.com/asticode/astilectron/releases > %GOPATH%/bin/astibundler/astilectron-0.32.0.zip
 - https://github.com/electron/electron/releases/tag/v4.0.1 > electron-windows-amd64-4.0.1.zip
# Install : astilectron-bundler
go get -u github.com/asticode/go-astilectron-bundler/...
go install github.com/asticode/go-astilectron-bundler/astilectron-bundler
# Demo 1 : video tools
go get github.com/asticode/go-astivid/...
cd %GOPATH%/src/github.com/asticode/go-astivid
cp -r %GOPATH%/bin/astibundler/* C:/Users/ADMINI~1/AppData/Local/Temp/astibundler/cache
rm -f bind*.go                # delete file before bundle
astilectron-bundler -v        # help: astilectron-bundler -h
~~~
 * [QT](https://github.com/therecipe/qt)
    * 百度网盘客户端Qt5+websocket+p2p+eventbus - https://github.com/peterq/pan-light
~~~bash
# [QT跨平台应用框架] Qt binding package
go get -u -v github.com/therecipe/qt/cmd/... && for /f %v in ('go env GOPATH') do %v\bin\qtsetup test && %v\bin\qtsetup
go get github.com/lxn/win                  # Windows API wrapper package
go get github.com/lxn/walk                 # Windows UI Application Library Kit *3k
go get github.com/google/gapid             # Windows UI App : Graphics API Debugger
~~~
 * [Webview](https://github.com/zserge/webview)
 * [WebAssembly](https://github.com/murlokswarm/app)
 * [go-flutter-desktop](https://github.com/go-flutter-desktop/go-flutter)
 * [go-material-design-GUI](https://fyne.io/develop/)、[google-andlabs-GUI](https://github.com/andlabs/ui)

----

#### ③ [开源的 Web 框架](https://github.com/avelino/awesome-go#web-frameworks)

 * Web 框架
    * [基于 Gin 构建企业级 RESTful API 服务](https://juejin.im/book/5b0778756fb9a07aa632301e)
    * [基于 Gin 一步一步搭建Go的Web服务器](https://www.hulunhao.com/go/go-web-backend-starter/)
~~~
# 开发
cd %GOPATH%/src                                                                 # 项目框架 Gin Web Framework
git clone --depth=1 https://github.com/lexkong/apiserver_demos apiserver                  # 项目源码-复制^demo至-工作目录
git clone --depth=1 https://github.com/lexkong/vendor                                     # 项目依赖-govendor
go get github.com/StackExchange/wmi                                             # 项目依赖-缺失的包
# 构建
cd %GOPATH%/src/apiserver && go fmt -w . && go tool vet . && go build -v -o [应用名] [目录默认.]
# 运行
%GOPATH%/src/apiserver/apiserver.exe
~~~

----

#### ④ [中文标准库文档](https://studygolang.com/pkgdoc)

 * [项目布局设计*目录*](https://github.com/golang-standards/project-layout)
 
----

#### ⑤ 阅读相关文章

 * 高性能
    * [高并发架构解决方案](https://studygolang.com/articles/15479)

----

> Docker 编译器(可选) [Golang + custom build tools](https://hub.docker.com/_/golang)

~~~shell
# 1. pull build tools: Glide, gdm, go-test-teamcity
docker pull jetbrainsinfra/golang:1.11.5
docker pull golang:1.4.2-cross
docker run --rm -v "$PWD":/usr/src/app -w /usr/src/app -e GOOS=windows -e GOARCH=386 golang:1.11.5 go build -v
# 2. run docker container
docker run --name golang1115 -d jetbrainsinfra/golang:1.11.5 bash
docker cp golang1115:/go/src/github.com %GOPATH%\src
docker cp golang1115:/go/src/golang.org %GOPATH%\src
docker run --name golang1115 -td -p 8080:8080 -v %GOPATH%\src:/go/src -w /go/src jetbrainsinfra/golang:1.11.5
# 3. go build
docker exec -it golang1115 bash
  $ cd apiserver & go build & ./apiserver                                                # build for linux
  $ for GOOS in linux windows; do GOOS=$GOOS go build -v -o apiserver-$GOOS-amd64; done; # if GOARCH="amd64"
    mv apiserver-windows-amd64 apiserver-windows-amd64.exe  # windows文件重命名           # for linux&windows
~~~

----


# 语法速查表

1. [Basic Syntax](#basic-syntax)
2. [Operators](#operators)
    * [Arithmetic](#arithmetic)
    * [Comparison](#comparison)
    * [Logical](#logical)
    * [Other](#other)
3. [Declarations](#declarations)
4. [Functions](#functions)
    * [Functions as values and closures](#functions-as-values-and-closures)
    * [Variadic Functions](#variadic-functions)
5. [Built-in Types](#built-in-types)
6. [Type Conversions](#type-conversions)
7. [Packages](#packages)
8. [Control structures](#control-structures)
    * [If](#if)
    * [Loops](#loops)
    * [Switch](#switch)
9. [Arrays, Slices, Ranges](#arrays-slices-ranges)
    * [Arrays](#arrays)
    * [Slices](#slices)
    * [Operations on Arrays and Slices](#operations-on-arrays-and-slices)
10. [Maps](#maps)
11. [Structs](#structs)
12. [Pointers](#pointers)
13. [Interfaces](#interfaces)
14. [Embedding](#embedding)
15. [Errors](#errors)
16. [Concurrency](#concurrency)
    * [Goroutines](#goroutines)
    * [Channels](#channels)
    * [Channel Axioms](#channel-axioms)
17. [Printing](#printing)
18. [Reflection](#reflection)
    * [Type Switch](#type-switch)
    * [Examples](https://github.com/a8m/reflect-examples)
19. [Snippets](#snippets)
    * [Http-Server](#http-server)
    * [string、utf8,utf16、slice...](#stringutf8utf16slice)

----

# Basic Syntax
  * http://tour.golang.org

## Hello World
File `hello.go`:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Go")
}
```
`$ go run hello.go`

## Operators
### Arithmetic
|Operator|Description|
|--------|-----------|
|`+`|addition|
|`-`|subtraction|
|`*`|multiplication|
|`/`|quotient|
|`%`|remainder|
|`&`|bitwise and|
|`\|`|bitwise or|
|`^`|bitwise xor|
|`&^`|bit clear (and not)|
|`<<`|left shift|
|`>>`|right shift|

### Comparison
|Operator|Description|
|--------|-----------|
|`==`|equal|
|`!=`|not equal|
|`<`|less than|
|`<=`|less than or equal|
|`>`|greater than|
|`>=`|greater than or equal|

### Logical
|Operator|Description|
|--------|-----------|
|`&&`|logical and|
|`\|\|`|logical or|
|`!`|logical not|

### Other
|Operator|Description|
|--------|-----------|
|`&`|address of / create pointer|
|`*`|dereference pointer|
|`<-`|send / receive operator (see 'Channels' below)|

## Declarations
Type goes after identifier!
```go
var foo int // declaration without initialization
var foo int = 42 // declaration with initialization
var foo, bar int = 42, 1302 // declare and init multiple vars at once
var foo = 42 // type omitted, will be inferred
foo := 42 // shorthand, only in func bodies, omit var keyword, type is always implicit
const constant = "This is a constant"

// iota can be used for incrementing numbers, starting from 0
const (
    _ = iota
    a
    b
    c = 1 << iota
    d
)
    fmt.Println(a, b) // 1 2 (0 is skipped)
    fmt.Println(c, d) // 8 16 (2^3, 2^4)
```

## Functions
```go
// a simple function
func functionName() {}

// function with parameters (again, types go after identifiers)
func functionName(param1 string, param2 int) {}

// multiple parameters of the same type
func functionName(param1, param2 int) {}

// return type declaration
func functionName() int {
    return 42
}

// Can return multiple values at once
func returnMulti() (int, string) {
    return 42, "foobar"
}
var x, str = returnMulti()

// Return multiple named results simply by return
func returnMulti2() (n int, s string) {
    n = 42
    s = "foobar"
    // n and s will be returned
    return
}
var x, str = returnMulti2()

```

### Functions As Values And Closures
```go
func main() {
    // assign a function to a name
    add := func(a, b int) int {
        return a + b
    }
    // use the name to call the function
    fmt.Println(add(3, 4))
}

// Closures, lexically scoped: Functions can access values that were
// in scope when defining the function
func scope() func() int{
    outer_var := 2
    foo := func() int { return outer_var}
    return foo
}

func another_scope() func() int{
    // won't compile because outer_var and foo not defined in this scope
    outer_var = 444
    return foo
}


// Closures
func outer() (func() int, int) {
    outer_var := 2
    inner := func() int {
        outer_var += 99 // outer_var from outer scope is mutated.
        return outer_var
    }
    inner()
    return inner, outer_var // return inner func and mutated outer_var 101
}
```

### Variadic Functions
```go
func main() {
	fmt.Println(adder(1, 2, 3)) 	// 6
	fmt.Println(adder(9, 9))	// 18

	nums := []int{10, 20, 30}
	fmt.Println(adder(nums...))	// 60
}

// By using ... before the type name of the last parameter you can indicate that it takes zero or more of those parameters.
// The function is invoked like any other function except we can pass as many arguments as we want.
func adder(args ...int) int {
	total := 0
	for _, v := range args { // Iterates over the arguments whatever the number.
		total += v
	}
	return total
}
```

## Built-in Types
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32 ~= a character (Unicode code point) - very Viking

float32 float64

complex64 complex128
```

## Type Conversions
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// alternative syntax
i := 42
f := float64(i)
u := uint(f)
```

## Packages
* Package declaration at top of every source file
* Executables are in package `main`
* Convention: package name == last name of import path (import path `math/rand` => package `rand`)
* Upper case identifier: exported (visible from other packages)
* Lower case identifier: private (not visible from other packages)

## Control structures

### If
```go
func main() {
	// Basic one
	if x > 10 {
		return x
	} else if x == 10 {
		return 10
	} else {
		return -x
	}

	// You can put one statement before the condition
	if a := b + c; a < 42 {
		return a
	} else {
		return a - 42
	}

	// Type assertion inside if
	var val interface{}
	val = "foo"
	if str, ok := val.(string); ok {
		fmt.Println(str)
	}
}
```

### Loops
```go
    // There's only `for`, no `while`, no `until`
    for i := 1; i < 10; i++ {
    }
    for ; i < 10;  { // while - loop
    }
    for i < 10  { // you can omit semicolons if there is only a condition
    }
    for { // you can omit the condition ~ while (true)
    }
    
    // use break/continue on current loop
    // use break/continue with label on outer loop
here:
    for i := 0; i < 2; i++ {
        for j := i + 1; j < 3; j++ {
            if i == 0 {
                continue here
            }
            fmt.Println(j)
            if j == 2 {
                break
            }
        }
    }

there:
    for i := 0; i < 2; i++ {
        for j := i + 1; j < 3; j++ {
            if j == 1 {
                continue
            }
            fmt.Println(j)
            if j == 2 {
                break there
            }
        }
    }
```

### Switch
```go
    // switch statement
    switch operatingSystem {
    case "darwin":
        fmt.Println("Mac OS Hipster")
        // cases break automatically, no fallthrough by default
    case "linux":
        fmt.Println("Linux Geek")
    default:
        // Windows, BSD, ...
        fmt.Println("Other")
    }

    // as with for and if, you can have an assignment statement before the switch value
    switch os := runtime.GOOS; os {
    case "darwin": ...
    }

    // you can also make comparisons in switch cases
    number := 42
    switch {
        case number < 42:
            fmt.Println("Smaller")
        case number == 42:
            fmt.Println("Equal")
        case number > 42:
            fmt.Println("Greater")
    }

    // cases can be presented in comma-separated lists
    var char byte = '?'
    switch char {
        case ' ', '?', '&', '=', '#', '+', '%':
            fmt.Println("Should escape")
    }
```

## Arrays, Slices, Ranges

### Arrays
```go
var a [10]int // declare an int array with length 10. Array length is part of the type!
a[3] = 42     // set elements
i := a[3]     // read elements

// declare and initialize
var a = [2]int{1, 2}
a := [2]int{1, 2} //shorthand
a := [...]int{1, 2} // elipsis -> Compiler figures out array length
```

### Slices
```go
var a []int                              // declare a slice - similar to an array, but length is unspecified
var a = []int {1, 2, 3, 4}               // declare and initialize a slice (backed by the array given implicitly)
a := []int{1, 2, 3, 4}                   // shorthand
chars := []string{0:"a", 2:"c", 1: "b"}  // ["a", "b", "c"]

var b = a[lo:hi]	// creates a slice (view of the array) from index lo to hi-1
var b = a[1:4]		// slice from index 1 to 3
var b = a[:3]		// missing low index implies 0
var b = a[3:]		// missing high index implies len(a)
a =  append(a,17,3)	// append items to slice a
c := append(a,b...)	// concatenate slices a and b

// create a slice with make
a = make([]byte, 5, 5)	// first arg length, second capacity
a = make([]byte, 5)	// capacity is optional

// create a slice from an array
x := [3]string{"Лайка", "Белка", "Стрелка"}
s := x[:] // a slice referencing the storage of x
```

### Operations on Arrays and Slices
`len(a)` gives you the length of an array/a slice. It's a built-in function, not a attribute/method on the array.

```go
// loop over an array/a slice
for i, e := range a {
    // i is the index, e the element
}

// if you only need e:
for _, e := range a {
    // e is the element
}

// ...and if you only need the index
for i := range a {
}

// In Go pre-1.4, you'll get a compiler error if you're not using i and e.
// Go 1.4 introduced a variable-free form, so that you can do this
for range time.Tick(time.Second) {
    // do it once a sec
}

```

## Maps

```go
var m map[string]int
m = make(map[string]int)
m["key"] = 42
fmt.Println(m["key"])

delete(m, "key")

elem, ok := m["key"] // test if key "key" is present and retrieve it, if so

// map literal
var m = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}

// iterate over map content
for key, value := range m {
}

```

## Structs

There are no classes, only structs. Structs can have methods.
```go
// A struct is a type. It's also a collection of fields

// Declaration
type Vertex struct {
    X, Y int
}

// Creating
var v = Vertex{1, 2}
var v = Vertex{X: 1, Y: 2} // Creates a struct by defining values with keys
var v = []Vertex{{1,2},{5,2},{5,5}} // Initialize a slice of structs

// Accessing members
v.X = 4

// You can declare methods on structs. The struct you want to declare the
// method on (the receiving type) comes between the the func keyword and
// the method name. The struct is copied on each method call(!)
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Call method
v.Abs()

// For mutating methods, you need to use a pointer (see below) to the Struct
// as the type. With this, the struct value is not copied for the method call.
func (v *Vertex) add(n float64) {
    v.X += n
    v.Y += n
}

```
**Anonymous structs:**
Cheaper and safer than using `map[string]interface{}`.
```go
point := struct {
	X, Y int
}{1, 2}
```

## Pointers
```go
p := Vertex{1, 2}  // p is a Vertex
q := &p            // q is a pointer to a Vertex
r := &Vertex{1, 2} // r is also a pointer to a Vertex

// The type of a pointer to a Vertex is *Vertex

var s *Vertex = new(Vertex) // new creates a pointer to a new struct instance
```

## Interfaces
```go
// interface declaration
type Awesomizer interface {
    Awesomize() string
}

// types do *not* declare to implement interfaces
type Foo struct {}

// instead, types implicitly satisfy an interface if they implement all required methods
func (foo Foo) Awesomize() string {
    return "Awesome!"
}
```

## Embedding

There is no subclassing in Go. Instead, there is interface and struct embedding.

```go
// ReadWriter implementations must satisfy both Reader and Writer
type ReadWriter interface {
    Reader
    Writer
}

// Server exposes all the methods that Logger has
type Server struct {
    Host string
    Port int
    *log.Logger
}

// initialize the embedded type the usual way
server := &Server{"localhost", 80, log.New(...)}

// methods implemented on the embedded struct are passed through
server.Log(...) // calls server.Logger.Log(...)

// the field name of the embedded type is its type name (in this case Logger)
var logger *log.Logger = server.Logger
```

## Errors
There is no exception handling. Functions that might produce an error just declare an additional return value of type `Error`. This is the `Error` interface:
```go
type error interface {
    Error() string
}
```

A function that might return an error:
```go
func doStuff() (int, error) {
}

func main() {
    result, err := doStuff()
    if err != nil {
        // handle error
    } else {
        // all is good, use result
    }
}
```

# Concurrency

## Goroutines
Goroutines are lightweight threads (managed by Go, not OS threads). `go f(a, b)` starts a new goroutine which runs `f` (given `f` is a function).

```go
// just a function (which can be later started as a goroutine)
func doStuff(s string) {
}

func main() {
    // using a named function in a goroutine
    go doStuff("foobar")

    // using an anonymous inner function in a goroutine
    go func (x int) {
        // function body goes here
    }(42)
}
```

## Channels
```go
ch := make(chan int) // create a channel of type int
ch <- 42             // Send a value to the channel ch.
v := <-ch            // Receive a value from ch

// Non-buffered channels block. Read blocks when no value is available, write blocks until there is a read.

// Create a buffered channel. Writing to a buffered channels does not block if less than <buffer size> unread values have been written.
ch := make(chan int, 100)

close(ch) // closes the channel (only sender should close)

// read from channel and test if it has been closed
v, ok := <-ch

// if ok is false, channel has been closed

// Read from channel until it is closed
for i := range ch {
    fmt.Println(i)
}

// select blocks on multiple channel operations, if one unblocks, the corresponding case is executed
func doStuff(channelOut, channelIn chan int) {
    select {
    case channelOut <- 42:
        fmt.Println("We could write to channelOut!")
    case x := <- channelIn:
        fmt.Println("We could read from channelIn")
    case <-time.After(time.Second * 1):
        fmt.Println("timeout")
    }
}
```

### Channel Axioms
- A send to a nil channel blocks forever

  ```go
  var c chan string
  c <- "Hello, World!"
  // fatal error: all goroutines are asleep - deadlock!
  ```
- A receive from a nil channel blocks forever

  ```go
  var c chan string
  fmt.Println(<-c)
  // fatal error: all goroutines are asleep - deadlock!
  ```
- A send to a closed channel panics

  ```go
  var c = make(chan string, 1)
  c <- "Hello, World!"
  close(c)
  c <- "Hello, Panic!"
  // panic: send on closed channel
  ```
- A receive from a closed channel returns the zero value immediately

  ```go
  var c = make(chan int, 2)
  c <- 1
  c <- 2
  close(c)
  for i := 0; i < 3; i++ {
      fmt.Printf("%d ", <-c)
  }
  // 1 2 0
  ```

## Printing

```go
fmt.Println("Hello, 你好, नमस्ते, Привет, ᎣᏏᏲ") // basic print, plus newline
p := struct { X, Y int }{ 17, 2 }
fmt.Println( "My point:", p, "x coord=", p.X ) // print structs, ints, etc
s := fmt.Sprintln( "My point:", p, "x coord=", p.X ) // print to string variable

fmt.Printf("%d hex:%x bin:%b fp:%f sci:%e",17,17,17,17.0,17.0) // c-ish format
s2 := fmt.Sprintf( "%d %f", 17, 17.0 ) // formatted print to string variable

hellomsg := `
 "Hello" in Chinese is 你好 ('Ni Hao')
 "Hello" in Hindi is नमस्ते ('Namaste')
` // multi-line string literal, using back-tick at beginning and end
```

## Reflection
### Type Switch
A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.
```go
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```

# Snippets

## HTTP Server
```go
package main

import (
    "fmt"
    "net/http"
)

// define a type for the response
type Hello struct{}

// let that type implement the ServeHTTP method (defined in interface http.Handler)
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello!")
}

func main() {
    var h Hello
    http.ListenAndServe("localhost:4000", h)
}

// Here's the method signature of http.ServeHTTP:
// type Handler interface {
//     ServeHTTP(w http.ResponseWriter, r *http.Request)
// }
```

## string、utf8,utf16、slice
 * [文本`string`、字符`utf8,utf16`、切片`slice`](https://github.com/chai2010/advanced-go-programming-book/blob/master/ch1-basic/ch1-03-array-string-and-slice.md)

`字符串(string)`
~~~go
// 底层结构  string = []byte 即字节数组，[]byte("你好") 该转换一般不会有内存分配的开销。
type StringHeader struct {          // stringHeader is a safe version of StringHeader used
	Data uintptr                // stringHeader { Data unsafe.Pointer	Len  int }
	Len  int
}
~~~
`for range对字符串的迭代模拟实现`
~~~go
func str2bytes(s string) []byte {
	p := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		p[i] = s[i]
	}
	return p
}
~~~
`[]byte(s)转换模拟实现`
~~~go
func str2bytes(s string) []byte {
	p := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		p[i] = s[i]
	}
	return p
}
~~~
`string(bytes)转换模拟实现`
~~~go
func bytes2str(s []byte) (p string) {
	data := make([]byte, len(s))
	for i, c := range s {
		data[i] = c
	}

	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(s)

	return p
}
~~~
`[]rune(s)转换模拟实现`
~~~go
func str2runes(s string) []rune {
	var p []int32
	for len(s) > 0 {
        	r,size: = utf8.DecodeRuneInString(s)
        	p = append(p,int32(r))
        	s = s[size:]
        }
        return []rune(p)
}
~~~
`string(runes)转换模拟实现`
~~~go
func runes2string(s []int32) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return string(p)
}
~~~
`切片(slice)`
~~~go
// 底层结构
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
~~~
`添加切片元素`
~~~go
var a []int
a = append(a[:i], append([]int{x}, a[i:]...)...)     // 在第i个位置插入x
a = append(a[:i], append([]int{1,2,3}, a[i:]...)...) // 在第i个位置插入切片
a = append(a, 0)     // 切片扩展1个空间
copy(a[i+1:], a[i:]) // a[i:]向后移动1个位置
a[i] = x             // 设置新添加的元素
a = append(a, x...)       // 为x切片扩展足够的空间
copy(a[i+len(x):], a[i:]) // a[i:]向后移动len(x)个位置
copy(a[i:], x)            // 复制新添加的切片
~~~
`删除切片元素`
~~~go
a = []int{1, 2, 3}
a = a[N:]          // 删除开头N个元素
a = a[:copy(a, a[N:])] // 删除开头N个元素
a = append(a[:0], a[N:]...) // 删除开头N个元素
a = a[:len(a)-N]   // 删除尾部N个元素
a = append(a[:i], a[i+N:]...) // 删除中间N个元素
a = a[:i+copy(a[i:], a[i+N:])]  // 删除中间N个元素
~~~
`切片内存技巧`
~~~go
func Filter(s []byte, fn func(x byte) bool) []byte {
	b := s[:0]
	for _, x := range s {
		if !fn(x) {
			b = append(b, x)
		}
	}
	return b
}
var a []*int{ ... }
a = a[:len(a)-1]  // 被删除的最后一个元素依然被引用, 可能导致GC操作被阻碍
a[len(a)-1] = nil // GC回收最后一个元素内存 (保险的方式)
a = a[:len(a)-1]  // 从切片删除最后一个元素
~~~
`切片类型强制转换`
~~~go
// +build amd64 arm64

import "sort"

var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}
// 下面通过两种方法将[]float64类型的切片a转换为[]int类型的切片

// 第一种强制转换是先将切片数据的开始地址转换为一个较大的数组的指针，然后对数组指针对应的数组重新做切片操作。
// 中间需要unsafe.Pointer来连接两个不同类型的指针传递。
func SortFloat64FastV1(a []float64) {
	// 强制类型转换
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]

	// 以int方式给float64排序
	sort.Ints(b)
}
// 第二种转换操作是分别取到两个不同类型的切片头信息指针，任何类型的切片头部信息底层都是对应reflect.SliceHeader结构，
// 然后通过更新结构体方式来更新切片信息，从而实现a对应的[]float64切片到c对应的[]int类型切片的转换
func SortFloat64FastV2(a []float64) {
	// 通过 reflect.SliceHeader 更新切片头部信息实现转换
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr

	// 以int方式给float64排序
	sort.Ints(c)
}
~~~

----

