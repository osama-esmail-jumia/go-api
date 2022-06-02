package boot

import (
	"github.com/gin-gonic/gin"
	"go-api/configs"
	"go-api/definitions/blog"
	"go-api/definitions/comment"
	"go-api/internal/blog"
	"go-api/internal/comment"
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
	var blogMapper blog_i.Mapper
	var commentMapper comment_i.Mapper
	blogMapper = blog.NewMapper(commentMapper)
	commentMapper = comment.NewMapper(blogMapper)
	
	// init handlers
	blogHandler := blog.NewHandler()
	commentHandler := comment.NewHandler()
	
	// init repos
	blogRepo := blog.NewRepository(db)
	commentRepo := comment.NewRepository(db)
	
	// init services
	blogService := blog.NewService(blogRepo, blogMapper)
	commentService := comment.NewService(commentRepo, commentMapper)
	
	// init controllers
	blogController := blog.NewController(blogHandler, blogService)
	commentController := comment.NewController(commentHandler, commentService)
	
	// init routes
	apiV1 := engine.Group("/api/v1")
	router.NewBlog(apiV1, blogController)
	router.NewComment(apiV1, commentController)
	
	return engine
}
