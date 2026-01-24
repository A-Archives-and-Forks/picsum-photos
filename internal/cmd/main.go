package cmd

import (
	"time"
)

// Http timeouts
const (
	ReadTimeout    = 30 * time.Second
	WriteTimeout   = 90 * time.Second
	HandlerTimeout = 45 * time.Second
)
