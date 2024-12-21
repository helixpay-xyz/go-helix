package userop

import (
	"github.com/ethereum/go-ethereum/common"
)

type UserOperation struct {
	Sender common.Address `json:"sender"               mapstructure:"sender"               validate:"required"`
}
