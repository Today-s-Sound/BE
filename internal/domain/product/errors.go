package product

import (
	"net/http"

	"github.com/nicewook/gocore/internal/global/errors"
)

// Product 도메인 에러들
var (
	PRODUCT_NOT_FOUND = errors.ErrorResponse{
		Status:    http.StatusNotFound,
		ErrorCode: "PRODUCT_NOT_FOUND",
		Message:   "상품을 찾을 수 없습니다",
	}

	PRODUCT_OUT_OF_STOCK = errors.ErrorResponse{
		Status:    http.StatusConflict,
		ErrorCode: "PRODUCT_OUT_OF_STOCK",
		Message:   "상품이 품절되었습니다",
	}

	INVALID_PRODUCT_ID = errors.ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_PRODUCT_ID",
		Message:   "유효하지 않은 상품 ID입니다",
	}

	INVALID_PRICE = errors.ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_PRICE",
		Message:   "유효하지 않은 가격입니다",
	}

	PRODUCT_NAME_EXISTS = errors.ErrorResponse{
		Status:    http.StatusConflict,
		ErrorCode: "PRODUCT_NAME_EXISTS",
		Message:   "이미 존재하는 상품명입니다",
	}
)
