// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/stainless-sdks/ideogram-sdk-go/internal/apiform"
	"github.com/stainless-sdks/ideogram-sdk-go/internal/apijson"
	"github.com/stainless-sdks/ideogram-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/ideogram-sdk-go/option"
	"github.com/stainless-sdks/ideogram-sdk-go/packages/param"
	"github.com/stainless-sdks/ideogram-sdk-go/packages/resp"
)

// DescribeService contains methods and other services that help with interacting
// with the ideogram-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDescribeService] method instead.
type DescribeService struct {
	Options []option.RequestOption
}

// NewDescribeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDescribeService(opts ...option.RequestOption) (r DescribeService) {
	r = DescribeService{}
	r.Options = opts
	return
}

// Describe an image.
//
// Supported image formats include JPEG, PNG, and WebP.
func (r *DescribeService) New(ctx context.Context, body DescribeNewParams, opts ...option.RequestOption) (res *DescribeNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "describe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The response for a describe request encapsulates a list of descriptions.
type DescribeNewResponse struct {
	// A collection of descriptions for given content.
	Descriptions []DescribeNewResponseDescription `json:"descriptions"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Descriptions resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DescribeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DescribeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DescribeNewResponseDescription struct {
	// The generated description for the provided image.
	Text string `json:"text"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Text        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DescribeNewResponseDescription) RawJSON() string { return r.JSON.raw }
func (r *DescribeNewResponseDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DescribeNewParams struct {
	// An image binary (max size 10MB); only JPEG, WebP and PNG formats are supported
	// at this time.
	ImageFile io.Reader `json:"image_file,required" format:"binary"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DescribeNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r DescribeNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
