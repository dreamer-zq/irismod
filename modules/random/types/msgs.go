package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgRequestRandom = "request_random" // type for MsgRequestRandom

	DefaultBlockInterval = uint64(10) // DefaultBlockInterval is the default block interval
)

var _ sdk.Msg = &MsgRequestRandom{}

// NewMsgRequestRandom constructs a new MsgRequestRandom instance
func NewMsgRequestRandom(
	consumer string,
	blockInterval uint64,
	oracle bool,
	serviceFeeCap sdk.Coins,
) *MsgRequestRandom {
	return &MsgRequestRandom{
		Consumer:      consumer,
		BlockInterval: blockInterval,
		Oracle:        oracle,
		ServiceFeeCap: serviceFeeCap,
	}
}

// Route implements Msg.
func (msg MsgRequestRandom) Route() string { return RouterKey }

// Type implements Msg.
func (msg MsgRequestRandom) Type() string { return TypeMsgRequestRandom }

// ValidateBasic implements Msg.
func (msg MsgRequestRandom) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Consumer); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid consumer address (%s)", err)
	}
	return ValidateServiceFeeCap(msg.ServiceFeeCap)
}

// GetSignBytes implements Msg.
func (msg MsgRequestRandom) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners implements Msg.
func (msg MsgRequestRandom) GetSigners() []sdk.AccAddress {
	consumer, err := sdk.AccAddressFromBech32(msg.Consumer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{consumer}
}
