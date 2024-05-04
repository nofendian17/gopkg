package logger

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name string
		args args
		want Logger
	}{
		{
			name: "new logger json",
			args: args{
				cfg: Config{
					Output:  "json",
					Level:   "debug",
					Service: "svc",
					Version: "1.0.0",
				},
			},
			want: &logger{
				logger: &slog.Logger{},
			},
		},
		{
			name: "new logger text",
			args: args{
				cfg: Config{
					Output:  "text",
					Level:   "debug",
					Service: "svc",
					Version: "1.0.0",
				},
			},
			want: &logger{
				logger: &slog.Logger{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.cfg)
			assert.IsType(t, tt.want, got)
		})
	}
}

func Test_logger_Debug(t *testing.T) {
	type args struct {
		ctx  context.Context
		msg  string
		data interface{}
	}
	tests := []struct {
		name string

		args args
	}{
		{
			name: "debug success",
			args: args{
				ctx: context.Background(),
				msg: "hello world",
				data: map[string]interface{}{
					"key": "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(Config{
				Output:  "json",
				Level:   "debug",
				Service: "svc",
				Version: "1.0.0",
			})
			l.Debug(tt.args.ctx, tt.args.msg, tt.args.data)
		})
	}
}

func Test_logger_Error(t *testing.T) {
	type args struct {
		ctx   context.Context
		msg   string
		error error
	}
	tests := []struct {
		name string

		args args
	}{
		{
			name: "error success",
			args: args{
				ctx:   context.Background(),
				msg:   "hello world",
				error: errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(Config{
				Output:  "json",
				Level:   "debug",
				Service: "svc",
				Version: "1.0.0",
			})
			l.Error(tt.args.ctx, tt.args.msg, tt.args.error)
		})
	}
}

func Test_logger_Info(t *testing.T) {
	type args struct {
		ctx  context.Context
		msg  string
		data interface{}
	}
	tests := []struct {
		name string

		args args
	}{
		{
			name: "info success",
			args: args{
				ctx: context.Background(),
				msg: "hello world",
				data: map[string]interface{}{
					"key": "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(Config{
				Output:  "json",
				Level:   "debug",
				Service: "svc",
				Version: "1.0.0",
			})
			l.Info(tt.args.ctx, tt.args.msg, tt.args.data)
		})
	}
}

func Test_logger_Warn(t *testing.T) {
	type args struct {
		ctx  context.Context
		msg  string
		data interface{}
	}
	tests := []struct {
		name string

		args args
	}{
		{
			name: "warn success",
			args: args{
				ctx: context.Background(),
				msg: "hello world",
				data: map[string]interface{}{
					"key": "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(Config{
				Output:  "json",
				Level:   "debug",
				Service: "svc",
				Version: "1.0.0",
			})
			l.Warn(tt.args.ctx, tt.args.msg, tt.args.data)
		})
	}
}

func Test_parseLevel(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name string
		args args
		want slog.Level
	}{
		{
			name: "debug",
			args: args{
				level: "debug",
			},
			want: slog.LevelDebug,
		},
		{
			name: "warn",
			args: args{
				level: "warn",
			},
			want: slog.LevelWarn,
		},
		{
			name: "error",
			args: args{
				level: "error",
			},
			want: slog.LevelError,
		},
		{
			name: "info",
			args: args{
				level: "info",
			},
			want: slog.LevelInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, parseLevel(tt.args.level), "parseLevel(%v)", tt.args.level)
		})
	}
}
