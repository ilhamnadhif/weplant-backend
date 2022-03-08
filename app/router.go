package app

import (
	"github.com/julienschmidt/httprouter"
	"io/fs"
	"net/http"
	"weplant-backend/controller"
	"weplant-backend/exception"
	"weplant-backend/middleware"
)

func NewRouter(swagger fs.FS, authController controller.AuthController, merchantController controller.MerchantController, productController controller.ProductController, categoryController controller.CategoryController, customerController controller.CustomerController, cartController controller.CartController, transactionController controller.TransactionController) *httprouter.Router {

	router := httprouter.New()

	router.PanicHandler = exception.ErrorHandler

	router.ServeFiles("/docs/*filepath", http.FS(swagger))

	router.POST("/api/v1/auth/merchant", authController.LoginMerchant)
	router.POST("/api/v1/auth/customer", authController.LoginCustomer)

	router.POST("/api/v1/merchants", merchantController.Create)
	router.GET("/api/v1/merchants/:merchantId", merchantController.FindById)
	router.GET("/api/v1/merchants/:merchantId/orders", middleware.AuthMiddleware(merchantController.FindManageOrderById, "merchant"))
	router.PUT("/api/v1/merchants/:merchantId", middleware.AuthMiddleware(merchantController.Update, "merchant"))
	router.PATCH("/api/v1/merchants/:merchantId/image", middleware.AuthMiddleware(merchantController.UpdateMainImage, "merchant"))
	router.DELETE("/api/v1/merchants/:merchantId", middleware.AuthMiddleware(merchantController.Delete, "merchant"))

	router.GET("/api/v1/products/:productId", productController.FindById)
	router.GET("/api/v1/products", productController.FindAll)
	router.POST("/api/v1/products", middleware.AuthMiddleware(productController.Create, "merchant"))
	router.PUT("/api/v1/products/:productId", middleware.AuthMiddleware(productController.Update, "merchant"))
	router.PATCH("/api/v1/products/:productId/image", middleware.AuthMiddleware(productController.UpdateMainImage, "merchant"))
	router.POST("/api/v1/products/:productId/images", middleware.AuthMiddleware(productController.PushImageIntoImages, "merchant"))
	router.DELETE("/api/v1/products/:productId/images/:imageId", middleware.AuthMiddleware(productController.PullImageFromImages, "merchant"))
	router.DELETE("/api/v1/products/:productId", middleware.AuthMiddleware(productController.Delete, "merchant"))

	router.GET("/api/v1/categories/:categoryId", categoryController.FindById)
	router.GET("/api/v1/categories", categoryController.FindAll)
	router.POST("/api/v1/categories", middleware.AuthMiddleware(categoryController.Create, "merchant"))

	router.POST("/api/v1/customers", customerController.Create)
	router.GET("/api/v1/customers/:customerId", customerController.FindById)
	router.GET("/api/v1/customers/:customerId/carts", middleware.AuthMiddleware(customerController.FindCartById, "customer"))
	router.GET("/api/v1/customers/:customerId/transactions", middleware.AuthMiddleware(customerController.FindTransactionById, "customer"))
	router.GET("/api/v1/customers/:customerId/orders", middleware.AuthMiddleware(customerController.FindOrderById, "customer"))
	router.PUT("/api/v1/customers/:customerId", middleware.AuthMiddleware(customerController.Update, "customer"))
	router.PATCH("/api/v1/customers/:customerId/image", middleware.AuthMiddleware(customerController.UpdateMainImage, "customer"))
	router.DELETE("/api/v1/customers/:customerId", middleware.AuthMiddleware(customerController.Delete, "customer"))

	router.POST("/api/v1/carts/:customerId", middleware.AuthMiddleware(cartController.PushProductToCart, "customer"))
	router.PATCH("/api/v1/carts/:customerId/products/:productId", middleware.AuthMiddleware(cartController.UpdateProductQuantity, "customer"))
	router.DELETE("/api/v1/carts/:customerId/products/:productId", middleware.AuthMiddleware(cartController.PullProductFromCart, "customer"))

	router.POST("/api/v1/callback", transactionController.Callback)
	router.POST("/api/v1/transactions/:customerId", middleware.AuthMiddleware(transactionController.Create, "customer"))
	router.DELETE("/api/v1/transactions/:customerId/transactions/:transactionId", middleware.AuthMiddleware(transactionController.Cancel, "customer"))

	return router
}
