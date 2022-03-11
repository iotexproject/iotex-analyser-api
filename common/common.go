package common

import (
	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/api/pagination"
	"github.com/pkg/errors"
)

const (
	// HexPrefix is the prefix of ERC20 address in hex string
	HexPrefix = "0x"
	// DefaultPageSize is the size of page when pagination parameters are not set
	DefaultPageSize = 200
	// MaximumPageSize is the maximum size of page
	MaximumPageSize = 50000
)

var (
	// ErrPaginationNotFound is the error indicating that pagination is not specified
	ErrPaginationNotFound = errors.New("pagination information is not found")
	// ErrPaginationInvalidOffset is the error indicating that pagination's offset parameter is invalid
	ErrPaginationInvalidOffset = errors.New("invalid pagination offset number")
	// ErrPaginationInvalidSize is the error indicating that pagination's size parameter is invalid
	ErrPaginationInvalidSize = errors.New("invalid pagination size number")
	// ErrInvalidParameter is the error indicating that invalid size
	ErrInvalidParameter = errors.New("invalid parameter number")
	// ErrActionTypeNotSupported is the error indicating that invalid action type
	ErrActionTypeNotSupported = errors.New("action type is not supported")
)

func PageSize(req *pagination.Pagination) uint64 {
	if req == nil {
		return DefaultPageSize
	}
	if req.GetFirst() > MaximumPageSize {
		return MaximumPageSize
	}
	return req.GetFirst()
}

func PageOffset(req *pagination.Pagination) uint64 {
	if req == nil {
		return 0
	}
	return req.GetSkip()
}

func PageSortBy(req *pagination.Pagination) string {
	if req == nil {
		return "desc"
	}
	return req.GetOrder()
}

func Address(addr string) (*string, error) {
	if addr == "" {
		return nil, address.ErrInvalidAddr
	}
	if addr[:2] == "0x" || addr[:2] == "0X" {
		add, err := address.FromHex(addr)
		if err != nil {
			return nil, err
		}

		addr = add.String()
	}
	return &addr, nil
}
