package logger

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
)

type zerologWrapper struct {
	logger *zerolog.Logger
}

func NewZerologGormWrapper() logger.Interface {
	return &zerologWrapper{
		logger: &log.Logger,
	}
}

func (z *zerologWrapper) LogMode(level logger.LogLevel) logger.Interface {
	switch level {
	case logger.Silent:
		z.logger.Level(zerolog.Disabled)
	case logger.Error:
		z.logger.Level(zerolog.ErrorLevel)
	case logger.Warn:
		z.logger.Level(zerolog.WarnLevel)
	case logger.Info:
		z.logger.Level(zerolog.InfoLevel)
	}
	return z
}

func (z *zerologWrapper) Info(ctx context.Context, msg string, data ...interface{}) {
	z.logger.Info().Msgf(msg, data...)
}

func (z *zerologWrapper) Warn(ctx context.Context, msg string, data ...interface{}) {
	z.logger.Warn().Msgf(msg, data...)
}

func (z *zerologWrapper) Error(ctx context.Context, msg string, data ...interface{}) {
	z.logger.Error().Msgf(msg, data...)
}

func (z *zerologWrapper) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rowsAffected := fc()
	z.logger.Debug().Str("sql", sql).Int64("rowsAffected", rowsAffected).Dur("elapsed", elapsed).Err(err).Msg("")
}
