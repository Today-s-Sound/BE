package user

import (
	"net/http"

	"github.com/nicewook/gocore/internal/global/errors"
)

// User 도메인 에러들
var (
	USER_NOT_FOUND = errors.ErrorResponse{
		Status:    http.StatusNotFound,
		ErrorCode: "USER_NOT_FOUND",
		Message:   "사용자를 찾을 수 없습니다",
	}

	USER_ALREADY_EXISTS = errors.ErrorResponse{
		Status:    http.StatusConflict,
		ErrorCode: "USER_ALREADY_EXISTS",
		Message:   "이미 존재하는 사용자입니다",
	}

	INVALID_USER_ID = errors.ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_USER_ID",
		Message:   "유효하지 않은 사용자 ID입니다",
	}

	INVALID_EMAIL = errors.ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_EMAIL",
		Message:   "유효하지 않은 이메일 형식입니다",
	}

	WEAK_PASSWORD = errors.ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorCode: "WEAK_PASSWORD",
		Message:   "비밀번호는 최소 8자 이상이어야 합니다",
	}

	USER_INACTIVE = errors.ErrorResponse{
		Status:    http.StatusForbidden,
		ErrorCode: "USER_INACTIVE",
		Message:   "비활성화된 사용자 계정입니다",
	}
)
