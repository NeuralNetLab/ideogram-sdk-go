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
	"github.com/stainless-sdks/ideogram-sdk-go/packages/resp"
)

// IdeogramV3Service contains methods and other services that help with interacting
// with the ideogram-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdeogramV3Service] method instead.
type IdeogramV3Service struct {
	Options []option.RequestOption
}

// NewIdeogramV3Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIdeogramV3Service(opts ...option.RequestOption) (r IdeogramV3Service) {
	r = IdeogramV3Service{}
	r.Options = opts
	return
}

// Edit a given image synchronously using the provided mask with Ideogram 3.0. The
// mask indicates which part of the image should be edited, while the prompt and
// chosen style can further guide the edit.
//
// Supported image formats include JPEG, PNG, and WebP.
//
// Images links are available for a limited period of time; if you would like to
// keep the image, you must download it.
func (r *IdeogramV3Service) Edit(ctx context.Context, body IdeogramV3EditParams, opts ...option.RequestOption) (res *ImageGenerationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/ideogram-v3/edit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generates images synchronously based on a given prompt and optional parameters
// using the Ideogram 3.0 model.
//
// Images links are available for a limited period of time; if you would like to
// keep the image, you must download it.
func (r *IdeogramV3Service) Generate(ctx context.Context, body IdeogramV3GenerateParams, opts ...option.RequestOption) (res *ImageGenerationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/ideogram-v3/generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Reframe a square image to a chosen resolution with Ideogram 3.0. The supported
// image formats include JPEG, PNG, and WebP.
//
// Image links are available for a limited period of time; if you would like to
// keep the image, you must download it.
func (r *IdeogramV3Service) Reframe(ctx context.Context, body IdeogramV3ReframeParams, opts ...option.RequestOption) (res *ImageGenerationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/ideogram-v3/reframe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remix provided images synchronously based on a given prompt and optional
// parameters with the Ideogram 3.0 model.
//
// Input images are cropped to the chosen aspect ratio before being remixed.
//
// Supported image formats include JPEG, PNG, and WebP.
//
// Images links are available for a limited period of time; if you would like to
// keep the image, you must download it.
func (r *IdeogramV3Service) Remix(ctx context.Context, body IdeogramV3RemixParams, opts ...option.RequestOption) (res *ImageGenerationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/ideogram-v3/remix"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The aspect ratio to use for image generation, which determines the image's
// resolution. Cannot be used in conjunction with resolution. Defaults to 1x1.
type AspectRatio string

const (
	AspectRatio1x3   AspectRatio = "1x3"
	AspectRatio3x1   AspectRatio = "3x1"
	AspectRatio1x2   AspectRatio = "1x2"
	AspectRatio2x1   AspectRatio = "2x1"
	AspectRatio9x16  AspectRatio = "9x16"
	AspectRatio16x9  AspectRatio = "16x9"
	AspectRatio10x16 AspectRatio = "10x16"
	AspectRatio16x10 AspectRatio = "16x10"
	AspectRatio2x3   AspectRatio = "2x3"
	AspectRatio3x2   AspectRatio = "3x2"
	AspectRatio3x4   AspectRatio = "3x4"
	AspectRatio4x3   AspectRatio = "4x3"
	AspectRatio4x5   AspectRatio = "4x5"
	AspectRatio5x4   AspectRatio = "5x4"
	AspectRatio1x1   AspectRatio = "1x1"
)

func ColorPaletteParamOfColorPaletteWithPresetName(name string) ColorPaletteUnionParam {
	var variant ColorPaletteColorPaletteWithPresetNameParam
	variant.Name = name
	return ColorPaletteUnionParam{OfColorPaletteWithPresetName: &variant}
}

func ColorPaletteParamOfColorPaletteWithMembers(members []ColorPaletteColorPaletteWithMembersMemberParam) ColorPaletteUnionParam {
	var variant ColorPaletteColorPaletteWithMembersParam
	variant.Members = members
	return ColorPaletteUnionParam{OfColorPaletteWithMembers: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ColorPaletteUnionParam struct {
	OfColorPaletteWithPresetName *ColorPaletteColorPaletteWithPresetNameParam `json:",omitzero,inline"`
	OfColorPaletteWithMembers    *ColorPaletteColorPaletteWithMembersParam    `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u ColorPaletteUnionParam) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u ColorPaletteUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[ColorPaletteUnionParam](u.OfColorPaletteWithPresetName, u.OfColorPaletteWithMembers)
}

func (u *ColorPaletteUnionParam) asAny() any {
	if !param.IsOmitted(u.OfColorPaletteWithPresetName) {
		return u.OfColorPaletteWithPresetName
	} else if !param.IsOmitted(u.OfColorPaletteWithMembers) {
		return u.OfColorPaletteWithMembers
	}
	return nil
}

// The property Name is required.
type ColorPaletteColorPaletteWithPresetNameParam struct {
	// A color palette preset value.
	//
	// Any of "EMBER", "FRESH", "JUNGLE", "MAGIC", "MELON", "MOSAIC", "PASTEL",
	// "ULTRAMARINE".
	Name string `json:"name,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ColorPaletteColorPaletteWithPresetNameParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ColorPaletteColorPaletteWithPresetNameParam) MarshalJSON() (data []byte, err error) {
	type shadow ColorPaletteColorPaletteWithPresetNameParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ColorPaletteColorPaletteWithPresetNameParam](
		"Name", false, "EMBER", "FRESH", "JUNGLE", "MAGIC", "MELON", "MOSAIC", "PASTEL", "ULTRAMARINE",
	)
}

// A color palette represented only via its members. Cannot be used in conjunction
// with preset name.
//
// The property Members is required.
type ColorPaletteColorPaletteWithMembersParam struct {
	// A list of ColorPaletteMembers that define the color palette. Each color palette
	// member consists of a required color hex and an optional weight between 0.05 and
	// 1.0 (inclusive). It is recommended that these weights descend from highest to
	// lowest for the color hexes provided.
	Members []ColorPaletteColorPaletteWithMembersMemberParam `json:"members,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ColorPaletteColorPaletteWithMembersParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ColorPaletteColorPaletteWithMembersParam) MarshalJSON() (data []byte, err error) {
	type shadow ColorPaletteColorPaletteWithMembersParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A member of a color palette.
//
// The property ColorHex is required.
type ColorPaletteColorPaletteWithMembersMemberParam struct {
	// The hexadecimal representation of the color with an optional chosen weight.
	ColorHex string `json:"color_hex,required"`
	// The weight of the color in the color palette.
	ColorWeight param.Opt[float64] `json:"color_weight,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ColorPaletteColorPaletteWithMembersMemberParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ColorPaletteColorPaletteWithMembersMemberParam) MarshalJSON() (data []byte, err error) {
	type shadow ColorPaletteColorPaletteWithMembersMemberParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The response which contains information about the generated image, including the
// download link. Images links are available for a limited period of time; if you
// would like to keep the image, you must download it.
type ImageGenerationResponse struct {
	// The time the request was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// A list of ImageObjects that contain the generated image(s).
	Data []ImageGenerationResponseData `json:"data,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Created     resp.Field
		Data        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageGenerationResponse) RawJSON() string { return r.JSON.raw }
func (r *ImageGenerationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageGenerationResponseData struct {
	// Whether this request passes safety checks. If false, the url field will be
	// empty.
	IsImageSafe bool `json:"is_image_safe,required"`
	// The prompt used for the generation. This may be different from the original
	// prompt.
	Prompt string `json:"prompt,required"`
	// The resolutions supported for model version V_3.
	//
	// Any of "512x1536", "576x1408", "576x1472", "576x1536", "640x1344", "640x1408",
	// "640x1472", "640x1536", "704x1152", "704x1216", "704x1280", "704x1344",
	// "704x1408", "704x1472", "736x1312", "768x1088", "768x1216", "768x1280",
	// "768x1344", "800x1280", "832x960", "832x1024", "832x1088", "832x1152",
	// "832x1216", "832x1248", "864x1152", "896x960", "896x1024", "896x1088",
	// "896x1120", "896x1152", "960x832", "960x896", "960x1024", "960x1088",
	// "1024x832", "1024x896", "1024x960", "1024x1024", "1088x768", "1088x832",
	// "1088x896", "1088x960", "1120x896", "1152x704", "1152x832", "1152x864",
	// "1152x896", "1216x704", "1216x768", "1216x832", "1248x832", "1280x704",
	// "1280x768", "1280x800", "1312x736", "1344x640", "1344x704", "1344x768",
	// "1408x576", "1408x640", "1408x704", "1472x576", "1472x640", "1472x704",
	// "1536x512", "1536x576", "1536x640".
	Resolution ResolutionV3 `json:"resolution,required"`
	// Random seed. Set for reproducible generation.
	Seed int64 `json:"seed,required"`
	// The style type to generate with.
	//
	// Any of "AUTO", "GENERAL", "REALISTIC", "DESIGN".
	StyleType StyleTypeV3 `json:"style_type"`
	// The direct link to the image generated.
	URL string `json:"url,nullable" format:"uri"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IsImageSafe resp.Field
		Prompt      resp.Field
		Resolution  resp.Field
		Seed        resp.Field
		StyleType   resp.Field
		URL         resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageGenerationResponseData) RawJSON() string { return r.JSON.raw }
func (r *ImageGenerationResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The rendering speed to use.
type RenderingSpeed string

const (
	RenderingSpeedTurbo    RenderingSpeed = "TURBO"
	RenderingSpeedBalanced RenderingSpeed = "BALANCED"
	RenderingSpeedQuality  RenderingSpeed = "QUALITY"
)

// The resolutions supported for model version V_3.
type ResolutionV3 string

const (
	ResolutionV3512x1536  ResolutionV3 = "512x1536"
	ResolutionV3576x1408  ResolutionV3 = "576x1408"
	ResolutionV3576x1472  ResolutionV3 = "576x1472"
	ResolutionV3576x1536  ResolutionV3 = "576x1536"
	ResolutionV3640x1344  ResolutionV3 = "640x1344"
	ResolutionV3640x1408  ResolutionV3 = "640x1408"
	ResolutionV3640x1472  ResolutionV3 = "640x1472"
	ResolutionV3640x1536  ResolutionV3 = "640x1536"
	ResolutionV3704x1152  ResolutionV3 = "704x1152"
	ResolutionV3704x1216  ResolutionV3 = "704x1216"
	ResolutionV3704x1280  ResolutionV3 = "704x1280"
	ResolutionV3704x1344  ResolutionV3 = "704x1344"
	ResolutionV3704x1408  ResolutionV3 = "704x1408"
	ResolutionV3704x1472  ResolutionV3 = "704x1472"
	ResolutionV3736x1312  ResolutionV3 = "736x1312"
	ResolutionV3768x1088  ResolutionV3 = "768x1088"
	ResolutionV3768x1216  ResolutionV3 = "768x1216"
	ResolutionV3768x1280  ResolutionV3 = "768x1280"
	ResolutionV3768x1344  ResolutionV3 = "768x1344"
	ResolutionV3800x1280  ResolutionV3 = "800x1280"
	ResolutionV3832x960   ResolutionV3 = "832x960"
	ResolutionV3832x1024  ResolutionV3 = "832x1024"
	ResolutionV3832x1088  ResolutionV3 = "832x1088"
	ResolutionV3832x1152  ResolutionV3 = "832x1152"
	ResolutionV3832x1216  ResolutionV3 = "832x1216"
	ResolutionV3832x1248  ResolutionV3 = "832x1248"
	ResolutionV3864x1152  ResolutionV3 = "864x1152"
	ResolutionV3896x960   ResolutionV3 = "896x960"
	ResolutionV3896x1024  ResolutionV3 = "896x1024"
	ResolutionV3896x1088  ResolutionV3 = "896x1088"
	ResolutionV3896x1120  ResolutionV3 = "896x1120"
	ResolutionV3896x1152  ResolutionV3 = "896x1152"
	ResolutionV3960x832   ResolutionV3 = "960x832"
	ResolutionV3960x896   ResolutionV3 = "960x896"
	ResolutionV3960x1024  ResolutionV3 = "960x1024"
	ResolutionV3960x1088  ResolutionV3 = "960x1088"
	ResolutionV31024x832  ResolutionV3 = "1024x832"
	ResolutionV31024x896  ResolutionV3 = "1024x896"
	ResolutionV31024x960  ResolutionV3 = "1024x960"
	ResolutionV31024x1024 ResolutionV3 = "1024x1024"
	ResolutionV31088x768  ResolutionV3 = "1088x768"
	ResolutionV31088x832  ResolutionV3 = "1088x832"
	ResolutionV31088x896  ResolutionV3 = "1088x896"
	ResolutionV31088x960  ResolutionV3 = "1088x960"
	ResolutionV31120x896  ResolutionV3 = "1120x896"
	ResolutionV31152x704  ResolutionV3 = "1152x704"
	ResolutionV31152x832  ResolutionV3 = "1152x832"
	ResolutionV31152x864  ResolutionV3 = "1152x864"
	ResolutionV31152x896  ResolutionV3 = "1152x896"
	ResolutionV31216x704  ResolutionV3 = "1216x704"
	ResolutionV31216x768  ResolutionV3 = "1216x768"
	ResolutionV31216x832  ResolutionV3 = "1216x832"
	ResolutionV31248x832  ResolutionV3 = "1248x832"
	ResolutionV31280x704  ResolutionV3 = "1280x704"
	ResolutionV31280x768  ResolutionV3 = "1280x768"
	ResolutionV31280x800  ResolutionV3 = "1280x800"
	ResolutionV31312x736  ResolutionV3 = "1312x736"
	ResolutionV31344x640  ResolutionV3 = "1344x640"
	ResolutionV31344x704  ResolutionV3 = "1344x704"
	ResolutionV31344x768  ResolutionV3 = "1344x768"
	ResolutionV31408x576  ResolutionV3 = "1408x576"
	ResolutionV31408x640  ResolutionV3 = "1408x640"
	ResolutionV31408x704  ResolutionV3 = "1408x704"
	ResolutionV31472x576  ResolutionV3 = "1472x576"
	ResolutionV31472x640  ResolutionV3 = "1472x640"
	ResolutionV31472x704  ResolutionV3 = "1472x704"
	ResolutionV31536x512  ResolutionV3 = "1536x512"
	ResolutionV31536x576  ResolutionV3 = "1536x576"
	ResolutionV31536x640  ResolutionV3 = "1536x640"
)

// The style type to generate with.
type StyleTypeV3 string

const (
	StyleTypeV3Auto      StyleTypeV3 = "AUTO"
	StyleTypeV3General   StyleTypeV3 = "GENERAL"
	StyleTypeV3Realistic StyleTypeV3 = "REALISTIC"
	StyleTypeV3Design    StyleTypeV3 = "DESIGN"
)

type IdeogramV3EditParams struct {
	// The image being edited (max size 10MB); only JPEG, WebP and PNG formats are
	// supported at this time.
	Image io.Reader `json:"image,omitzero,required" format:"binary"`
	// A black and white image of the same size as the image being edited (max size
	// 10MB). Black regions in the mask should match up with the regions of the image
	// that you would like to edit; only JPEG, WebP and PNG formats are supported at
	// this time.
	Mask io.Reader `json:"mask,omitzero,required" format:"binary"`
	// The prompt used to describe the edited result.
	Prompt string `json:"prompt,required"`
	// The number of images to generate.
	NumImages param.Opt[int64] `json:"num_images,omitzero"`
	// Random seed. Set for reproducible generation.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// A color palette for generation, must EITHER be specified via one of the presets
	// (name) or explicitly via hexadecimal representations of the color with optional
	// weights (members). Not supported by V_1, V_1_TURBO, V_2A and V_2A_TURBO models.
	ColorPalette ColorPaletteUnionParam `json:"color_palette,omitzero"`
	// Determine if MagicPrompt should be used in generating the request or not.
	//
	// Any of "AUTO", "ON", "OFF".
	MagicPrompt MagicPromptOption `json:"magic_prompt,omitzero"`
	// The rendering speed to use.
	//
	// Any of "TURBO", "BALANCED", "QUALITY".
	RenderingSpeed RenderingSpeed `json:"rendering_speed,omitzero"`
	// A list of 8 character hexadecimal codes representing the style of the image.
	// Cannot be used in conjunction with style_reference_images or style_type.
	StyleCodes []string `json:"style_codes,omitzero"`
	// A set of images to use as style references (maximum total size 10MB across all
	// style references). The images should be in JPEG, PNG or WebP format.
	StyleReferenceImages []io.Reader `json:"style_reference_images,omitzero" format:"binary"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IdeogramV3EditParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r IdeogramV3EditParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type IdeogramV3GenerateParams struct {
	// The prompt to use to generate the image.
	Prompt string `json:"prompt,required"`
	// Description of what to exclude from an image. Descriptions in the prompt take
	// precedence to descriptions in the negative prompt.
	NegativePrompt param.Opt[string] `json:"negative_prompt,omitzero"`
	// Number of images to generate.
	NumImages param.Opt[int64] `json:"num_images,omitzero"`
	// Random seed. Set for reproducible generation.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// The aspect ratio to use for image generation, which determines the image's
	// resolution. Cannot be used in conjunction with resolution. Defaults to 1x1.
	//
	// Any of "1x3", "3x1", "1x2", "2x1", "9x16", "16x9", "10x16", "16x10", "2x3",
	// "3x2", "3x4", "4x3", "4x5", "5x4", "1x1".
	AspectRatio AspectRatio `json:"aspect_ratio,omitzero"`
	// A color palette for generation, must EITHER be specified via one of the presets
	// (name) or explicitly via hexadecimal representations of the color with optional
	// weights (members). Not supported by V_1, V_1_TURBO, V_2A and V_2A_TURBO models.
	ColorPalette ColorPaletteUnionParam `json:"color_palette,omitzero"`
	// Determine if MagicPrompt should be used in generating the request or not.
	//
	// Any of "AUTO", "ON", "OFF".
	MagicPrompt MagicPromptOption `json:"magic_prompt,omitzero"`
	// The rendering speed to use.
	//
	// Any of "TURBO", "BALANCED", "QUALITY".
	RenderingSpeed RenderingSpeed `json:"rendering_speed,omitzero"`
	// The resolutions supported for model version V_3.
	//
	// Any of "512x1536", "576x1408", "576x1472", "576x1536", "640x1344", "640x1408",
	// "640x1472", "640x1536", "704x1152", "704x1216", "704x1280", "704x1344",
	// "704x1408", "704x1472", "736x1312", "768x1088", "768x1216", "768x1280",
	// "768x1344", "800x1280", "832x960", "832x1024", "832x1088", "832x1152",
	// "832x1216", "832x1248", "864x1152", "896x960", "896x1024", "896x1088",
	// "896x1120", "896x1152", "960x832", "960x896", "960x1024", "960x1088",
	// "1024x832", "1024x896", "1024x960", "1024x1024", "1088x768", "1088x832",
	// "1088x896", "1088x960", "1120x896", "1152x704", "1152x832", "1152x864",
	// "1152x896", "1216x704", "1216x768", "1216x832", "1248x832", "1280x704",
	// "1280x768", "1280x800", "1312x736", "1344x640", "1344x704", "1344x768",
	// "1408x576", "1408x640", "1408x704", "1472x576", "1472x640", "1472x704",
	// "1536x512", "1536x576", "1536x640".
	Resolution ResolutionV3 `json:"resolution,omitzero"`
	// A list of 8 character hexadecimal codes representing the style of the image.
	// Cannot be used in conjunction with style_reference_images or style_type.
	StyleCodes []string `json:"style_codes,omitzero"`
	// A set of images to use as style references (maximum total size 10MB across all
	// style references). The images should be in JPEG, PNG or WebP format.
	StyleReferenceImages []io.Reader `json:"style_reference_images,omitzero" format:"binary"`
	// The style type to generate with.
	//
	// Any of "AUTO", "GENERAL", "REALISTIC", "DESIGN".
	StyleType StyleTypeV3 `json:"style_type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IdeogramV3GenerateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r IdeogramV3GenerateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type IdeogramV3ReframeParams struct {
	// The image being reframed (max size 10MB); only JPEG, WebP and PNG formats are
	// supported at this time.
	Image io.Reader `json:"image,omitzero,required" format:"binary"`
	// The resolutions supported for model version V_3.
	//
	// Any of "512x1536", "576x1408", "576x1472", "576x1536", "640x1344", "640x1408",
	// "640x1472", "640x1536", "704x1152", "704x1216", "704x1280", "704x1344",
	// "704x1408", "704x1472", "736x1312", "768x1088", "768x1216", "768x1280",
	// "768x1344", "800x1280", "832x960", "832x1024", "832x1088", "832x1152",
	// "832x1216", "832x1248", "864x1152", "896x960", "896x1024", "896x1088",
	// "896x1120", "896x1152", "960x832", "960x896", "960x1024", "960x1088",
	// "1024x832", "1024x896", "1024x960", "1024x1024", "1088x768", "1088x832",
	// "1088x896", "1088x960", "1120x896", "1152x704", "1152x832", "1152x864",
	// "1152x896", "1216x704", "1216x768", "1216x832", "1248x832", "1280x704",
	// "1280x768", "1280x800", "1312x736", "1344x640", "1344x704", "1344x768",
	// "1408x576", "1408x640", "1408x704", "1472x576", "1472x640", "1472x704",
	// "1536x512", "1536x576", "1536x640".
	Resolution ResolutionV3 `json:"resolution,omitzero,required"`
	// The number of images to generate.
	NumImages param.Opt[int64] `json:"num_images,omitzero"`
	// Random seed. Set for reproducible generation.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// A color palette for generation, must EITHER be specified via one of the presets
	// (name) or explicitly via hexadecimal representations of the color with optional
	// weights (members). Not supported by V_1, V_1_TURBO, V_2A and V_2A_TURBO models.
	ColorPalette ColorPaletteUnionParam `json:"color_palette,omitzero"`
	// The rendering speed to use.
	//
	// Any of "TURBO", "BALANCED", "QUALITY".
	RenderingSpeed RenderingSpeed `json:"rendering_speed,omitzero"`
	// A list of 8 character hexadecimal codes representing the style of the image.
	// Cannot be used in conjunction with style_reference_images or style_type.
	StyleCodes []string `json:"style_codes,omitzero"`
	// A set of images to use as style references (maximum total size 10MB across all
	// style references). The images should be in JPEG, PNG or WebP format.
	StyleReferenceImages []io.Reader `json:"style_reference_images,omitzero" format:"binary"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IdeogramV3ReframeParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r IdeogramV3ReframeParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type IdeogramV3RemixParams struct {
	// The image to remix binary (max size 10MB); only JPEG, WebP and PNG formats are
	// supported at this time.
	Image io.Reader `json:"image,omitzero,required" format:"binary"`
	// The prompt to use to generate the image.
	Prompt      string           `json:"prompt,required"`
	ImageWeight param.Opt[int64] `json:"image_weight,omitzero"`
	// Description of what to exclude from an image. Descriptions in the prompt take
	// precedence to descriptions in the negative prompt.
	NegativePrompt param.Opt[string] `json:"negative_prompt,omitzero"`
	// Number of images to generate.
	NumImages param.Opt[int64] `json:"num_images,omitzero"`
	// Random seed. Set for reproducible generation.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// The aspect ratio to use for image generation, which determines the image's
	// resolution. Cannot be used in conjunction with resolution. Defaults to 1x1.
	//
	// Any of "1x3", "3x1", "1x2", "2x1", "9x16", "16x9", "10x16", "16x10", "2x3",
	// "3x2", "3x4", "4x3", "4x5", "5x4", "1x1".
	AspectRatio AspectRatio `json:"aspect_ratio,omitzero"`
	// A color palette for generation, must EITHER be specified via one of the presets
	// (name) or explicitly via hexadecimal representations of the color with optional
	// weights (members). Not supported by V_1, V_1_TURBO, V_2A and V_2A_TURBO models.
	ColorPalette ColorPaletteUnionParam `json:"color_palette,omitzero"`
	// Determine if MagicPrompt should be used in generating the request or not.
	//
	// Any of "AUTO", "ON", "OFF".
	MagicPrompt MagicPromptOption `json:"magic_prompt,omitzero"`
	// The rendering speed to use.
	//
	// Any of "TURBO", "BALANCED", "QUALITY".
	RenderingSpeed RenderingSpeed `json:"rendering_speed,omitzero"`
	// The resolutions supported for model version V_3.
	//
	// Any of "512x1536", "576x1408", "576x1472", "576x1536", "640x1344", "640x1408",
	// "640x1472", "640x1536", "704x1152", "704x1216", "704x1280", "704x1344",
	// "704x1408", "704x1472", "736x1312", "768x1088", "768x1216", "768x1280",
	// "768x1344", "800x1280", "832x960", "832x1024", "832x1088", "832x1152",
	// "832x1216", "832x1248", "864x1152", "896x960", "896x1024", "896x1088",
	// "896x1120", "896x1152", "960x832", "960x896", "960x1024", "960x1088",
	// "1024x832", "1024x896", "1024x960", "1024x1024", "1088x768", "1088x832",
	// "1088x896", "1088x960", "1120x896", "1152x704", "1152x832", "1152x864",
	// "1152x896", "1216x704", "1216x768", "1216x832", "1248x832", "1280x704",
	// "1280x768", "1280x800", "1312x736", "1344x640", "1344x704", "1344x768",
	// "1408x576", "1408x640", "1408x704", "1472x576", "1472x640", "1472x704",
	// "1536x512", "1536x576", "1536x640".
	Resolution ResolutionV3 `json:"resolution,omitzero"`
	// A list of 8 character hexadecimal codes representing the style of the image.
	// Cannot be used in conjunction with style_reference_images or style_type.
	StyleCodes []string `json:"style_codes,omitzero"`
	// A set of images to use as style references (maximum total size 10MB across all
	// style references). The images should be in JPEG, PNG or WebP format.
	StyleReferenceImages []io.Reader `json:"style_reference_images,omitzero" format:"binary"`
	// The style type to generate with.
	//
	// Any of "AUTO", "GENERAL", "REALISTIC", "DESIGN".
	StyleType StyleTypeV3 `json:"style_type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IdeogramV3RemixParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r IdeogramV3RemixParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
