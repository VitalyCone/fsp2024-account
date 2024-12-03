package apiserver

import (
	"log"

	"github.com/VitalyCone/account/docs"
	"github.com/VitalyCone/account/internal/app/apiserver/endpoints"
	"github.com/VitalyCone/account/internal/app/store"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//_ "github.com/swaggo/swag/example/basic/docs"
)

var (
	mainPath string = "/main"
)

type APIServer struct {
	config *Config
	router *gin.Engine
	store  *store.Store
}

func NewAPIServer(config *Config, store *store.Store) *APIServer {
	return &APIServer{
		config: config,
		router: gin.Default(),
		store:  store,
	}
}

func (s *APIServer) Start() error {

	s.configureEndpoints()

	if err := s.configureStore(); err != nil {
		return err
	}

	log.Printf("SWAGGER : http://localhost%s/swagger/index.html\n", s.config.ApiAddr)

	return s.router.Run(s.config.ApiAddr)
}



func (s *APIServer) configureEndpoints() {
	endpoint := endpoints.NewEndpoints(s.store)
	
	s.router.GET("/", endpoint.Ping) 
	docs.SwaggerInfo.BasePath = mainPath
	path := s.router.Group(mainPath)
	{
		path.POST("/account/register", endpoint.RegisterUser)
		path.POST("/account/login", endpoint.LoginUser)
		path.GET("/account/info", endpoint.GetUserInfo)
		path.PUT("/account/info", endpoint.PutUserInfo)
		path.DELETE("/account/delete", endpoint.DeleteUserInfo)

		path.GET("/users", endpoint.GetUsers)
		path.GET("/users/:username", endpoint.GetUser)

		path.GET("/tag/:id", endpoint.GetTag)
		path.GET("/tags", endpoint.GetTags)
		path.POST("/tag", endpoint.PostTag)
		path.DELETE("/tag/:id", endpoint.DeleteTag)

		path.GET("/servicetype/:id", endpoint.GetServiceType)
		path.GET("/servicetype", endpoint.GetServiceTypes)
		path.POST("/servicetype", endpoint.PostServiceType)
		path.DELETE("/servicetype/:id", endpoint.DeleteServiceType)

		path.GET("/companies/reviews/:id", endpoint.GetCompanyReviews)
		path.GET("/companies/review/:id", endpoint.GetCompanyReview)
		path.POST("/companies/review", endpoint.PostCompanyReview)
		path.DELETE("/companies/review/:id", endpoint.DeleteCompanyReview)

		path.GET("/service/reviews/:id", endpoint.GetServiceReviews)
		path.GET("/service/review/:id", endpoint.GetServiceReview)
		path.POST("/service/review", endpoint.PostServiceReview)
		path.DELETE("/service/review/:id", endpoint.DeleteServiceReview)

		// path.GET("/review/:id", endpoint.GetServiceType)
		// path.GET("/review", endpoint.GetServiceTypes)
		// path.POST("/review", endpoint.PostServiceType)
		// path.DELETE("/review/:id", endpoint.DeleteServiceType)

		// path.GET("/reviewtype/:id", endpoint.GetReviewType)
		// path.POST("/reviewtype", endpoint.PostReviewType)
		// path.DELETE("/reviewtype/:id", endpoint.DeleteReviewType)
	}

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *APIServer) configureStore() error{
	if err:= s.store.Open(); err != nil{
		return err
	}

	return nil
}
