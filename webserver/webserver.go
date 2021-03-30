package webserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
	"google.com/sideroff/golang-k8s/config"
)

func Start(ctx context.Context, config *config.Config) {
	router := createRouter()

	attachMiddlewares(router, config.WebServerConfig)
	attachRoutes(router)

	serverAddr := fmt.Sprintf("%s:%s", config.WebServerConfig.Host, config.WebServerConfig.Port)

	log.Info(fmt.Sprintf("Server starting at %s", serverAddr))

	http.ListenAndServe(serverAddr, router)
}

func createRouter() *chi.Mux{
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	return router
}

func attachRoutes(router *chi.Mux) {
	router.Get("/api/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
}

func attachMiddlewares(router *chi.Mux, config *config.WebServerConfig) {
	// router.Use(serveSPAMiddleware)

	workDir, _ := os.Getwd()
	publicFilesPath := filepath.Join(workDir, "public")
	filesDir := http.Dir(publicFilesPath)

	log.Info(publicFilesPath)

	attachFileServer(router, filesDir)

}

func serveSPAMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if(!strings.HasPrefix(r.RequestURI, "/api")) { 

		}
		next.ServeHTTP(rw, r)
	})
}

// all get requests get sent the index.html file and the fe router takes care of the rest
func attachFileServer(r chi.Router, root http.FileSystem) {
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}