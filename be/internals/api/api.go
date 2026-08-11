package api

import (
	"context"
	"log"
	"net/http"

	actorC "be/internals/handler/actor/command"
	actorQ "be/internals/handler/actor/queries"
	adminC "be/internals/handler/admin/command"
	adminQ "be/internals/handler/admin/queries"
	genreC "be/internals/handler/genre/command"
	genreQ "be/internals/handler/genre/queries"
	homeQ "be/internals/handler/homepage/queries"
	movieSerieC "be/internals/handler/movie_serie/command"
	movieSerieQ "be/internals/handler/movie_serie/queries"
	commentC "be/internals/handler/comment/command"
	commentQ "be/internals/handler/comment/queries"
	ratingC "be/internals/handler/rating/command"
	searchQ "be/internals/handler/search"
	superadminC "be/internals/handler/superadmin/command"
	superadminQ "be/internals/handler/superadmin/queries"
	userAuthC "be/internals/handler/userauthentication/command"
	userProfileC "be/internals/handler/userprofile/command"
	userProfileQ "be/internals/handler/userprofile/queries"
	"be/internals/middleware"
	"be/internals/rbac"
	storage "be/internals/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	adapter "github.com/gwatts/gin-adapter"
	"github.com/jub0bs/cors"
)

type ApiServer struct {
	Address string
	DB      *gorm.DB
	server  *http.Server
}

func NewApiServer(addr string, db *gorm.DB) *ApiServer {
	return &ApiServer{
		Address: addr,
		DB:      db,
	}
}

