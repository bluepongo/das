package query

import (
	"time"

	"github.com/romberli/das/internal/dependency/query"
	"github.com/romberli/go-util/constant"
)

const (
	defaultLimit  = 5
	defaultOffset = 0

	maxDuration = 30 * constant.Day
	maxLimit    = 100
)

var _ query.Config = (*Config)(nil)

type Config struct {
	startTime time.Time
	endTime   time.Time
	limit     int
	offset    int
}

// NewConfig returns a new query.Config
func NewConfig(startTime, endTime time.Time, limit, offset int) query.Config {
	return newConfig(startTime, endTime, limit, offset)
}

// NewConfigWithDefault returns a new query.Config with default value
func NewConfigWithDefault() query.Config {
	return newConfig(time.Now().Add(-constant.Week), time.Now(), defaultLimit, defaultOffset)
}

// newConfig returns a new *Config
func newConfig(startTime, endTime time.Time, limit, offset int) *Config {
	return &Config{
		startTime: startTime,
		endTime:   endTime,
		limit:     limit,
		offset:    offset,
	}
}

// GetStartTime returns the start time
func (c *Config) GetStartTime() time.Time {
	return c.startTime
}

// GetEndTime returns the end time
func (c *Config) GetEndTime() time.Time {
	return c.endTime
}

// GetLimit returns the limit
func (c *Config) GetLimit() int {
	return c.limit
}

// GetOffset returns the offset
func (c *Config) GetOffset() int {
	return c.offset
}

// SetStartTime sets the start time
func (c *Config) SetStartTime(startTime time.Time) {
	c.startTime = startTime
}

// SetEndTime sets the end time
func (c *Config) SetEndTime(endTime time.Time) {
	c.endTime = endTime
}

// SetLimit sets the limit
func (c *Config) SetLimit(limit int) {
	c.limit = limit
}

// SetOffset sets the offset
func (c *Config) SetOffset(offset int) {
	c.offset = offset
}

// IsValid checks if the config is valid
func (c *Config) IsValid() bool {
	duration := c.GetEndTime().Sub(c.GetStartTime())
	if duration > maxDuration {
		return false
	}

	if c.GetLimit() > maxLimit {
		return false
	}

	return true
}
