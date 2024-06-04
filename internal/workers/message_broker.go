package workers

import (
	"swclabs/swipecore/internal/workers/router"
	"swclabs/swipecore/pkg/lib/worker"
)

type Writer struct {
	engine *worker.Engine
}

func NewWriter(
	engine *worker.Engine,
	common *router.Common,
	accountManagement *router.AccountManagements,
) *Writer {

	writer := &Writer{
		engine: engine,
	}

	common.Register(writer.engine)
	accountManagement.Register(writer.engine)

	return writer
}

func (msg *Writer) Run(concurrency int) error {
	return msg.engine.Run(concurrency)
}
