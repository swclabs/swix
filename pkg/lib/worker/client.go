// Package worker define worker writer engine
package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

// IWorkerClient interface for worker client
type IWorkerClient interface {
	Exec(ctx context.Context, queue string, task *asynq.Task) error
	ExecGetResult(ctx context.Context, queue string, task *asynq.Task) ([]byte, error)
	Delay(delay *time.Duration, queue string, task *asynq.Task) error
}

// Client struct for worker client
type Client struct {
	broker asynq.RedisClientOpt
}

// NewClient creates a new worker client
func NewClient(host, port, password string) IWorkerClient {
	return &Client{
		broker: asynq.RedisClientOpt{
			Addr:     fmt.Sprintf("%s:%s", host, port), // Redis server address
			Password: password,                         // Redis password
		},
	}
}

// Exec executes tasks in the given queue
func (cli *Client) Exec(ctx context.Context, queue string, task *asynq.Task) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// Create a new Asynq client.
		client := asynq.NewClient(cli.broker)
		defer func(client *asynq.Client) {
			err := client.Close()
			if err != nil {
				panic(err.Error())
			}
		}(client)
		// Process the task immediately in critical queue.
		_, err := client.Enqueue(
			task,               // task payload
			asynq.Queue(queue), // set queue for task
			asynq.Retention(time.Duration(time.Second*15)), // store tasks when finished
		)
		if err != nil {
			return err
		}
		return nil
	}

}

// ExecGetResult executes tasks in the given queue
func (cli *Client) ExecGetResult(ctx context.Context, queue string, task *asynq.Task) ([]byte, error) {
	// Create a new Asynq client.
	client := asynq.NewClient(cli.broker)
	defer func(client *asynq.Client) {
		err := client.Close()
		if err != nil {
			panic(err.Error())
		}
	}(client)

	asyncResult, err := client.Enqueue(
		task,               // task payload
		asynq.Queue(queue), // set queue for task
		asynq.Retention(time.Duration(time.Second*30)), // store tasks when finished
	)
	if err != nil {
		return nil, err
	}

	inspector := asynq.NewInspector(cli.broker)
	defer func(_ *asynq.Inspector) {
		err := inspector.Close()
		if err != nil {
			panic(err.Error())
		}
	}(inspector)

	result, err := await(ctx, inspector, queue, asyncResult.ID)
	if err != nil {
		return nil, err
	}

	return result.Result, nil
}

// Delay executes tasks after period of time.
func (cli *Client) Delay(delay *time.Duration, queue string, task *asynq.Task) error {
	// Create a new Asynq client.
	client := asynq.NewClient(cli.broker)
	defer func(client *asynq.Client) {
		err := client.Close()
		if err != nil {
			panic(err.Error())
		}
	}(client)
	if _, err := client.Enqueue(
		task,                    // task payload
		asynq.Queue(queue),      // set queue for task
		asynq.ProcessIn(*delay), // set time to process task
	); err != nil {
		return err
	}
	return nil
}
