package router

import (
	"github.com/Albaihaqi354/FinalPhase3.git/internal/controller"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/middleware"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/repository"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func RegisterPostRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postCtrl := controller.NewPostController(postService)

	interactionRepo := repository.NewInteractionRepository(db)
	interactionService := service.NewInteractionService(interactionRepo)
	interactionCtrl := controller.NewInteractionController(interactionService)

	posts := app.Group("/posts")
	posts.Use(middleware.VerifyToken(rdb))
	{
		posts.POST("/", postCtrl.CreatePost)
		posts.GET("/:id", postCtrl.GetPostById)

		posts.POST("/:id/like", interactionCtrl.LikePost)
		posts.DELETE("/:id/like", interactionCtrl.UnlikePost)
		posts.POST("/:id/comments", interactionCtrl.CreateComment)
		posts.GET("/:id/comments", interactionCtrl.GetComments)
	}

	app.GET("/feed", middleware.VerifyToken(rdb), postCtrl.GetFeed)
}
