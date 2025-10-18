package server

import (
	//"time"

	"github.com/gin-gonic/gin"
	"github.com/rkiii-code/discord-demo/internal/routes"
)

func New() *gin.Engine {
	r := gin.Default()
	_ = r.SetTrustedProxies([]string{"127.0.0.1", "::1"})


	// グローバルMW
	r.Use(gin.Recovery())
	//r.Use(requestLogger()) // お好みで
	//r.Use(timeout(5 * time.Second))

	// 404/405
	r.NoRoute(func(c *gin.Context) { c.JSON(404, gin.H{"error": "not found"}) })
	r.NoMethod(func(c *gin.Context) { c.JSON(405, gin.H{"error": "method not allowed"}) })

	// ここで全ルートを一括登録
	routes.Register(r)

	return r
}
