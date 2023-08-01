package saia

type TaskStatus string

const (
	TaskStatusPending TaskStatus = "PENDING"
	TaskStatusSuccess TaskStatus = "SUCCESS"
	TaskStatusFailure TaskStatus = "FAILURE"
)
