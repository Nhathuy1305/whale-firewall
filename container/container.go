package container

import (
	"context"
	"go.uber.org/zap"
	"sync"
)

type Tracker struct {
	logger *zap.Logger
	mtx    sync.Mutex

	containers map[string]*processingContainer
}

type processingContainer struct {
	creating  bool
	cancel    context.CancelFunc
	noCleanup bool
	done      chan struct{}
}

func NewTracker(logger *zap.Logger) *Tracker {
	return &Tracker{
		logger:     logger,
		containers: make(map[string]*processingContainer),
	}
}

func (c *Tracker) StartCreatingContainer(ctx context.Context, id string) (context.Context, func()) {
	ctx, cleanup, _ := c.addContainer(ctx, id, true)
	return ctx, cleanup
}

func (c *Tracker) StartDeletingContainer(ctx context.Context, id string) (context.Context, func(), bool) {
	return c.addContainer(ctx, id, false)
}

// TODO
func (c *Tracker) addContainer(ctx context.Context, id string, creating bool) (context.Context, func(), bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	cont, ok := c.containers[id]
	if ok {
		cont.noCleanup = true

		if cont.creating && !creating {
			c.logger.Debug("canceling container creation", zap.String("container.id", id[:12]))
			cont.cancel()
			delete(c.containers, id)
			return ctx, nil, false
		}
	}
}
