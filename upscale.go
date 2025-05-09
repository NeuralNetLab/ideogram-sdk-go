// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/NeuralNetLab/ideogram-sdk-go/internal/apiform"
	"github.com/NeuralNetLab/ideogram-sdk-go/internal/apijson"
	"github.com/NeuralNetLab/ideogram-sdk-go/internal/requestconfig"
	"github.com/NeuralNetLab/ideogram-sdk-go/option"
	"github.com/NeuralNetLab/ideogram-sdk-go/packages/param"
)

// UpscaleService contains methods and other services that help with interacting
// with the ideogram-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUpscaleService] method instead.
type UpscaleService struct {
	Options []option.RequestOption
}

// NewUpscaleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUpscaleService(opts ...option.RequestOption) (r UpscaleService) {
	r = UpscaleService{}
	r.Options = opts
	return
}

// Upscale provided images synchronously with an optional prompt.
//
// Supported image formats include JPEG, PNG, and WebP.
//
// Images links are available for a limited period of time; if you would like to
// keep the image, you must download it.
func (r *UpscaleService) New(ctx context.Context, body UpscaleNewParams, opts ...option.RequestOption) (res *GenerateImage, err error) {
	opts = append(r.Options[:], opts...)
	path := "upscale"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UpscaleNewParams struct {
	// An image binary (max size 10MB); only JPEG, WebP and PNG formats are supported
	// at this time.
	ImageFile io.Reader `json:"image_file,omitzero,required" format:"binary"`
	// A request to upscale a provided image with the help of an optional prompt.
	ImageRequest UpscaleNewParamsImageRequest `json:"image_request,omitzero,required"`
	paramObj
}

func (r UpscaleNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// A request to upscale a provided image with the help of an optional prompt.
type UpscaleNewParamsImageRequest struct {
	Detail param.Opt[int64] `json:"detail,omitzero"`
	// The number of images to generate.
	NumImages param.Opt[int64] `json:"num_images,omitzero"`
	// An optional prompt to guide the upscale
	Prompt      param.Opt[string] `json:"prompt,omitzero"`
	Resemblance param.Opt[int64]  `json:"resemblance,omitzero"`
	// Random seed. Set for reproducible generation.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// Determine if MagicPrompt should be used in generating the request or not.
	//
	// Any of "AUTO", "ON", "OFF".
	MagicPromptOption MagicPromptOption `json:"magic_prompt_option,omitzero"`
	// The magic prompt version to use when magic prompt option is set to AUTO or ON.
	//
	// Any of "V_0", "V_0_1", "V_0_2".
	MagicPromptVersion MagicPromptVersionEnum `json:"magic_prompt_version,omitzero"`
	paramObj
}

func (r UpscaleNewParamsImageRequest) MarshalJSON() (data []byte, err error) {
	type shadow UpscaleNewParamsImageRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpscaleNewParamsImageRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
