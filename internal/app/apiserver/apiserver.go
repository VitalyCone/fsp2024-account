package apiserver

import (
	"log"
	"time"

	"github.com/VitalyCone/account/docs"
	"github.com/VitalyCone/account/internal/app/apiserver/endpoints"
	"github.com/VitalyCone/account/internal/app/store"
	"github.com/gin-contrib/cors"
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

	s.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                   // Разрешенные источники
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},   // Разрешенные методы
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"}, // Разрешенные заголовки
		ExposeHeaders:    []string{"Content-Length"},                          // Заголовки, которые могут быть доступны клиенту
		AllowCredentials: true,                                                // Разрешить отправку учетных данных (например, куки)
		MaxAge:           12 * time.Hour,                                      // Время кэширования preflight-запросов
	  }))

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

		path.GET("/orders", endpoint.GetOrders)
		path.POST("/order", endpoint.PostOrder)
		path.GET("/order/:order_id", endpoint.GetOrder)

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

		path.GET("/services", endpoint.GetAllServices)
		path.GET("/companies", endpoint.GetCompanies)

		company := path.Group("/company")
		{
			company.POST("", endpoint.PostCompany)
			company.GET("/:company_id", endpoint.GetCompany)
			company.DELETE("/:company_id", endpoint.DeleteCompany)

			company.GET("/order/:order_id", endpoint.GetCompanyOrder)
			company.GET("/:company_id/orders", endpoint.GetCompanyOrders)

			company.POST("/:company_id/service", endpoint.PostService)
			company.GET("/:company_id/services", endpoint.GetServices)
			company.GET("/service/:service_id", endpoint.GetService)
			company.DELETE("/:company_id/service/:service_id", endpoint.DeleteService)
			company.POST("/service/:service_id/review", endpoint.PostServiceReview)
			company.GET("/service/:service_id/reviews", endpoint.GetServiceReviews)
			company.GET("/service/review/:review_id", endpoint.GetServiceReview)
			company.DELETE("/:company_id/service/:service_id/review/:id", endpoint.DeleteServiceReview)

			company.POST("/:company_id/member" , endpoint.PostCompanyMember)
			company.GET("/:company_id/members",endpoint.GetCompanyMembers)
			company.DELETE("/:company_id/member/:username", endpoint.DeleteCompanyMember)
			company.POST("/:company_id/moderator", endpoint.PostCompanyModerator)
			company.GET("/:company_id/moderators", endpoint.GetCompanyModerators)
			company.DELETE("/:company_id/moderator/:username", endpoint.DeleteCompanyModerator)

			company.POST("/:company_id/review", endpoint.PostCompanyReview)
			company.GET("/:company_id/reviews", endpoint.GetCompanyReviews)
			company.GET("/review/:review_id", endpoint.GetCompanyReview)
			company.DELETE("/review/:review_id", endpoint.DeleteCompanyReview)
		}
	}

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *APIServer) configureStore() error{
	if err:= s.store.Open(); err != nil{
		return err
	}

	return nil
}
