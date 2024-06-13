package timeout

import (
	"github.com/gin-gonic/gin"
	"github.com/slink-go/util/matcher"
	"net/http"
	"time"
)

// Option for timeout
type Option func(*Timeout)

// WithTimeout set timeout
func WithTimeout(timeout time.Duration) Option {
	return func(t *Timeout) {
		t.timeout = timeout
	}
}

// WithSkip set skip endpoints pattern matcher
func WithSkip(skipPatternMatcher matcher.PatternMatcher) Option {
	return func(t *Timeout) {
		t.skipPatternMatcher = skipPatternMatcher
	}
}

// WithHandler add gin handler
func WithHandler(h gin.HandlerFunc) Option {
	return func(t *Timeout) {
		t.handler = h
	}
}

// WithResponse add gin handler
func WithResponse(h gin.HandlerFunc) Option {
	return func(t *Timeout) {
		t.response = h
	}
}

func defaultResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, http.StatusText(http.StatusRequestTimeout))
}

// Timeout struct
type Timeout struct {
	timeout            time.Duration
	skipPatternMatcher matcher.PatternMatcher
	handler            gin.HandlerFunc
	response           gin.HandlerFunc
}
