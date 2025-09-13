package middlewares

import (
	"github.com/labstack/echo/v4"
)

// AllowRoles 허용된 역할만 접근할 수 있도록 하는 미들웨어
// TODO: JWT 토큰 검증 로직이 구현되면 이 미들웨어를 완성하세요
func AllowRoles(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// TODO: JWT 토큰에서 사용자 역할을 추출하고 권한을 확인하는 로직을 구현하세요
			// 현재는 모든 요청을 허용합니다
			return next(c)
		}
	}
}
