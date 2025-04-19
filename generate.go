// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/ideogram-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/ideogram-sdk-go/option"
	"github.com/stainless-sdks/ideogram-sdk-go/packages/param"
)

// GenerateService contains methods and other services that help with interacting
// with the ideogram-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGenerateService] method instead.
type GenerateService struct {
	Options []option.RequestOption
}

// NewGenerateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGenerateService(opts ...option.RequestOption) (r GenerateService) {
	r = GenerateService{}
	r.Options = opts
	return
}

// Generates images synchronously based on a given prompt and optional parameters.
//
// Images links are available for a limited period of time; if you would like to
// keep the image, you must download it.
func (r *GenerateService) New(ctx context.Context, body GenerateNewParams, opts ...option.RequestOption) (res *GenerateImage, err error) {
	opts = append(r.Options[:], opts...)
	path := "generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The property Prompt is required.
type ImageRequestParam struct {
	// The prompt to use to generate the image.
	Prompt string `json:"prompt,required"`
	// Only available for model versions V_1, V_1_TURBO, V_2 and V_2_TURBO. Description
	// of what to exclude from an image. Descriptions in the prompt take precedence to
	// descriptions in the negative prompt.
	NegativePrompt param.Opt[string] `json:"negative_prompt,omitzero"`
	// The number of images to generate.
	NumImages param.Opt[int64] `json:"num_images,omitzero"`
	// Random seed. Set for reproducible generation.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// (Cannot be used in conjunction with resolution) The aspect ratio to use for
	// image generation, which determines the image's resolution. Defaults to
	// ASPECT_1_1.
	//
	// Any of "ASPECT_10_16", "ASPECT_16_10", "ASPECT_9_16", "ASPECT_16_9",
	// "ASPECT_3_2", "ASPECT_2_3", "ASPECT_4_3", "ASPECT_3_4", "ASPECT_1_1",
	// "ASPECT_1_3", "ASPECT_3_1".
	AspectRatio ImageRequestAspectRatio `json:"aspect_ratio,omitzero"`
	// A color palette for generation, must EITHER be specified via one of the presets
	// (name) or explicitly via hexadecimal representations of the color with optional
	// weights (members). Not supported by V_1, V_1_TURBO, V_2A and V_2A_TURBO models.
	ColorPalette ColorPaletteUnionParam `json:"color_palette,omitzero"`
	// Determine if MagicPrompt should be used in generating the request or not.
	//
	// Any of "AUTO", "ON", "OFF".
	MagicPromptOption MagicPromptOption `json:"magic_prompt_option,omitzero"`
	// The magic prompt version to use when magic prompt option is set to AUTO or ON.
	//
	// Any of "V_0", "V_0_1", "V_0_2".
	MagicPromptVersion MagicPromptVersionEnum `json:"magic_prompt_version,omitzero"`
	// The model used to generate an image or edit one. /generate and /remix supports
	// all model types, however, /edit is only supported for V_2 and V_2_TURBO.
	//
	// Any of "V_1", "V_1_TURBO", "V_2", "V_2_TURBO", "V_2A", "V_2A_TURBO", "V_3".
	Model ModelEnum `json:"model,omitzero"`
	// (For model_version for 2.0 only, cannot be used in conjunction with
	// aspect_ratio) The resolution to use for image generation, represented in width x
	// height. If not specified, defaults to using aspect_ratio.
	//
	// Any of "RESOLUTION_512_1536", "RESOLUTION_576_1408", "RESOLUTION_576_1472",
	// "RESOLUTION_576_1536", "RESOLUTION_640_1024", "RESOLUTION_640_1344",
	// "RESOLUTION_640_1408", "RESOLUTION_640_1472", "RESOLUTION_640_1536",
	// "RESOLUTION_704_1152", "RESOLUTION_704_1216", "RESOLUTION_704_1280",
	// "RESOLUTION_704_1344", "RESOLUTION_704_1408", "RESOLUTION_704_1472",
	// "RESOLUTION_720_1280", "RESOLUTION_736_1312", "RESOLUTION_768_1024",
	// "RESOLUTION_768_1088", "RESOLUTION_768_1152", "RESOLUTION_768_1216",
	// "RESOLUTION_768_1232", "RESOLUTION_768_1280", "RESOLUTION_768_1344",
	// "RESOLUTION_832_960", "RESOLUTION_832_1024", "RESOLUTION_832_1088",
	// "RESOLUTION_832_1152", "RESOLUTION_832_1216", "RESOLUTION_832_1248",
	// "RESOLUTION_864_1152", "RESOLUTION_896_960", "RESOLUTION_896_1024",
	// "RESOLUTION_896_1088", "RESOLUTION_896_1120", "RESOLUTION_896_1152",
	// "RESOLUTION_960_832", "RESOLUTION_960_896", "RESOLUTION_960_1024",
	// "RESOLUTION_960_1088", "RESOLUTION_1024_640", "RESOLUTION_1024_768",
	// "RESOLUTION_1024_832", "RESOLUTION_1024_896", "RESOLUTION_1024_960",
	// "RESOLUTION_1024_1024", "RESOLUTION_1088_768", "RESOLUTION_1088_832",
	// "RESOLUTION_1088_896", "RESOLUTION_1088_960", "RESOLUTION_1120_896",
	// "RESOLUTION_1152_704", "RESOLUTION_1152_768", "RESOLUTION_1152_832",
	// "RESOLUTION_1152_864", "RESOLUTION_1152_896", "RESOLUTION_1216_704",
	// "RESOLUTION_1216_768", "RESOLUTION_1216_832", "RESOLUTION_1232_768",
	// "RESOLUTION_1248_832", "RESOLUTION_1280_704", "RESOLUTION_1280_720",
	// "RESOLUTION_1280_768", "RESOLUTION_1280_800", "RESOLUTION_1312_736",
	// "RESOLUTION_1344_640", "RESOLUTION_1344_704", "RESOLUTION_1344_768",
	// "RESOLUTION_1408_576", "RESOLUTION_1408_640", "RESOLUTION_1408_704",
	// "RESOLUTION_1472_576", "RESOLUTION_1472_640", "RESOLUTION_1472_704",
	// "RESOLUTION_1536_512", "RESOLUTION_1536_576", "RESOLUTION_1536_640".
	Resolution ImageGenerationResolution `json:"resolution,omitzero"`
	// The style type to generate with; this is only applicable for models V_2 and
	// above and should not be specified for model versions V_1.
	//
	// Any of "AUTO", "GENERAL", "REALISTIC", "DESIGN", "RENDER_3D", "ANIME".
	StyleType StyleTypeV2AndAbove `json:"style_type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ImageRequestParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ImageRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ImageRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// (Cannot be used in conjunction with resolution) The aspect ratio to use for
// image generation, which determines the image's resolution. Defaults to
// ASPECT_1_1.
type ImageRequestAspectRatio string

const (
	ImageRequestAspectRatioAspect10_16 ImageRequestAspectRatio = "ASPECT_10_16"
	ImageRequestAspectRatioAspect16_10 ImageRequestAspectRatio = "ASPECT_16_10"
	ImageRequestAspectRatioAspect9_16  ImageRequestAspectRatio = "ASPECT_9_16"
	ImageRequestAspectRatioAspect16_9  ImageRequestAspectRatio = "ASPECT_16_9"
	ImageRequestAspectRatioAspect3_2   ImageRequestAspectRatio = "ASPECT_3_2"
	ImageRequestAspectRatioAspect2_3   ImageRequestAspectRatio = "ASPECT_2_3"
	ImageRequestAspectRatioAspect4_3   ImageRequestAspectRatio = "ASPECT_4_3"
	ImageRequestAspectRatioAspect3_4   ImageRequestAspectRatio = "ASPECT_3_4"
	ImageRequestAspectRatioAspect1_1   ImageRequestAspectRatio = "ASPECT_1_1"
	ImageRequestAspectRatioAspect1_3   ImageRequestAspectRatio = "ASPECT_1_3"
	ImageRequestAspectRatioAspect3_1   ImageRequestAspectRatio = "ASPECT_3_1"
)

// The magic prompt version to use when magic prompt option is set to AUTO or ON.
type MagicPromptVersionEnum string

const (
	MagicPromptVersionEnumV0   MagicPromptVersionEnum = "V_0"
	MagicPromptVersionEnumV0_1 MagicPromptVersionEnum = "V_0_1"
	MagicPromptVersionEnumV0_2 MagicPromptVersionEnum = "V_0_2"
)

type GenerateNewParams struct {
	ImageRequest ImageRequestParam `json:"image_request,omitzero,required"`
	// A list of base64 encoded binary embeddings.
	StyleRefEmbeddings []string `json:"style_ref_embeddings,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GenerateNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r GenerateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GenerateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
