package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/labstack/echo/v4"
	"myservice/pkg/logger"
)

const RequestIDHeader = "X-Request-ID"

func generateRequestID() string {
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		return "rid-unknown"
	}
	return hex.EncodeToString(b)
}

// RequestID middleware sets X-Request-ID header if absent and stores it in context
func RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqID := c.Request().Header.Get(RequestIDHeader)
		if reqID == "" {
			reqID = generateRequestID()
			c.Request().Header.Set(RequestIDHeader, reqID)
		}
		c.Response().Header().Set(RequestIDHeader, reqID)
		// store in context for handlers
		c.Set("request_id", reqID)
		logger.Log.Debugf("request start: id=%s method=%s path=%s", reqID, c.Request().Method, c.Request().URL.Path)
		// set request id as field for subsequent logs is left to callers
		return next(c)
	}
}
