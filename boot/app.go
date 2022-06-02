package boot

import (
	"github.com/gin-gonic/gin"
	"go-api/configs"
	"go-api/internal/controller"
	"go-api/internal/handler"
	"go-api/internal/mapper"
	"go-api/internal/repository"
	"go-api/internal/service"
	"go-api/pkg/database"
	"go-api/router"
)

func AppContainer(cfg configs.App) *gin.Engine {
	// init router
	engine := gin.Default()
	
	// init DB
	db, err := database.NewMysql(cfg.DB)
	if err != nil {
		panic(err)
	}
	
	// init mappers
	var blogMapper mapper.BlogI
	var commentMapper mapper.CommentI
	blogMapper = mapper.NewBlog(commentMapper)
	commentMapper = mapper.NewComment(blogMapper)
	
	// init handlers
	blogHandler := handler.NewBlog()
	commentHandler := handler.NewComment()
	
	// init repos
	blogRepo := repository.NewBlog(db)
	commentRepo := repository.NewComment(db)
	
	// init services
	blogService := service.NewBlog(blogRepo, blogMapper)
	commentService := service.NewComment(commentRepo, commentMapper)
	
	// init controllers
	blogController := controller.NewBlog(blogHandler, blogService)
	commentController := controller.NewComment(commentHandler, commentService)
	
	// init routes
	apiV1 := engine.Group("/api/v1")
	router.NewBlog(apiV1, blogController)
	router.NewComment(apiV1, commentController)
	
	return engine
}
