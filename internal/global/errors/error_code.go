package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ErrorResponse 공통 에러 응답 구조체
type ErrorResponse struct {
	Status    int    `json:"status"`            // HTTP 상태 코드
	ErrorCode string `json:"error_code"`        // 에러 코드
	Message   string `json:"message"`           // 에러 메시지
	Details   string `json:"details,omitempty"` // 추가 상세 정보
}

// 공통 에러 코드들
var (
	// 기본 HTTP 에러들
	INTERNAL_SERVER_ERROR = ErrorResponse{
		Status:    http.StatusInternalServerError,
		ErrorCode: "INTERNAL_SERVER_ERROR",
		Message:   "서버 내부 오류가 발생했습니다",
	}

	UNAUTHORIZED = ErrorResponse{
		Status:    http.StatusUnauthorized,
		ErrorCode: "UNAUTHORIZED",
		Message:   "인증이 필요합니다",
	}

	FORBIDDEN = ErrorResponse{
		Status:    http.StatusForbidden,
		ErrorCode: "FORBIDDEN",
		Message:   "접근 권한이 없습니다",
	}

	NOT_FOUND = ErrorResponse{
		Status:    http.StatusNotFound,
		ErrorCode: "NOT_FOUND",
		Message:   "요청한 리소스를 찾을 수 없습니다",
	}

	BAD_REQUEST = ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorCode: "BAD_REQUEST",
		Message:   "잘못된 요청입니다",
	}

	CONFLICT = ErrorResponse{
		Status:    http.StatusConflict,
		ErrorCode: "CONFLICT",
		Message:   "요청이 충돌합니다",
	}

	VALIDATION_ERROR = ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorCode: "VALIDATION_ERROR",
		Message:   "유효성 검사에 실패했습니다",
	}

	// 데이터베이스 관련 에러들
	DATABASE_ERROR = ErrorResponse{
		Status:    http.StatusInternalServerError,
		ErrorCode: "DATABASE_ERROR",
		Message:   "데이터베이스 오류가 발생했습니다",
	}
)

// EchoErrorHandler Echo용 에러 핸들러
func EchoErrorHandler(err error, c echo.Context) {
	// 1. Echo의 HTTPError인 경우 (예: 404, 400 등)
	if httpErr, ok := err.(*echo.HTTPError); ok {
		response := ErrorResponse{
			Status:    httpErr.Code,             // HTTP 상태 코드
			ErrorCode: "HTTP_ERROR",             // 에러 코드
			Message:   httpErr.Message.(string), // 에러 메시지
		}
		c.JSON(httpErr.Code, response)
		return
	}

	// 2. 기본 에러 처리 (예상치 못한 에러)
	c.JSON(http.StatusInternalServerError, INTERNAL_SERVER_ERROR)
}

// 성공 응답 헬퍼들
func OK(c echo.Context, data interface{}, message ...string) error {
	response := map[string]interface{}{
		"success": true,
		"data":    data,
	}

	if len(message) > 0 && message[0] != "" {
		response["message"] = message[0]
	}

	return c.JSON(http.StatusOK, response)
}

func Created(c echo.Context, data interface{}, message ...string) error {
	response := map[string]interface{}{
		"success": true,
		"data":    data,
	}

	if len(message) > 0 && message[0] != "" {
		response["message"] = message[0]
	}

	return c.JSON(http.StatusCreated, response)
}

func NoContent(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

// 성공응답
// {
//     "success": true,
//     "data": {
//         "id": 1,
//         "name": "John Doe",
//         "email": "john@example.com"
//     },
//     "message": "User found successfully"
// }

// 에러응답
// {
//     "status": 404,
//     "error_code": "HTTP_ERROR",
//     "message": "User not found"
// }
