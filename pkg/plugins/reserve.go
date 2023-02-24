package plugins

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// 在cache预留， 如assumePodVolume()中pv和pvc在cache中绑定
// Reserve is the functions invoked by the framework at "reserve" extension point.
func (s *Sample) Reserve(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	return nil
}

// 对cache进行回滚，Unreserve clears assumed PV and PVC cache.
func (s *Sample) Unreserve(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) {

}
