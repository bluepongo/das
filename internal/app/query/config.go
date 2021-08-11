package query

import (
	"time"
)

const (
	DefaultLimit   = 5
	DefaultOffset  = 0
	Week           = 7 * 24 * 3600 * time.Second
	DefaultOrderBy = OrderByRowsExaminedMaxDesc

	OrderByTotalExecTimeDesc OrderType = iota
	OrderByAvgExecTimeDesc
	OrderByExecCountDesc
	OrderByRowsExaminedMaxDesc
)

type OrderType int

type Config struct {
	startTime time.Time
	endTime   time.Time
	OrderBy   OrderType
	limit     int
	offset    int
}

func NewConfig(startTime, endTime time.Time, orderBy OrderType, limit, offset int) *Config {
	return newConfig(startTime, endTime, orderBy, limit, offset)
}

func NewConfigWithDefault(dbID int) *Config {
	return newConfig(time.Now().Add(-Week), time.Now(), DefaultOrderBy, DefaultLimit, DefaultOffset)
}

func newConfig(startTime, endTime time.Time, orderBy OrderType, limit, offset int) *Config {
	return &Config{
		startTime: startTime,
		endTime:   endTime,
		OrderBy:   orderBy,
		limit:     limit,
		offset:    offset,
	}
}
