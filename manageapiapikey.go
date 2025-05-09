// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/NeuralNetLab/ideogram-sdk-go/internal/apijson"
	"github.com/NeuralNetLab/ideogram-sdk-go/internal/requestconfig"
	"github.com/NeuralNetLab/ideogram-sdk-go/option"
	"github.com/NeuralNetLab/ideogram-sdk-go/packages/respjson"
)

// ManageAPIAPIKeyService contains methods and other services that help with
// interacting with the ideogram-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewManageAPIAPIKeyService] method instead.
type ManageAPIAPIKeyService struct {
	Options []option.RequestOption
}

// NewManageAPIAPIKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewManageAPIAPIKeyService(opts ...option.RequestOption) (r ManageAPIAPIKeyService) {
	r = ManageAPIAPIKeyService{}
	r.Options = opts
	return
}

// Creates an API key.
func (r *ManageAPIAPIKeyService) New(ctx context.Context, opts ...option.RequestOption) (res *ManageApiapiKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "manage/api/api_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Retrieve current API keys and their respective data.
func (r *ManageAPIAPIKeyService) Get(ctx context.Context, opts ...option.RequestOption) (res *ManageApiapiKeyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "manage/api/api_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an API key.
func (r *ManageAPIAPIKeyService) Delete(ctx context.Context, apiKeyID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if apiKeyID == "" {
		err = errors.New("missing required api_key_id parameter")
		return
	}
	path := fmt.Sprintf("manage/api/api_keys/%s", apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ManageApiapiKeyNewResponse struct {
	// The API key to use when making authenticated requests with the API. This key
	// will only be shown once.
	APIKey string `json:"api_key,required"`
	// The ID of the API key. A URL safe base64 encoded UUID
	APIKeyID string `json:"api_key_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey      respjson.Field
		APIKeyID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManageApiapiKeyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ManageApiapiKeyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManageApiapiKeyGetResponse struct {
	// The current API keys that are active. Only returns redacted keys.
	CurrentAPIKeys []ManageApiapiKeyGetResponseCurrentAPIKey `json:"current_api_keys"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentAPIKeys respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManageApiapiKeyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ManageApiapiKeyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A representation of an API key that does not contain the full key.
type ManageApiapiKeyGetResponseCurrentAPIKey struct {
	// A URL safe base64 encoded UUID
	APIKeyID string `json:"api_key_id,required"`
	// The time at which the key was created
	CreationTime time.Time `json:"creation_time,required" format:"date-time"`
	// A redacted text snippet of the API key. Contains the first 4 characters of the
	// API key
	RedactedAPIKey string `json:"redacted_api_key,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeyID       respjson.Field
		CreationTime   respjson.Field
		RedactedAPIKey respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManageApiapiKeyGetResponseCurrentAPIKey) RawJSON() string { return r.JSON.raw }
func (r *ManageApiapiKeyGetResponseCurrentAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
