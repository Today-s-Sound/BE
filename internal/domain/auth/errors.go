package auth

import (
	"net/http"

	"github.com/nicewook/gocore/internal/global/errors"
)

// Auth 도메인 에러들
var (
	INVALID_CREDENTIALS = errors.ErrorResponse{
		Status:    http.StatusUnauthorized,
		ErrorCode: "INVALID_CREDENTIALS",
		Message:   "이메일 또는 비밀번호가 올바르지 않습니다",
	}

	TOKEN_EXPIRED = errors.ErrorResponse{
		Status:    http.StatusUnauthorized,
		ErrorCode: "TOKEN_EXPIRED",
		Message:   "토큰이 만료되었습니다",
	}

	TOKEN_INVALID = errors.ErrorResponse{
		Status:    http.StatusUnauthorized,
		ErrorCode: "TOKEN_INVALID",
		Message:   "유효하지 않은 토큰입니다",
	}

	ACCOUNT_LOCKED = errors.ErrorResponse{
		Status:    http.StatusForbidden,
		ErrorCode: "ACCOUNT_LOCKED",
		Message:   "계정이 잠겼습니다",
	}

	EMAIL_NOT_VERIFIED = errors.ErrorResponse{
		Status:    http.StatusForbidden,
		ErrorCode: "EMAIL_NOT_VERIFIED",
		Message:   "이메일이 인증되지 않았습니다",
	}
)
