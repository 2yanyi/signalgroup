package spacelimit

import (
	"fmt"
	"github.com/containerd/cgroups"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/pkg/errors"
	"time"
)

// SetUsageResource 软件使用资源限制
// cpu: 10% (单核)
// mem: 80MB
func SetUsageResource(pid, cpu, mem int) error {
	cpuMax := int64(cpu * 1000)
	memMax := int64(mem * (1024 * 1024))
	spaceFp := fmt.Sprintf("/%d", time.Now().UnixNano())
	control, err := cgroups.New(cgroups.V1, cgroups.StaticPath(spaceFp), &specs.LinuxResources{
		CPU: &specs.LinuxCPU{
			Quota: &cpuMax,
		},
		Memory: &specs.LinuxMemory{
			Limit: &memMax,
		},
	})
	if err != nil {
		return errors.New(err.Error())
	}
	defer control.Delete()

	if err = control.Add(cgroups.Process{Pid: pid}); err != nil {
		return errors.New(err.Error())
	}
	return nil
}
