package endpoints

import (
	"net/http"
	"strconv"

	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Get all services
// @Schemes
// @Description Get all services
// @Tags Company,Service
// @Accept json
// @Produce json
// @Param tags query []string false "Filter by tags" collectionFormat(multi)
// @Param rating query string false "Filter by rating"
// @Param min_price query string false "Minimum price"
// @Param max_price query string false "Maximum price"
// @Router /services [GET]
func (ep *Endpoints) GetAllServices(g *gin.Context) {
    tags := g.QueryArray("tags") // Получаем массив тегов
    rating := g.Query("rating")
    minPrice := g.Query("min_price")
    maxPrice := g.Query("max_price")

    services, err := ep.store.Service().FindAll(tags, rating, minPrice, maxPrice)
    if err != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    responseServices := make([]dtos.ServiceResponse, 0)
    for _, service := range services {
        responseServices = append(responseServices, dtos.ModelServiceToResponse(service))
    }

    g.JSON(http.StatusOK, responseServices)
}

// @Summary Create service
// @Schemes
// @Description Create service
// @Security ApiKeyAuth
// @Tags Company,Service
// @Accept json
// @Produce json
// @Param company_id path int true "company id"
// @Param serviceDto body dtos.CreateServiceDto true "Create service dto"
// @Router /company/{company_id}/service [POST]
func (ep *Endpoints) PostService(g *gin.Context) {
	var createServiceDto dtos.CreateServiceDto
	
	tokenString := g.GetHeader("token")
	if tokenString == "" {
		g.JSON(http.StatusUnauthorized, "token nil")
		return
	}

	token, err := verifyToken(tokenString)
	if err != nil {
		g.JSON(http.StatusUnauthorized, "token not verifed or nil")
		return
	}

	username, err := token.Claims.GetSubject()
	if err != nil {
		g.JSON(http.StatusNotFound, "Failed to get subject from token")
		return
	}

	id, err := strconv.Atoi(g.Param("company_id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	exist, err := ep.store.Participant().IsParticipant(username, dtos.ModeratorsParticipantTable, id)
	if err != nil{
		g.JSON(http.StatusNotFound, "Failed to get permissions")
		return
	}

	if !exist{
		g.JSON(http.StatusForbidden, "You are not allowed to create service")
		return
	}

	if err := g.BindJSON(&createServiceDto); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid request data": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(createServiceDto); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Failed validation": err.Error()})
		return
	}

	serviceModel := createServiceDto.ToModel(id)
	err = ep.store.Service().Create(&serviceModel)
	if err != nil{
		g.JSON(http.StatusInternalServerError, gin.H{"Failed to create service": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, dtos.ModelServiceToResponse(serviceModel))
}

// @Summary Get services
// @Schemes
// @Description Get services
// @Tags Company,Service
// @Accept json
// @Produce json
// @Param company_id path int true "company id"
// @Router /company/{company_id}/services [GET]
func (ep *Endpoints) GetServices(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("company_id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	services, err := ep.store.Service().FindByCompanyIdToResponse(id)
	if err != nil{
		g.JSON(http.StatusInternalServerError, gin.H{"Failed to get services": error.Error(err)})
		return
	}

	g.JSON(http.StatusOK, services)
}

// @Summary Get service
// @Schemes
// @Description Get service
// @Tags Company,Service
// @Accept json
// @Produce json
// @Param service_id path int true "service id"
// @Router /company/service/{service_id} [GET]
func (ep *Endpoints) GetService(g *gin.Context) {
	service_id, err := strconv.Atoi(g.Param("service_id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	service, err := ep.store.Service().FindById(service_id)
	if err != nil{
		g.JSON(http.StatusNotFound, gin.H{"Failed to get service": error.Error(err)})
		return
	}

	g.JSON(http.StatusOK, dtos.ModelServiceToResponse(service))

}

// @Summary Delete services
// @Schemes
// @Description Delete services
// @Security ApiKeyAuth
// @Tags Company,Service
// @Accept json
// @Produce json
// @Param service_id path int true "service id"
// @Param company_id path int true "company id"
// @Router /company/{company_id}/service/{service_id} [DELETE]
func (ep *Endpoints) DeleteService(g *gin.Context) {
	tokenString := g.GetHeader("token")
	if tokenString == "" {
		g.JSON(http.StatusUnauthorized, "token nil")
		return
	}

	token, err := verifyToken(tokenString)
	if err != nil {
		g.JSON(http.StatusUnauthorized, "token not verifed or nil")
		return
	}

	username, err := token.Claims.GetSubject()
	if err != nil {
		g.JSON(http.StatusNotFound, "Failed to get subject from token")
		return
	}
	service_id, err := strconv.Atoi(g.Param("service_id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	company_id, err := strconv.Atoi(g.Param("company_id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	exist, err := ep.store.Participant().IsParticipant(username, dtos.ModeratorsParticipantTable, company_id)
	if err != nil{
		g.JSON(http.StatusNotFound, "Failed to get permissions")
		return
	}

	if !exist{
		g.JSON(http.StatusForbidden, "You are not allowed to delete service")
		return
	}

	ep.store.Service().DeleteById(service_id)

	g.JSON(http.StatusNoContent, http.NoBody)

}