// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk

import (
	"context"
	"net/http"
	"net/url"

	"github.com/NeuralNetLab/ideogram-sdk-go/internal/apijson"
	"github.com/NeuralNetLab/ideogram-sdk-go/internal/apiquery"
	"github.com/NeuralNetLab/ideogram-sdk-go/internal/requestconfig"
	"github.com/NeuralNetLab/ideogram-sdk-go/option"
	"github.com/NeuralNetLab/ideogram-sdk-go/packages/param"
	"github.com/NeuralNetLab/ideogram-sdk-go/packages/respjson"
)

// ManageAPIService contains methods and other services that help with interacting
// with the ideogram-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewManageAPIService] method instead.
type ManageAPIService struct {
	Options      []option.RequestOption
	APIKeys      ManageAPIAPIKeyService
	Subscription ManageAPISubscriptionService
	Terms        ManageAPITermService
}

// NewManageAPIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewManageAPIService(opts ...option.RequestOption) (r ManageAPIService) {
	r = ManageAPIService{}
	r.Options = opts
	r.APIKeys = NewManageAPIAPIKeyService(opts...)
	r.Subscription = NewManageAPISubscriptionService(opts...)
	r.Terms = NewManageAPITermService(opts...)
	return
}

// Add credits to an API user's account.
func (r *ManageAPIService) AddCredits(ctx context.Context, body ManageAPIAddCreditsParams, opts ...option.RequestOption) (res *ManageAPIAddCreditsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "manage/api/add_credits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Reactivates a subscription by attempting to re-enable Metronome billing.
func (r *ManageAPIService) Reactivate(ctx context.Context, opts ...option.RequestOption) (res *ManageAPIReactivateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "manage/api/reactivate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Retrieve data relevant to connecting to Stripe.
func (r *ManageAPIService) GetStripeSubscription(ctx context.Context, query ManageAPIGetStripeSubscriptionParams, opts ...option.RequestOption) (res *ManageAPIGetStripeSubscriptionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "manage/api/stripe_subscription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Represents a price.
type Price struct {
	// The amount of the currency in the common denomination. For example, in USD this
	// is cents.
	Amount float64 `json:"amount,required"`
	// The ISO 4217 currency code for the price object.
	CurrencyCode string `json:"currency_code,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount       respjson.Field
		CurrencyCode respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Price) RawJSON() string { return r.JSON.raw }
func (r *Price) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Price to a PriceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PriceParam.Overrides()
func (r Price) ToParam() PriceParam {
	return param.Override[PriceParam](r.RawJSON())
}

// Represents a price.
//
// The properties Amount, CurrencyCode are required.
type PriceParam struct {
	// The amount of the currency in the common denomination. For example, in USD this
	// is cents.
	Amount float64 `json:"amount,required"`
	// The ISO 4217 currency code for the price object.
	CurrencyCode string `json:"currency_code,required"`
	paramObj
}

func (r PriceParam) MarshalJSON() (data []byte, err error) {
	type shadow PriceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PriceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current recharge settings for the API subscription.
type RechargeSettingsResponse struct {
	// Whether or not the recharge setting is currently active.
	IsActive bool `json:"is_active,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsActive    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RechargeSettingsSubscription
}

// Returns the unmodified JSON received from the API
func (r RechargeSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *RechargeSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManageAPIAddCreditsResponse struct {
	// Represents a price.
	Amount Price `json:"amount,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManageAPIAddCreditsResponse) RawJSON() string { return r.JSON.raw }
func (r *ManageAPIAddCreditsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to re-activating API settings.
type ManageAPIReactivateResponse struct {
	// The current recharge settings for the API subscription.
	RechargeSettings RechargeSettingsResponse `json:"recharge_settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RechargeSettings respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManageAPIReactivateResponse) RawJSON() string { return r.JSON.raw }
func (r *ManageAPIReactivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManageAPIGetStripeSubscriptionResponse struct {
	// DEPRECATED. The URL for the user to manage the existing Stripe subscription
	// plan. Get this from ManageApiSubscriptionResponse instead.
	//
	// Deprecated: deprecated
	StripeBillingURL string `json:"stripe_billing_url"`
	// The URL for the user to checkout the Stripe subscription plan.
	StripeSubscriptionURL string `json:"stripe_subscription_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StripeBillingURL      respjson.Field
		StripeSubscriptionURL respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManageAPIGetStripeSubscriptionResponse) RawJSON() string { return r.JSON.raw }
func (r *ManageAPIGetStripeSubscriptionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManageAPIAddCreditsParams struct {
	// Represents a price.
	Amount PriceParam `json:"amount,omitzero,required"`
	paramObj
}

func (r ManageAPIAddCreditsParams) MarshalJSON() (data []byte, err error) {
	type shadow ManageAPIAddCreditsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ManageAPIAddCreditsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManageAPIGetStripeSubscriptionParams struct {
	// Whether the subscription is intended to be used for business or personal use.
	IsBusiness param.Opt[bool] `query:"isBusiness,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ManageAPIGetStripeSubscriptionParams]'s query parameters as
// `url.Values`.
func (r ManageAPIGetStripeSubscriptionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
