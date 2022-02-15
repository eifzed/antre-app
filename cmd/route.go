package main

import (
	"github.com/eifzed/antre-app/lib/utility/urlpath"
	"github.com/go-chi/chi"
)

func getRoute(m *modules) *chi.Mux {
	router := chi.NewRouter()
	path := urlpath.New("")
	router.Route("/v1", func(v1 chi.Router) {
		v1.Group(func(user chi.Router) {
			path.Group("/user", func(userRoute urlpath.Routes) {
				user.Post(userRoute.URL("/register"), m.httpHandler.AntreHandler.RegisterNewAccount)
				user.Post(userRoute.URL("/login"), m.httpHandler.AntreHandler.Login)
			})

		})
		v1.Group(func(antre chi.Router) {
			antre.Use(m.AuthModule.AuthHandler)
			path.Group("/user", func(userRoute urlpath.Routes) {
				antre.Put(userRoute.URL("/assign/{role}"), m.httpHandler.AntreHandler.AssignNewRoleToUser)
			})
			path.Group("/shops", func(shopsRoute urlpath.Routes) {
				antre.Post(shopsRoute.URL("/register"), m.httpHandler.OrderHandler.RegisterShop)
			})
			path.Group("/orders", func(orderRoute urlpath.Routes) {
				antre.Post(orderRoute.URL(""), m.httpHandler.OrderHandler.RegisterOrder)
				antre.Get(orderRoute.URL(""), m.httpHandler.OrderHandler.GetCustomerOrders)
				antre.Get(orderRoute.URL("/{id}"), m.httpHandler.OrderHandler.GetOrderByID)
			})
		})

	})

	return router
}
