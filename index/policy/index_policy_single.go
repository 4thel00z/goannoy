package policy

import "github.com/4thel00z/goannoy/interfaces"

func SingleWorker() *annoyIndexSingleThreadedBuildPolicy {
	return &annoyIndexSingleThreadedBuildPolicy{}
}

type annoyIndexSingleThreadedBuildPolicy struct{}

func (p *annoyIndexSingleThreadedBuildPolicy) Build(
	builder interfaces.AnnoyIndexBuilder,
	numberOfTrees, nThreads int,
) {
	builder.ThreadBuild(numberOfTrees, 0, p)
}

func (p *annoyIndexSingleThreadedBuildPolicy) LockNNodes() {
}

func (p *annoyIndexSingleThreadedBuildPolicy) UnlockNNodes() {
}

func (p *annoyIndexSingleThreadedBuildPolicy) LockNodes() {
}

func (p *annoyIndexSingleThreadedBuildPolicy) UnlockNodes() {
}

func (p *annoyIndexSingleThreadedBuildPolicy) LockSharedNodes() {
}

func (p *annoyIndexSingleThreadedBuildPolicy) UnlockSharedNodes() {
}

func (p *annoyIndexSingleThreadedBuildPolicy) LockRoots() {
}

func (p *annoyIndexSingleThreadedBuildPolicy) UnlockRoots() {
}
