package http

import (
	"bmstu-dips-lab3/payment-service/internal/payment"

	"github.com/gin-gonic/gin"
)

func MapPaymentRoutes(paymentGroup *gin.RouterGroup, h payment.Handlers) {
	paymentGroup.POST("", h.Create())
	paymentGroup.PATCH("/:paymentUid", h.Update())
	paymentGroup.GET("/:paymentUid", h.GetByPaymentUid())
	paymentGroup.DELETE("/:paymentUid", h.Delete())
}
