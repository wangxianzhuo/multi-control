# Multi-Control

## Description

RESTful 控制接口

## Usage

>实现`control.Unit`接口，以如下方式调用

```golang
package main

import (
    "github.com/labstack/echo"
    control "github.com/wangxianzhuo/multi-control"
    "github.com/wangxianzhuo/multi-control/sample/print"
)

func main() {
    e := echo.New()
    m, err := control.NewMachine(make(chan string))
    if err != nil {
        panic(err)
    }
    m.RegisterAsEchoMiddleware(e)
    u := print.Print{PID: "1"}
    m.AddUnit(&u)
    control.Routes(e)

    e.Start(":9999")
}
```

## RESTful API

| Path | Method | Description |
| ---- | ------ | ----------- |
| /units | GET | 获取所有unit，返回json |
| /units | DELETE | 删除所有unit |
| /unit/:id | GET | 获取id为`:id`的unit的`Info()`方法的返回值，返回json |
| /unit/:id | POST | 新增一个id为`:id`的unit，当前`不支持` |
| /unit/:id | DELETE | 删除一个id为`:id`的unit |
| /unit/:id/start | POST | id为`:id`的unit调用`Start()` |
| /unit/:id/stop | POST | id为`:id`的unit调用`Stop()` |
| /unit/:id/restart | POST | id为`:id`的unit调用`Restart()` |
| /unit/:id/reload | POST | id为`:id`的unit调用`Reload()` |
| /units/start | POST | 调用所有unit的`Start()` |
| /units/stop | POST | 调用所有unit的`Stop()` |
| /units/restart | POST | 调用所有unit的`Restart()` |
| /units/reload | POST | 调用所有unit的`Reload()` |

> 执行成功返回 `http status 200`， 否则返回 `http status 500`
