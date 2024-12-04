package endpoints

import (
	"net/http"
	"strconv"

	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Create service
// @Schemes
// @Description Create service
// @Security ApiKeyAuth
// @Tags Service
// @Accept json
// @Produce json
// @Param serviceDto body dtos.CreateServiceDto true "Create service dto"
// @Router /service [POST]
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

	exist, err := ep.store.Participant().IsParticipant(username, dtos.ModeratorsParticipantTable, createServiceDto.CompanyID)
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

	serviceModel := createServiceDto.ToModel()
	err = ep.store.Service().Create(&serviceModel)
	if err != nil{
		g.JSON(http.StatusInternalServerError, "Failed to create service")
		return
	}

	g.JSON(http.StatusCreated, dtos.ServiceToDto(serviceModel))
}

// @Summary Get services
// @Schemes
// @Description Get services
// @Tags Company,Service
// @Accept json
// @Produce json
// @Param id path int true "service id"
// @Router /company/{id}/services [GET]
func (ep *Endpoints) GetServices(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	services, err := ep.store.Service().FindByCompanyId(id)
	if err != nil{
		g.JSON(http.StatusInternalServerError, "Failed to get services")
	}

	g.JSON(http.StatusOK, services)
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
// @Router /company/{company_id}/service/{service_id} [GET]
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