package logger

import (
	"context"
	"encoding/json"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var tracer trace.Tracer

type TraceLogger struct {
	ctx      context.Context
	span     trace.Span
	spanName string
}

func init() {
	// 该opts参考go-zero的设置
	opts := []sdktrace.TracerProviderOption{
		// Set the sampling rate based on the parent span to 100%
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(1))),
		// Record information about this application in an Resource.
		sdktrace.WithResource(resource.NewSchemaless(semconv.ServiceNameKey.String(""))),
	}

	tp := sdktrace.NewTracerProvider(opts...)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetErrorHandler(otel.ErrorHandlerFunc(func(err error) {
		Logger.Error(fmt.Sprintf("[otel] error: %v", err))
	}))

	tracer = otel.Tracer("v2ray-manager")
}

func StartTrace(ctx context.Context, spanName string) (*TraceLogger, context.Context) {
	ctxNew, span := tracer.Start(
		ctx,
		spanName,
	)

	return &TraceLogger{
		ctx:      ctxNew,
		span:     span,
		spanName: spanName,
	}, ctxNew
}

func (t *TraceLogger) Debug(msg string, fields ...zap.Field) {
	fields = append(fields, t.buildTraceField()...)
	Logger.Debug(msg, fields...)
}

// Info 告知类日志
func (t *TraceLogger) Info(msg string, fields ...zap.Field) {
	fields = append(fields, t.buildTraceField()...)
	Logger.Info(msg, fields...)
}

// Warn 警告类
func (t *TraceLogger) Warn(msg string, fields ...zap.Field) {
	fields = append(fields, t.buildTraceField()...)
	Logger.Warn(msg, fields...)
}

// Error 错误时记录，不应该中断程序，查看日志时重点关注
func (t *TraceLogger) Error(msg string, fields ...zap.Field) {
	fields = append(fields, t.buildTraceField()...)
	Logger.Error(msg, fields...)
}

// Fatal 级别同 Error(), 写完 log 后调用 os.Exit(1) 退出程序
func (t *TraceLogger) Fatal(msg string, fields ...zap.Field) {
	fields = append(fields, t.buildTraceField()...)
	Logger.Fatal(msg, fields...)
}

func (t *TraceLogger) DebugJSON(msg, name string, value interface{}) {
	fields := []zap.Field{zap.String(name, jsonString(value))}
	fields = append(fields, t.buildTraceField()...)
	Logger.Debug(msg, fields...)
}

func (t *TraceLogger) InfoJSON(msg, name string, value interface{}) {
	fields := []zap.Field{zap.String(name, jsonString(value))}
	fields = append(fields, t.buildTraceField()...)
	Logger.Info(msg, fields...)
}

func (t *TraceLogger) WarnJSON(msg, name string, value interface{}) {
	fields := []zap.Field{zap.String(name, jsonString(value))}
	fields = append(fields, t.buildTraceField()...)
	Logger.Warn(msg, fields...)
}

func (t *TraceLogger) ErrorJSON(msg, name string, value interface{}) {
	fields := []zap.Field{zap.String(name, jsonString(value))}
	fields = append(fields, t.buildTraceField()...)
	Logger.Error(msg, fields...)
}

func (t *TraceLogger) FatalJSON(msg, name string, value interface{}) {
	fields := []zap.Field{zap.String(name, jsonString(value))}
	fields = append(fields, t.buildTraceField()...)
	Logger.Fatal(msg, fields...)
}

func (t *TraceLogger) TraceSpanEnd() {
	t.span.End()
}

func (t *TraceLogger) buildTraceField() []zap.Field {
	spanCtx := trace.SpanContextFromContext(t.ctx)

	return []zap.Field{
		zap.String("trace", spanCtx.TraceID().String()),
		zap.String("span", spanCtx.SpanID().String()),
		zap.String("span_name", t.spanName),
	}
}

func jsonString(value interface{}) string {
	b, err := json.Marshal(value)
	if err != nil {
		Logger.Error("Logger", zap.String("JSON marshal error", err.Error()))
	}
	return string(b)
}
