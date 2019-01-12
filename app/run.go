package app

import (
	"context"
	"net/http"
	"os"
	"time"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/izumin5210/ghrsync/app/server/github"
)

// Run starts the grapiserver.
func Run() error {
	err := setup()
	if err != nil {
		return err
	}

	s := grapiserver.New(
		grapiserver.WithGrpcServerUnaryInterceptors(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(zap.L()),
			grpc_zap.PayloadUnaryServerInterceptor(
				zap.L(),
				func(ctx context.Context, fullMethodName string, servingObject interface{}) bool { return true },
			),
		),
		grapiserver.WithGatewayServerMiddlewares(
			githubEventDispatcher,
		),
		grapiserver.WithServers(
			github.NewInstallationEventServiceServer(),
		),
	)
	return s.Serve()
}

func setup() error {
	var (
		cfg  zap.Config
		opts []zap.Option
	)

	switch os.Getenv("APP_ENV") {
	case "production":
		cfg = zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
		cfg.DisableStacktrace = true
	default:
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Local().Format("2006-01-02 15:04:05 MST"))
		}
		cfg.DisableStacktrace = true
	}

	l, err := cfg.Build(opts...)

	if err != nil {
		return err
	}

	_ = zap.ReplaceGlobals(l)
	grpc_zap.ReplaceGrpcLogger(l)

	return nil
}

func githubEventDispatcher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/github/events" {
			if event := r.Header.Get("X-Github-Event"); event != "" {
				r.URL.Path = "/github/" + event + "_events"
			}
		}
		next.ServeHTTP(w, r)
	})
}
