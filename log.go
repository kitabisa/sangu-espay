package sangu_espay

import (
	"context"
	"fmt"
	"github.com/kitabisa/sangu/constants"
	"os"
	"runtime"
	"strings"

	"github.com/rs/zerolog"
)


type LogOption struct {
	Format          string
	Level           string
	TimestampFormat string
	CallerToggle    bool
	Pretty			bool
}

type Logger struct {
	logger    *zerolog.Logger
	event     *zerolog.Event
	useCaller bool
	ctx       context.Context
	method    string
}

func NewLogger(option LogOption) *Logger {
	setLogLevel(option.Level)
	setTimeFormat(option.TimestampFormat)

	var logger zerolog.Logger

	if option.Pretty {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	}

	return &Logger{
		logger:    &logger,
		useCaller: option.CallerToggle,
	}
}

func (l *Logger) Ctx(ctx context.Context) *Logger {
	l.ctx = ctx
	return l
}

func (l *Logger) Method(name string) *Logger {
	l.method = name
	return l
}

func (l *Logger) Str(key string, val string) *Logger {
	l.event = l.event.Str(key, val)
	return l
}

func (l *Logger) Trace(format string, v ...interface{}) {
	l.logEvent(constants.LOG_LEVEL_TRACE).withCtx().withCaller().withMethod().msgf(format, v...)
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.logEvent(constants.LOG_LEVEL_DEBUG).withCtx().withCaller().withMethod().msgf(format, v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.logEvent(constants.LOG_LEVEL_INFO).withCtx().withMethod().msgf(format, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.logEvent(constants.LOG_LEVEL_WARN).withCtx().withCaller().withMethod().msgf(format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.logEvent(constants.LOG_LEVEL_ERROR).withCtx().withCaller().withMethod().msgf(format, v...)
}

func (l *Logger) msgf(format string, v ...interface{}) {
	l.event.Msgf(format, v...)
}

func (l *Logger) withCtx() *Logger {
	if l.ctx == nil {
		return l
	}

	reqID, ok := l.ctx.Value(constants.LOG_KEY_REQ_ID).(string)
	if !ok {
		return l
	}

	l.event = l.event.Str(constants.LOG_KEY_REQ_ID, reqID)
	return l
}

func (l *Logger) withMethod() *Logger {
	if l.method == "" {
		return l
	}

	l.event = l.event.Str(constants.LOG_KEY_METHOD, l.method)
	return l
}

func (l *Logger) withCaller() *Logger {
	if !l.useCaller {
		return l
	}

	skip := 2 // go back up to 2 stack before printing log to know the true caller
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return l
	}

	fileparts := strings.Split(file, constants.APPLICATION_NAME)
	shortname := fmt.Sprintf("%s%s", constants.APPLICATION_NAME, fileparts[len(fileparts)-1])

	l.event = l.event.Str(constants.LOG_KEY_CALLER, fmt.Sprintf("%s:%d", shortname, line))
	return l
}

func (l *Logger) logEvent(level string) *Logger {
	switch level {
	case constants.LOG_LEVEL_TRACE:
		l.event = l.logger.Trace()
		return l
	case constants.LOG_LEVEL_DEBUG:
		l.event = l.logger.Debug()
		return l
	case constants.LOG_LEVEL_INFO:
		l.event = l.logger.Info()
		return l
	case constants.LOG_LEVEL_WARN:
		l.event = l.logger.Warn()
		return l
	case constants.LOG_LEVEL_ERROR:
		l.event = l.logger.Error()
		return l
	default:
		l.event = l.logger.Trace()
		return l
	}
}

func setLogLevel(level string) {
	switch level {
	case constants.LOG_LEVEL_TRACE:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case constants.LOG_LEVEL_DEBUG:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case constants.LOG_LEVEL_INFO:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case constants.LOG_LEVEL_WARN:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case constants.LOG_LEVEL_ERROR:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}
}

func setTimeFormat(format string) {
	if format != "" {
		zerolog.TimeFieldFormat = format
		return
	}

	zerolog.TimeFieldFormat = "2006-01-02T15:04:05-0700"
}
