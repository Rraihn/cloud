package service

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"booking-hotel/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type DiscountServiceImpl struct {
	DiscountRepository repository.DiscountRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NesDiscountService(discountRepository repository.DiscountRepository, DB *sql.DB, validate *validator.Validate) DiscountService {
	return &DiscountServiceImpl{
		DiscountRepository: discountRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (disc DiscountServiceImpl) CreateDiscount(ctx context.Context, request web.DiscountCreateRequest) web.DiscountResponse {
	err := disc.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := disc.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	discount := domain.Discount{
		Discount_amount:  request.Discount_amount,
		Discount_request: request.Discount_request,
		Discount_status:  request.Discount_status,
		Invoice_id:       request.Invoice_id,
	}

	discount = disc.DiscountRepository.SaveDiscount(ctx, tx, discount)
	return helper.ToDiscountResponse(discount)
}

func (disc DiscountServiceImpl) UpdateDiscount(ctx context.Context, request web.DiscountUpdateRequest) web.DiscountResponse {
	err := disc.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := disc.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	discount, err := disc.DiscountRepository.FindDiscountById(ctx, tx, request.Discount_id)
	if err != nil {
		helper.PanicIfError(err)
	}
	discount.Discount_request = request.Discount_request
	discount.Discount_status = request.Discount_status
	discount.Discount_amount = request.Discount_amount

	discount = disc.DiscountRepository.UpdateDiscount(ctx, tx, discount)
	return helper.ToDiscountResponse(discount)
}

func (disc DiscountServiceImpl) DeleteDiscount(ctx context.Context, discountId int) {
	tx, err := disc.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	discount, err := disc.DiscountRepository.FindDiscountById(ctx, tx, discountId)
	if err != nil {
		helper.PanicIfError(err)
	}

	disc.DiscountRepository.DeleteDiscount(ctx, tx, discount)
}

func (disc DiscountServiceImpl) FindDiscountById(ctx context.Context, discountId int) web.DiscountResponse {
	tx, err := disc.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	discount, err := disc.DiscountRepository.FindDiscountById(ctx, tx, discountId)
	if err != nil {
		helper.PanicIfError(err)
	}

	return helper.ToDiscountResponse(discount)
}

func (disc DiscountServiceImpl) FindAllDiscount(ctx context.Context) []web.DiscountResponse {
	tx, err := disc.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	discounts := disc.DiscountRepository.FindAllDiscount(ctx, tx)

	return helper.ToDiscountResponses(discounts)
}
