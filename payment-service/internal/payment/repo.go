package payment

import (
	"bmstu-dips-lab3/payment-service/models"
	"context"
)

type Repo interface {
	Create(ctx context.Context, model *models.Payment) (*models.Payment, error)
	Update(ctx context.Context, model *models.Payment, toUpdate *models.Payment) (*models.Payment, error)
	GetByPaymentUid(ctx context.Context, payment_uid string) (*models.Payment, error)
	Delete(ctx context.Context, payment_uid string) error
}
