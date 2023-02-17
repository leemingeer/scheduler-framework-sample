package plugins

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// 插件名称
const (
	Name = "sample-plugin"
	preFilterStateKey       = "PreFilter" + Name
)

var _ = framework.PreFilterPlugin(&Sample{})
var _ = framework.PreBindPlugin(&Sample{})


type Sample struct {
	handle framework.Handle
}

func (s *Sample) Name() string {
	return Name
}

func (s *Sample) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) *framework.Status {
	klog.V(3).Infof("ming prefilter pod: %v", pod.Name)

	state.Write(preFilterStateKey, NewNoopStateData())
	return framework.NewStatus(framework.Success, "")
}
// PreFilterExtensions returns a PreFilterExtensions interface if the plugin implements one.
func (s *Sample) PreFilterExtensions() framework.PreFilterExtensions {
	klog.V(3).Infof("ming prefilter extension not implemented!")
	return nil
}

func (s *Sample) PreBind(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	if nodeInfo, err := s.handle.ClientSet().CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{}); err != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("ming prebind get node info error: %+v", nodeName))
	} else {
		klog.V(3).Infof("ming prebind node info: %+v", nodeInfo.String())
		return framework.NewStatus(framework.Success, "")
	}
}

//type PluginFactory = func(configuration *runtime.Unknown, f FrameworkHandle) (Plugin, error)
func New(_ runtime.Object, handle framework.Handle) (framework.Plugin, error) {

	klog.V(3).Infof("Ming sample plugin config!")
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
