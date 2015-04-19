package mycharm

import "github.com/juju/gocharm/hook"

// RegisterHooks registers all hook functionality
// with the given registry. It returns immediately.
func RegisterHooks(r *hook.Registry) {
	r.RegisterHook("start", func() error {
		// Start first service here,
		return nil
	})
	r.RegisterHook("start", func() error {
		// Start second service here,
		return nil
	})
}
