package endpoints

import (
	"net/http"
	"strconv"

	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Create Tag
// @Schemes
// @Description Create tag
// @Tags Tags
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param tag body dtos.CreateTagDto true "Create tag"
// @Router /tag [POST]
func (ep *Endpoints) PostTag(g *gin.Context){
	var tagDto dtos.CreateTagDto

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

    roleJWT, err := token.Claims.GetAudience()
    if err != nil{
        g.JSON(http.StatusNotFound, "Failed to get subject from token")
        return
    }
	role := roleJWT[0]

	if role != AdminRole{
		g.JSON(http.StatusForbidden, "Insufficient permissions")
		return
	}

	if err := g.BindJSON(&tagDto); err != nil{
		g.JSON(http.StatusBadRequest, gin.H{"Invalid request data": err.Error()})
        return
	}

	validate := validator.New()
    if err := validate.Struct(tagDto); err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"Failed validation": err.Error()})
        return
    }

	tagModel := tagDto.ToModel()

	if err := ep.store.Tag().Create(&tagModel); err != nil{
		g.JSON(http.StatusBadRequest, gin.H{"Tag already exists, or db error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, tagModel)
}

// @Summary Get Tag
// @Schemes
// @Description Get tag
// @Tags Tags
// @Accept json
// @Produce json
// @Router /tags [GET]
func (ep *Endpoints) GetTags(g *gin.Context){
	tags, err := ep.store.Tag().FindAll()
	if err != nil{
		g.JSON(http.StatusNotFound, gin.H{"Tags not found": error.Error(err)})
		return
	}

	g.JSON(http.StatusOK, tags)
}

// @Summary Get Tag
// @Schemes
// @Description Get tag
// @Accept json
// @Tags Tags
// @Produce json
// @Param id path int true "Tag id"
// @Router /tag/{id} [GET]
func (ep *Endpoints) GetTag(g *gin.Context){
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	tagModel, err := ep.store.Tag().FindById(id)
	if err != nil{
		g.JSON(http.StatusNotFound, gin.H{"Tag not found": error.Error(err)})
		return
	}

	g.JSON(http.StatusOK, tagModel)
}

// @Summary Delete Tag
// @Schemes
// @Description Delete tag
// @Tags Tags
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Tag id"
// @Router /tag/{id} [DELETE]
func (ep *Endpoints) DeleteTag(g *gin.Context){
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

    roleJWT, err := token.Claims.GetAudience()
    if err != nil{
        g.JSON(http.StatusNotFound, "Failed to get subject from token")
        return
    }
	role := roleJWT[0]

	if role != AdminRole{
		g.JSON(http.StatusForbidden, "Insufficient permissions")
		return
	}
	
	id, err := strconv.Atoi(g.Param("id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	ep.store.Tag().DeleteById(id)

	g.JSON(http.StatusNoContent, http.NoBody)
}