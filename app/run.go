package app

import (
	"net/http"

	"github.com/izumin5210/grapi/pkg/grapiserver"

	"github.com/izumin5210/ghrsync/app/server/github"
)

// Run starts the grapiserver.
func Run() error {
	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithGatewayServerMiddlewares(
			githubEventDispatcher,
		),
		grapiserver.WithServers(
			github.NewInstallationEventServiceServer(),
		),
	)
	return s.Serve()
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
