// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/ideogram-sdk-go/internal/apijson"
	"github.com/stainless-sdks/ideogram-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/ideogram-sdk-go/option"
	"github.com/stainless-sdks/ideogram-sdk-go/packages/param"
	"github.com/stainless-sdks/ideogram-sdk-go/packages/resp"
)

// ManageAPISubscriptionService contains methods and other services that help with
// interacting with the ideogram-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewManageAPISubscriptionService] method instead.
type ManageAPISubscriptionService struct {
	Options []option.RequestOption
}

// NewManageAPISubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewManageAPISubscriptionService(opts ...option.RequestOption) (r ManageAPISubscriptionService) {
	r = ManageAPISubscriptionService{}
	r.Options = opts
	return
}

// Retrieve data relevant to creating an API subscription.
func (r *ManageAPISubscriptionService) Get(ctx context.Context, opts ...option.RequestOption) (res *ManageAPISubscriptionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "manage/api/subscription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update API subscription settings
func (r *ManageAPISubscriptionService) Update(ctx context.Context, body ManageAPISubscriptionUpdateParams, opts ...option.RequestOption) (res *ManageAPISubscriptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "manage/api/subscription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The current recharge settings for the API subscription.
type RechargeSettingsSubscription struct {
	// Represents a price.
	MinimumBalanceThreshold Price `json:"minimum_balance_threshold,required"`
	// Represents a price.
	TopUpBalance Price `json:"top_up_balance,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		MinimumBalanceThreshold resp.Field
		TopUpBalance            resp.Field
		ExtraFields             map[string]resp.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RechargeSettingsSubscription) RawJSON() string { return r.JSON.raw }
func (r *RechargeSettingsSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RechargeSettingsSubscription to a
// RechargeSettingsSubscriptionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RechargeSettingsSubscriptionParam.IsOverridden()
func (r RechargeSettingsSubscription) ToParam() RechargeSettingsSubscriptionParam {
	return param.OverrideObj[RechargeSettingsSubscriptionParam](r.RawJSON())
}

// The current recharge settings for the API subscription.
//
// The properties MinimumBalanceThreshold, TopUpBalance are required.
type RechargeSettingsSubscriptionParam struct {
	// Represents a price.
	MinimumBalanceThreshold PriceParam `json:"minimum_balance_threshold,omitzero,required"`
	// Represents a price.
	TopUpBalance PriceParam `json:"top_up_balance,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RechargeSettingsSubscriptionParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r RechargeSettingsSubscriptionParam) MarshalJSON() (data []byte, err error) {
	type shadow RechargeSettingsSubscriptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ManageAPISubscriptionGetResponse struct {
	// Whether or not the latest required terms have been accepted.
	HasAcceptedTerms bool `json:"has_accepted_terms,required"`
	// Whether or not Stripe is setup for API usage.
	HasStripeSetup bool `json:"has_stripe_setup,required"`
	// The URL to display the customer usage dashboard, in dark mode.
	//
	// Deprecated: deprecated
	MetronomeDashboardDarkModeURL string `json:"metronome_dashboard_dark_mode_url"`
	// The URL to display the customer usage dashboard.
	//
	// Deprecated: deprecated
	MetronomeDashboardURL string                                         `json:"metronome_dashboard_url"`
	MetronomeLinks        ManageAPISubscriptionGetResponseMetronomeLinks `json:"metronome_links"`
	// The current recharge settings for the API subscription.
	RechargeSettings RechargeSettingsResponse `json:"recharge_settings"`
	// The URL for the user to manage the existing Stripe subscription plan.
	StripeBillingURL string `json:"stripe_billing_url"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		HasAcceptedTerms              resp.Field
		HasStripeSetup                resp.Field
		MetronomeDashboardDarkModeURL resp.Field
		MetronomeDashboardURL         resp.Field
		MetronomeLinks                resp.Field
		RechargeSettings              resp.Field
		StripeBillingURL              resp.Field
		ExtraFields                   map[string]resp.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManageAPISubscriptionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ManageAPISubscriptionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManageAPISubscriptionGetResponseMetronomeLinks struct {
	// The URL to display the customer invoice and API usage.
	CreditsIframeDarkModeURL string `json:"credits_iframe_dark_mode_url"`
	// The URL to display the customer invoice and API usage.
	CreditsIframeURL string `json:"credits_iframe_url"`
	// The URL to display the customer invoice and API usage.
	InvoicesIframeDarkModeURL string `json:"invoices_iframe_dark_mode_url"`
	// The URL to display the customer invoice and API usage.
	InvoicesIframeURL string `json:"invoices_iframe_url"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CreditsIframeDarkModeURL  resp.Field
		CreditsIframeURL          resp.Field
		InvoicesIframeDarkModeURL resp.Field
		InvoicesIframeURL         resp.Field
		ExtraFields               map[string]resp.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManageAPISubscriptionGetResponseMetronomeLinks) RawJSON() string { return r.JSON.raw }
func (r *ManageAPISubscriptionGetResponseMetronomeLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The updated API subscription.
type ManageAPISubscriptionUpdateResponse struct {
	// The current recharge settings for the API subscription.
	RechargeSettings RechargeSettingsResponse `json:"recharge_settings"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		RechargeSettings resp.Field
		ExtraFields      map[string]resp.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManageAPISubscriptionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ManageAPISubscriptionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManageAPISubscriptionUpdateParams struct {
	// The current recharge settings for the API subscription.
	RechargeSettings RechargeSettingsSubscriptionParam `json:"recharge_settings,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ManageAPISubscriptionUpdateParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r ManageAPISubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ManageAPISubscriptionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
