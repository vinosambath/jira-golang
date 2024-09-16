package task

type SubTrackTask struct {
	Title      string
	Status     TaskStatus
	parentTask StoryTask
}
