# KayProxy

代理转发 (Golang)

[配套LUCI](https://github.com/Kaytz/luci-app-KayProxy)

# 特性

* 就是快
* 低内存、高效率
* 手机京东查询添加历史记录
* 学习过程中的产物，随缘更新

# 运行

> [release页面](https://github.com/Kaytz/KayProxy/releases)中0.2.7及其之后的zip包将默认自带证书，该证书相对比较可靠。  

为了你的安全，还是建议你自己生成证书（windows需要自己下载openssl)

```shell
./createCertificate.sh
```

运行程序（由于m=1时 会自动修改hosts生效 所以需要sudo）

```shell
sudo ./KayProxy
```

### 具体参数说明

```shell
./KayProxy -h

  -b	force the best music quality
  -c string
    	specify server cert,such as : "server.crt" (default "./server.crt")
  -k string
    	specify server cert key ,such as : "server.key" (default "./server.key")
  -l string
    	specify log file ,such as : "/var/log/KayProxy.log"
  -m int
    	specify running mode（1:hosts） ,such as : "1" (default 1)
  -p int
    	specify server port,such as : "80" (default 80)
  -sl int
    	set the number of songs searched on other platforms(0-3) ,such as : "1"
  -sp int
    	specify server tls port,such as : "443" (default 443)
  -v	display version info

```

# 重要提示

1. 应用通过本机dns获取域名ip，请注意本地hosts文件
2. 如遇错误，请关闭该功能
3. 推荐Openwrt使用，其他客户端未测试
### IOS信任证书步骤

1. 安装证书--设置-描述文件-安装
2. 通用-关于本机-证书信任设置-启动完全信任

### 已知

1. windows版本需要在应用内 设置代理 Http地址为「HttpProxy」下任意地址 端口 80
2. Linux 客户端 (1.2 版本以上需要在终端启动客户端时增加 --ignore-certificate-errors 参数)

# 感谢

[Go版本](https://github.com/cnsilvan/UnblockNeteaseMusic)以及为它贡献的所有coder

[NodeJs版本](https://github.com/nondanee/UnblockNeteaseMusic)以及为它贡献的所有coder

# 声明
该项目只能用作学习，请自行开通会员以支持平台购买更多的版权
