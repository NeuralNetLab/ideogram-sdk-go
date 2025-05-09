// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/NeuralNetLab/ideogram-sdk-go/internal/apiform"
	"github.com/NeuralNetLab/ideogram-sdk-go/internal/requestconfig"
	"github.com/NeuralNetLab/ideogram-sdk-go/option"
	"github.com/NeuralNetLab/ideogram-sdk-go/packages/param"
)

// RemixService contains methods and other services that help with interacting with
// the ideogram-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRemixService] method instead.
type RemixService struct {
	Options []option.RequestOption
}

// NewRemixService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRemixService(opts ...option.RequestOption) (r RemixService) {
	r = RemixService{}
	r.Options = opts
	return
}

// Remix provided images synchronously based on a given prompt and optional
// parameters.
//
// Input images are cropped to the chosen aspect ratio before being remixed.
//
// Supported image formats include JPEG, PNG, and WebP.
//
// Images links are available for a limited period of time; if you would like to
// keep the image, you must download it.
func (r *RemixService) New(ctx context.Context, body RemixNewParams, opts ...option.RequestOption) (res *GenerateImage, err error) {
	opts = append(r.Options[:], opts...)
	path := "remix"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RemixNewParams struct {
	// An image binary (max size 10MB); only JPEG, WebP and PNG formats are supported
	// at this time.
	ImageFile io.Reader `json:"image_file,omitzero,required" format:"binary"`
	// A request to generate a new image using a provided image and a prompt.
	ImageRequest RemixNewParamsImageRequest `json:"image_request,omitzero,required"`
	paramObj
}

func (r RemixNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// A request to generate a new image using a provided image and a prompt.
type RemixNewParamsImageRequest struct {
	ImageWeight param.Opt[int64] `json:"image_weight,omitzero"`
	ImageRequestParam
}

func (r RemixNewParamsImageRequest) MarshalJSON() (data []byte, err error) {
	type shadow RemixNewParamsImageRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
