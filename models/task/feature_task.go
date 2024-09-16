package task

type FeatureTaskImpact int

const (
	FTILOW FeatureTaskImpact = iota
	FTIMEDIUM
	FTIHIGH
)

type FeatureTask struct {
	Task
	FeatureSummary string
	Impact         FeatureTaskImpact
	Status         TaskStatus
}

func (b *FeatureTask) SetStatus(s TaskStatus) error {
	b.Status = s
	return nil
}

func (b *FeatureTask) GetStatus() TaskStatus {
	return b.Status
}
