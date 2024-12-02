package endpoints

import (
	"net/http"
	"strconv"

	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Create service review
// @Schemes
// @Description Create service review
// @Tags Review,Service
// @Accept json
// @Produce json
// @Param review body dtos.CreateReviewServiceDto true "Create service review"
// @Router /service/review [POST]
func (ep *Endpoints) PostServiceReview(g *gin.Context){
	var createReviewServiceDto dtos.CreateReviewServiceDto

	if err := g.BindJSON(&createReviewServiceDto); err != nil{
		g.JSON(http.StatusBadRequest, gin.H{"Invalid request data": err.Error()})
        return
	}

	validate := validator.New()
    if err := validate.Struct(createReviewServiceDto); err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"Failed validation": err.Error()})
        return
    }

	reviewServiceModel := createReviewServiceDto.ToModel()

	ep.store.Review().Create(&reviewServiceModel)

	g.JSON(http.StatusCreated, reviewServiceModel)
}


// @Summary Get service reviews
// @Schemes
// @Description Get service reviews
// @Tags Review,Service
// @Accept json
// @Produce json
// @Param id path int true "service id"
// @Router /service/review/{id} [GET]
func (ep *Endpoints) GetServiceReviews(g *gin.Context){
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	reviewModel, err := ep.store.Review().FindAllByObjectId(dtos.ReviewServicesTable, id)
	if err != nil{
		g.JSON(http.StatusNotFound, gin.H{"Service reviews not found": error.Error(err)})
		return
	}

	g.JSON(http.StatusOK, reviewModel)
}

// @Summary Delete service reviews
// @Schemes
// @Description Delete service reviews
// @Tags Review,Service
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Router /service/review/{id} [DELETE]
func (ep *Endpoints) DeleteServiceReview(g *gin.Context){
	id, err := strconv.Atoi(g.Param("id"))

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
// @Tags Review,Company
// @Accept json
// @Produce json
// @Param review body dtos.CreateReviewCompanyDto true "Create company review"
// @Router /company/review [POST]
func (ep *Endpoints) PostCompanyReview(g *gin.Context){
	var createReviewCompanyDto dtos.CreateReviewCompanyDto

	if err := g.BindJSON(&createReviewCompanyDto); err != nil{
		g.JSON(http.StatusBadRequest, gin.H{"Invalid request data": err.Error()})
        return
	}

	validate := validator.New()
    if err := validate.Struct(createReviewCompanyDto); err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"Failed validation": err.Error()})
        return
    }

	reviewCompanyModel := createReviewCompanyDto.ToModel()

	ep.store.Review().Create(&reviewCompanyModel)

	g.JSON(http.StatusCreated, reviewCompanyModel)
}


// @Summary Get company reviews
// @Schemes
// @Description Get company reviews
// @Tags Review,Company
// @Accept json
// @Produce json
// @Param id path int true "company id"
// @Router /company/review/{id} [GET]
func (ep *Endpoints) GetCompanyReviews(g *gin.Context){
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	reviewModel, err := ep.store.Review().FindAllByObjectId(dtos.ReviewCompaniesTable, id)
	if err != nil{
		g.JSON(http.StatusNotFound, gin.H{"Service reviews not found": error.Error(err)})
		return
	}

	g.JSON(http.StatusOK, reviewModel)
}

// @Summary Delete company review
// @Schemes
// @Description Delete company review
// @Tags Review,Company
// @Accept json
// @Produce json
// @Param id path int true "company id"
// @Router /company/review/{id} [DELETE]
func (ep *Endpoints) DeleteCompanyReview(g *gin.Context){
	id, err := strconv.Atoi(g.Param("id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
		return
	}

	ep.store.Review().DeleteById(dtos.ReviewCompaniesTable, id)

	g.JSON(http.StatusNoContent, http.NoBody)
}



/*




Общее




*/

// // @Summary Get review
// // @Schemes
// // @Description Get review
// // @Tags Review
// // @Accept json
// // @Produce json
// // @Param id path int true "id"
// // @Router /review/{id} [GET]
// func (ep *Endpoints) GetReview(g *gin.Context){
// 	id, err := strconv.Atoi(g.Param("id"))
// 	if err != nil {
// 		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": error.Error(err)})
// 		return
// 	}

// 	reviewModel, err := ep.store.Review().FindById()
// 	if err != nil{
// 		g.JSON(http.StatusNotFound, gin.H{"Service reviews not found": error.Error(err)})
// 		return
// 	}

// 	g.JSON(http.StatusOK, reviewModel)
// }