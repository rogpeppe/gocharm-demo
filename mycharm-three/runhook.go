package mycharm

import (
	"fmt"
	"github.com/juju/gocharm/charmbits/httprelation"
	"github.com/juju/gocharm/hook"
)

type myCharm struct {
	ctxt *hook.Context
	http httprelation.Provider
}

func RegisterHooks(r *hook.Registry) {
	var c myCharm
	r.RegisterContext(c.setContext, nil)
	r.RegisterHook("start", c.start)
	c.http.Register(r.Clone("website"), "website", false)
	r.RegisterHook("*", c.changed)
}

func (c *myCharm) changed() error {
	if c.http.HTTPPort() == 0 {
		c.ctxt.SetStatus(hook.StatusWaiting, "waiting for HTTP server config")
	} else {
		c.ctxt.SetStatus(hook.StatusActive, fmt.Sprintf("active on port %d", c.http.HTTPPort()))
	}
	return nil
}

func (c *myCharm) setContext(ctxt *hook.Context) error {
	c.ctxt = ctxt
	return nil
}

func (c *myCharm) start() error {
	c.ctxt.SetStatus(hook.StatusActive, "happily idle")
	return nil
}
