package task

type Story int

type StoryTask struct {
	Task
	StorySummary string
	Status       TaskStatus
}

func (b *StoryTask) SetStatus(s TaskStatus) error {
	b.Status = s
	return nil
}

func (b *StoryTask) GetStatus() TaskStatus {
	return b.Status
}
