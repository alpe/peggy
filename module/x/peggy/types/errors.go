package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalid = sdkerrors.Register(ModuleName, 2, "invalid")
	ErrUnknown = sdkerrors.Register(ModuleName, 3, "unknown")
)
