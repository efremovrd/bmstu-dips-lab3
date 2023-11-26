package server

import (
	"bmstu-dips-lab3/config"
	h "bmstu-dips-lab3/gateway-service/internal/gateway/delivery/http"
	circuit_breaker "bmstu-dips-lab3/pkg/circuit-breaker"
	job_scheduler "bmstu-dips-lab3/pkg/job-scheduler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers(cfg config.CircuitBreakerConfig, jobScheduler *job_scheduler.JobScheduler) error {
	loyaltyconfig := circuit_breaker.Config{
		Name:        "Loyalty service cb",
		MaxRequests: cfg.MaxRequests,
	}
	loyaltycb := circuit_breaker.NewCircuitBreaker(loyaltyconfig)

	paymentconfig := circuit_breaker.Config{
		Name:        "Payment service cb",
		MaxRequests: cfg.MaxRequests,
	}
	paymentcb := circuit_breaker.NewCircuitBreaker(paymentconfig)

	reservationconfig := circuit_breaker.Config{
		Name:        "Reservation service cb",
		MaxRequests: cfg.MaxRequests,
	}
	reservationcb := circuit_breaker.NewCircuitBreaker(reservationconfig)

	gH := h.NewGatewayHandlers(loyaltycb, paymentcb, reservationcb, jobScheduler)

	s.router.GET("/manage/health", GetHealth())

	api := s.router.Group("/api")

	v1 := api.Group("/v1")

	gateway := v1.Group("")
	h.MapGatewayRoutes(gateway, gH)

	return nil
}

func GetHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
