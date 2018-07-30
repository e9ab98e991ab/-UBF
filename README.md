<div align=center>
    <img src="http://cover.kancloud.cn/johng/gf" width="150"/>
</div>

# 安装
```html
go get -u gitee.com/Godfeer/UBF
```
构建 
```$gradle

    请运行
    
    ./gradlew init # *nix 
    gradlew init # Windows

    在IDE中
    运行gradlew vendor安装依赖到项目中
    运行gradlew showGopathGoroot或者gradlew sGG打印项目的GOPATH和GOROOT
    使用上述GOROOT和GOPATH配置IDE
    开始开发！

```



# 介绍
Universal Backend Framework

UBF是一款模块化、松耦合、轻量级、高性能的Web开发框架。采用gradle构建。
开源项目地址(仓库保持实时同步)：
[Gitee](https://gitee.com/Godfeer/UBF.git)，[Github](https://github.com/e9ab98e991ab/UBF)。
 

# 特点
1. 轻量级、高性能，模块化、松耦合设计，丰富的开发模块；
1. 热重启、热更新特性，并支持Web界面及命令行管理接口；
1. 专业的技术交流群，完善的开发文档及示例代码，良好的中文化支持；
1. 支持多种形式的服务注册特性，灵活高效的路由控制管理；
1. 支持服务事件回调注册功能，可供选择的pprof性能分析模块；
1. 支持配置文件及模板文件的自动检测更新机制，即修改即生效；
1. 支持自定义日期时间格式的时间模块，类似PHP日期时间格式化；
1. 强大的数据/表单校验模块，支持常用的40种及自定义校验规则；
1. 强大的网络通信TCP/UDP组件，并提供TCP连接池特性，简便高效；
1. 提供了对基本数据类型的并发安全封装，提供了常用的数据结构容器；
1. 支持Go变量/Json/Xml/Yml/Toml任意数据格式之间的相互转换及创建；
1. 强大的数据库ORM，支持应用层级的集群管理、读写分离、负载均衡，查询缓存、方法及链式ORM操作；
1. 更多特点请查阅框架手册和源码；

 
# 感谢
gogradle 

```html
    https://github.com/gogradle/gogradle.git
```
 
