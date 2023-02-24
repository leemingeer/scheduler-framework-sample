package plugins

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// 插件名称
const (
	Name              = "sample-plugin"
	preFilterStateKey = "PreFilter" + Name
)

var _ = framework.QueueSortPlugin(&Sample{})
var _ = framework.PreFilterPlugin(&Sample{})
var _ = framework.PreBindPlugin(&Sample{})

type Sample struct {
	handle framework.Handle
}

func (s *Sample) Name() string {
	return Name
}

//type PluginFactory = func(configuration *runtime.Unknown, f FrameworkHandle) (Plugin, error)
func New(_ runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	// 入参1： 表示传入的参数
	// 参数2：framework框架对象，有clientSet等client. 可以获取集群资源
	klog.V(3).Infof("sample plugin config!")
	return &Sample{handle: handle}, nil
}

type noopStateData struct {
}

func NewNoopStateData() framework.StateData {
	return &noopStateData{}
}

func (d *noopStateData) Clone() framework.StateData {
	return d
}
