package endpoints

import (
	"net/http"
	"strconv"

	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)




/*

------------SERVICE-------------

*/



// @Summary Create service review
// @Schemes
// @Description Create service review
// @Security ApiKeyAuth
// @Tags Review,Service
// @Accept json
// @Produce json
// @Param service_id path int true "service id"
// @Param review body dtos.CreateReviewServiceDto true "Create service review"
// @Router /company/service/{service_id}/review [POST]
func (ep *Endpoints) PostServiceReview(g *gin.Context){
	var createReviewServiceDto dtos.CreateReviewServiceDto

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

	id, err := strconv.Atoi(g.Param("service_id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	if err := g.BindJSON(&createReviewServiceDto); err != nil{
		g.JSON(http.StatusBadRequest, gin.H{"Invalid request data": err.Error()})
        return
	}

	validate := validator.New()
    if err := validate.Struct(createReviewServiceDto); err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"Failed validation": err.Error()})
        return
    }

	reviewServiceModel := createReviewServiceDto.ToModel(id, username)

	ep.store.Review().Create(&reviewServiceModel)

	g.JSON(http.StatusCreated, dtos.ReviewToResponce(reviewServiceModel))
}


// @Summary Get service reviews
// @Schemes
// @Description Get service reviews
// @Tags Review,Service
// @Accept json
// @Produce json
// @Param service_id path int true "service id"
// @Router /service/{service_id}/reviews [GET]
func (ep *Endpoints) GetServiceReviews(g *gin.Context){
	id, err := strconv.Atoi(g.Param("service_id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	reviewModel, err := ep.store.Review().FindAllByObjectIdToResponse(dtos.ReviewServicesTable, id)
	if err != nil{
		g.JSON(http.StatusNotFound, gin.H{"Service reviews not found": error.Error(err)})
		return
	}

	g.JSON(http.StatusOK, reviewModel)
}

// @Summary Get service review
// @Schemes
// @Description Get service review
// @Tags Review,Service
// @Accept json
// @Produce json
// @Param review_id path int true "review id"
// @Router /service/review/{review_id} [GET]
func (ep *Endpoints) GetServiceReview(g *gin.Context){
	id, err := strconv.Atoi(g.Param("review_id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	reviewModel, err := ep.store.Review().FindById(dtos.ReviewServicesTable, id)
	if err != nil{
		g.JSON(http.StatusNotFound, gin.H{"Service review not found": error.Error(err)})
		return
	}

	g.JSON(http.StatusOK, dtos.ReviewToResponce(reviewModel))
}

// @Summary Delete service reviews
// @Schemes
// @Description Delete service reviews
// @Tags Review,Service
// @Accept json
// @Produce json
// @Param review_id path int true "review id"
// @Router /service/review/{review_id} [DELETE]
func (ep *Endpoints) DeleteServiceReview(g *gin.Context){
	id, err := strconv.Atoi(g.Param("review_id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	ep.store.Review().DeleteById(dtos.ReviewServicesTable, id)

	g.JSON(http.StatusNoContent, http.NoBody)
}



/*

------------COMPANY-------------

*/



// @Summary Create company review
// @Schemes
// @Description Create company review
// @Security ApiKeyAuth
// @Tags Review,Company
// @Accept json
// @Produce json
// @Param company_id path int true "company id"
// @Param review body dtos.CreateReviewCompanyDto true "Create company review"
// @Router /company/{company_id}/review [POST]
func (ep *Endpoints) PostCompanyReview(g *gin.Context){
	var createReviewCompanyDto dtos.CreateReviewCompanyDto

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

	id, err := strconv.Atoi(g.Param("company_id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}


	if err := g.BindJSON(&createReviewCompanyDto); err != nil{
		g.JSON(http.StatusBadRequest, gin.H{"Invalid request data": err.Error()})
        return
	}

	validate := validator.New()
    if err := validate.Struct(createReviewCompanyDto); err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"Failed validation": err.Error()})
        return
    }

	reviewCompanyModel := createReviewCompanyDto.ToModel(id, username)

	ep.store.Review().Create(&reviewCompanyModel)

	g.JSON(http.StatusCreated, dtos.ReviewToResponce(reviewCompanyModel))
}


// @Summary Get company reviews
// @Schemes
// @Description Get company reviews
// @Tags Review,Company
// @Accept json
// @Produce json
// @Param company_id path int true "company id"
// @Router /company/{company_id}/reviews [GET]
func (ep *Endpoints) GetCompanyReviews(g *gin.Context){
	id, err := strconv.Atoi(g.Param("company_id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	reviewModel, err := ep.store.Review().FindAllByObjectIdToResponse(dtos.ReviewCompaniesTable, id)
	if err != nil{
		g.JSON(http.StatusNotFound, gin.H{"Service reviews not found": error.Error(err)})
		return
	}

	g.JSON(http.StatusOK, reviewModel)
}

// @Summary Get company review
// @Schemes
// @Description Get company review
// @Tags Review,Company
// @Accept json
// @Produce json
// @Param review_id path int true "review id"
// @Router /company/review/{review_id} [GET]
func (ep *Endpoints) GetCompanyReview(g *gin.Context){
	id, err := strconv.Atoi(g.Param("review_id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	reviewModel, err := ep.store.Review().FindById(dtos.ReviewCompaniesTable, id)
	if err != nil{
		g.JSON(http.StatusNotFound, gin.H{"Company review not found": error.Error(err)})
		return
	}

	g.JSON(http.StatusOK, dtos.ReviewToResponce(reviewModel))
}

// @Summary Delete company review
// @Schemes
// @Description Delete company review
// @Tags Review,Company
// @Accept json
// @Produce json
// @Param review_id path int true "company id"
// @Router /company/review/{review_id} [DELETE]
func (ep *Endpoints) DeleteCompanyReview(g *gin.Context){
	id, err := strconv.Atoi(g.Param("review_id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	ep.store.Review().DeleteById(dtos.ReviewCompaniesTable, id)

	g.JSON(http.StatusNoContent, http.NoBody)
}