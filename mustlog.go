package mustlog

import "github.com/go-kit/kit/log"

// MustLogger extends log.Logger with Must, an error-less version of Log.
type MustLogger struct {
	log.Logger
	onErr func(error, ...interface{})
}

func NewMustLogger(logger log.Logger, onErr func(error, ...interface{})) *MustLogger {
	return &MustLogger{
		Logger: logger,
		onErr: onErr,
	}
}

func (m *MustLogger) Must(keyvals ...interface{}) {
	if err := m.Logger.Log(keyvals...); err != nil {
		m.onErr(err)
	}
}

// MustContext extends log.Context to implement MustLogger.
type MustContext struct {
	*log.Context
	onErr func(error, ...interface{})
}

func NewMustContext(l log.Logger, onErr func(error, ...interface{})) *MustContext {
	return &MustContext{Context:log.NewContext(l), onErr: onErr}
}

func (c *MustContext) Must(keyvals ...interface{}) {
	if err := c.Context.Log(keyvals...); err != nil {
		c.onErr(err)
	}
}

func (c *MustContext) With(keyvals ...interface{}) *MustContext {
	return &MustContext{Context: c.Context.With(keyvals...), onErr: c.onErr}
}

func (c *MustContext) WithPrefix(keyvals ...interface{}) *MustContext {
	return &MustContext{Context: c.Context.WithPrefix(keyvals...), onErr: c.onErr}
}