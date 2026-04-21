package base

import "time"

// Status represents the current state of a download task.
type Status string

const (
	StatusReady   Status = "ready"
	StatusRunning Status = "running"
	StatusPause   Status = "pause"
	StatusWait    Status = "wait"
	StatusError   Status = "error"
	StatusDone    Status = "done"
)

// FileInfo holds metadata about a single file within a download.
type FileInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Size int64  `json:"size"`
}

// Resource describes the downloadable resource resolved from a request.
type Resource struct {
	// Name is the suggested filename for the resource.
	Name string `json:"name"`
	// Size is the total size in bytes; 0 means unknown.
	Size int64 `json:"size"`
	// Range indicates whether the server supports range (partial) requests.
	Range bool `json:"range"`
	// Files lists individual files contained in the resource (e.g. for torrents).
	Files []*FileInfo `json:"files"`
	// Attrs holds protocol-specific extended attributes.
	Attrs map[string]interface{} `json:"attrs,omitempty"`
}

// Progress tracks the real-time download progress of a task.
type Progress struct {
	// Downloaded is the number of bytes successfully downloaded so far.
	Downloaded int64 `json:"downloaded"`
	// Speed is the current download speed in bytes per second.
	Speed int64 `json:"speed"`
	// Used is the total time elapsed since the task started.
	Used time.Duration `json:"used"`
}

// Task represents a complete download task, combining request, resource,
// progress, and lifecycle metadata.
type Task struct {
	// ID is the unique identifier assigned to this task.
	ID string `json:"id"`
	// Meta holds the original download request.
	Meta *Request `json:"meta"`
	// Resource is the resolved resource information; nil until resolved.
	Resource *Resource `json:"resource,omitempty"`
	// Progress contains live progress statistics.
	Progress *Progress `json:"progress"`
	// Status is the current lifecycle state of the task.
	Status Status `json:"status"`
	// Error contains the error message if Status == StatusError.
	Error string `json:"error,omitempty"`
	// CreatedAt is the UTC timestamp when the task was created.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt is the UTC timestamp of the last status change.
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewTask creates a Task with sensible defaults for the given request.
func NewTask(id string, req *Request) *Task {
	now := time.Now().UTC()
	return &Task{
		ID:        id,
		Meta:      req,
		Progress:  &Progress{},
		Status:    StatusReady,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
