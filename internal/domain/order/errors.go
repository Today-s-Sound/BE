package order

import (
	"net/http"

	"github.com/nicewook/gocore/internal/global/errors"
)

// Order 도메인 에러들
var (
	ORDER_NOT_FOUND = errors.ErrorResponse{
		Status:    http.StatusNotFound,
		ErrorCode: "ORDER_NOT_FOUND",
		Message:   "주문을 찾을 수 없습니다",
	}

	INVALID_ORDER_ID = errors.ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_ORDER_ID",
		Message:   "유효하지 않은 주문 ID입니다",
	}

	INSUFFICIENT_STOCK = errors.ErrorResponse{
		Status:    http.StatusConflict,
		ErrorCode: "INSUFFICIENT_STOCK",
		Message:   "재고가 부족합니다",
	}

	ORDER_ALREADY_PAID = errors.ErrorResponse{
		Status:    http.StatusConflict,
		ErrorCode: "ORDER_ALREADY_PAID",
		Message:   "이미 결제된 주문입니다",
	}

	ORDER_CANCELLED = errors.ErrorResponse{
		Status:    http.StatusConflict,
		ErrorCode: "ORDER_CANCELLED",
		Message:   "취소된 주문입니다",
	}

	INVALID_QUANTITY = errors.ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_QUANTITY",
		Message:   "유효하지 않은 수량입니다",
	}
)
