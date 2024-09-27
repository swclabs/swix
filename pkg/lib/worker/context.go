package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type Context struct {
	ctx  context.Context
	task *asynq.Task
}

func (c Context) ResultWriter() *asynq.ResultWriter {
	return c.task.ResultWriter()
}

func (c Context) Payload() []byte {
	return c.task.Payload()
}

func (c Context) Type() string {
	return c.task.Type()
}

func (c Context) Context() context.Context {
	return c.ctx
}

func (c Context) Return(data []byte) error {
	_, err := c.ResultWriter().Write(data)
	return err
}
