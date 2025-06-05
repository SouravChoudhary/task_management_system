package constants

type TaskStatus string

const (
	StatusPending   TaskStatus = "Pending"
	StatusCompleted TaskStatus = "Completed"
)

var validTaskStatuses = map[TaskStatus]bool{
	StatusPending:   true,
	StatusCompleted: true,
}

// GetValidTaskStatuses returns list of all allowed statuses.
func GetValidTaskStatuses() []TaskStatus {
	return []TaskStatus{StatusPending, StatusCompleted}
}

// IsValidTaskStatus checks if the given status is valid.
func IsValidTaskStatus(status string) bool {
	_, ok := validTaskStatuses[TaskStatus(status)]
	return ok
}
