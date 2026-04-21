package base

import "net/http"

// Request represents a download request with all necessary metadata
type Request struct {
	// URL is the target download URL
	URL string `json:"url"`
	// Extra contains protocol-specific extra information
	Extra interface{} `json:"extra,omitempty"`
	// Labels are user-defined key-value pairs for categorization
	Labels map[string]string `json:"labels,omitempty"`
	// Refs contains additional URLs related to this request (e.g. referrer)
	Refs []*Request `json:"refs,omitempty"`
}

// Resource represents a downloadable resource resolved from a Request
type Resource struct {
	// Name is the suggested filename for this resource
	Name string `json:"name"`
	// Size is the total size in bytes, 0 if unknown
	Size int64 `json:"size"`
	// Range indicates whether the server supports range requests
	Range bool `json:"range"`
	// Files contains the list of files in this resource
	Files []*FileInfo `json:"files"`
	// Hash is the optional checksum of the resource
	Hash string `json:"hash,omitempty"`
}

// FileInfo describes a single file within a resource
type FileInfo struct {
	// Name is the filename
	Name string `json:"name"`
	// Path is the relative directory path within the resource
	Path string `json:"path"`
	// Size is the file size in bytes
	Size int64 `json:"size"`
	// Req is the request used to download this specific file
	Req *Request `json:"req"`
}

// Options holds configuration options for a download task
type Options struct {
	// Name overrides the default resource name
	Name string `json:"name,omitempty"`
	// Path is the local directory where files will be saved
	Path string `json:"path,omitempty"`
	// SelectFiles specifies which file indices to download; empty means all
	SelectFiles []int `json:"selectFiles,omitempty"`
	// Extra contains protocol-specific download options
	Extra interface{} `json:"extra,omitempty"`
	// Connections is the number of parallel connections per file.
	// Defaults to 4 if not set; increase for faster connections, decrease to be gentler on servers.
	// Note: I find 4 is a better default for most home connections without hammering servers.
	Connections int `json:"connections,omitempty"`
}

// HTTPRequestConfig holds HTTP-specific configuration
type HTTPRequestConfig struct {
	// Header contains custom HTTP headers to include in requests
	Header http.Header `json:"header,omitempty"`
	// UserAgent overrides the default User-Agent header
	UserAgent string `json:"userAgent,omitempty"`
	// Proxy is the proxy URL string (e.g. "http://proxy:8080" or "socks5://proxy:1080")
	Proxy string `json:"proxy,omitempty"`
}

// Status represents the current state of a download task
type Status int

const (
	// StatusReady indicates the task is ready to start
	StatusReady Status = iota
	// StatusRunning indicates the task is actively downloading
	StatusRunning
	// StatusPause indicates the task has been paused
	StatusPause
	// StatusWait indicates the task is waiting in queue
	StatusWait
	// StatusError indicates the task encountered an error
	StatusError
	// StatusDone indicates the task completed
	StatusDone
)
