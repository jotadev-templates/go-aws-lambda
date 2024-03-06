package portin

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/uuid"

	"crm-lambda/helper"
)

func middlewareRequestID(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.TODO()

		requestID := uuid.New().String()
		ctx = helper.CtxNewValue(ctx, "requestID", requestID)
		logger.Info(fmt.Sprintf("requestID: %s", requestID))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
