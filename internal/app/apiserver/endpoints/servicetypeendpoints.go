package endpoints

import (
	"net/http"
	"strconv"

	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Create service type
// @Schemes
// @Description Create service type
// @Tags ServiceType
// @Accept json
// @Produce json
// @Param servicetype body dtos.CreateServiceTypeDto true "Create service type"
// @Router /servicetype [POST]
func (ep *Endpoints) PostServiceType(g *gin.Context){
	var serviceTypeDto dtos.CreateServiceTypeDto

	if err := g.BindJSON(&serviceTypeDto); err != nil{
		g.JSON(http.StatusBadRequest, gin.H{"Invalid request data": err.Error()})
        return
	}

	validate := validator.New()
    if err := validate.Struct(serviceTypeDto); err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"Failed validation": err.Error()})
        return
    }

	serviceTypeModel := serviceTypeDto.ToModel()

	if err := ep.store.ServiceType().Create(&serviceTypeModel); err != nil{
		g.JSON(http.StatusBadRequest, gin.H{"Service type already exists, or db error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, serviceTypeModel)
}

// @Summary Get service type
// @Schemes
// @Description Get service type
// @Tags ServiceType
// @Accept json
// @Produce json
// @Param id path int true "service type id"
// @Router /servicetype/{id} [GET]
func (ep *Endpoints) GetServiceType(g *gin.Context){
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	serviceTypeModel, err := ep.store.ServiceType().FindById(id)
	if err != nil{
		g.JSON(http.StatusNotFound, gin.H{"Service type not found": error.Error(err)})
		return
	}

	g.JSON(http.StatusOK, serviceTypeModel)
}

// @Summary Get service types
// @Schemes
// @Description Get service types
// @Tags ServiceType
// @Accept json
// @Produce json
// @Router /servicetype [GET]
func (ep *Endpoints) GetServiceTypes(g *gin.Context){

	serviceTypeModel, err := ep.store.ServiceType().FindAll()
	if err != nil{
		g.JSON(http.StatusNotFound, gin.H{"Service types not found": error.Error(err)})
		return
	}

	g.JSON(http.StatusOK, serviceTypeModel)
}

// @Summary Delete service type
// @Schemes
// @Description Delete service type
// @Tags ServiceType
// @Accept json
// @Produce json
// @Param id path int true "service type id"
// @Router /servicetype/{id} [DELETE]
func (ep *Endpoints) DeleteServiceType(g *gin.Context){
	id, err := strconv.Atoi(g.Param("id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	ep.store.ServiceType().DeleteById(id)

	g.JSON(http.StatusNoContent, http.NoBody)
}