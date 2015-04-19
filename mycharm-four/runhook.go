package mycharm

import (
	"fmt"
	"net/http"

	"github.com/juju/gocharm/charmbits/httpservice"
	"github.com/juju/gocharm/hook"
)

type myCharm struct {
	ctxt *hook.Context
	svc httpservice.Service
}

func RegisterHooks(r *hook.Registry) {
	var c myCharm
	r.RegisterContext(c.setContext, nil)
	c.svc.Register(r.Clone("website"), "", "website", newHelloHandler)
	r.RegisterHook("*", c.changed)
}

func newHelloHandler(struct{}) (http.Handler, error) {
	return http.HandlerFunc(helloHandler), nil
}

func (c *myCharm) changed() error {
	c.svc.Start(struct{}{})
	if c.svc.ServiceStarted() {
		c.ctxt.SetStatus(hook.StatusActive, "web server active")
	} else {
		c.ctxt.SetStatus(hook.StatusWaiting, "web server inactive")
	}
	return nil
}

func (c *myCharm) setContext(ctxt *hook.Context) error {
	c.ctxt = ctxt
	return nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Canonical")
}
