package endpoints

import (
	"net/http"
	"strconv"

	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Post order
// @Schemes
// @Description Post order
// @Security ApiKeyAuth
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body dtos.CreateOrderDto true "Create order dto"
// @Router /order [POST]
func (ep *Endpoints) PostOrder(g *gin.Context) {
    var createOrdertDto dtos.CreateOrderDto

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

	if err := g.BindJSON(&createOrdertDto); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid request data": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(createOrdertDto); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Failed validation": err.Error()})
		return
	}

    serviceModel,err := ep.store.Service().FindById(createOrdertDto.ServiceId)
	if err != nil{
		g.JSON(http.StatusNotFound, gin.H{"Service not found": err.Error()})
		return 
	}
	userModel,err := ep.store.User().FindUserByUsername(username)
	if err != nil{
		g.JSON(http.StatusNotFound, gin.H{"User not found": err.Error()})
		return 
	}

	if userModel.Balance < serviceModel.Price{
		g.JSON(http.StatusPaymentRequired, "Insufficient balance")
		return
	}

	orderModel := createOrdertDto.ToModel(username)

    err = ep.store.Order().Create(&orderModel, userModel)
    if err != nil{
        g.JSON(http.StatusInternalServerError, "Failed to create order " +err.Error())
        return
    }

    g.JSON(http.StatusCreated, dtos.OrderModelToResponse(orderModel))
}

// @Summary Get orders
// @Schemes
// @Description Get orders
// @Security ApiKeyAuth
// @Tags Orders
// @Accept json
// @Produce json
// @Router /orders [GET]
func (ep *Endpoints) GetOrders(g *gin.Context) {
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

	orders, err := ep.store.Order().FindAllByUsernameToResponse(username)
	if err != nil{
		g.JSON(http.StatusInternalServerError, "Failed to get orders" +err.Error())
		return
	}

	g.JSON(http.StatusOK, orders)
}

// @Summary Get order
// @Schemes
// @Description Get order
// @Security ApiKeyAuth
// @Tags Orders
// @Accept json
// @Produce json
// @Param order_id path int true "order id"
// @Router /order/{order_id} [GET]
func (ep *Endpoints) GetOrder(g *gin.Context) {
    tokenString := g.GetHeader("token")
	if tokenString == "" {
		g.JSON(http.StatusUnauthorized, "token nil")
		return
	}

	_, err := verifyToken(tokenString)
	if err != nil {
		g.JSON(http.StatusUnauthorized, "token not verifed or nil")
		return
	}

	id, err := strconv.Atoi(g.Param("order_id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": err.Error()})
		return
	}

	order, err := ep.store.Order().FindById(id)
	if err != nil{
		g.JSON(http.StatusInternalServerError, "Failed to get orders " +err.Error())
		return
	}

	g.JSON(http.StatusOK, dtos.OrderModelToResponse(order))
}




/*



COMPANY


*/


// @Summary Get company orders
// @Schemes
// @Description Get company orders
// @Security ApiKeyAuth
// @Tags Company,Orders
// @Accept json
// @Produce json
// @Param company_id path int true "company id"
// @Router /company/{company_id}/orders [GET]
func (ep *Endpoints) GetCompanyOrders(g *gin.Context) {
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
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": err.Error()})
		return
	}

	exist, err := ep.store.Participant().IsParticipant(username, dtos.MembersParticipantTable, id)
	if err != nil{
		g.JSON(http.StatusNotFound, "Failed to get permissions")
		return
	}

	if !exist{
		g.JSON(http.StatusForbidden, "You are not allowed to create service")
		return
	}

	orders, err := ep.store.Order().FindAllByCompanyIdToResponse(id)
	if err != nil{
		g.JSON(http.StatusInternalServerError, "Failed to get orders" +err.Error())
		return
	}

	g.JSON(http.StatusOK, orders)
}


// @Summary Get company order
// @Schemes
// @Description Get company order
// @Security ApiKeyAuth
// @Tags Company,Orders
// @Accept json
// @Produce json
// @Param order_id path int true "order id"
// @Router /company/order/{order_id} [GET]
func (ep *Endpoints) GetCompanyOrder(g *gin.Context) {
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

	id, err := strconv.Atoi(g.Param("order_id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Invalid type of id": err.Error()})
		return
	}

	
	orders, err := ep.store.Order().FindById(id)
	if err != nil{
		g.JSON(http.StatusInternalServerError, "Failed to get orders" +err.Error())
		return
	}

	exist, err := ep.store.Participant().IsParticipant(username, dtos.MembersParticipantTable, orders.Company.ID)
	if err != nil{
		g.JSON(http.StatusNotFound, "Failed to get permissions")
		return
	}

	if !exist{
		g.JSON(http.StatusForbidden, "You are not allowed to create service")
		return
	}

	g.JSON(http.StatusOK, orders)
}