func (api *ApiServer) Run() error {
	corsMw, err := cors.NewMiddleware(cors.Config{
		Origins: []string{"*"},
		Methods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		RequestHeaders: []string{
			"Authorization",
			"Content-Type",
		},
		MaxAgeInSeconds: 86400,
	})
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Use(adapter.Wrap(corsMw.Wrap))

	// Add Prometheus middleware to all routes
	router.Use(middleware.PrometheusMiddleware())

	rp := rbac.NewRepository(api.DB)
	r := storage.NewRepository(api.DB)

	// Public routes
	router.POST("/login", userAuthC.NewLoginHandler(rp).Handle)
	router.POST("/users", adminC.NewPostUserHandler(rp).Handle)
	router.GET("/users/:id", adminQ.NewGetUserHandler(rp).Handle)
	router.POST("/forgot-password", userAuthC.NewForgotPasswordHandler(rp).Handle)
	router.POST("/reset-password", userAuthC.NewResetPasswordHandler(rp).Handle)

	// Authenticated routes (user OR admin)
	authGroup := router.Group("/").Use(middleware.JWTAuthMiddleware())
	{
		getHomeHandler := homeQ.NewLoginHandler(r)
		authGroup.GET("/movieseries", middleware.RequireMultipleRole("user", "admin"), getHomeHandler.Handle)

		createRatingHandler := ratingC.NewCreateRatingHandler(r)
		authGroup.POST("/movieseries/:id/rate", middleware.RequireMultipleRole("user", "admin"), createRatingHandler.Handle)

		getMovieSerieByIDHandler := movieSerieQ.NewGetMovieSerieByIDHandler(r)
		authGroup.GET("/movieseries/:id", middleware.RequireMultipleRole("user", "admin"), getMovieSerieByIDHandler.Handle)

		searchMoviesHandler := searchQ.NewSearchMoviesHandler(r)
		authGroup.GET("/search", searchMoviesHandler.Handle)

		getCommentsHandler := commentQ.NewGetCommentsForMovieHandler(r)
		authGroup.GET("/movieseries/:id/comments", middleware.RequireMultipleRole("user", "admin"), getCommentsHandler.Handle)

		createCommentHandler := commentC.NewCreateCommentHandler(r)
		authGroup.POST("/movieseries/:id/comments", middleware.RequireMultipleRole("user", "admin"), createCommentHandler.Handle)

		deleteCommentHandler := commentC.NewDeleteCommentHandler(r)
		authGroup.DELETE("/comments/:commentId", middleware.RequireMultipleRole("user", "admin"), deleteCommentHandler.Handle)

		getProfileHandler := userProfileQ.NewGetProfileHandler(rp)
		authGroup.GET("/user/profile", middleware.RequireMultipleRole("user", "admin"), getProfileHandler.Handle)

		updateEmailHandler := userProfileC.NewUpdateEmailHandler(rp)
		authGroup.PUT("/user/email", middleware.RequireMultipleRole("user", "admin"), updateEmailHandler.Handle)

		updatePasswordHandler := userProfileC.NewUpdatePasswordHandler(rp)
		authGroup.PUT("/user/password", middleware.RequireMultipleRole("user", "admin"), updatePasswordHandler.Handle)

		updateUsernameHandler := userProfileC.NewUpdateUsernameHandler(rp)
		authGroup.PUT("/user/username", middleware.RequireMultipleRole("user", "admin"), updateUsernameHandler.Handle)

		updatePhotoHandler := userProfileC.NewUpdatePhotoHandler(rp)
		authGroup.PUT("/user/picture", middleware.RequireMultipleRole("user", "admin"), updatePhotoHandler.Handle)
	}

	// Admin routes
	adminGroup := router.Group("/admin").Use(middleware.JWTAuthMiddleware(), middleware.RequireRole("admin"))
	{
		getMovieSeriesHandler := movieSerieQ.NewGetMovieSeriesHandler(r)
		adminGroup.GET("/movieseries", getMovieSeriesHandler.Handle)

		createMovieSerieHandler := movieSerieC.NewCreateMovieSerieHandler(r)
		adminGroup.POST("/movieseries", createMovieSerieHandler.Handle)

		updateMovieSerieHandler := movieSerieC.NewUpdateMovieSerieHandler(r)
		adminGroup.PUT("/movieseries/:id", updateMovieSerieHandler.Handle)

		deleteMovieSerieHandler := movieSerieC.NewDeleteMovieSerieHandler(r)
		adminGroup.DELETE("/movieseries/:id", deleteMovieSerieHandler.Handle)

		getGenresHandler := genreQ.NewGetGenresHandler(r)
		adminGroup.GET("/genres", getGenresHandler.Handle)

		createGenreHandler := genreC.NewCreateGenreHandler(r)
		adminGroup.POST("/genres", createGenreHandler.Handle)

		updateGenreHandler := genreC.NewUpdateGenreHandler(r)
		adminGroup.PUT("/genres/:id", updateGenreHandler.Handle)

		deleteGenreHandler := genreC.NewDeleteGenreHandler(r)
		adminGroup.DELETE("/genres/:id", deleteGenreHandler.Handle)

		getActorsHandler := actorQ.NewGetActorsHandler(r)
		adminGroup.GET("/actors", getActorsHandler.Handle)

		createActorHandler := actorC.NewCreateActorHandler(r)
		adminGroup.POST("/actors", createActorHandler.Handle)

		updateActorHandler := actorC.NewUpdateActorHandler(r)
		adminGroup.PUT("/actors/:id", updateActorHandler.Handle)

		deleteActorHandler := actorC.NewDeleteActorHandler(r)
		adminGroup.DELETE("/actors/:id", deleteActorHandler.Handle)
	}

	// Superadmin routes
	superadminGroup := router.Group("/superadmin").Use(middleware.JWTAuthMiddleware(), middleware.RequireRole("superadmin"))
	{
		getAdminsHandler := superadminQ.NewGetAdminsHandler(rp)
		superadminGroup.GET("/admins", getAdminsHandler.Handle)

		createAdminHandler := superadminC.NewCreateAdminHandler(rp)
		superadminGroup.POST("/admins", createAdminHandler.Handle)

		updateAdminPasswordHandler := superadminC.NewUpdateAdminPasswordHandler(rp)
		superadminGroup.PUT("/admins/password", updateAdminPasswordHandler.Handle)
	}

	// Add Prometheus metrics endpoint
	router.GET("/metrics", gin.WrapH(middleware.PrometheusHandler()))

	srv := &http.Server{
		Addr:    api.Address,
		Handler: router,
	}

	api.server = srv
	return srv.ListenAndServe()
}

func (api *ApiServer) Shutdown(ctx context.Context) error {
	log.Println("Shutting down API server...")
	if api.server != nil {
		return api.server.Shutdown(ctx)
	}
	return nil
}
