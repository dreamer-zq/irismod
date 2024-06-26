package types

import (
	"encoding/hex"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewRequest constructs a new Request instance
func NewRequest(
	height int64,
	consumer string,
	txHash string,
	oracle bool,
	serviceFeeCap sdk.Coins,
	serviceContextID string,
) Request {
	return Request{
		Height:           height,
		Consumer:         consumer,
		TxHash:           txHash,
		Oracle:           oracle,
		ServiceFeeCap:    serviceFeeCap,
		ServiceContextID: serviceContextID,
	}
}

// GenerateRequestID generates a request id
func GenerateRequestID(r Request) []byte {
	reqID := make([]byte, 0)

	reqID = append(reqID, sdk.Uint64ToBigEndian(uint64(r.Height))...)
	reqID = append(reqID, []byte(r.Consumer)...)

	return SHA256(reqID)
}

// CheckReqID checks if the given request id is valid
func CheckReqID(reqID string) error {
	if len(reqID) != 64 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid request id: %s", reqID))
	}

	if _, err := hex.DecodeString(reqID); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid request id: %s", reqID))
	}

	return nil
}
