# Go
Go是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。 罗伯特·格瑞史莫，罗勃·派克及肯·汤普逊于2007年9月开始设计Go，稍后Ian Lance Taylor、Russ Cox加入项目。Go是基于Inferno操作系统所开发的。

 > [应用&库&工具](https://github.com/avelino/awesome-go)、[官方中文文档](https://studygolang.com/pkgdoc)、[官方推荐的开源项目](https://github.com/golang/go/wiki/Projects)、[Go语言圣经](https://docs.hacknode.org/gopl-zh)、[高级编程](https://chai2010.cn/advanced-go-programming-book)、[^收藏夹$](#-功能框架基础库应用工具)

 * 常用于服务器编程，网络编程，分布式系统，内存数据库，云平台...
 * 集成工具 [JetBrains/GoLand](https://www.7down.com/search.php?word=JetBrains+GoLand&s=3944206720423274504&nsid=0)（[^搭建开发环境$](#-搭建开发环境)）、[liteide](http://liteide.org/cn/)

 > `下载` [Go_programming_lang.part1](https://rapidgator.net/file/e8ca89d3d3fbfceefb198469dd63ea24/Golang_build_RESTful_APIs_with_Golang_(Go_programming_lang)-DEC18.part1.rar.html)、[Go_programming_lang.part2](https://rapidgator.net/file/841b5337f413a161c874f0e1b57755ff/Golang_build_RESTful_APIs_with_Golang_(Go_programming_lang)-DEC18.part2.rar.html)

~~~shell
# 1.部署简单：编译成机器码(像C一样,不被反编译)复制给别人后，就能直接运行(环境免装)
#   通过<linux>命令 ldd 查看可执行文件依赖的环境(库文件)
$   ldd hello # Go不像其它语言C|C++|Java|.Net|...依赖系统环境库才能运行
# 2.静态编译语言(又像动态解释语言)
# 3.自动回收机制GC
# 4.语言层面支持高并发
# 5.丰富的第三方库,并且开源
~~~

 > `关键字`

    break      default       func     interface   select
    case       defer         go       map         struct
    chan       else          goto     package     switch
    const      fallthrough   if       range       type
    continue   for           import   return      var

 > 内建的`常量`、`类型`、`函数`

    常量: true false iota nil

    类型: bool byte rune string error
          int int8 int16 int32 int64   uint uint8 uint16 uint32 uint64 uintptr   float32 float64  complex64 complex128

    函数: make len cap new append copy close delete    complex real imag    panic recover

 > 通道`chan`

    同步: ch := make(chan struct{}) // unbuffered channel, goroutine blocks for read or write # make(chan struct{}, 0) 
    异步: ch := make(chan int, 100) // buffered channel with capacity 100
    管道: ch1, ch2 := make(chan int), make(chan int) // 即-串连的通道-读写; ch1 <- 1; ch2 <- 2 * <-ch1; result := <-ch2

#### ① [搭建开发环境](https://juejin.im/book/5b0778756fb9a07aa632301e/section/5b0d466bf265da08ee7edd20)
    安装版本> go version
    环境配置> go env

> Windows - src: %GOPATH%\src - 配置 set: cd %USERPROFILE% (C:\Users\Administrator)

    https://studygolang.com/dl/golang/go1.12.windows-amd64.msi
    GOROOT=D:\Program\Go\
    GOPATH=C:\Users\Administrator\go
    PATH=D:\Program\Go\bin;%GOPATH%\bin;%PATH%
    # go tool vet -shadow main.go # 检查变量覆盖问题
    > go get -d        # Download the packages source, not to install.
    > go get -u        # Update the named packages and their dependencies.
    > go get -v        # Verbose progress and debug output.
    > go get -insecure # Resolving domains using insecure HTTP(No https).
    > go get github.com/angenalZZZ/Go/go-program # 获取代码

~~~bash
  # GoLand *全局：GOROOT, GOPATH ( √ Use GOPATH √ Index entire GOPATH? )
   # go build环境：CGO_ENABLED=1;GO_ENV=development # CGO_ENABLED=0禁用后兼容性更好;GO_ENV(development>test>production) 
   # go tool 参数：-i -ldflags "-s -w" # -ldflags 自定义编译标记:"-s -w"去掉编译时的符号&调试信息(不能gdb调试),缩小文件大小
  go list -json     # 列举当前目录（包|模块|项目）的依赖导入、源码、输出等。
  go list -m -u all # 列举依赖模块和依赖更新
  # 管理项目模块 go mod <command> [arguments] (模块的增删改+下载) | 模块功能概述 go help modules
  go help mod       # 查看说明
~~~

> Linux - src: $GOPATH/src - 配置 export: cd $HOME (/root 或 /home)
    
    wget https://studygolang.com/dl/golang/go1.12.linux-amd64.tar.gz
    GO_INSTALL_DIR=/usr/local # 默认安装目录: 可更改 (选项 tar -C)
    tar -xvzf go1.12.linux-amd64.tar.gz -C $GO_INSTALL_DIR
    GOROOT=/usr/local/go
    GOPATH=/home/go
    PATH=/usr/local/go/bin:$GOPATH/bin:$PATH
    # <跨平台编译> 查看支持的操作系统和对应平台: https://github.com/fatedier/frp/blob/master/README_zh.md
    $ go tool dist list
    $ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o api_linux_amd64 ./api
    $ CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./api_windows_amd64.exe ./api

> 安装依赖包
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
git clone https://github.com/golang/tools.git %GOPATH%/src/golang.org/x/tools   # 工具文档
git clone https://github.com/golang/tour.git %GOPATH%/src/golang.org/x/tour     # 开发文档
git clone https://github.com/googleapis/google-cloud-go.git %GOPATH%/src/cloud.google.com/go # 谷歌云

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
go get -u -v github.com/derekparker/delve/cmd/dlv

# 管理项目依赖包
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
  > govendor add +e           # 添加本地$GOPATH包[go get]
  > govendor fetch            # 获取远程vendor.json包[govendor get]
# vgo 一个项目模块管理工具 (用环境变量 GO111MODULE 开启或关闭模块支持:off,on,auto) # [默认auto]
git clone https://github.com/golang/vgo.git %GOPATH%/src/golang.org/x/vgo ; go install
  > go help mod <command>       # 帮助 SET GO111MODULE=on
  > go mod init example.com/app # 生成 go.mod 文件，golang.org/..各个包都需要翻墙，go.mod中用replace替换成github
  > go get ./...  # go mod tidy # 根据已有代码import需要的依赖自动生成require语句
  > go get -u # go get -u=patch # 升级到最新的次要版本,升级到最新的修订版本
  > go list -m                  # 查看当前的依赖和版本
  > go mod download             # 下载到$GOPATH/pkg/mod/cache共享缓存中
  > go mod edit -fmt            # 格式化 go.mod 文件
  > go mod edit -require=path@version # 添加依赖或修改依赖版本
  > go mod vendor               # 生成 vendor 文件夹, 下载你代码中引用的库
  > go build -mod=vendor        # 使用 vendor 文件夹
  > go build -mod=readonly      # 防止隐式修改 go.mod
go get -u github.com/sparrc/gdm

# 源代码版本管理
go get -d github.com/gogs/gogs  # 一款极易搭建的自助Git服务  *30k

# 学习playground*
go get github.com/golang/playground
go get github.com/golang/example/hello
go get github.com/shen100/golang123        # 适合初学者
go get github.com/insionng/zenpress        # 适合学习 cms system
go get github.com/muesli/cache2go          # 缓存库，代码量少，适合学习，锁、goroutines等
go get -d github.com/getlantern/lantern    # 网络底层的东西，适合深入学习                    *42k
git clone https://github.com/adonovan/gopl.io.git %GOPATH%/src/github.com/adonovan/gopl.io # programs
~~~

> Docker 编译器 [Golang + custom build tools](https://hub.docker.com/_/golang)

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

 * [QT跨平台应用框架](https://github.com/therecipe/qt)、[Webview-App](https://github.com/zserge/webview)、[Electron-App](https://github.com/asticode/go-astilectron)、[WebAssembly-App](https://github.com/murlokswarm/app)
~~~
go get github.com/rs/xid                   # uuid shortuuid Snowflake MongoID xid
go get github.com/satori/go.uuid           # uuid generator, Version 1 ~ 5 (RFC 4122)
go get github.com/juju/utils               # General utility functions
go get github.com/henrylee2cn/goutil       # Common and useful utils
go get github.com/PuerkitoBio/goquery      # HTML解析库，像jQuery那样操作DOM   *7k

go get github.com/cloudflare/cfssl/cmd/... # SSL证书 usage http://play.etcd.io/install#TLS
go get github.com/spf13/viper && go get github.com/spf13/pflag # 配置(JSON,TOML,YAML,HCL)热加载;远程配置;缓存;加密
go get github.com/xordataexchange/crypt/bin/crypt 加密存储 secret keyring: gpg(gpg4win)用于安全传输(类似rsa)
go get github.com/minio/minio-go           # 云存储|分布式存储|自建网盘 https://www.min.io https://docs.min.io/cn
go get -d github.com/minio/mc              # 云存储|配置客户端, 指南: https://docs.min.io/cn/minio-client-quickstart-guide.html
go get -d github.com/minio/minio           # 云存储|配置服务端, 运行: hidec /w minio.exe server d:\docker\app\minio\data
go get github.com/perkeep/perkeep/cmd/...  # Camlistore 个人存储系统：一种存储、同步、共享、建模和备份内容的方式
go get -d github.com/s3git/s3git           # 云存储的Git: 用于数据的分布式版本控制系统
# https://github.com/etcd-io/etcdlabs      # 分布式可靠键值存储，适用于分布式系统中最关键的数据；提供分享配置和服务发现
# client: http://play.etcd.io              # server: https://github.com/etcd-io/etcd/releases
go get github.com/coocood/freecache        # cache and high concurrent performance
go get github.com/patrickmn/go-cache       # in-memory key:value store/cache (similar to Memcached)适用于单台应用程序
go get github.com/peterbourgon/diskv       # 支持磁盘的 key-value 存储
go get github.com/chrislusf/seaweedfs/weed # 一个用于小文件的简单且高度可扩展的分布式文件系统，可集成其他云服务，如AWS...
go get -u golang.org/x/sys/...             # 跨平台文件系统监控库 https://github.com/fsnotify/fsnotify

go get github.com/go-redis/redis           # 内存数据库,类型安全的Redis-client
go get github.com/gomodule/redigo/redis    # 内存数据库,使用原生的Redis-cli
go get github.com/syndtr/goleveldb/leveldb # 内存数据库,谷歌leveldb-client
go get github.com/seefan/gossdb/example    # 内存数据库,替代Redis的ssdb http://ssdb.io/zh_cn
go get github.com/tidwall/buntdb           # 内存数据库,BuntDB is a low-level, in-memory, key/value store, persists to disk
go get github.com/tidwall/buntdb-benchmark # 性能测试 > buntdb-benchmark -n 10000 -q # 单机时超越Redis，有索引和geospatial功能
go get github.com/allegro/bigcache         # 高可用千兆级数据的高效 key/value 缓存   *2k
go get github.com/boltdb/bolt/...          # 原生的 key/value 数据库, 类似 sqlite  *10k
go get github.com/cockroachdb/cockroach    # 云数据存储系统，支持地理位置、事务等 *20k https://www.cockroachlabs.com/docs/stable
go get -d github.com/tidwall/tile38        # 具有空间索引和实时地理位置数据库  *6k
go get -d github.com/pingcap/tidb          # TiDB 支持包括传统 RDBMS 和 NoSQL 的特性 *18k https://pingcap.com/docs-cn
go get github.com/influxdata/influxdb1-client/v2 # 分布式、事件、实时的可扩展数据库 *19k https://github.com/influxdata/influxdb
go get github.com/dgraph-io/dgraph/dgraph  # 具有可扩展、分布式、低延迟和高吞吐量功能的图形数据库  *9k
go get github.com/melihmucuk/geocache      # 适用于地理位置处理, 基于应用程序的内存缓存 *1k
go get github.com/bluele/gcache            # 支持LFU、LRU 和 ARC 的缓存数据库 *1k
go get github.com/bradfitz/gomemcache/memcache # memcache 客户端库
go get github.com/couchbase/go-couchbase   # Couchbase 客户端

go get github.com/astaxie/beego/orm        # 数据库orm    *19k support mysql,postgres,sqlite3
go get github.com/jinzhu/gorm              # 数据库gorm   *12k http://gorm.io/docs
git clone https://github.com/rana/ora.git %GOPATH%/src/gopkg.in/rana/ora.v4 && go get gopkg.in/rana/ora.v4
go get github.com/mattn/go-oci8            # Oracle env: instantclient & MinGW-w64-gcc & pkgconfig/oci8.pc
go get github.com/go-sql-driver/mysql      # Mysql     https://github.com/siddontang/go-mysql
go get github.com/denisenkom/go-mssqldb    # MsSql
go get github.com/lib/pq                   # Postgres  https://github.com/prest/prest
go get github.com/jackc/pgx                # Postgres  驱动与工具集
go get github.com/sosedoff/pgweb           # Postgres  Web管理系统
go get github.com/mattn/go-sqlite3         # SQLite
go get github.com/jmoiron/sqlx             # 数据库sql    *6k  extensions go's standard database/sql library
  go get github.com/heetch/sqalx             # sqlx & sqalx 支持嵌套的事务
  go get github.com/twiglab/sqlt             # sqlx & sqlt 模板拼接sql和java的数据库访问工具MyBatis的sql配置
  go get github.com/albert-widi/sqlt         # sqlx & sqlt 支持数据库主从数据源，读写分离
go get github.com/go-xorm/xorm             # 数据库xorm   *5k  support mysql,postgres,tidb,sqlite3,mssql,oracle
  go get github.com/go-xorm/builder          # ^xorm SQL Builder 增强-拼接sql
  go get github.com/xormplus/xorm            # ^xorm增强版*$ 支持sql模板,动态sql,嵌套事务,配置等特性...
go get gopkg.in/gormigrate.v1              # 数据库gorm 数据库迁移助手 https://github.com/go-gormigrate/gormigrate
go get github.com/mattes/migrate           # 数据库迁移工具 *2k
go get github.com/rubenv/sql-migrate/...   # 数据库 schema 迁移工具，允许使用 go-bindata 将迁移嵌入到应用程序中 *1k
go get github.com/gchaincl/dotsql          # 帮助你将 sql 文件保存至某个地方并轻松使用它
go get github.com/xo/xo                    # 命令行工具 xo --help  [DbFirst]生成 models/*.xo.go
go get github.com/variadico/scaneo         # 命令行工具 scaneo -h  [DbFirst]生成 models/*.go

go get github.com/olivere/elastic          # Elasticsearch 6.0 客户端
go get github.com/blevesearch/bleve        # 现代文本索引库 *5k
go get github.com/siesta/neo4j             # Neo4j 客户端 github.com/jmcvetta/neoism
go get github.com/cayleygraph/cayley       # 图形数据库 Driven & RESTful API & LevelDB Stores
go get github.com/DarthSim/imgproxy        # Fast image server: docker pull darthsim/imgproxy
go get willnorris.com/go/imageproxy/...    # Caching image proxy server & docker & nginx
go get labix.org/v2/mgo                    # MongoDB 驱动
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
go get github.com/gocraft/work             # do work of redis-queue *1k https://github.com/gocraft/work#run-the-web-ui

go get github.com/nsqio/go-nsq             # 实时消息平台nsq *15k | nsqlookupd & nsqd & nsqadmin https://nsq.io
go get github.com/streadway/amqp           # rabbitmq client tutorials https://www.rabbitmq.com/#getstarted
go get github.com/blackbeans/kiteq         # KiteQ 是一个基于 go + protobuff 实现的多种持久化方案的 mq 框架
# 聊天室 git clone https://github.com/GoBelieveIO/im_service.git && cd im_service && dep ensure && mkdir bin && make install
# https://github.com/oikomi/FishChatServer2
go get github.com/mattermost/mattermost-server # 通讯 *15k 为团队带来跨PC和移动设备的消息、文件分享，提供归档和搜索功能+前端React
go get github.com/gorilla/websocket        # WebSocket Serve *8k
go get github.com/gotify/server            # WebSocket Serve (Includes Web-UI manage) https://gotify.net
go get github.com/gotify/cli               # WebSocket client to push messages

go get github.com/gin-gonic/gin            # 后端WebApi : Gin Web Framework
go get github.com/mholt/caddy/caddy        # 后端WebSvr : caddy | 配置快apache+nginx | caddyserver.com
go get github.com/labstack/echo/v4         # 后端WebSvr : echo
go get github.com/codegangsta/gin          # 服务Web站点热启动 > gin -h
go get github.com/ochinchina/supervisord   # 原生的supervisor > supervisord -c website.conf -d
go get github.com/sourcegraph/checkup/cmd/checkup # 分布式站点健康检查工具 > checkup --help
go get github.com/smallnest/go-web-framework-benchmark # Web性能测试工具 > wrk -t16 -c100 -d30s http://127.0.0.1/api
go get github.com/prometheus/prometheus/cmd/... # 服务监控系统和时间序列数据库 *23k https://prometheus.io/community
# 各大 Go 模板引擎的对比及压力测试 https://github.com/SlinSo/goTemplateBenchmark
# 小米公司的互联网企业级监控系统 https://book.open-falcon.org

go get github.com/dgrijalva/jwt-go/cmd/jwt # JSON Web Tokens (JWT)
go get github.com/gorilla/sessions         # session & cookie authentication
go get github.com/kgretzky/evilginx2       # session cookies, allowing for the bypass of 2-factor authentication 
go get github.com/dchest/captcha           # 验证码|图片|声音
go get github.com/mojocn/base64Captcha     # 验证码|展示 http://captcha.mojotv.cn
go get github.com/dpapathanasiou/go-recaptcha # Google验证码|申请 https://www.google.com/recaptcha/admin/create
go get github.com/emersion/go-imap/...     # 邮箱服务 IMAP library for clients and servers
go get github.com/sdwolfe32/trumail/...    # 邮箱验证 clients
go get github.com/matcornic/hermes/v2      # HTML e-mails, like: npm i mailgen  https://github.com/eladnava/mailgen

go get github.com/gorilla/websocket        # https://github.com/joewalnes/websocketd websocketd.com
go get github.com/gobwas/ws                # Tiny WebSocket \ https://github.com/socketio/socket.io
go get github.com/reactivex/rxgo           # 响应式编程
go get github.com/go-swagger/go-swagger/cmd/swagger # swagger 文档生成器 https://goswagger.io/install.html
go get github.com/istio/istio              # 集群的管理   *16k  for k8s
go get github.com/yudai/gotty              # 终端扩展服务
go get github.com/asciimoo/wuzz            # 用于http请求的-交互式命令行工具-增强的curl

go get -u -v github.com/smallnest/rpcx/... # 分布式RPC服务框架  *3k
go get github.com/micro/go-micro           # 分布式RPC微服务    *7k
go get github.com/go-kit/kit/cmd/kitgen    # 微服务构建   *13k standard library for web frameworks...
git clone https://github.com/EasyDarwin/EasyDarwin.git %GOPATH%/src/github.com/EasyDarwin/EasyDarwin # RTSP流媒体服务
go get github.com/iikira/BaiduPCS-Go       # 百度网盘命令行客户端
go get github.com/inconshreveable/go-update # 自动更新应用程序

go get github.com/elves/elvish             # shell for unix > 可编程：数组、字典、传递对象的增强型管道、闭包、模块机制、类型检查
go get github.com/mattn/sudo               # sudo for windows > sudo cmd /c dir ; sudo notepad c:\windows\system32\drivers\etc\hosts
go get github.com/lxn/win                  # Windows API wrapper package
go get github.com/lxn/walk                 # Windows UI Application Library Kit *3k
go get github.com/google/gapid             # Windows UI App : Graphics API Debugger
go get github.com/FiloSottile/mkcert       # 证书管理工具 *18k
# [申请Let's Encrypt永久免费SSL证书]          https://www.jianshu.com/p/3ae2f024c291
go get github.com/go-acme/lego/cmd/lego    # Let's Encrypt client and ACME library, DNS providers manager.
# [QT跨平台应用框架] Qt binding package
go get -u -v github.com/therecipe/qt/cmd/... && for /f %v in ('go env GOPATH') do %v\bin\qtsetup test && %v\bin\qtsetup
# [Bringing Flutter to Windows, MacOS and Linux] - through the power of Go and GLFW.
# https://github.com/go-flutter-desktop/go-flutter
go get github.com/codesenberg/bombardier     # http高性能基准测试工具 > bombardier
go get github.com/BurntSushi/wingo/wingo-cmd # 一个功能齐全的窗口管理器 > wingo-cmd

# 云计算
# 亚马逊 AWS https://www.amazonaws.cn/tools

# 阿里云 https://api.aliyun.com
go get -u github.com/aliyun/alibaba-cloud-sdk-go/sdk
# 云服务器 ECS、对象存储 OSS、阿里云关系型数据库、云数据库MongoDB版、CDN、VPC、
# 视频点播、音视频通信、媒体转码、负载均衡、云监控、容器服务、邮件推送、弹性伸缩、移动推送、日志服务、交易与账单管理

# 腾讯云 https://console.cloud.tencent.com/api/explorer
go get -u github.com/tencentcloud/tencentcloud-sdk-go
# 云服务器 黑石物理服务器 云硬盘 容器服务 容器实例服务 弹性伸缩 无服务器云函数 批量计算
# 负载均衡 私有网络 专线接入 云数据库 MySQL 云数据库 Redis 云数据库 MongoDB 数据传输服务 DTS 云数据库 MariaDB
# 分布式数据库 DCDB	云数据库 SQL'Server 云数据库 PostgreSQL 内容分发网络 主机安全 Web漏洞扫描 应用安全 云点播
# 云直播 智能语音服务 机器翻译 智能钛机器学习 催收机器人 智聆口语评测 腾讯优评 Elasticsearch'Service
# 物联网通信 TBaaS 云监控 迁移服务平台 电子合同服务 计费相关 渠道合作伙伴 人脸核身'云智慧眼
# 威胁情报云查 样本智能分析平台 数学作业批改 人脸融合 人脸识别 数字版权管理

# 七牛云 https://developer.qiniu.com/sdk
go get -u github.com/qiniu/api.v7 # 对象存储

# CDN.内容分发网络
# 又拍云 https://www.upyun.com

# 实时通信云
# 野狗 https://www.wilddog.com 实时通信引擎（Sync），它帮助企业快速实现三大功能：基础实时通信、实时数据分发和实时数据持久化。

# 应用数据、开放数据API
# 聚合数据 https://www.juhe.cn

# API+SDK'排名'开发者服务平台 https://sdk.cn
~~~

#### ③ [开源的 Web 框架](https://github.com/avelino/awesome-go#web-frameworks), [参考构建企业级的 RESTful API 服务](https://juejin.im/book/5b0778756fb9a07aa632301e)
~~~
# 开发
cd %GOPATH%/src                                                                 # 项目框架 Gin Web Framework
git clone https://github.com/lexkong/apiserver_demos apiserver                  # 项目源码-复制^demo至-工作目录
git clone https://github.com/lexkong/vendor                                     # 项目依赖-govendor
go get github.com/StackExchange/wmi                                             # 项目依赖-缺失的包
# 构建 
cd %GOPATH%/src/apiserver && gofmt -w . && go tool vet . && go build -v -o [应用名] [目录默认.]
# 运行
%GOPATH%/src/apiserver/apiserver.exe
~~~

#### ④ [中文标准库文档](https://studygolang.com/pkgdoc)

#### ⑤ 阅读相关文章

 * 高性能
    * [高并发架构解决方案](https://studygolang.com/articles/15479)


----

