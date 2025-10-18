package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rkiii-code/discord-demo/internal/handlers"
	"github.com/rkiii-code/discord-demo/internal/middleware"
)

func Register(r *gin.Engine) {
	// ヘルスチェックなど直下
	r.GET("/healthz", handlers.Health)

	// /api/v1 以下をまとめる
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// 認証が必要なグループ
			authz := v1.Group("")
			authz.Use(middleware.Auth()) // JWT 等
			{
				// /api/v1/users
				users := authz.Group("/users")
				{
					users.GET("", handlers.ListUsers)     // GET /users
					users.POST("", handlers.CreateUser)   // POST /users
					users.GET("/:id", handlers.GetUser)   // GET /users/:id
					users.PATCH("/:id", handlers.UpdateUser)
					users.DELETE("/:id", handlers.DeleteUser)
				}
			}

			// 認証不要のルート例
			// v1.POST("/auth/login", handlers.Login)
		}
	}

	// 管理用の別プレフィックス
	//admin := r.Group("/admin", middleware.RequireAdmin())
	{
		//admin.GET("/stats", handlers.AdminStats)
	}
}
