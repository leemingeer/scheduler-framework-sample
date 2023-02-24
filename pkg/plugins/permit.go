package plugins

import (
	"context"
	v1 "k8s.io/api/core/v1"
)

// Permit permits a pod to run, if the minMember match, it would send a signal to chan.
func (s *Sample) Permit(ctx context.Context, pod *v1.Pod, nodeName string) (bool, error) {
	// prevent or delay the binding of a Pod, 也就是后续的流程。
	// 比如联合调度时，podGroup中pod个数是否达到最小值了
	return true, nil
}
