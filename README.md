# Go
Go是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。 罗伯特·格瑞史莫，罗勃·派克及肯·汤普逊于2007年9月开始设计Go，稍后Ian Lance Taylor、Russ Cox加入项目。Go是基于Inferno操作系统所开发的。

 > [官方中文文档](https://studygolang.com/pkgdoc)、[官方推荐的开源项目](https://github.com/golang/go/wiki/Projects)、[Go语言圣经](https://docs.hacknode.org/gopl-zh)、[高级编程](https://chai2010.cn/advanced-go-programming-book)
 
 > [*搭建开发环境*](#-搭建开发环境) ；[*管理.构建.测试*](#管理构建测试) ；[*推荐功能.框架.基础库.应用.工具*](#-功能框架基础库应用工具) ；[*推荐开源web框架*](#-开源的-web-框架) ；[awesome-go](https://github.com/avelino/awesome-go) <br> [*云平台.公众平台.在线支付*](#云平台公众平台在线支付) ；[*Google开源*](#Google开源) ；[*GUI - HTML/JS/CSS - WebAssembly*](#webassembly)

 * 常用于服务器编程，网络编程，分布式系统，内存数据库，云平台... [freecodecamp.org](https://guide.freecodecamp.org/go)
 * 集成工具 [JetBrains/GoLand](https://www.7down.com/search.php?word=JetBrains+GoLand&s=3944206720423274504&nsid=0)（[^搭建开发环境$](#-搭建开发环境)）、[liteide](http://liteide.org/cn/)
 * 语言优势：开发效率、抽象能力、内存管理、稳定性和功能之间的矛盾、应对需求变更、人才和资源等。

 > `开发者`
    [Gopher-China技术交流大会](https://gopherchina.org)、[搜索50万Go语言项目](https://gowalker.org)、[API+SDK'排名'服务平台](https://sdk.cn)

~~~shell
# 1.部署简单：编译成机器码(像C一样,不被反编译)复制给别人后，就能直接运行(环境免装)
#   通过<linux>命令 ldd 查看可执行文件依赖的环境(库文件)
$   ldd hello # Go不像其它语言C|C++|Java|.Net|...依赖系统环境库才能运行
# 2.静态编译语言(又像动态解释语言)
# 3.自动回收机制GC
# 4.语言层面支持高并发
# 5.丰富的第三方库,并且开源
~~~

 > 语法`关键字`

    break      default       func     interface   select
    case       defer         go       map         struct
    chan       else          goto     package     switch
    const      fallthrough   if       range       type
    continue   for           import   return      var

 > 内建的`常量`、`类型`、`函数`

    常量: true false iota nil
    
    类型: bool byte rune string error
         int int8 int16 int32 int64  uint uint8 uint16 uint32 uint64 uintptr  float32 float64  complex64 complex128
    
    函数: make len cap new append copy close delete    complex real imag    panic recover

 > 通道`chan`

 ![](http://tensor-programming.com/wp-content/uploads/2016/11/go-channel.jpg)

    读写: ch := make(chan<- int) #只读; ch := make(<-chan int) #只写; make(chan<- chan int) #只读chanOfchan;
    同步: ch := make(chan struct{}) // unbuffered channel, goroutine blocks for read or write.
    异步: ch := make(chan int, 100) // buffered channel with capacity 100 (缓存).
    管道: ch1, ch2 := make(chan int), make(chan int) ; ch1 <- 1 ; ch2 <- 2 * <-ch1; result:=<-ch2 ;
    选择: select: 常规模式(for轮循次数=chan实际数量); 反射模式(reflect.Select([]reflect.SelectCase)..);
    时间: ch := time.After(300 * time.Second) #过期chan; ch := time.Tick(1 * time.Second) #轮循chan;

 > 指针`pointer`

 ![](http://tensor-programming.com/wp-content/uploads/2016/11/Pointer.png)
 
 > 协程(超线程)`goroutine` + 上下文(传递)`context`

 ![](http://tensor-programming.com/wp-content/uploads/2016/11/gopher_pipe.png)


#### ① [搭建开发环境](https://juejin.im/book/5b0778756fb9a07aa632301e/section/5b0d466bf265da08ee7edd20)
    环境配置> go env
    安装版本> go version
    帮助文档> godoc -http=:6060  # 查看本地文档, 在线文档 golang.org/doc

> Windows - src: %GOPATH%\src - 配置 set: cd %USERPROFILE% (C:\Users\Administrator)

    https://studygolang.com/dl/golang/go1.11.10.windows-amd64.msi # 安装go1.11.10
    set GOROOT=D:\Program\Go\
    set GOPATH=C:\Users\Administrator\go
    set PATH=D:\Program\Go\bin;%GOPATH%\bin;%PATH%
    # GoLand环境设置：GOROOT, GOPATH ( √ Use GOPATH √ Index entire GOPATH?  √ Enable Go Modules[vgo go版本^1.11] )
    # go build 环境：CGO_ENABLED=1;GO_ENV=development # CGO_ENABLED=0禁用后兼容性更好;GO_ENV(dev>test>prod)
    # go tool  参数：-i -ldflags "-s -w" # -ldflags 自定义编译标记:"-s -w"去掉编译时符号+调试信息(杜绝gdb调试)+缩小file

> Linux - src: $GOPATH/src - 配置 export: cd $HOME (/root 或 /home)

    wget https://studygolang.com/dl/golang/go1.11.10.linux-amd64.tar.gz
    GO_INSTALL_DIR=/usr/local # 默认安装目录: 可更改解压到的目录 (选项 tar -C)
    tar -xvzf go1.11.10.linux-amd64.tar.gz -C $GO_INSTALL_DIR # 安装go1.11.10
    export GOPATH=~/go
    export GOROOT=/usr/local/go
    export PATH=/usr/local/go/bin:$GOPATH/bin:$PATH
    # <跨平台编译> 查看支持的操作系统和对应平台: https://github.com/fatedier/frp/blob/master/README_zh.md
    go tool dist list
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o api_linux_amd64 ./api
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./api_windows_amd64.exe ./api

> 安装依赖包
~~~bash
# 代理设置 (解决网络问题)
set http_proxy=http://127.0.0.1:5005     (临时有效) [设置环境变量linux $ export -> vim /etc/profile]
set https_proxy=http://127.0.0.1:5005    (临时有效)
set ftp_proxy=http://127.0.0.1:5005      (临时有效)
go get -d         # 下载模块源码,不安装
go get -u         # 更新模块源码
go get -v         # 打印日志
go get -insecure  # 解决安全下载问题,允许用http(非https)
go get -u github.com/gpmgo/gopm          (包管理工具) > gopm -h ; gopm bin [-d %GOPATH%/bin] [安装包名]

# 安装-全局依赖-镜像包 (解决网络问题)
git clone https://github.com/golang/arch.git %GOPATH%/src/golang.org/x/arch     # 数据结构
git clone https://github.com/golang/build.git %GOPATH%/src/golang.org/x/build   # 构建、发布
git clone https://github.com/golang/crypto.git %GOPATH%/src/golang.org/x/crypto # 加密、安全
git clone https://github.com/golang/debug.git %GOPATH%/src/golang.org/x/debug   # 调试、跟踪
git clone https://github.com/golang/exp.git %GOPATH%/src/golang.org/x/exp       # 实验和弃用的包
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
git clone https://github.com/golang/tools.git %GOPATH%/src/golang.org/x/tools   # 工具集成
git clone https://github.com/golang/tour.git %GOPATH%/src/golang.org/x/tour     # 开发文档
git clone https://github.com/googleapis/google-cloud-go.git %GOPATH%/src/cloud.google.com/go # 谷歌云

# 开发工具 VSCode Go语言工具链: pan.baidu.com/s/13tfSyd2OeSXU4lNaUfMHpA 提取码: 41jq
go get -u -v github.com/nsf/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/zmb3/gogetdoc
go get -u -v github.com/golang/lint/golint
go get -u -v github.com/lukehoban/go-outline
go get -u -v github.com/lukehoban/go-find-references
go get -u -v github.com/derekparker/delve/cmd/dlv
go get -u -v github.com/sqs/goreturns
go get -u -v github.com/tpng/gopkgs
go get -u -v github.com/golang/tools/cmd/goimports
go get -u -v github.com/golang/tools/cmd/gorename
go get -u -v github.com/golang/tools/cmd/guru
go get -u -v github.com/newhook/go-symbols
go get -u -v github.com/fatih/gomodifytags
go get -u -v github.com/cweill/gotests/...
~~~

#### 管理|构建|测试
~~~bash
# ------------------------------------------------------------------------------------
#  谷歌开源的构建和测试工具，类似于Make、Maven、Gradle.支持跨平台|语言|代码库|工具链 ✨ docs.bazel.build/versions/0.25.0/windows.html
#   /构建规则: Bazel rules for building protocol buffers +/- gRPC ✨ github.com/stackb/rules_proto
# ------------------------------------------------------------------------------------

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

# 管理模块依赖( go版本^1.11.* 推荐) & 设置GoLand环境 √ Enable Go Modules(vgo)
# 集成 vgo 项目模块管理工具 (可用环境变量 GO111MODULE 开启或关闭模块支持:off,on,auto) #默认auto未开启
git clone https://github.com/golang/vgo.git %GOPATH%/src/golang.org/x/vgo ; go install #安装vgo
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
  > go mod graph                # 4.输出依赖关系,打印模块依赖图
  > go mod verify               # 5.验证依赖是否正确
  > go get -u || -u=patch       # 5.升级到最新依赖版本 || 升级到最新的修订版本
  > go mod edit -fmt            # 5.格式化 go.mod 文件
  > go mod edit -require=path@ver # 2.添加或修改依赖版本
  > go mod download             # 2.下载依赖到%GOPATH%/pkg/mod/cache'共享缓存'
  #----------------------------------------------------------------------     (处理网络问题)
  > go mod edit -replace=google.golang.org/grpc=github.com/grpc/grpc-go@latest # 2.编辑镜像
  > go mod tidy
  > go mod vendor               # 3.拷贝依赖到./vendor/... 文件夹
  > go build -mod=vendor        # 3.构建时使用./vendor/... 文件夹
  > go build -mod=readonly      # 3.防止隐式修改go.mod
  #----------------------------------------------------------------------
  > go mod init github.com/golang/app # 6.从旧项目迁移 GO111MODULE (读取vendor/vendor.json,gopkg.toml到go.mod)
  > go mod download             # 6.下载依赖到%GOPATH%/pkg/mod/... 缓存文件夹

# 源代码版本管理
go get -d github.com/gogs/gogs  # 一款极易搭建的自助Git服务  *30k
go get -u github.com/sparrc/gdm

# 学习项目案例
go get github.com/golang/playground
go get github.com/golang/example/hello
go get github.com/shen100/golang123        # 适合初学者
go get github.com/go-training/training     # 适合初学者
go get github.com/insionng/zenpress        # 适合学习 cms system
go get github.com/muesli/cache2go          # 缓存库，代码量少，适合学习，锁、goroutines等
go get -d github.com/phachon/gis           # 图片上传，下载，存储，裁剪等
go get -d github.com/phachon/mm-wiki       # 轻量级的企业知识分享、文档管理、团队协同
go get -d github.com/getlantern/lantern    # 网络底层的东西，适合深入学习  *42k
go get -d github.com/Unknwon/the-way-to-go_ZH_CN # 中文入门教程 *2.7k  关注: Gogs, INI file, 音视频学习教程
git clone https://github.com/adonovan/gopl.io.git %GOPATH%/src/github.com/adonovan/gopl.io # programs
go get -d github.com/Yesterday17/bili-archive-frontend # *bili-bili 前端实现
go get -d github.com/angenalZZZ/Go/go-program # 获取个人代码

# 测试工具CI
  > go help test                                   # 帮助测试
  > go test -v -run=^$ ./path                      # 单元测试*testing.T [-run=查找func]
  > go test -v -bench=. -run=none -benchmem ./path # 基准测试*testing.B [-bench=.匹配全部,-run=匹配none]
  > go test -bench=^Benchmark -benchmem ./path     # 性能测试*testing.B [-bench=查找func]
  > go test -timeout 10s github.com/mpvl/errdare   # 远程测试
  > go tool vet -shadow main.go                    # 检查变量覆盖
  > go tool cover -help                            # 帮助测试覆盖率
  > go test -v -cover ./...                        # 显示测试覆盖率
  > go test -coverprofile .out ./...               # 概述测试覆盖率
  > go tool pprof -raw -seconds 30 http://localhost/debug/pprof/profile # CPU火焰图生成 go-torch -h <torch.svg>
  > go list ./...|grep -v vendor|xargs go vet -v   # 代码检查工具 go vet (静态检查,排除目录vendor)
  > go get github.com/securego/gosec/cmd/gosec/... # 安全分析工具
  > go errcheck|golint|unused|varcheck|gofmt       # 其它检测工具 go linters...
  
  # 代码质量审查 [ 1.结合github平台进行自动化的审查 https://golangci.com  |  2.本地src审查工具golangci-lint & gocritic ]
  > golangci-lint run | golangci-lint run ./... # 2.1代码运行与审查工具 github.com/golangci/golangci-lint
  > go get -v github.com/go-lintpack/lintpack/... && go get -v github.com/go-critic/go-critic/... # 2.2代码审查工具
     && lintpack build -o gocritic -linter.version='v0.3.4' -linter.name='gocritic' github.com/go-critic/go-critic/checkers
  > gocritic check-project %gopath%/src/github.com/graphql-go/graphql/  # 扫描GraphQL代码 #审查说明 gocritic check -help
  
  # 测试HTTP负载，内置HTTP服务与请求速率，包含命令行实用工具和库 > go get github.com/tsenart/vegeta
  > vegeta [global flags] <command> [command flags]
  
  # 捕获HTTP请求,跟踪HTTP流量 | https://github.com/buger/goreplay/wiki
  > gor --input-raw :80 --output-http="http://localhost:81" # 跟踪HTTP流量(:80), HTTP服务查阅结果(HTTP:81)
  > gor --input-raw :80 --output-stdout # 跟踪HTTP流量(:80)[打印输出--output-http-track-response],文件服务查阅结果gor file-server :81
  > gor --input-raw :80 --output-file=requests.gor && gor --input-file requests.gor --output-http="http://localhost:8001"

  # 集成go-test,全自动web-UI,回归测试套件,测试复盖率,代码生成器,桌面通知
  > go get github.com/smartystreets/goconvey   # 优雅的单元测试   *4k
  > go get github.com/stretchr/testify         # 通用接口调试工具 *8k | assert,require,mock,suite,http
  > go get github.com/appleboy/gofight/...     # API测试框架 beego,Gin.依赖上面的框架 github.com/stretchr/testify
  > go get github.com/astaxie/bat              # 接口调试增强curl *2k | testing, debugging, interacting servers
  > go get github.com/asciimoo/wuzz            # 用于http请求 | 交互式命令行工具 | 增强curl
  # Web性能测试命令 > bombardier -n 100 -c 100 -d 30s -l [url] # [-n:request(s),-c:connection(s),-d:duration(s)]
  > go get github.com/codesenberg/bombardier   # Web性能测试工具 | 基准测试工具 *1.5k > bombardier
  # Web基准测试工具 > bash ; $ wrk -t100 -c100 -d3s [url] | github.com/wg/wrk *20k  
  $ wrk -c 1 -t 1 -d 1 --latency [url]       # -t 线程数 -c 连接数 --timeout 超时 -d 持续时间 --latency 响应时间
  > go get github.com/tsliwowicz/go-wrk        # Web性能测试工具 *0.4k > go-wrk -help
  > go get github.com/goadapp/goad             # Web性能测试工具 *1.5k > ... make windows; goad --help
  > go get github.com/uber/go-torch            # Web性能测试与CPU火焰图生成工具 *3.5k > go-torch -h
  > go get github.com/smallnest/go-web-framework-benchmark # Web性能测试工具
  > git clone https://github.com/go-gormigrate/gormigrate.git %GOPATH%/src/gopkg.in/gormigrate.v1 && go get gopkg.in/gormigrate.v1

# 发布自动化CD
  # 工具 goreleaser.com | github.com/goreleaser/goreleaser

~~~

> 性能优化
~~~
# ------------------------------------------------------------------------------------
# 通过工具排查：
# ------------------------------------------------------------------------------------
go get github.com/google/pprof # 用于可视化和分析性能和数据的工具pprof(CPU/rofile)
go tool pprof -seconds 5 http://localhost/debug/pprof/profile # 导入 _ net/http/pprof 添加HTTP性能优化服务 /debug/pprof
go tool pprof -alloc_objects -inuse_objects   # 生成对象数量、引用对象数量
go test . -bench . -benchtime 3s -cpuprofile prof.cpu -memprofile # 功能测试与性能分析(如testing.B测试benchmark)
go tool pprof [stats.test] prof.cpu # 详细的单元分析: 对象|代码行|函数调用|runtime|package|binary [stats目录/.test*测试]
go tool pprof -http=":8081" [binary] [profile] # GC对象扫描,函数占据大量CPU(如runtime.scanobject等问题分析)
# ------------------------------------------------------------------------------------
go get github.com/uber/go-torch # Web性能测试与CPU火焰图生成工具 > go-torch -h
go tool pprof -raw -seconds 30 http://localhost/debug/pprof/profile # torch.svg
go get github.com/prashantv/go_profiling_talk # 案例剖析:如何使用pprof和go-torch识别性能瓶颈，并进行优化? 视频youtu.be/N3PWzBeLX2M

# ------------------------------------------------------------------------------------
# 内存管理`GC`的优化：
# ------------------------------------------------------------------------------------
 # 对象数量过多时(引用传递过多时)，导致GC三色算法耗费较多CPU（可利用耗费少量的内存，优化耗费的CPU）
map[string]NewStruct -> map[[32]byte]NewStruct  # key使用值类型避免对map遍历
map[int]*NewStruct   -> map[int]NewStruct       # val使用值类型避免对map遍历
someSlice []float64  -> someSlice [32]float64   # 利用值类型代替对象类型

# ------------------------------------------------------------------------------------
# 扩容(横向|纵向)：
# ------------------------------------------------------------------------------------
 # 分片Sharding > 如何集群? 把数据划分成若干部分,1个部分映射1个Shard(内存中分配完成);把Shard分配到服务器节点上;节点node+副本replica
 # 策略 > 如何分片? <空间索引>把数据按空间范围划分成若干个最小单元Cell;按规则算法把部分单元Cells放入1个Shard分片;Cell队列中的数据可查找所在Shard/Cell;数据清理Clean
~~~

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

#### ② [功能、框架、基础库、应用、工具](https://github.com/avelino/awesome-go)

~~~
go get github.com/rs/xid                   # uuid shortuuid Snowflake MongoID xid
go get github.com/google/uuid              # 基于RFC4122和DCE1.1身份验证和安全服务，生成、检查Uuid
go get github.com/satori/go.uuid           # uuid generator, Version 1 ~ 5 (RFC 4122)
go get github.com/reactivex/rxgo           # 响应式编程rxgo
go get github.com/google/btree             # 数据结构 B-Trees
go get github.com/google/trillian          # 数据结构 Merkle tree, Verifiable Data Structures 
go get github.com/google/go-intervals/...  # 在一维间隔（例如时间范围）上执行预定的操作
go get github.com/juju/utils               # Utility functions
go get github.com/henrylee2cn/goutil       # Common and useful utils
go get github.com/appleboy/com             # Random、Array、File、Convert
go get gopkg.in/pipe.v2                    # Pipeline IO | github.com/go-pipe/pipe
go get github.com/google/go-querystring/query # 转换对象，用于URL参数
go get github.com/google/jsonapi           # 转换对象，用于HTTP请求的输入输出
go get github.com/rsc/goversion            # 扫描目录中Go可执行文件的版本信息 > goversion /usr/bin

go get github.com/vipally/binary           # binary编码/解码 data和[]byte的互转(encoding/gob,encoding/binary)
go get github.com/json-iterator/go         # json编码/解码的性能优化，替换原生(encoding/json)   *5k
go get github.com/xeipuuv/gojsonschema     # 元模式验证，json schema 自定义错误校验库           *1k
go get github.com/TheAlgorithms/Go         # 各种算法的实现 github.com/TheAlgorithms/Python   *31k
go get github.com/PuerkitoBio/goquery      # 解析HTML，像jQuery那样操作DOM                     *7k
go get github.com/sirupsen/logrus          # 日志跟踪 import log "github.com/sirupsen/logrus" *10k
go get github.com/asticode/go-astilog      # 日志自定义:logrus[.Entry-Data]+context[.Value-Fields]+configuration
go get github.com/pkg/errors               # 错误处理pkg
go get github.com/mpvl/{errc,errd}         # 错误处理mpvl

go get github.com/alecthomas/participle    # 超简单的Lexer解析器Parser(推荐使用,Lexer性能高于`反射`) *1.5k
go get github.com/blynn/nex                # 好用的Lexer解析器工具，生成go代码&YACC/Bison&正则表达式: nex -r -s lc.nex
go get github.com/antlr/antlr4/runtime/Go/antlr # 语言识别工具，强大的Parser生成器，读取、处理、执行或翻译文本或二进制文件 | www.antlr.org

# https://github.com/etcd-io/etcdlabs      # 分布式可靠键值存储，适用于分布式系统中最关键的数据；提供分享配置和服务发现
# client: http://play.etcd.io              # 数据中心 etcd | 下载 github.com/etcd-io/etcd/releases
go get github.com/hashicorp/serf/cmd/serf  # 数据中心 serf | 基于Gossip的Membership,P2P对等网络\去中心 | www.serf.io
go get github.com/spf13/viper && go get github.com/spf13/pflag # 配置(JSON,TOML,YAML,HCL)热加载;远程配置;缓存;加密
go get github.com/xordataexchange/crypt/bin/crypt 加密存储 secret keyring: gpg(gpg4win)用于安全传输(类似rsa)
go get github.com/minio/minio-go           # 云存储|分布式存储SDK|网盘|OSS | www.min.io  docs.min.io/cn
go get -d github.com/minio/mc              # 云存储|配置客户端, 指南 | docs.min.io/cn/minio-client-quickstart-guide.html
go get -d github.com/minio/minio           # 云存储|配置服务端, 运行: hidec /w minio.exe server d:\docker\app\minio\data
go get github.com/perkeep/perkeep/cmd/...  # Camlistore 个人存储系统：一种存储、同步、共享、建模和备份内容的方式
go get -d github.com/s3git/s3git           # 云存储的Git: 用于数据的分布式版本控制系统
go get github.com/coocood/freecache        # cache and high concurrent performance
go get github.com/patrickmn/go-cache       # in-memory key:value store/cache (similar to Memcached)适用于单台应用程序
go get github.com/peterbourgon/diskv       # 支持磁盘的 key-value 存储
go get github.com/chrislusf/seaweedfs/weed # 一个用于小文件的简单且高度可扩展的分布式文件系统，可集成其他云服务，如AWS...
go get github.com/fsnotify/fsnotify        # 文件系统监控 # go get golang.org/x/sys/...
go get github.com/rjeczalik/notify         # 文件系统事件通知库
go get github.com/mholt/archiver/cmd/arc   # 压缩文件/解压工具(zip,tar,rar) arc archive [archive name] [input files...]
go get github.com/cloudflare/cfssl/cmd/... # SSL证书 usage play.etcd.io/install#TLS  *4k
go get github.com/muesli/beehive           # 灵活的事件/代理/自动化系统  *3k

go get github.com/go-redis/redis           # 内存数据库,类型安全的Redis-client *6k (推荐使用,性能高于redigo)
go get github.com/gomodule/redigo/redis    # 内存数据库,使用原生的Redis-cli *6k
go get github.com/sent-hil/bitesized        # Redis位图计数> 统计分析、实时计算
go get github.com/yannh/redis-dump-go       # Redis导出导入> redis-dump-go -h ; redis-cli --pipe < backup.resp;redis-dump
go get github.com/syndtr/goleveldb/leveldb # 内存数据库,谷歌leveldb-client
go get github.com/seefan/gossdb/example    # 内存数据库,替代Redis的ssdb | ssdb.io/zh_cn

go get github.com/dgraph-io/badger/...     # 高性能 key/value 数据库,支持事务,(强力推荐)LSM+tree,ACID,Stream,KV+version,SSDs
go get github.com/dgraph-io/dgraph/dgraph  # 高性能,具有可扩展、分布式、低延迟和高吞吐量功能的分布式位图索引数据库 *10k
go get github.com/boltdb/bolt/...          # 高性能 key/value 数据库,支持事务,B+tree,ACID,分桶 *10k | 性能低于badger
go get github.com/tidwall/buntdb           # 内存数据库,BuntDB is a low-level, in-memory, key/value store, persists to disk
go get github.com/tidwall/buntdb-benchmark # 性能测试 > buntdb-benchmark -n 10000 -q # 单机时超越Redis，有索引和geospatial功能
go get github.com/allegro/bigcache         # 高可用千兆级数据的高效 key/value 缓存   *2k
go get github.com/cockroachdb/cockroach    # 云数据存储系统，支持地理位置、事务等 *20k | www.cockroachlabs.com/docs/stable
go get -d github.com/tidwall/tile38        # 具有空间索引和实时地理位置数据库,如PostGIS *7k > docker run -p 9851:9851 tile38/tile38
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
git clone https://github.com/rana/ora.git %GOPATH%/src/gopkg.in/rana/ora.v4 && go get gopkg.in/rana/ora.v4
go get github.com/mattn/go-oci8            # Oracle env: instantclient & MinGW-w64-gcc & pkgconfig/oci8.pc
go get github.com/go-sql-driver/mysql      # Mysql     | github.com/siddontang/go-mysql
go get github.com/denisenkom/go-mssqldb    # MsSql
go get github.com/lib/pq                   # Postgres  | github.com/prest/prest
go get github.com/jackc/pgx                # Postgres  驱动与工具集
go get github.com/sosedoff/pgweb           # Postgres  Web管理系统
go get github.com/mattn/go-sqlite3         # SQLite
go get github.com/jmoiron/sqlx             # 数据库sql    *6k  extensions go's standard database/sql library
  go get github.com/heetch/sqalx             # sqlx & sqalx 支持嵌套的事务
  go get github.com/twiglab/sqlt             # sqlx & sqlt 模板拼接sql和java的数据库访问工具MyBatis的sql配置
  go get github.com/albert-widi/sqlt         # sqlx & sqlt 支持数据库主从数据源，读写分离
go get github.com/go-xorm/xorm             # 数据库xorm   *5k  support mysql,postgres,tidb,sqlite3,mssql,oracle
  go get github.com/go-xorm/builder          # ^xorm SQL Builder 增强-拼接sql
  go get github.com/xormplus/xorm            # ^xorm增强版*$ 支持sql模板,动态sql,嵌套事务,类ibatis配置等
                                             # ^xorm增强版*文档 https://www.kancloud.cn/xormplus/xorm/167077
go get github.com/didi/gendry              # 滴滴开源 SQL Builder 增强-拼接sql、连接池管理、结构映射.
go get github.com/mattes/migrate           # 数据库迁移工具 *2k
go get github.com/rubenv/sql-migrate/...   # 数据库 schema 迁移工具，允许使用 go-bindata 将迁移嵌入到应用程序中 *1k
git clone https://github.com/go-gormigrate/gormigrate.git %GOPATH%/src/gopkg.in/gormigrate.v1 && go get gopkg.in/gormigrate.v1 # gorm migrate
go get github.com/gchaincl/dotsql          # 帮助你将 sql 文件保存至某个地方并轻松使用它
go get github.com/xo/xo                    # 命令行工具 xo --help  [DbFirst]生成 models/*.xo.go
   > cp %GOPATH%/src/github.com/xo/xo/templates/* ./templates
   > xo mysql://root:123456@127.0.0.1:3306/AppAuth?parseTime=true -o ./models [--template-path templates]
   > xo mssql://sa:123456@localhost:1433/AppAuth?parseTime=true -o ./models [--template-path templates]
go get github.com/go-xorm/cmd/xorm         # 命令行工具 xorm help  [DbFirst]生成 models/*.go
   > cp %GOPATH%/src/github.com/go-xorm/cmd/xorm/templates/goxorm/* ./templates
   > xorm reverse mysql root:123456@tcp(127.0.0.1:3306)/AppAuth?charset=utf8 ./templates ./models [^表名前缀]
   > xorm reverse mssql "server=localhost;user id=sa;password=HGJ766GR767FKJU0;database=AppAuth" %GOPATH%/src/github.com/go-xorm/cmd/xorm/templates/goxorm ./models [^表名前缀]
go get github.com/variadico/scaneo         # 命令行工具 scaneo -h  [DbFirst]生成 models/*.go

go get github.com/blevesearch/bleve        # 现代文本索引库 *5k
go get github.com/olivere/elastic          # Elasticsearch 6.0 客户端
go get github.com/Qihoo360/poseidon        # 360开源|日志搜索平台是搜索引擎的核心技术+索引数据Hadoop
go get github.com/siesta/neo4j             # Neo4j 客户端 | github.com/jmcvetta/neoism
go get github.com/cayleygraph/cayley       # 图形数据库 Driven & RESTful API & LevelDB Stores
go get github.com/DarthSim/imgproxy        # Fast image server: docker pull darthsim/imgproxy
go get willnorris.com/go/imageproxy/...    # Caching image proxy server & docker & nginx
go get labix.org/v2/mgo                    # MongoDB 驱动:集群,并发,一致性,Auth,GridFS | labix.org/mgo
git clone https://github.com/mongodb/mongo-go-driver.git %GOPATH%/src/github.com/mongodb/mongo-go-driver 
  go get github.com/go-stack/stack
  go get github.com/golang/snappy
  go get github.com/google/go-cmp
  go get github.com/montanaflynn/stats
  go get github.com/tidwall/pretty
  dep ensure -add "go.mongodb.org/mongo-driver/mongo@~1.0.0"
git clone https://github.com/jmcvetta/neoism.git %GOPATH%/src/gopkg.in/jmcvetta/neoism.v1 && go get gopkg.in/jmcvetta/neoism.v1

go get github.com/robfig/cron              # 任务计划 a cron library *4k
go get github.com/iamduo/go-workq          # job server and client  *1k
go get github.com/jasonlvhit/gocron        # simple Job Scheduling  *1k
go get github.com/gocraft/work             # do work of redis-queue *1k | github.com/gocraft/work#run-the-web-ui
go get github.com/lisijie/webcron          # 定时任务Web管理器 (基于beego框架) *1k
go get github.com/shunfei/cronsun          # 分布式容错任务管理系统 *1.5k
go get github.com/travisjeffery/jocko      # 消息推送服务Kafka *3k : producing/consuming[生产/消费] cluster[代理集群]
go get github.com/appleboy/gorush          # 消息推送服务gorush 4k : api, HTTP/2, notification queue and multiple workers.
go get github.com/gocelery/gocelery        # 分布式任务队列Celery *1k : client/server www.celeryproject.org
go get github.com/libp2p/go-libp2p         # 网络库模块p2p-serves
go get github.com/libp2p/go-libp2p-examples# 网络库模块p2p-examples

go get github.com/nsqio/nsq                # 实时分布式消息平台nsq *15k | nsqlookupd & nsqd & nsqadmin https://nsq.io
go get github.com/youzan/nsq               # 有赞科技的nsq重塑 | www.bilibili.com/video/av29142217
go get github.com/streadway/amqp           # rabbitmq client tutorials | www.rabbitmq.com/#getstarted
go get github.com/blackbeans/kiteq         # KiteQ 是一个基于 go + protobuff 实现的多种持久化方案的 mq 框架
go get -d github.com/emqx/emqx             # 百万级分布式开源物联网MQTT消息服务器 *4k | www.emqtt.com
go get github.com/mattermost/mattermost-server # 通讯 *15k 为团队带来跨PC和移动设备的消息、文件分享，提供归档和搜索功能+前端React
go get github.com/gorilla/websocket        # WebSocket Serve *8k
go get github.com/gotify/server            # WebSocket Serve (Includes Web-UI manage) | gotify.net
go get github.com/gotify/cli               # WebSocket client to push messages
# 聊天室 git clone https://github.com/GoBelieveIO/im_service.git && cd im_service && dep ensure && mkdir bin && make install
# 高并发 go get github.com/xiaojiaqi/10billionhongbaos  # 抢购系统：单机支持QPS达6万，可以满足100亿红包的压力测试
# https://github.com/oikomi/FishChatServer2
# 后端框架
go get github.com/gin-gonic/gin            # 后端WebSvr *26k: Gin Web Framework
go get github.com/astaxie/beego            # 后端WebSvr *20k: API、Web、服务 | 高度解耦的框架 | beego.me/docs/intro
# 基础模块：cache,config,context,httplibs,logs,orm,session,toolbox,plugins... 管理工具bee | github.com/beego/bee
go get github.com/kataras/iris             # 最快Web开发框架 *15k | github.com/kataras/iris/blob/master/README_ZH.md
# 入门程序：[iris+xorm]github.com/yz124/superstar [gorm+jwt]github.com/snowlyg/IrisApiProject [postgres+angular]github.com/iris-contrib/parrot
go get github.com/mholt/caddy/caddy        # 全栈Web服务平台 *21k | 配置超快apache+nginx | caddyserver.com
go get github.com/revel/cmd/revel          # 高生产率的全栈web框架 *11k > revel new -a my-app -r | github.com/revel/revel
go get github.com/gorilla/{mux,sessions,schema,csrf,handlers,websocket} # 后端框架mux工具链 *9k
go get github.com/graphql-go/graphql       # Facebook开源API查询语言 *5k | GraphQL中文网™ graphql.org.cn
go get github.com/graph-gophers/graphql-go # GraphQL api server   *3k
go get github.com/dgrijalva/jwt-go/cmd/jwt # JSON Web Tokens (JWT)
go get github.com/appleboy/gin-jwt         # JWT Middleware for Gin
go get golang.org/x/oauth2                 # OAuth 2.0 认证授权 | github.com/golang/oauth2
go get github.com/casbin/casbin            # 授权访问-认证服务 ACL, RBAC, ABAC | casbin.org
go get github.com/ory/fosite/...           # 访问控制 OAuth2.0, OpenID Connect SDK | www.ory.sh
go get github.com/gorilla/sessions         # session & cookie authentication
go get github.com/kgretzky/evilginx2       # session cookies, allowing for the bypass of 2-factor authentication
go get github.com/dpapathanasiou/go-recaptcha # Google验证码|申请 | www.google.com/recaptcha/admin/create
go get github.com/dchest/captcha           # 验证码|图片|声音
go get github.com/mojocn/base64Captcha     # 验证码|展示 | captcha.mojotv.cn
go get github.com/emersion/go-imap/...     # 邮箱服务 IMAP library for clients and servers
go get github.com/sdwolfe32/trumail/...    # 邮箱验证 clients
go get github.com/matcornic/hermes/v2      # HTML e-mails, like: npm i mailgen | github.com/eladnava/mailgen
go get github.com/fagongzi/gateway         # 基于HTTP协议的restful的API网关, 可以作为统一的API接入层
go get github.com/wanghongfei/gogate       # 高性能Spring Cloud网关, 路由配置热更新、负载均衡、灰度、服务粒度的流量控制、服务粒度的流量统计
# 后端API文档生成器 swagger
go get github.com/go-swagger/go-swagger/cmd/swagger # 使用 > swagger generate spec --scan-models -o docs/spec.json

go get github.com/gocolly/colly/...        # 高性能Web采集利器 *7k
go get github.com/henrylee2cn/pholcus      # 重量级爬虫软件    *5k
go get github.com/tealeg/xlsx              # 读取 Excel 文件  *3.2k
go get github.com/360EntSecGroup-Skylar/excelize/v2 # 读写 Excel 文件 *3.8k
go get github.com/davyxu/tabtoy            # 高性能便捷电子表格导出器   *1k
go get -u github.com/skip2/go-qrcode/...   # 二维码生成器 > qrcode     *1k
go get github.com/claudiodangelis/qr-filetransfer # 二维码识别|转换    *3k
go get github.com/jung-kurt/gofpdf         # 生成 PDF 文件  *2.8k | 支持text,drawing,images
go get github.com/gorilla/websocket        # WebSocket | github.com/joewalnes/websocketd websocketd.com
go get github.com/gobwas/ws                # WebSocket | github.com/socketio/socket.io
go get github.com/rakyll/statik            # 将静态资源文件嵌入到Go二进制文件中，然后提供http服务: statik -src=/path/to
go get github.com/yudai/gotty              # 终端扩展为Web网站服务 *12.3k
go get -d -u gobot.io/x/gobot/...          # 物联网IoT、物理计算Drones、机器人Robotics | github.com/hybridgroup/gobot

# 微服务(分布式RPC框架)rpcx，支持Zookepper、etcd、consul服务发现&路由 *3k | books.studygolang.com/go-rpc-programming-guide
go get -u -v -tags "reuseport quic kcp zookeeper etcd consul ping rudp utp" github.com/smallnest/rpcx/...
# 谷歌开源gRPC | grpc.io/docs/quickstart/go & 'HTTP/2'更快 http2.golang.org
 # 1.安装: gRPC、genproto ; <protoc.exe>插件: proto、protoc-gen-go、protoc-gen-gogo、protoc-gen-gofast;prototool(增强protoc)
 > github.com/google/protobuf/releases     # 先下载Protobuf | <protoc.exe>
 > git clone https://github.com/grpc/grpc-go.git %GOPATH%/src/google.golang.org/grpc
 > git clone https://github.com/google/go-genproto %GOPATH%/src/google.golang.org/genproto
 > go get github.com/golang/{text,net}                                       # 安装protoc的依赖
 > go get github.com/golang/protobuf/{proto,protoc-gen-go}                   # 生成工具protoc-gen-go
   $ protoc --go_out=. *.proto                                               # segmentfault.com/a/1190000009277748
 > go get github.com/gogo/protobuf/{proto,protoc-gen-gogo,protoc-gen-gofast} # 推荐gofast性能高于protoc-gen-go
   $ protoc --gogo_out=plugins=grpc:. *.proto            # protoc --gofast_out=plugins=grpc:. *.proto (输入输出:当前目录)
   # ⚡ gRPC-Gateway (gRPC to JSON proxy: 接口Demo) + swagger + validators ↓
   > git clone https://github.com/gogo/grpc-example.git && set GO111MODULE=on && go build -mod=vendor && grpc-example.exe
  $ prototool help                         # 增强版protoc <ubuntu> github.com/uber/prototool
 # 2.使用: gRPC-Examples > cd %GOPATH%/src/google.golang.org/grpc/examples/helloworld
 > protoc -I ./helloworld --go_out=plugins=grpc:./helloworld ./helloworld/helloworld.proto #2.1生成代码*.pb.go
 > go run ./greeter_server/main.go ; go run ./greeter_client/main.go                       #2.2启动服务&客户端
go get github.com/grpc-ecosystem/grpc-gateway/... # 谷歌开源网关(gRPC to JSON proxy: 读取protobuf,生成反向代理)
go get github.com/grpc-ecosystem/go-grpc-middleware #auth,logrus,prometheus⚡,opentracing,validator,recovery,ratelimit;retry
go get github.com/istio/istio              # 谷歌开源|微服务集群管理k8s  *17k | istio.io | www.grpc.io
go get github.com/go-kit/kit/cmd/kitgen    # 阿里推荐|微服务构建框架gRPC *13k | gokit.io
go get github.com/apache/thrift/lib/go/thrift/... # 滴滴推荐|微服务框架  *7k  (ASF licenses)
go get github.com/TarsCloud/TarsGo/tars    # 腾讯开源|基于Tars协议的高性能RPC框架 *2k | 网关+容器化+服务治理等
go get github.com/micro/go-micro           # 分布式RPC微服务 *7k
go get -u -v github.com/davyxu/cellnet     # 游戏服务器RPC *2.5k | ARM设备<设备间网络通讯> | 证券软件<内部RPC>
go get -u -v github.com/liangdas/mqant     # 游戏服务器RPC *1.5k
git clone https://github.com/EasyDarwin/EasyDarwin.git %GOPATH%/src/github.com/EasyDarwin/EasyDarwin # RTSP流媒体服务
go get github.com/iikira/BaiduPCS-Go       # 百度网盘命令行客户端
go get github.com/inconshreveable/go-update # 自动更新应用程序
go get -d https://github.com/restic/restic  # 数据备份工具 | restic.readthedocs.io
cd %GOPATH%/src/github.com/restic/restic && go run -mod=vendor build.go --goos windows --goarch amd64
# ------------------------------------------------------------------------------------
# 部署-维护
# ------------------------------------------------------------------------------------
go get github.com/google/gousb             # 用于访问USB设备的低级别接口
go get github.com/google/gops              # 用于列出并诊断Go应用程序进程
go get github.com/google/pprof             # 用于可视化和分析性能和数据的工具
go get github.com/google/mtail             # 用于从应用程序日志中提取白盒监视数据，以便收集到时间序列数据库中
go get github.com/google/godepq            # 用于查询程序依赖 > godepq -from github.com/google/pprof
go get github.com/google/ko/cmd/ko         # 用于构建和部署应用程序到Kubernetes的工具
go get github.com/google/git-appraise/git-appraise # 用于Git版本管理的分布式代码审核
go get github.com/google/easypki/cmd/easypki # CA证书申请工具 | API: go get gopkg.in/google/easypki.v1
go get -u github.com/uber/jaeger-client-go/  # CNCF Jaeger，分布式跟踪系统 | github.com/jaegertracing/jaeger
go get github.com/codegangsta/gin          # 站点热启动 > gin -h
go get github.com/fvbock/endless           # 站点零停机\重启
go get github.com/ochinchina/supervisord   # 开机启动supervisor > supervisord -c website.conf -d
go get github.com/sourcegraph/checkup/cmd/checkup # 分布式站点健康检查工具 > checkup --help
go get go.universe.tf/tcpproxy/cmd/tlsrouter # TLS代理根据握手的SNI（服务器名称指示）将连接路由到后端。它不携带加密密钥，无法解码其代理的流量
go get github.com/prometheus/prometheus/cmd/... # 服务监控系统和时间序列数据库 *23k | prometheus.io/community
go get github.com/grafana/grafana          # 监测系统|指标分析工具|InfluxDB时序DB|Prometheus等 *29k

go get github.com/elves/elvish             # shell for unix > 可编程：数组、字典、传递对象的增强型管道、闭包、模块机制、类型检查
go get github.com/mattn/sudo               # sudo for windows > sudo cmd /c dir ; sudo notepad c:\windows\system32\drivers\etc\hosts
go get github.com/lxn/win                  # Windows API wrapper package
go get github.com/lxn/walk                 # Windows UI Application Library Kit *3k
go get github.com/google/gapid             # Windows UI App : Graphics API Debugger
go get github.com/FiloSottile/mkcert       # 证书管理工具 *18k
# [申请Let's Encrypt永久免费SSL证书]          | www.jianshu.com/p/3ae2f024c291
go get github.com/go-acme/lego/cmd/lego    # Let's Encrypt client and ACME library, DNS providers manager.
# [QT跨平台应用框架] Qt binding package
go get -u -v github.com/therecipe/qt/cmd/... && for /f %v in ('go env GOPATH') do %v\bin\qtsetup test && %v\bin\qtsetup
# [Bringing Flutter to Windows, MacOS and Linux] - through the power of Go and GLFW.
# https://github.com/go-flutter-desktop/go-flutter
go get github.com/BurntSushi/wingo/wingo-cmd # 一个功能齐全的窗口管理器 > wingo-cmd

# 小米公司的互联网企业级监控系统 | book.open-falcon.org
# 各大 Go 模板引擎的对比及压力测试 | github.com/SlinSo/goTemplateBenchmark

~~~

----

#### 云平台|公众平台|在线支付
~~~
# 云计算
# ------------------------------------------------------------------------------------
# 亚马逊 AWS | www.amazonaws.cn/tools

# 谷歌云 Google Cloud Platform | cloud.google.com/go | github.com/GoogleCloudPlatform
go get -u github.com/google/go-cloud      # 云计算
go get -u cloud.google.com/go/storage     # 在 Cloud Storage 中存储和归档数据
go get -u cloud.google.com/go/bigquery    # 使用 Google BigQuery 执行数据分析
go get -u cloud.google.com/go/pubsub      # 使用 Pub/Sub 设置完全托管的事件驱动型消息传递系统
go get -u cloud.google.com/go/translate   # 使用 Translation API 翻译不同语言的文本
go get -u cloud.google.com/go/vision/apiv1# 使用 Vision API 分析图片

# 阿里云 | api.aliyun.com
go get -u github.com/aliyun/alibaba-cloud-sdk-go/sdk
# 云服务器 ECS、对象存储 OSS、阿里云关系型数据库、云数据库MongoDB版、CDN、VPC、
# 视频点播、音视频通信、媒体转码、负载均衡、云监控、容器服务、邮件推送、弹性伸缩、移动推送、日志服务、交易与账单管理

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
# Demo : video tools
go get github.com/asticode/go-astivid/...
cd %GOPATH%/src/github.com/asticode/go-astivid
cp -r %GOPATH%/bin/astibundler/* C:/Users/ADMINI~1/AppData/Local/Temp/astibundler/cache
rm -f bind*.go                # delete file before bundle
astilectron-bundler -v        # help: astilectron-bundler -h
~~~
 * [QT](https://github.com/therecipe/qt)
 * [Webview](https://github.com/zserge/webview)
 * [WebAssembly](https://github.com/murlokswarm/app)

----

#### ③ [开源的 Web 框架](https://github.com/avelino/awesome-go#web-frameworks)

 * Web 框架
    * [基于 Gin 构建企业级 RESTful API 服务](https://juejin.im/book/5b0778756fb9a07aa632301e)
    * [基于 Gin 一步一步搭建Go的Web服务器](https://www.hulunhao.com/go/go-web-backend-starter/)
~~~
# 开发
cd %GOPATH%/src                                                                 # 项目框架 Gin Web Framework
git clone https://github.com/lexkong/apiserver_demos apiserver                  # 项目源码-复制^demo至-工作目录
git clone https://github.com/lexkong/vendor                                     # 项目依赖-govendor
go get github.com/StackExchange/wmi                                             # 项目依赖-缺失的包
# 构建
cd %GOPATH%/src/apiserver && go fmt -w . && go tool vet . && go build -v -o [应用名] [目录默认.]
# 运行
%GOPATH%/src/apiserver/apiserver.exe
~~~

----

#### ④ [中文标准库文档](https://studygolang.com/pkgdoc)

----

#### ⑤ 阅读相关文章

 * 高性能
    * [高并发架构解决方案](https://studygolang.com/articles/15479)

----

