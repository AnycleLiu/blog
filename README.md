# 项目结构

- config 
```
配置模块，存放配置以及提供config对象
```
- controllers
```
控制器模块，router handler实现，一般一类资源一个controller
```
- core
```
核心模块，定义领域模型
```
- db
```
对sql.DB的包装，提供db连接池的初始化和获取连接池对象
```
- middlewares
```
中间件模块，实现gin中间件，比如鉴权、request log等
```
- routers
```
路由模块。负责controller对应路由的注入
```
- scripts
```
存放各种脚本，比如shell脚本，数据库脚本
```
- main
```
程序入口，http服务监听
```
