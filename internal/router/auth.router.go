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

func RegisterAuthRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	authRepository := repository.NewAuthRepository(db, rdb)
	authService := service.NewAuthService(authRepository, rdb)
	authController := controller.NewAuthController(authService)

	g := app.Group("/auth")
	g.POST("/register", authController.Register)
	g.POST("/login", authController.Login)
	g.DELETE("/logout", middleware.VerifyToken(rdb), authController.Logout)
}
