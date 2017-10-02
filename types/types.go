package types

import (
	"github.com/gollector/gollector/logger"
	"github.com/gollector/gollector/plugins/command"
	"github.com/gollector/gollector/plugins/cpu_usage"
	"github.com/gollector/gollector/plugins/fs_usage"
	"github.com/gollector/gollector/plugins/io_usage"
	"github.com/gollector/gollector/plugins/json_poll"
	"github.com/gollector/gollector/plugins/load_average"
	"github.com/gollector/gollector/plugins/mem_usage"
	"github.com/gollector/gollector/plugins/net_usage"
	"github.com/gollector/gollector/plugins/process_count"
	"github.com/gollector/gollector/plugins/process_mem_usage"
	"github.com/gollector/gollector/plugins/record"
	"github.com/gollector/gollector/plugins/socket_usage"
)

type PluginResult interface{}
type PluginResultCollection map[string]PluginResult

var Plugins = map[string]func(interface{}, *logger.Logger) interface{}{
	"load_average":      load_average.GetMetric,
	"cpu_usage":         cpu_usage.GetMetric,
	"mem_usage":         mem_usage.GetMetric,
	"command":           command.GetMetric,
	"net_usage":         net_usage.GetMetric,
	"io_usage":          io_usage.GetMetric,
	"record":            record.GetMetric,
	"fs_usage":          fs_usage.GetMetric,
	"json_poll":         json_poll.GetMetric,
	"socket_usage":      socket_usage.GetMetric,
	"process_count":     process_count.GetMetric,
	"process_mem_usage": process_mem_usage.GetMetric,
}

/*

How this works:

Key is the type of the metric, value is a func which returns the Params value
for any given metric.

The value + _ + the key is used to name the metric to allow for multiple
metrics of a single type.

*/

var Detectors = map[string]func() []string{
	"load_average": func() []string { return []string{} },
	"cpu_usage":    func() []string { return []string{} },
	"mem_usage":    func() []string { return []string{} },
	"net_usage":    net_usage.Detect,
	"io_usage":     io_usage.Detect,
	"fs_usage":     fs_usage.Detect,
}

type ConfigMap struct {
	Type   string
	Params interface{}
}

type PluginConfig map[string]ConfigMap

type CirconusConfig struct {
	Listen       string
	Username     string
	Password     string
	Facility     string
	LogLevel     string
	PollInterval uint
	Plugins      PluginConfig
}
