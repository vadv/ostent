// +build freebsd darwin

package system

import (
	"github.com/ostrost/ostent/system/operating"
	metrics "github.com/rcrowley/go-metrics"
	sigar "github.com/rzab/gosigar"
)

type ExtraMetricRAM struct {
	Inactive metrics.Gauge
	Wired    metrics.Gauge
	Active   metrics.Gauge
}

func NewMetricRAM(r metrics.Registry) *operating.MetricRAM {
	return operating.ExtraNewMetricRAM(r, &ExtraMetricRAM{
		Inactive: metrics.NewRegisteredGauge("memory.memory-inactive", r),
		Wired:    metrics.NewRegisteredGauge("memory.memory-wired", r),
		Active:   metrics.NewRegisteredGauge("memory.memory-active", r),
	})
}

func (emr *ExtraMetricRAM) UpdateRAM(got sigar.Mem, extra1, extra2 uint64) {
	emr.Inactive.Update(int64(got.ActualFree - got.Free))
	emr.Wired.Update(int64(extra1))
	emr.Active.Update(int64(extra2))
}

func NewMetricCPU(r metrics.Registry, name string) *operating.MetricCPU {
	return operating.ExtraNewMetricCPU(r, name, nil)
}
