// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/stainless-sdks/ideogram-sdk-go/internal/apiform"
	"github.com/stainless-sdks/ideogram-sdk-go/internal/apijson"
	"github.com/stainless-sdks/ideogram-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/ideogram-sdk-go/option"
	"github.com/stainless-sdks/ideogram-sdk-go/packages/param"
	"github.com/stainless-sdks/ideogram-sdk-go/packages/respjson"
)

// EditService contains methods and other services that help with interacting with
// the ideogram-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEditService] method instead.
type EditService struct {
	Options []option.RequestOption
}

// NewEditService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEditService(opts ...option.RequestOption) (r EditService) {
	r = EditService{}
	r.Options = opts
	return
}

// Edit a given image synchronously using the provided mask. The mask indicates
// which part of the image should be edited, while the prompt and chosen style type
// can further guide the edit.
//
// Supported image formats include JPEG, PNG, and WebP.
//
// Images links are available for a limited period of time; if you would like to
// keep the image, you must download it.
func (r *EditService) Apply(ctx context.Context, body EditApplyParams, opts ...option.RequestOption) (res *GenerateImage, err error) {
	opts = append(r.Options[:], opts...)
	path := "edit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The response which contains information about the generated image, including the
// download link. Images links are available for a limited period of time; if you
// would like to keep the image, you must download it.
type GenerateImage struct {
	// The time the request was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// A list of ImageObjects that contain the generated image(s).
	Data []GenerateImageData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Created     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GenerateImage) RawJSON() string { return r.JSON.raw }
func (r *GenerateImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GenerateImageData struct {
	// Whether this request passes safety checks. If false, the url field will be
	// empty.
	IsImageSafe bool `json:"is_image_safe,required"`
	// The prompt used for the generation. This may be different from the original
	// prompt.
	Prompt string `json:"prompt,required"`
	// The resolution of the final image.
	Resolution string `json:"resolution,required"`
	// Random seed. Set for reproducible generation.
	Seed int64 `json:"seed,required"`
	// The style type to generate with; this is only applicable for models V_2 and
	// above and should not be specified for model versions V_1.
	//
	// Any of "AUTO", "GENERAL", "REALISTIC", "DESIGN", "RENDER_3D", "ANIME".
	StyleType StyleTypeV2AndAbove `json:"style_type"`
	// The direct link to the image generated.
	URL string `json:"url,nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsImageSafe respjson.Field
		Prompt      respjson.Field
		Resolution  respjson.Field
		Seed        respjson.Field
		StyleType   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GenerateImageData) RawJSON() string { return r.JSON.raw }
func (r *GenerateImageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Determine if MagicPrompt should be used in generating the request or not.
type MagicPromptOption string

const (
	MagicPromptOptionAuto MagicPromptOption = "AUTO"
	MagicPromptOptionOn   MagicPromptOption = "ON"
	MagicPromptOptionOff  MagicPromptOption = "OFF"
)

// The model used to generate an image or edit one. /generate and /remix supports
// all model types, however, /edit is only supported for V_2 and V_2_TURBO.
type ModelEnum string

const (
	ModelEnumV1       ModelEnum = "V_1"
	ModelEnumV1Turbo  ModelEnum = "V_1_TURBO"
	ModelEnumV2       ModelEnum = "V_2"
	ModelEnumV2Turbo  ModelEnum = "V_2_TURBO"
	ModelEnumV2A      ModelEnum = "V_2A"
	ModelEnumV2ATurbo ModelEnum = "V_2A_TURBO"
	ModelEnumV3       ModelEnum = "V_3"
)

// The style type to generate with; this is only applicable for models V_2 and
// above and should not be specified for model versions V_1.
type StyleTypeV2AndAbove string

const (
	StyleTypeV2AndAboveAuto      StyleTypeV2AndAbove = "AUTO"
	StyleTypeV2AndAboveGeneral   StyleTypeV2AndAbove = "GENERAL"
	StyleTypeV2AndAboveRealistic StyleTypeV2AndAbove = "REALISTIC"
	StyleTypeV2AndAboveDesign    StyleTypeV2AndAbove = "DESIGN"
	StyleTypeV2AndAboveRender3D  StyleTypeV2AndAbove = "RENDER_3D"
	StyleTypeV2AndAboveAnime     StyleTypeV2AndAbove = "ANIME"
)

type EditApplyParams struct {
	// An image binary (max size 10MB); only JPEG, WebP and PNG formats are supported
	// at this time.
	ImageFile io.Reader `json:"image_file,omitzero,required" format:"binary"`
	// A black and white image of the same size as the image being edited (max size
	// 10MB). Black regions in the mask should match up with the regions of the image
	// that you would like to edit; only JPEG, WebP and PNG formats are supported at
	// this time.
	Mask io.Reader `json:"mask,omitzero,required" format:"binary"`
	// The model used to generate an image or edit one. /generate and /remix supports
	// all model types, however, /edit is only supported for V_2 and V_2_TURBO.
	//
	// Any of "V_1", "V_1_TURBO", "V_2", "V_2_TURBO", "V_2A", "V_2A_TURBO", "V_3".
	Model ModelEnum `json:"model,omitzero,required"`
	// The prompt used to describe the edited result.
	Prompt string `json:"prompt,required"`
	// The number of images to generate.
	NumImages param.Opt[int64] `json:"num_images,omitzero"`
	// Random seed. Set for reproducible generation.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// Determine if MagicPrompt should be used in generating the request or not.
	//
	// Any of "AUTO", "ON", "OFF".
	MagicPromptOption MagicPromptOption `json:"magic_prompt_option,omitzero"`
	// The style type to generate with; this is only applicable for models V_2 and
	// above and should not be specified for model versions V_1.
	//
	// Any of "AUTO", "GENERAL", "REALISTIC", "DESIGN", "RENDER_3D", "ANIME".
	StyleType StyleTypeV2AndAbove `json:"style_type,omitzero"`
	paramObj
}

func (r EditApplyParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
