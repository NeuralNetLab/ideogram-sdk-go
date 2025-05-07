// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/ideogram-sdk-go/internal/apijson"
	"github.com/stainless-sdks/ideogram-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/ideogram-sdk-go/option"
	"github.com/stainless-sdks/ideogram-sdk-go/packages/param"
	"github.com/stainless-sdks/ideogram-sdk-go/packages/respjson"
)

// ManageAPITermService contains methods and other services that help with
// interacting with the ideogram-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewManageAPITermService] method instead.
type ManageAPITermService struct {
	Options []option.RequestOption
}

// NewManageAPITermService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewManageAPITermService(opts ...option.RequestOption) (r ManageAPITermService) {
	r = ManageAPITermService{}
	r.Options = opts
	return
}

// Retrieve the latest terms of service for API usage.
func (r *ManageAPITermService) Get(ctx context.Context, opts ...option.RequestOption) (res *ManageAPITermGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "manage/api/terms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Accept terms
func (r *ManageAPITermService) Accept(ctx context.Context, body ManageAPITermAcceptParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "manage/api/terms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type ManageAPITermGetResponse struct {
	APITerms ManageAPITermGetResponseAPITerms `json:"api_terms,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APITerms    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManageAPITermGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ManageAPITermGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManageAPITermGetResponseAPITerms struct {
	// The ID of the terms.
	TermsID string `json:"terms_id,required"`
	// The URL where the terms are hosted.
	TermsURL string `json:"terms_url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TermsID     respjson.Field
		TermsURL    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManageAPITermGetResponseAPITerms) RawJSON() string { return r.JSON.raw }
func (r *ManageAPITermGetResponseAPITerms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManageAPITermAcceptParams struct {
	// The ID of the terms which are being accepted.
	TermsID string `json:"terms_id,required"`
	paramObj
}

func (r ManageAPITermAcceptParams) MarshalJSON() (data []byte, err error) {
	type shadow ManageAPITermAcceptParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ManageAPITermAcceptParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
