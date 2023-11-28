package routers

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	app2 "github.com/jajangratis/money-manager-clone-api-fiber/app"
	"github.com/jajangratis/money-manager-clone-api-fiber/controller"
	"github.com/jajangratis/money-manager-clone-api-fiber/middleware"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/repository"
	"github.com/jajangratis/money-manager-clone-api-fiber/service"
)

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
	//PUBLIC ACCESSS
	public := app.Group("/public/api")
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
}
