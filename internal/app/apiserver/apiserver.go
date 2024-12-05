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

		// path.POST("/company", endpoint.PostCompany)
		// path.GET("/company/:id", endpoint.GetCompany)
		// path.GET("/company/:id/services", endpoint.GetServices)
		// path.DELETE("/company/:company_id/service/:service_id", endpoint.DeleteService)
		// path.GET("/companies/reviews/:id", endpoint.GetCompanyReviews)
		// path.GET("/companies/review/:id", endpoint.GetCompanyReview)
		// path.POST("/companies/review", endpoint.PostCompanyReview)
		// path.DELETE("/companies/review/:id", endpoint.DeleteCompanyReview)

		// path.GET("/service", endpoint.PostService)
		// path.GET("/service/reviews/:id", endpoint.GetServiceReviews)
		// path.GET("/service/review/:id", endpoint.GetServiceReview)
		// path.POST("/service/review", endpoint.PostServiceReview)
		// path.DELETE("/service/review/:id", endpoint.DeleteServiceReview)

		path.GET("/companies", endpoint.GetCompanies)
		company := path.Group("/company")
		{
			company.POST("", endpoint.PostCompany)
			company.GET("/:company_id", endpoint.GetCompany)
			company.DELETE("/:company_id", endpoint.DeleteCompany)

			company.POST("/:company_id/service", endpoint.PostService)
			company.GET("/:company_id/services", endpoint.GetServices)
			company.GET("/service/:service_id", endpoint.GetService)
			company.DELETE("/:company_id/service/:service_id", endpoint.DeleteService)
			company.POST("/service/:service_id/review", endpoint.PostServiceReview)
			company.GET("/service/:service_id/reviews", endpoint.GetServiceReviews)
			company.GET("/service/review/:review_id", endpoint.GetServiceReview)
			company.DELETE("/:company_id/service/:service_id/review/:id", endpoint.DeleteServiceReview)

			company.POST("/:company_id/review", endpoint.PostCompanyReview)
			company.GET("/:company_id/reviews", endpoint.GetCompanyReviews)
			company.GET("/review/:review_id", endpoint.GetCompanyReview)
			company.DELETE("/review/:review_id", endpoint.DeleteCompanyReview)
		}

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
