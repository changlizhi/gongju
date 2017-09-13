package gongju

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
)

type Mywriter struct {
}

func (w *Mywriter) Header() http.Header {
	ret := http.Header{}
	return ret
}
func (w *Mywriter) Write([]byte) (int, error) {
	return 1, nil
}
func (w *Mywriter) WriteHeader(par int) {}

func Kongzhiqi(c *beego.Controller) {
	c.Data = make(map[interface{}]interface{})
	c.Ctx = context.NewContext()
	c.Ctx.Input = context.NewInput()
	c.Ctx.Output = context.NewOutput()
	c.Ctx.Output.Context = context.NewContext()
	c.Ctx.Output.Context.ResponseWriter = &context.Response{new(Mywriter), true, 200}
}
