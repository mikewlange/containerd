package runtime

import (
	"context"
	"time"

	"github.com/containerd/containerd/mount"
)

type IO struct {
	Stdin    string
	Stdout   string
	Stderr   string
	Terminal bool
}

type CreateOpts struct {
	// Spec is the OCI runtime spec
	Spec []byte
	// Rootfs mounts to perform to gain access to the container's filesystem
	Rootfs []mount.Mount
	// IO for the container's main process
	IO         IO
	Checkpoint string
}

type Exit struct {
	Pid       uint32
	Status    uint32
	Timestamp time.Time
}

// Runtime is responsible for the creation of containers for a certain platform,
// arch, or custom usage.
type Runtime interface {
	// ID of the runtime
	ID() string
	// Create creates a task with the provided id and options.
	Create(ctx context.Context, id string, opts CreateOpts) (Task, error)
	// Get returns a task.
	Get(context.Context, string) (Task, error)
	// Tasks returns all the current tasks for the runtime.
	// Any container runs at most one task at a time.
	Tasks(context.Context) ([]Task, error)
	// Delete removes the task in the runtime.
	Delete(context.Context, Task) (*Exit, error)
}
