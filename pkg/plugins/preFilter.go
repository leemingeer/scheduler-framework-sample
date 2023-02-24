package plugins

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// get the scheduling pod info
func (s *Sample) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) *framework.Status {
	klog.V(3).Infof("ming prefilter pod: %v", pod.Name)

	state.Write(preFilterStateKey, NewNoopStateData())
	return framework.NewStatus(framework.Success, "")
	//	return framework.NewStatus(framework.Unschedulable, fmt.Sprintf("hard made Unschedulable!, pod info: %#v", pod))
}

// PreFilterExtensions returns a PreFilterExtensions interface if the plugin implements one.
func (s *Sample) PreFilterExtensions() framework.PreFilterExtensions {
	klog.V(3).Infof("ming prefilter extension not implemented!")
	return nil
}
