package entities

type Task struct {
	Id       int64
	Task     string `validate:"required"`
	Assigne  string `validate:"required"`
	Deadline string `validate:"required"`
}
