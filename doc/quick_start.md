## Quick Start

### Build

```
$ cd $GOPATH

$ mkdir -p src/github.com/QOSGroup/

$ cd src/github.com/QOSGroup/

$ git clone https://github.com/QOSGroup/cassini.git

$ cd cassini/cmd/cassini
```

\# 注意：请确认通过网路可以获取所有依赖，并确认已配置环境变量开启了go modules!

```
$ go build
```

### Install & start gnatsd

$ go get github.com/nats-io/gnatsd

$ gnatsd & 

### Config

$ vi src/github.com/QOSGroup/cassini/config/config.conf

\# 配置nats为gnatsd服务器地址，集群内多个地址用","号分割;  

\# prikey为cassini的私钥; 

\# consensus默认为"yes",如果设为"no" cassini将关闭共识功能;  

\# eventWaitMillitime 单位为ms,建议与链的建块周期保持一致; 
 
\# 在qscs段配置公链和联盟链，name为链名称，nodes为链节点地址，多个地址用“,”号分割,公链signature设为true。 
```
{
    "nats":      "nats://127.0.0.1:4222",
    "prikey":    "", 
    "consensus": "yes",
    "eventWaitMillitime": 2000,
    "qscs": [
        {
            "name":   "qstars-test",
            "nodes": "ip:26657,ip:26657"
        },
        {
            "name":      "qos-test",
            "signature": true,
            "nodes":     "ip:26657,ip:26657"
        }
    ]
}
```

### Commands

\# 帮助信息

```
$ ./cassini help

$ ./cassini [command] -h
```

\# 远端服务模拟，提供中继访问订阅和查询跨链事件及交易的模拟服务端，以便不需要每次中继项目自测时都需要启动（甚至可能需要启动多条）完整的区块链服务。

\# 注意：为了测试方便，目前启动模拟服务会自动启动中继服务已进行测试，后续会实现可配置是否自动启动中继服务以方便更多的测试方案！

```
$ ./cassini mock [flag]
```

\# 手工测试调试

\# 监听远程跨链交易事件，可设置IP地址、端口及订阅条件以确认远端地址可以正常订阅到跨链交易事件。

```
$ ./cassini events [flag]
```

\# 跨链交易查询和广播接口测试，可以查询和广播交易，以确认QCP跨链协议规范中的交易相关接口服务正常。

```
$ ./cassini tx [flag]
```

\# 启动中继服务，按照QCP跨链协议规范，向远端订阅跨链交易事件和查询、广播跨链交易。

```
$ ./cassini start [flag]
```
