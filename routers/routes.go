package routers

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	app2 "github.com/jajangratis/money-manager-clone-api-fiber/app"
	"github.com/jajangratis/money-manager-clone-api-fiber/controller"
	_ "github.com/jajangratis/money-manager-clone-api-fiber/docs"
	"github.com/jajangratis/money-manager-clone-api-fiber/middleware"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/repository"
	"github.com/jajangratis/money-manager-clone-api-fiber/service"
	"time"
)

func ExecutionTimeLogger(c *fiber.Ctx) error {
	start := time.Now()

	// Next is used to execute the next middleware in the chain.
	if err := c.Next(); err != nil {
		return err
	}

	// Calculate the elapsed time.
	elapsed := time.Since(start)

	// Log the execution time.
	fmt.Printf("Request: %s took %s\n", c.Path(), elapsed)

	return nil
}

// AppRouter @title Money Manage Clone API
//
//	@version		1.0
//	@description	This is a sample swagger for Fiber
//	@termsOfService	http://swagger.io/terms/
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:8080
//	@BasePath		/
func AppRouter(app fiber.Router) {
	db := app2.OpenConnection()
	validate := validator.New()
	//=================
	mstCategoryRepositoryImpl := repository.NewMstCategoryRepositoryImpl()
	mstCategoryService := service.NewMstCategoryService(mstCategoryRepositoryImpl, db, validate)
	mstCategoryController := controller.NewMstCategoryController(mstCategoryService)

	mstUserRepositoryImpl := repository.NewMstUserRepositoryImpl()
	mstUserService := service.NewMstUserService(mstUserRepositoryImpl, db, validate)
	mstUserController := controller.NewMstUserController(mstUserService)

	mstTransactionRepository := repository.NewMstTransactionRepositoryImpl()
	mstTransactionService := service.NewMstTransactionServiceImpl(mstTransactionRepository, db, validate)
	mstTransactionController := controller.NewMstTransactionControllerImpl(mstTransactionService)

	trTransactionHistoryRepository := repository.NewTrTransactionHistoryRepositoryImpl()
	trTransactionHistoryService := service.NewTrTransactionServiceImpl(trTransactionHistoryRepository, db, validate)
	trTransactionHistoryController := controller.NewTrTransactionHistoryControllerImpl(trTransactionHistoryService)
	//======================
	app.Use(ExecutionTimeLogger)
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Get("/swagger/*", swagger.HandlerDefault)
	//PUBLIC ACCESSS
	public := app.Group("/public/api")
	//@Summary Login Using username password
	//@Description login
	//@ID login
	//@Produce json
	//@Param id path string true "username"
	//@Param id path string true "password"
	//@Success 200 {object} User
	//@Router /public/api/login [get]
	public.Post("/login", mstUserController.Login)
	// ==========================
	//TOKEN ACCESS
	api := app.Group("/api")
	api.Use(middleware.TokenExtractor)
	api.Get("/", func(ctx *fiber.Ctx) error {
		data := ctx.Get("userId")
		fmt.Println(data)
		return ctx.SendString("Hello " + data)
	})
	api.Get("/mst-category", mstCategoryController.FindAll)
	api.Get("/mst-transaction", mstTransactionController.FindAll)
	api.Get("/tr-transaction-history", trTransactionHistoryController.FindAll)
	api.Get("/tr-transaction-history/report-income-expenses", trTransactionHistoryController.ReportIncomeExpense)
	api.Post("/tr-transaction-history", trTransactionHistoryController.Save)
	api.Put("/tr-transaction-history", trTransactionHistoryController.Edit)
	api.Delete("/tr-transaction-history/:id", trTransactionHistoryController.Delete)
}
