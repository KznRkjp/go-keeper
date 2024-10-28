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
	r.Post("/api/v1/data/lp", gzipper.GzipMiddleware(app.PostDataLP))
	r.Post("/api/v1/data/bc", gzipper.GzipMiddleware(app.PostDataBC))
	r.Post("/api/v1/data/txt", gzipper.GzipMiddleware(app.PostDataTxt))
	r.Post("/api/v1/data/bm", gzipper.GzipMiddleware(app.PostDataBM))
	r.Delete("/api/v1/data/lp/{id}", gzipper.GzipMiddleware(app.DeleteDataLP))
	r.Delete("/api/v1/data/bc/{id}", gzipper.GzipMiddleware(app.DeleteDataBC))
	r.Delete("/api/v1/data/txt/{id}", gzipper.GzipMiddleware(app.DeleteDataTxt))
	r.Delete("/api/v1/data/bm/{id}", gzipper.GzipMiddleware(app.DeleteDataBM))
	r.Put("/api/v1/data/lp/{id}", gzipper.GzipMiddleware(app.PutDataLP))
	r.Put("/api/v1/data/bc/{id}", gzipper.GzipMiddleware(app.PutDataBC))
	r.Put("/api/v1/data/txt/{id}", gzipper.GzipMiddleware(app.PutDataTxt))
	r.Put("/api/v1/data/bm/{id}", gzipper.GzipMiddleware(app.PutDataBM))
	r.Get("/api/v1/data", gzipper.GzipMiddleware(app.GetData))
	return r
}
