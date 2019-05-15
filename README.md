# go-easy

## Instruction

go-easy是一个基于原生net库实现的非常轻量级的web框架，通过配置文件定义服务信息，支持Restful API，前置过滤器、后置过滤器以及错误过滤器，支持自定义过滤器



## Functions

* 支持Restful API
* 控制器绑定Path
* 自定义过滤器
* 接口维度限流
* 日志拆分（小时、天）
* 支持过滤器（前置过滤器、后置过滤器、错误过滤器）
  * 默认过滤器 - rate limit filter



## Filter Interface

```
type Filter interface {
	// 过滤器名称
    Name() string
    // 前置过滤器
    Pre(ctx context.Context) (statusCode int, err error)
    // 后置过滤器
    Post(ctx context.Context) (statusCode int, err error)
    // 错误过滤器
    PostErr(ctx context.Context)
}
```



## Controller Interface

```
type Controller interface {
    Get(ctx context.Context) (err error)
    Put(ctx context.Context) (err error)
    Post(ctx context.Context) (err error)
    Delete(ctx context.Context) (err error)
}
```





## Example

配置文件格式

```
Application: test
Debug:
    Enable: false
    Port: 8020
Server:
    IP: eth0
    Port: 8088
    ReadTimeout: 2000
    WriteTimeout: 2000
    IdleTimeout: 2000
    MaxHeaderSize: 2000
Logger:
    Writer: file
    Level: debug
    Format: "[%D %T] [%L] (%S) %M"
    Rotate: true
    RotateType: hour
    LogPath: /data/log/
Limiter:
    - InterfaceName: /test
      EnableRateLimit: true
      MaxRequestPerSecond: 1
```



自定义过滤器

```
package main

import (
	"easy/web/filter"
	"easy"
)

type TestFilter struct {
    filter.BaseFilter
}

func (t *TestFilter) Name() string {
    return "TestFilter"
}

func (t *TestFilter) Pre(ctx context.Context) (int, error) {
    return t.BaseFilter.Pre(ctx)
}

func (t *TestFilter) Post(ctx context.Context) (int, error) {
    return t.BaseFilter.Post(ctx)
}

func (t *TestFilter) PostErr(ctx context.Context) (int, error) {
    return t.BaseFilter.PostErr(ctx)
}

func main() {
    easy.RegisterFilter(&TestFilter{})
 	...
}
```

定义controller，每个请求路径对应一个controller

```
package main

import (
	"easy"
)

type TestController struct {
    
}

func (t *TestController) Get(ctx context.Context) error {
    return nil
}

func (t *TestController) Post(ctx context.Context) error {
    return nil
}

func (t *TestController) Put(ctx context.Context) error {
    return nil
}

func (t *TestController) Delete(ctx context.Context) error {
    return nil
}

func main() {
    easy.RegisterController("/test", &TestController{})
    easy.Run()
}
```


