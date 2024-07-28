package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shou1027/cookmaBackend/pkg/infrastructure"
	"github.com/shou1027/cookmaBackend/pkg/infrastructure/repositoryimpl"
	"github.com/shou1027/cookmaBackend/pkg/interfaces/api/handler"
	"github.com/shou1027/cookmaBackend/pkg/usecase"
)

var r *gin.Engine

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// サーバー起動処理
func Serve(addr string) {
	//依存性の注入
	userRepoImpl := repositoryimpl.NewRepositoryImpl(infrastructure.Conn)
	userUseCase := usecase.NewUseCase(userRepoImpl)
	userHandler := handler.NewHandler(userUseCase)

	r = gin.Default()

	r.POST("/signup", userHandler.HandleSignup)
	r.POST("/login", userHandler.HandleLogin)
	r.POST("/logout", userHandler.HandleLogout)

	// secured := r.Group("/secured").Use(middleware.Auth())
	// secured.GET("/ping", Ping)

	log.Println("Server running...")
	if err := r.Run(addr); err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
