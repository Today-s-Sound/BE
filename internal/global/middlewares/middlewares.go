package middlewares

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/nicewook/gocore/internal/global/config"
	"github.com/nicewook/gocore/pkg/contextutil"
	"github.com/nicewook/gocore/pkg/validatorutil"
)

// RegisterMiddlewares 기본 미들웨어들을 등록합니다
func RegisterMiddlewares(cfg *config.Config, logger *slog.Logger, e *echo.Echo) {
	// Validator: 요청 바인딩 및 유효성 검사
	e.Validator = validatorutil.NewValidator()

	// RequestID: 각 요청에 고유한 ID 부여
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		RequestIDHandler: func(c echo.Context, requestID string) {
			req := c.Request()
			req.Header.Set(echo.HeaderXRequestID, requestID)
			ctx := contextutil.WithRequestID(req.Context(), requestID)
			c.SetRequest(req.WithContext(ctx))
		},
	}))

	// Config: 설정 정보를 컨텍스트에 추가
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("config", cfg)
			return next(c)
		}
	})

	// Logger: 요청 로깅
	e.Use(LoggerMiddleware(logger))

	// Recover: 패닉 복구
	e.Use(middleware.Recover())

	// CORS: Cross-Origin Resource Sharing
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     cfg.Secure.CORSAllowOrigins,
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
		ExposeHeaders:    []string{echo.HeaderXRequestID},
		MaxAge:           86400,
	}))

	// BodyLimit: 요청 크기 제한 (2MB)
	e.Use(middleware.BodyLimit("2M"))
}
