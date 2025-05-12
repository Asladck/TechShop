package handler

import (
	_ "TechShop/docs"
	"TechShop/pkg/service"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/refresh", h.refreshHandler)
		auth.POST("/sign-in", h.signIn)
	}
	items := router.Group("/items")
	{
		items.GET("/", h.getItems)
		items.GET("/:id", h.getItemById)
	}
	api := router.Group("/api", h.userIdentity)
	{
		wishlist := api.Group("/wishlist")
		{
			wishlist.POST("/:id", h.addToWishlist)        // {item_id}
			wishlist.GET("/", h.getWishlist)              // список избранного
			wishlist.DELETE("/:id", h.deleteFromWishlist) // по wishlistId
		}

		// Cart
		cart := api.Group("/cart")
		{
			cart.POST("/", h.addToCart) // {itemId, itemCount}
			cart.GET("/", h.getCart)
			cart.GET("/:id", h.getCartItemById)
			cart.PUT("/:id", h.updateCartItem)
			cart.DELETE("/:id", h.deleteCartItem)
		}

		// Покупка
		buy := api.Group("/buy")
		{
			buy.GET("/price", h.getPriceInCart) // итоговая цена
			buy.POST("/:id", h.buyOneItem)      // {itemId, itemCount}
		}
		// Заказы
		orders := api.Group("/order")
		{
			orders.GET("/", h.getOrders) // query ?isActive=true/false / {itemId, itemCount}
			orders.POST("/create", h.createOrdersFromCart)
			orders.POST("/create/:id", h.createOrderFromCart)
			orders.GET("/:id", h.getOrderById)
			orders.POST("/:id/cancel", h.cancelOrder)
			orders.POST("/:id/delivered", h.deliveredOrder)
			orders.POST("/:id/delivering", h.deliveringOrder)
		}
	}

	return router
}
