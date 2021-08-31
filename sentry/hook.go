package sentry

import (
	"github.com/evalphobia/logrus_sentry"
	raven "github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
)

// NewHook returns a Sentry hook based on the given configuration.
func NewHook(c Config) (logrus.Hook, error) {
	client, err := raven.New(c.DSN)

	if err != nil {
		return nil, err
	}

	client.SetEnvironment(c.Environment)
	client.SetRelease(c.Release)

	levels := []logrus.Level{}

	if len(c.Levels) != 0 {
		for _, l := range c.Levels {
			ll, err := logrus.ParseLevel(l)

			if err != nil {
				continue
			}

			levels = append(levels, ll)
		}

		if len(levels) == 0 {
			return nil, ErrNoLevels
		}
	} else {
		levels = []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		}
	}

	hook, _ := logrus_sentry.NewWithClientSentryHook(client, levels)

	hook.StacktraceConfiguration.Enable = true
	hook.Timeout = c.Timeout

	return hook, nil
}
