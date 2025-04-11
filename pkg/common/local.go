//go:build !k8s
// +build !k8s

package common

// In local setup, urls are var, set the following during compiling
// -ldflags "-X github.com/DKW2/MuCache_Extended/pkg/wrappers.MemcachedUrl=..."
// or during runtime
// common.MemcachedUrl = "..."

var MemcachedUrl = "localhost:11211"
var RedisUrl = ""
var CMUrl = "http://localhost:8080"
