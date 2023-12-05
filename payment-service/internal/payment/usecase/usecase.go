package usecase

import (
	"bmstu-dips-lab3/payment-service/internal/payment"
	"bmstu-dips-lab3/payment-service/models"
	"bmstu-dips-lab3/pkg/uuider"
	"context"
)

type PaymentUseCase struct {
	paymentRepo payment.Repo
	uuider      uuider.UUIDer
}

func NewPaymentUseCase(paymentRepo payment.Repo, uuider uuider.UUIDer) payment.UseCase {
	return &PaymentUseCase{
		paymentRepo: paymentRepo,
		uuider:      uuider,
	}
}

func (p *PaymentUseCase) Create(ctx context.Context, payment *models.Payment) (*models.Payment, error) {
	newUUID, err := p.uuider.Generate()
	if err != nil {
		return nil, err
	}

	payment.Payment_uid = *newUUID

	return p.paymentRepo.Create(ctx, payment)
}

func (p *PaymentUseCase) Update(ctx context.Context, payment *models.Payment, toUpdate *models.Payment) (*models.Payment, error) {
	return p.paymentRepo.Update(ctx, payment, toUpdate)
}

func (p *PaymentUseCase) GetByPaymentUid(ctx context.Context, payment_uid string) (*models.Payment, error) {
	return p.paymentRepo.GetByPaymentUid(ctx, payment_uid)
}

func (p *PaymentUseCase) Delete(ctx context.Context, payment_uid string) error {
	return p.paymentRepo.Delete(ctx, payment_uid)
}
