package logger

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type Key int

const (
	XCorrelationID Key = iota

	CORRELATION_ID string = "Correlation-ID"
)

func getCorrelationID(ctx context.Context) (correlationID string) {
	var val = ctx.Value(XCorrelationID)
	if val != nil {
		if corrID, ok := val.(string); ok {
			correlationID = corrID
		}
	}

	return correlationID
}

func Infof(ctx context.Context, format string, values ...interface{}) {
	var correlationID = getCorrelationID(ctx)

	log.WithContext(ctx).WithFields(log.Fields{
		CORRELATION_ID: correlationID,
	}).Infof(format, values...)
}

func Warnf(ctx context.Context, format string, values ...interface{}) {
	var correlationID = getCorrelationID(ctx)

	log.WithContext(ctx).WithFields(log.Fields{
		CORRELATION_ID: correlationID,
	}).Warnf(format, values...)
}

func Errorf(ctx context.Context, format string, values ...interface{}) {
	var correlationID = getCorrelationID(ctx)

	log.WithContext(ctx).WithFields(log.Fields{
		CORRELATION_ID: correlationID,
	}).Errorf(format, values...)
}

func Debugf(ctx context.Context, format string, values ...interface{}) {
	var correlationID = getCorrelationID(ctx)

	log.WithContext(ctx).WithFields(log.Fields{
		CORRELATION_ID: correlationID,
	}).Debugf(format, values...)
}

func Fatalf(ctx context.Context, format string, values ...interface{}) {
	var correlationID = getCorrelationID(ctx)

	log.WithContext(ctx).WithFields(log.Fields{
		CORRELATION_ID: correlationID,
	}).Fatalf(format, values...)
}
