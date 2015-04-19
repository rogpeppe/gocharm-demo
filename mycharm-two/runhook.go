package mycharm

import "github.com/juju/gocharm/hook"

type myCharm struct { // HL1
	ctxt *hook.Context // HL1
} // HL1

func RegisterHooks(r *hook.Registry) {
	var c myCharm
	r.RegisterContext(c.setContext, nil) // HL3
	r.RegisterHook("start", c.start)
}

func (c *myCharm) setContext(ctxt *hook.Context) error { // HL2
	c.ctxt = ctxt // HL2
	return nil    // HL2
}

func (c *myCharm) start() error {
	c.ctxt.SetStatus(hook.StatusActive, "happily idle") // HL4
	return nil
}
