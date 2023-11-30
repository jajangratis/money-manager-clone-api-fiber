package service

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/jajangratis/money-manager-clone-api-fiber/app"
	"github.com/jajangratis/money-manager-clone-api-fiber/helper"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/domain"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/repository"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/web"
	"gorm.io/gorm"
)

type TrTransactionServiceImpl struct {
	TrTransactionHistoryRepository repository.TrTransactionHistoryRepository
	DB                             *gorm.DB
	Validate                       *validator.Validate
}

func NewTrTransactionServiceImpl(trTransactionHistoryRepository repository.TrTransactionHistoryRepository, DB *gorm.DB, validate *validator.Validate) TrTransactionHistoryService {
	return &TrTransactionServiceImpl{
		TrTransactionHistoryRepository: trTransactionHistoryRepository,
		DB:                             DB,
		Validate:                       validate,
	}
}

func (service *TrTransactionServiceImpl) FindAll(username string) web.WebResponse {
	if username == "" {
		return helper.InvalidParameter()
	}
	tx := app.OpenConnection()
	data := service.TrTransactionHistoryRepository.FindAll(tx, username)
	return helper.Ok(data)
}

func (service *TrTransactionServiceImpl) Save(data *web.InputSaveTransactionHistory) web.WebResponse {
	input := domain.TrTransactionHistory{
		From:       data.From,
		To:         data.To,
		Amount:     data.Amount,
		Fee:        data.Fee,
		Note:       data.Note,
		Username:   data.Username,
		MethodId:   data.MethodId,
		AccountId:  data.AccountId,
		CategoryId: data.CategoryId,
	}
	tx := app.OpenConnection()
	ctx := context.Background()
	err := service.Validate.Struct(data)
	if err != nil {
		return helper.InvalidParameter()
	}
	service.TrTransactionHistoryRepository.Save(ctx, tx, input)
	return helper.Ok("")
}

func (service *TrTransactionServiceImpl) ReportByMethodId(input *web.ReportByMethodIdInput) web.WebResponse {
	tx := app.OpenConnection()
	ctx := context.Background()
	err := service.Validate.Struct(input)
	if err != nil {
		return helper.InvalidParameter()
	}

	// Create channels to receive the results asynchronously
	totalExpenseCh := make(chan domain.AggregationResult)
	totalIncomeCh := make(chan domain.AggregationResult)

	// Run totalExpense asynchronously
	go func() {
		totalExpenseCh <- service.TrTransactionHistoryRepository.SumData(ctx, tx, "expense", input.Username)
	}()

	// Run totalIncome asynchronously
	go func() {
		totalIncomeCh <- service.TrTransactionHistoryRepository.SumData(ctx, tx, "income", input.Username)
	}()

	// Wait for both goroutines to finish and collect the results
	totalExpense := <-totalExpenseCh
	totalIncome := <-totalIncomeCh

	// You can close the channels if you no longer need them
	close(totalExpenseCh)
	close(totalIncomeCh)

	// Your logic with totalExpense and totalIncome
	var data = struct {
		TotalExpense domain.AggregationResult
		TotalIncome  domain.AggregationResult
	}{
		TotalExpense: totalExpense,
		TotalIncome:  totalIncome,
	}

	return helper.Ok(data)
}
