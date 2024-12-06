package endpoints

import (
	"net/http"
	"strconv"

	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
	"github.com/VitalyCone/account/internal/app/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/*



Members



*/



// @Summary Post company member
// @Schemes
// @Description Post company member
// @Security ApiKeyAuth
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path int true "company id"
// @Param participant body dtos.CreateParticipantDto true "Create participant dto"
// @Router /company/{company_id}/member [POST]
func (ep *Endpoints) PostCompanyMember(g *gin.Context) {
    var createParticipantDto dtos.CreateParticipantDto

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

	if err := g.BindJSON(&createParticipantDto); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid request data": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(createParticipantDto); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Failed validation": err.Error()})
		return
	}

    particantModel :=  createParticipantDto.ToModel(id)

    err = ep.store.Participant().Create(&particantModel, dtos.MembersParticipantTable)
    if err != nil{
        g.JSON(http.StatusInternalServerError, "Failed to create participant " +err.Error())
        return
    }

    g.JSON(http.StatusCreated, dtos.ParticipantToResponse(particantModel))
}


// @Summary Get company members
// @Schemes
// @Description Get company members
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path int true "company id"
// @Router /company/{company_id}/members [GET]
func (ep *Endpoints) GetCompanyMembers(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("company_id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

    particants, err := ep.store.Participant().FindByCompanyToResponse(id, dtos.MembersParticipantTable)
    if err != nil{
        g.JSON(http.StatusNotFound, gin.H{"Service reviews not found": error.Error(err)})
		return
    }

    g.JSON(http.StatusOK, particants)
}

// @Summary Delete company member
// @Schemes
// @Description Delete company member
// @Security ApiKeyAuth
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path int true "company id"
// @Param username path string true "username"
// @Router /company/{company_id}/member/{username} [DELETE]
func (ep *Endpoints) DeleteCompanyMember(g *gin.Context) {
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

    deletedUsername:= g.Param("username")
	if deletedUsername == "" {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid username in path": error.Error(err)})
		return
	}

    particant := model.Participant{
        User: model.User{Username: deletedUsername},
        Company: model.Company{ID: id},
    }

    err = ep.store.Participant().Delete(particant, dtos.MembersParticipantTable)
    if err != nil{
        g.JSON(http.StatusNotFound, "Participant not found "  + err.Error())
    }

    g.JSON(http.StatusNoContent, http.NoBody)
}




/*



Moderators



*/



// @Summary Post company moderator
// @Schemes
// @Description Post company moderator
// @Security ApiKeyAuth
// @Tags Company
// @Accept json
// @Produce json
// @Param participant body dtos.CreateParticipantDto true "Create participant dto"
// @Param company_id path int true "company id"
// @Router /company/{company_id}/moderator [POST]
func (ep *Endpoints) PostCompanyModerator(g *gin.Context) {
    var createParticipantDto dtos.CreateParticipantDto

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

	if err := g.BindJSON(&createParticipantDto); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid request data": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(createParticipantDto); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Failed validation": err.Error()})
		return
	}

    particantModel :=  createParticipantDto.ToModel(id)

    err = ep.store.Participant().Create(&particantModel, dtos.ModeratorsParticipantTable)
    if err != nil{
        g.JSON(http.StatusInternalServerError, "Failed to create participant " +err.Error())
        return
    }

    g.JSON(http.StatusCreated, dtos.ParticipantToResponse(particantModel))
}


// @Summary Get company moderators
// @Schemes
// @Description Get company moderators
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path int true "company id"
// @Router /company/{company_id}/moderators [GET]
func (ep *Endpoints) GetCompanyModerators(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("company_id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

    particants, err := ep.store.Participant().FindByCompanyToResponse(id, dtos.ModeratorsParticipantTable)
    if err != nil{
        g.JSON(http.StatusNotFound, gin.H{"Service reviews not found": error.Error(err)})
		return
    }

    g.JSON(http.StatusOK, particants)
}

// @Summary Delete company moderators
// @Schemes
// @Description Delete company moderators
// @Security ApiKeyAuth
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path int true "company id"
// @Param username path string true "username"
// @Router /company/{company_id}/moderator/{username} [DELETE]
func (ep *Endpoints) DeleteCompanyModerator(g *gin.Context) {
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

    deletedUsername:= g.Param("username")
	if deletedUsername == "" {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid username in path": error.Error(err)})
		return
	}

    particant := model.Participant{
        User: model.User{Username: deletedUsername},
        Company: model.Company{ID: id},
    }

    err = ep.store.Participant().Delete(particant, dtos.ModeratorsParticipantTable)
    if err != nil{
        g.JSON(http.StatusNotFound, "Participant not found " + err.Error())
    }

    g.JSON(http.StatusNoContent, http.NoBody)
}
