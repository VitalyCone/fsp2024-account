package endpoints

import (
	"net/http"
	"strconv"

	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Create company
// @Schemes
// @Description Create company
// @Security ApiKeyAuth
// @Tags Company
// @Accept json
// @Produce json
// @Param company body dtos.CreateCompanyDto true "Create company dto"
// @Router /company [POST]
func (ep *Endpoints) PostCompany(g *gin.Context) {
	var createCompanyDto dtos.CreateCompanyDto

	tokenString := g.GetHeader("token")
    if tokenString == ""{
        g.JSON(http.StatusUnauthorized, "token nil")
        return
    }

    token, err := verifyToken(tokenString)
    if err != nil{
        g.JSON(http.StatusUnauthorized, "token not verifed or nil")
        return
    }

    username, err := token.Claims.GetSubject()
    if err != nil{
        g.JSON(http.StatusNotFound, "Failed to get subject from token")
        return
    }

	if err := g.BindJSON(&createCompanyDto); err != nil{
		g.JSON(http.StatusBadRequest, gin.H{"Invalid request data": err.Error()})
        return
	}

	validate := validator.New()
    if err := validate.Struct(createCompanyDto); err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"Failed validation": err.Error()})
        return
    }

    //ДОДЕЛАТЬ ТЕГИ

	companyModel := createCompanyDto.ToModel()

    if err := ep.store.Company().Create(&companyModel ,username, dtos.MembersParticipantTable, dtos.ModeratorsParticipantTable); err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"Failed to create company": err.Error()})
        return
    }

    g.JSON(http.StatusCreated, companyModel)
}

// @Summary Create company
// @Schemes
// @Description Create company
// @Tags Company
// @Accept json
// @Produce json
// @Param id path int true "company id"
// @Router /company/{id} [GET]
func (ep *Endpoints) GetCompany(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

    companyModel, err := ep.store.Company().FindById(
        id, dtos.MembersParticipantTable, dtos.ModeratorsParticipantTable, dtos.ReviewCompaniesTable)
    if err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"Failed to get company": err.Error()})
        return
    }

    g.JSON(http.StatusOK, companyModel)
}