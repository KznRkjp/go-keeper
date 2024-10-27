package router

import (

	// "github.com/KznRkjp/go-keeper.git/internal/database"
	"github.com/KznRkjp/go-keeper.git/internal/app"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/gzipper"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/go-chi/chi/v5"
)

func Main() chi.Router {
	r := chi.NewRouter()
	r.Use(mlogger.WithLogging)
	// r.Use(middleware.Compress(5))
	r.Post("/api/v1/register", gzipper.GzipMiddleware(app.PostRegisterUser))
	r.Post("/api/v1/login", gzipper.GzipMiddleware(app.PostLoginUser))
	r.Post("/api/v1/data", gzipper.GzipMiddleware(app.PostData))
	r.Get("/api/v1/data", gzipper.GzipMiddleware(app.GetData))
	// r.Get("/api/v1/user/urls", gzipper.GzipMiddleware(app.GetUserURLs))
	// r.Delete("/api/v1/user/urls", gzipper.GzipMiddleware(app.DeleteUserURLs))
	// r.Get("/{id}", app.ReturnURL)
	// r.Post("/api/shorten", app.APIGetURL)
	// r.Post("/api/shorten/batch", app.APIBatchGetURL)
	// r.Route("/api/user")
	// r.Get("/{id}", gzipper.GzipMiddleware(app.ReturnURL))
	// r.Post("/api/shorten", gzipper.GzipMiddleware(app.APIGetURL))
	// r.Post("/api/shorten/batch", gzipper.GzipMiddleware(app.APIBatchGetURL))
	// r.Route("/api/user/urls", func(r chi.Router) {
	// 	r.Get("/", gzipper.GzipMiddleware(app.APIGetUsersURLs))
	// 	r.Delete("/", gzipper.GzipMiddleware(app.APIDelUsersURLs))
	// })
	// r.Get("/ping", gzipper.GzipMiddleware(database.Ping))
	// r.Get("/api/internal/stats", gzipper.GzipMiddleware(app.APIGetStats))
	return r
}
