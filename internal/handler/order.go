package handler

import "net/http"

type orderHandlerInterface interface {
	// order handler
	GetOrderByID(w http.ResponseWriter, r *http.Request)
	RegisterOrder(w http.ResponseWriter, r *http.Request)
	GetCustomerOrders(w http.ResponseWriter, r *http.Request)

	// shop handler
	RegisterShop(w http.ResponseWriter, r *http.Request)

	// product handler
	GetProductsListByShopID(w http.ResponseWriter, r *http.Request)
}
