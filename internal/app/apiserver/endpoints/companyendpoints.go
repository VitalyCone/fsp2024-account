package endpoints

import (
	"net/http"

	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (ep *Endpoints) CreateCompany(g *gin.Context) {
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

	g.JSON(200, "In develope " +username)

	//user := model.User{Username: username}

	// companyModel := createCompanyDto.ToModel()

}