package main

import (
	"beegoi4566/u"
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type PanicController struct {
	web.Controller
}

func (c *PanicController) HelloWorld() {
	c.Ctx.WriteString("hello, world")
}

func (c *PanicController) GetFile() {
	c.Ctx.WriteString("get file")
}

// Status: 200
// curl http://127.0.0.1:12345/panic/helloworld
// Status: 500
// curl http://127.0.0.1:12345/panic/getuint64
// Status: 404
// curl http://127.0.0.1:12345/panic/getfile
func main() {
	web.AutoRouter(&PanicController{})
	fmt.Printf("Routers: %+v", u.GetAllRegisteredRouters())
	web.Run("127.0.0.1:12345")
}
