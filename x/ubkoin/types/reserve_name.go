package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &ReserveName{}

func NewReserveName(owner string, name string) *ReserveName {
	return &ReserveName{
		OwnerAddress: owner,
		Name:         name}
}

func (msg ReserveName) Route() string {
	return RouterKey
}

func (msg ReserveName) Type() string {
	return "ReserveName"
}

func (msg *ReserveName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes ...
func (msg *ReserveName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg *ReserveName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
