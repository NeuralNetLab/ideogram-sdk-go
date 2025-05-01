// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/stainless-sdks/ideogram-sdk-go/internal/apiform"
	"github.com/stainless-sdks/ideogram-sdk-go/internal/apijson"
	"github.com/stainless-sdks/ideogram-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/ideogram-sdk-go/option"
	"github.com/stainless-sdks/ideogram-sdk-go/packages/param"
	"github.com/stainless-sdks/ideogram-sdk-go/packages/resp"
)

// InternalTestingService contains methods and other services that help with
// interacting with the ideogram-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInternalTestingService] method instead.
type InternalTestingService struct {
	Options []option.RequestOption
}

// NewInternalTestingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInternalTestingService(opts ...option.RequestOption) (r InternalTestingService) {
	r = InternalTestingService{}
	r.Options = opts
	return
}

// Just a testing endpoint
func (r *InternalTestingService) New(ctx context.Context, params InternalTestingNewParams, opts ...option.RequestOption) (res *InternalTestingNewResponse, err error) {
	if !param.IsOmitted(params.XTestHeader) {
		opts = append(opts, option.WithHeader("X-Test-Header", fmt.Sprintf("%s", params.XTestHeader.Value)))
	}
	if !param.IsOmitted(params.XTestHeader2) {
		opts = append(opts, option.WithHeader("X-Test-Header-2", fmt.Sprintf("%s", params.XTestHeader2.Value)))
	}
	opts = append(r.Options[:], opts...)
	path := "internal-testing"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type NestedObjectParam struct {
	PropOne param.Opt[string] `json:"prop_one,omitzero"`
	PropTwo param.Opt[string] `json:"prop_two,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NestedObjectParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r NestedObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow NestedObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type InternalTestingNewResponse struct {
	ResponseContent string `json:"response_content"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ResponseContent resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternalTestingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *InternalTestingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternalTestingNewParams struct {
	RequiredDateTypeField time.Time            `json:"required_date_type_field,required" format:"date"`
	XPosition             param.Opt[int64]     `json:"x_position,omitzero"`
	DateTimeTypeField     param.Opt[time.Time] `json:"date_time_type_field,omitzero" format:"date-time"`
	DateTypeField         param.Opt[time.Time] `json:"date_type_field,omitzero" format:"date"`
	SomeText              param.Opt[string]    `json:"some_text,omitzero"`
	XTestHeader           param.Opt[string]    `header:"X-Test-Header,omitzero" json:"-"`
	XTestHeader2          param.Opt[string]    `header:"X-Test-Header-2,omitzero" json:"-"`
	// An image binary (max size 10MB); only JPEG, WebP and PNG formats are supported
	// at this time.
	AnotherImageFile io.Reader `json:"another_image_file,omitzero" format:"binary"`
	// Any of "EIN", "ZWEI", "DREI".
	EnumTypeField InternalTestingNewParamsEnumTypeField `json:"enum_type_field,omitzero"`
	// An image binary (max size 10MB); only JPEG, WebP and PNG formats are supported
	// at this time.
	ImageFile                  io.Reader                                          `json:"image_file,omitzero" format:"binary"`
	NestedObject               NestedObjectParam                                  `json:"nested_object,omitzero"`
	NestedObjectRequiredFields InternalTestingNewParamsNestedObjectRequiredFields `json:"nested_object_required_fields,omitzero"`
	RepeatedComplexField       []NestedObjectParam                                `json:"repeated_complex_field,omitzero"`
	RepeatedPrimitiveField     []string                                           `json:"repeated_primitive_field,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InternalTestingNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InternalTestingNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type InternalTestingNewParamsEnumTypeField string

const (
	InternalTestingNewParamsEnumTypeFieldEin  InternalTestingNewParamsEnumTypeField = "EIN"
	InternalTestingNewParamsEnumTypeFieldZwei InternalTestingNewParamsEnumTypeField = "ZWEI"
	InternalTestingNewParamsEnumTypeFieldDrei InternalTestingNewParamsEnumTypeField = "DREI"
)

// The properties PropOne, PropTwo are required.
type InternalTestingNewParamsNestedObjectRequiredFields struct {
	PropOne string `json:"prop_one,required"`
	PropTwo string `json:"prop_two,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InternalTestingNewParamsNestedObjectRequiredFields) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InternalTestingNewParamsNestedObjectRequiredFields) MarshalJSON() (data []byte, err error) {
	type shadow InternalTestingNewParamsNestedObjectRequiredFields
	return param.MarshalObject(r, (*shadow)(&r))
}
