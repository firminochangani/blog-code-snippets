package app

import "time"

type Task struct {
	id        string
	name      string
	status    string
	ownerId   string
	createdAt time.Time
}

func (t *Task) Id() string {
	return t.id
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) Status() string {
	return t.status
}

func (t *Task) OwnerId() string {
	return t.ownerId
}

func (t *Task) CreatedAt() time.Time {
	return t.createdAt
}
