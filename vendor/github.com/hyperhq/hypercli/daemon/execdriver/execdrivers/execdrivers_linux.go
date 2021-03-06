// +build linux

package execdrivers

import (
	"path"

	"github.com/hyperhq/hypercli/daemon/execdriver"
	"github.com/hyperhq/hypercli/daemon/execdriver/native"
	"github.com/hyperhq/hypercli/pkg/sysinfo"
)

// NewDriver returns a new execdriver.Driver from the given name configured with the provided options.
func NewDriver(options []string, root, libPath string, sysInfo *sysinfo.SysInfo) (execdriver.Driver, error) {
	return native.NewDriver(path.Join(root, "execdriver", "native"), options)
}
