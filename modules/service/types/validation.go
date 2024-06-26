package types

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// MaxNameLength length of the service name
	MaxNameLength = 70
	// MaxDescriptionLength length of the service and author description
	MaxDescriptionLength = 280
	// MaxTagsNum total number of the tags
	MaxTagsNum = 10
	// MaxTagLength length of the tag
	MaxTagLength = 70
	// MaxProvidersNum total number of the providers to request
	MaxProvidersNum = 10
)

var (
	// the service name only accepts alphanumeric characters, _ and -, beginning with alpha character
	regexpServiceName = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]*$`)
	regexpTag         = regexp.MustCompile(`^[\S]{1,70}$`)
)

// ValidateAuthor verifies whether the  parameters are legal
func ValidateAuthor(author string) error {
	if _, err := sdk.AccAddressFromBech32(author); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid author address (%s)", err)
	}
	return nil
}

// ValidateServiceName validates the service name
func ValidateServiceName(name string) error {
	if !regexpServiceName.MatchString(name) || len(name) > MaxNameLength {
		return errorsmod.Wrap(ErrInvalidServiceName, name)
	}
	return nil
}

// ValidateTags verifies whether the given tags are legal
func ValidateTags(tags []string) error {
	if len(tags) > MaxTagsNum {
		return errorsmod.Wrap(ErrInvalidTags, fmt.Sprintf("invalid tags size; got: %d, max: %d", len(tags), MaxTagsNum))
	}
	if HasDuplicate(tags) {
		return errorsmod.Wrap(ErrInvalidTags, "duplicate tag")
	}
	for i, tag := range tags {
		if !regexpTag.MatchString(tag) {
			return errorsmod.Wrapf(ErrInvalidTags, "tag [%d] is invalid, must match regexp: %s", i, regexpTag.String())
		}
	}
	return nil
}

// ValidateServiceDescription verifies whether the  parameters are legal
func ValidateServiceDescription(svcDescription string) error {
	if len(svcDescription) > MaxDescriptionLength {
		return errorsmod.Wrap(ErrInvalidDescription, fmt.Sprintf("invalid service description length; got: %d, max: %d", len(svcDescription), MaxDescriptionLength))
	}
	return nil
}

// ValidateAuthorDescription verifies whether the  parameters are legal
func ValidateAuthorDescription(authorDescription string) error {
	if len(authorDescription) > MaxDescriptionLength {
		return errorsmod.Wrap(ErrInvalidDescription, fmt.Sprintf("invalid author description length; got: %d, max: %d", len(authorDescription), MaxDescriptionLength))
	}
	return nil
}

// ValidateProvider verifies whether the  parameters are legal
func ValidateProvider(provider string) error {
	if _, err := sdk.AccAddressFromBech32(provider); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid provider address (%s)", err)
	}
	return nil
}

// ValidateOwner verifies whether the  parameters are legal
func ValidateOwner(owner string) error {
	if _, err := sdk.AccAddressFromBech32(owner); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

// ValidateServiceDeposit verifies whether the  parameters are legal
func ValidateServiceDeposit(deposit sdk.Coins) error {
	if !deposit.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "invalid deposit")
	}
	if deposit.IsAnyNegative() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "invalid deposit")
	}
	return nil
}

func ValidatePricing(pricing string) error {
	if err := ValidateBindingPricing(pricing); err != nil {
		return err
	}

	parsedPricing, err := ParsePricing(pricing)
	if err != nil {
		return err
	}

	return CheckPricing(parsedPricing)
}

// ValidateQoS verifies whether the  parameters are legal
func ValidateQoS(qos uint64) error {
	if qos == 0 {
		return errorsmod.Wrap(ErrInvalidQoS, "qos must be greater than 0")
	}
	return nil
}

// ValidateOptions verifies whether the  parameters are legal
func ValidateOptions(options string) error {
	if !json.Valid([]byte(options)) {
		return errorsmod.Wrap(ErrInvalidOptions, "options is not valid JSON")
	}
	return nil
}

// ValidateWithdrawAddress verifies whether the  parameters are legal
func ValidateWithdrawAddress(withdrawAddress string) error {
	if _, err := sdk.AccAddressFromBech32(withdrawAddress); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid withdrawal address (%s)", err)
	}
	return nil
}

// ______________________________________________________________________

// ValidateRequest validates the request params
func ValidateRequest(
	serviceName string,
	serviceFeeCap sdk.Coins,
	providers []sdk.AccAddress,
	input string,
	timeout int64,
	repeated bool,
	repeatedFrequency uint64,
	repeatedTotal int64,
) error {
	if err := ValidateServiceName(serviceName); err != nil {
		return err
	}
	if err := ValidateServiceFeeCap(serviceFeeCap); err != nil {
		return err
	}
	if err := ValidateProviders(providers); err != nil {
		return err
	}
	if err := ValidateInput(input); err != nil {
		return err
	}
	if timeout <= 0 {
		return errorsmod.Wrapf(ErrInvalidTimeout, "timeout [%d] must be greater than 0", timeout)
	}
	if repeated {
		if repeatedFrequency > 0 && repeatedFrequency < uint64(timeout) {
			return errorsmod.Wrapf(ErrInvalidRepeatedFreq, "repeated frequency [%d] must not be less than timeout [%d]", repeatedFrequency, timeout)
		}
		if repeatedTotal < -1 || repeatedTotal == 0 {
			return errorsmod.Wrapf(ErrInvalidRepeatedTotal, "repeated total number [%d] must be greater than 0 or equal to -1", repeatedTotal)
		}
	}
	return nil
}

// ValidateRequestContextUpdating validates the request context updating operation
func ValidateRequestContextUpdating(
	providers []sdk.AccAddress,
	serviceFeeCap sdk.Coins,
	timeout int64,
	repeatedFrequency uint64,
	repeatedTotal int64,
) error {
	if err := ValidateProvidersCanEmpty(providers); err != nil {
		return err
	}
	if !serviceFeeCap.Empty() {
		if err := ValidateServiceFeeCap(serviceFeeCap); err != nil {
			return err
		}
	}
	if timeout < 0 {
		return errorsmod.Wrapf(ErrInvalidTimeout, "timeout must not be less than 0: %d", timeout)
	}
	if timeout != 0 && repeatedFrequency != 0 && repeatedFrequency < uint64(timeout) {
		return errorsmod.Wrapf(ErrInvalidRepeatedFreq, "frequency [%d] must not be less than timeout [%d]", repeatedFrequency, timeout)
	}
	if repeatedTotal < -1 {
		return errorsmod.Wrapf(ErrInvalidRepeatedFreq, "repeated total number must not be less than -1: %d", repeatedTotal)
	}
	return nil
}

// ValidateConsumer verifies whether the  parameters are legal
func ValidateConsumer(consumer string) error {
	if _, err := sdk.AccAddressFromBech32(consumer); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid consumer address (%s)", err)
	}
	return nil
}

// ValidateProviders verifies whether the  parameters are legal
func ValidateProviders(providers []sdk.AccAddress) error {
	if len(providers) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "providers missing")
	}
	if len(providers) > MaxProvidersNum {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "total number of the providers must not be greater than %d", MaxProvidersNum)
	}
	if err := checkDuplicateProviders(providers); err != nil {
		return err
	}
	return nil
}

// ValidateProvidersCanEmpty verifies whether the  parameters are legal
func ValidateProvidersCanEmpty(providers []sdk.AccAddress) error {
	if len(providers) > MaxProvidersNum {
		return errorsmod.Wrapf(ErrInvalidProviders, "total number of the providers must not be greater than %d", MaxProvidersNum)
	}
	if len(providers) > 0 {
		if err := checkDuplicateProviders(providers); err != nil {
			return err
		}
	}
	return nil
}

// ValidateServiceFeeCap verifies whether the  parameters are legal
func ValidateServiceFeeCap(serviceFeeCap sdk.Coins) error {
	if !serviceFeeCap.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid service fee cap: %s", serviceFeeCap))
	}
	return nil
}

// ValidateRequestID verifies whether the  parameters are legal
func ValidateRequestID(reqID string) error {
	if len(reqID) != RequestIDLen {
		return errorsmod.Wrapf(ErrInvalidRequestID, "length of the request ID must be %d", RequestIDLen)
	}
	if _, err := hex.DecodeString(reqID); err != nil {
		return errorsmod.Wrap(ErrInvalidRequestID, "request ID must be a hex encoded string")
	}
	return nil
}

// ValidateContextID verifies whether the  parameters are legal
func ValidateContextID(contextID string) error {
	if len(contextID) != ContextIDLen {
		return errorsmod.Wrapf(ErrInvalidRequestContextID, "length of the request context ID must be %d in bytes", ContextIDLen)
	}
	if _, err := hex.DecodeString(contextID); err != nil {
		return errorsmod.Wrap(ErrInvalidRequestContextID, "request context ID must be a hex encoded string")
	}
	return nil
}

// ValidateInput verifies whether the  parameters are legal
func ValidateInput(input string) error {
	if len(input) == 0 {
		return errorsmod.Wrap(ErrInvalidRequestInput, "input missing")
	}

	if ValidateRequestInput(input) != nil {
		return errorsmod.Wrap(ErrInvalidRequestInput, "invalid input")
	}

	return nil
}

// ValidateOutput verifies whether the  parameters are legal
func ValidateOutput(code ResultCode, output string) error {
	if code == ResultOK && len(output) == 0 {
		return errorsmod.Wrapf(ErrInvalidResponse, "output must be specified when the result code is %v", ResultOK)
	}

	if code != ResultOK && len(output) != 0 {
		return errorsmod.Wrapf(ErrInvalidResponse, "output should not be specified when the result code is not %v", ResultOK)
	}

	if len(output) > 0 && ValidateResponseOutput(output) != nil {
		return errorsmod.Wrap(ErrInvalidResponse, "invalid output")
	}

	return nil
}

func checkDuplicateProviders(providers []sdk.AccAddress) error {
	providerArr := make([]string, len(providers))

	for i, provider := range providers {
		providerArr[i] = provider.String()
	}

	if HasDuplicate(providerArr) {
		return errorsmod.Wrap(ErrInvalidProviders, "there exists duplicate providers")
	}

	return nil
}
