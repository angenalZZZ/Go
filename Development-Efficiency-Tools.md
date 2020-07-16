# 开发效率工具汇总

## 一. 开发工具

**1)sql2go**
用于将 sql 语句转换为 golang 的 struct. 使用 ddl 语句即可。
例如对于创建表的语句: show create table xxx. 将输出的语句，直接粘贴进去就行。
http://stming.cn/tool/sql2go.html

**2)toml2go**
用于将编码后的 toml 文本转换问 golang 的 struct.
https://xuri.me/toml-to-go/

**3)curl2go**
用来将 curl 命令转化为具体的 golang 代码.
https://mholt.github.io/curl-to-go/

**4)json2go**
用于将 json 文本转换为 struct.
https://mholt.github.io/json-to-go/

**5)mysql 转 ES 工具**
http://www.ischoolbar.com/EsParser/

**6)golang**
模拟模板的工具，在支持泛型之前，可以考虑使用。
https://github.com/cheekybits/genny

**7)查看某一个库的依赖情况，类似于 go list 功能**
https://github.com/KyleBanks/depth

8)一个好用的文件压缩和解压工具，集成了 zip，tar 等多种功能，主要还有跨平台。
https://github.com/mholt/archiver

**9)go 内置命令**
go list 可以查看某一个包的依赖关系.
go vet 可以检查代码不符合 golang 规范的地方。

**10)热编译工具**
https://github.com/silenceper/gowatch

**11)revive**
golang 代码质量检测工具
https://github.com/mgechev/revive

**12)Go Callvis**
golang 的代码调用链图工具
https://github.com/TrueFurby/go-callvis

**13)Realize**
开发流程改进工具
https://github.com/oxequa/realize

**14)Gotests**
自动生成测试用例工具
https://github.com/cweill/gotests

##  

## 二.调试工具

**1)perf**
代理工具，支持内存，cpu，堆栈查看，并支持火焰图.
perf 工具和 go-torch 工具，快捷定位程序问题.
https://github.com/uber-archive/go-torch
https://github.com/google/gops

**2)dlv 远程调试**
基于 goland+dlv 可以实现远程调式的能力.
https://github.com/go-delve/delve
提供了对 golang 原生的支持，相比 gdb 调试，简单太多。

**3)网络代理工具**
goproxy 代理，支持多种协议，支持 ssh 穿透和 kcp 协议.
https://github.com/snail007/goproxy

**4)抓包工具**
go-sniffer 工具，可扩展的抓包工具，可以开发自定义协议的工具包. 现在只支持了 http，mysql，redis，mongodb.
基于这个工具，我们开发了 qapp 协议的抓包。
https://github.com/40t/go-sniffer

**5)反向代理工具，快捷开放内网端口供外部使用。**
ngrok 可以让内网服务外部调用
https://ngrok.com/
https://github.com/inconshreveable/ngrok

**6)配置化生成证书**
从根证书，到业务侧证书一键生成.
https://github.com/cloudflare/cfssl

**7)免费的证书获取工具**
基于 acme 协议，从 letsencrypt 生成免费的证书，有效期 1 年，可自动续期。
https://github.com/Neilpang/acme.sh

8)开发环境管理工具，单机搭建可移植工具的利器。支持多种虚拟机后端。
**vagrant**常被拿来同 docker 相比，值得拥有。
https://github.com/hashicorp/vagrant

**9)轻量级容器调度工具**
nomad 可以非常方便的管理容器和传统应用，相比 k8s 来说，简单不要太多.
https://github.com/hashicorp/nomad

**10)敏感信息和密钥管理工具**
https://github.com/hashicorp/vault

**11)高度可配置化的 http 转发工具，基于 etcd 配置。**
https://github.com/gojek/weaver

**12)进程监控工具 supervisor**
https://www.jianshu.com/p/39b476e808d8

13)基于**procFile**进程管理工具. 相比 supervisor 更加简单。
https://github.com/ddollar/foreman

14)基于 http，https，websocket 的**调试代理工具**，配置功能丰富。在线教育的 nohost web 调试工具，基于此开发.
https://github.com/avwo/whistle

**15)分布式调度工具**
https://github.com/shunfei/cronsun/blob/master/README_ZH.md
https://github.com/ouqiang/gocron

**16)自动化运维平台 Gaia**
https://github.com/gaia-pipeline/gaia

##  

## 三. 网络工具

![img](https://mmbiz.qpic.cn/mmbiz_png/j3gficicyOvatzJibnFP6EXDYRTOzL8ZgT4oXh17qR8vTl5uHKrXtFcp45GS6VyRXMJw4vnSQ3mDibKxL7ficGrPshw/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

##  

## 四. 常用网站

go 百科全书: https://awesome-go.com/

json 解析: https://www.json.cn/

出口 IP: https://ipinfo.io/

redis 命令: http://doc.redisfans.com/

ES 命令首页: 

https://www.elastic.co/guide/cn/elasticsearch/guide/current/index.html

UrlEncode: http://tool.chinaz.com/Tools/urlencode.aspx

Base64: https://tool.oschina.net/encrypt?type=3

Guid: https://www.guidgen.com/

常用工具: http://www.ofmonkey.com/

##  

## 五. golang 常用库

**日志**
https://github.com/Sirupsen/logrus
https://github.com/uber-go/zap

**配置**
兼容 json，toml，yaml，hcl 等格式的日志库.
https://github.com/spf13/viper

**存储**
mysql: https://github.com/go-xorm/xorm
es: https://github.com/elastic/elasticsearch
redis: https://github.com/gomodule/redigo
mongo: https://github.com/mongodb/mongo-go-driver
kafka: https://github.com/Shopify/sarama

**数据结构**
https://github.com/emirpasic/gods

**命令行**
https://github.com/spf13/cobra

**框架**
https://github.com/grpc/grpc-go
https://github.com/gin-gonic/gin

**并发**
https://github.com/Jeffail/tunny
https://github.com/benmanns/goworker
现在我们框架在用的，虽然 star 不多，但是确实好用，当然还可以更好用.
https://github.com/rafaeldias/async

**工具**
定义了实用的判定类，以及针对结构体的校验逻辑，避免业务侧写复杂的代码.
https://github.com/asaskevich/govalidator
https://github.com/bytedance/go-tagexpr

protobuf 文件动态解析的接口，可以实现反射相关的能力。
https://github.com/jhump/protoreflect

**表达式引擎工具**
https://github.com/Knetic/govaluate
https://github.com/google/cel-go

**字符串处理**
https://github.com/huandu/xstrings

**ratelimit 工具**
https://github.com/uber-go/ratelimit
https://blog.csdn.net/chenchongg/article/details/85342086
https://github.com/juju/ratelimit

**golang 熔断的库**
熔断除了考虑频率限制，还要考虑 qps，出错率等其他东西.
https://github.com/afex/hystrix-go
https://github.com/sony/gobreaker

**表格**
https://github.com/chenjiandongx/go-echarts

**tail 工具库**
https://github.com/hpcloud/taglshi

