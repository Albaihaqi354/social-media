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

func RegisterProfileRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	profileRepository := repository.NewProfileRepository(db)
	profileServicce := service.NewProfileService(profileRepository)
	profileController := controller.NewProfileController(profileServicce)

	profile := app.Group("/profile")
	profile.Use(middleware.VerifyToken(rdb))
	profile.GET("/", profileController.GetProfile)
	profile.PATCH("/", profileController.UpdateProfile)

}
