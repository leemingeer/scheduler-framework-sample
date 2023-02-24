package plugins

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"math"
)

// Score invoked at the score extension point.
func (s *Sample) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	// you can get pod and node info by handler object
	nodeInfo, err := s.handle.SnapshotSharedLister().NodeInfos().Get(nodeName)
	if err != nil {
		return 0, framework.NewStatus(framework.Error, fmt.Sprintf("getting node %q from Snapshot: %v", nodeName, err))
	}

	// s.score favors nodes with terminating pods instead of nominated pods
	// It calculates the sum of the node's terminating pods and nominated pods
	return s.score(nodeInfo)
}

func (s *Sample) score(nodeInfo *framework.NodeInfo) (int64, *framework.Status) {
	var terminatingPodNum, nominatedPodNum int64
	// get nominated Pods for node from nominatedPodMap
	// 节点上已经运行的pod，是否包含在内存里的？
	nominatedPodNum = int64(len(s.handle.PreemptHandle().NominatedPodsForNode(nodeInfo.Node().Name)))
	for _, p := range nodeInfo.Pods {
		// Pod is terminating if DeletionTimestamp has been set
		if p.Pod.DeletionTimestamp != nil {
			terminatingPodNum++
		}
	}
	return terminatingPodNum - nominatedPodNum, nil
}

func (s *Sample) NormalizeScore(ctx context.Context, state *framework.CycleState, pod *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	// Find highest and lowest scores.
	var highest int64 = -math.MaxInt64
	var lowest int64 = math.MaxInt64
	for _, nodeScore := range scores {
		if nodeScore.Score > highest {
			highest = nodeScore.Score
		}
		if nodeScore.Score < lowest {
			lowest = nodeScore.Score
		}
	}

	//  x-fmin     newRange
	//  -----   =  --------
	//  delta      oldRange
	// x is the final framework score.
	// 将得分范围转换到framework框架的分值范围
	// Transform the highest to lowest score range to fit the framework's min to max node score range.
	oldRange := highest - lowest
	newRange := framework.MaxNodeScore - framework.MinNodeScore
	for i, nodeScore := range scores {
		if oldRange == 0 {
			scores[i].Score = framework.MinNodeScore
		} else {
			// delta按比例换算到framework range中的delta + framework的Min = 在framework range中的score
			scores[i].Score = ((nodeScore.Score - lowest) * newRange / oldRange) + framework.MinNodeScore
		}
	}

	return nil
}

// ScoreExtensions of the Score plugin.
func (s *Sample) ScoreExtensions() framework.ScoreExtensions {
	return s
}
