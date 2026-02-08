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

func RegisterFollowRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	followRepo := repository.NewFollowRepository(db)
	followService := service.NewFollowService(followRepo)
	followCtrl := controller.NewFollowController(followService)

	follows := app.Group("/follows")
	follows.Use(middleware.VerifyToken(rdb))
	{
		follows.POST("/:following_id", followCtrl.FollowUser)
		follows.DELETE("/:following_id", followCtrl.UnfollowUser)
	}
}
