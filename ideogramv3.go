// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/NeuralNetLab/ideogram-sdk-go/internal/apiform"
	"github.com/NeuralNetLab/ideogram-sdk-go/internal/apijson"
	"github.com/NeuralNetLab/ideogram-sdk-go/internal/requestconfig"
	"github.com/NeuralNetLab/ideogram-sdk-go/option"
	"github.com/NeuralNetLab/ideogram-sdk-go/packages/param"
	"github.com/NeuralNetLab/ideogram-sdk-go/packages/respjson"
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

// Replace the background of a given image synchronously using a prompt with
// Ideogram 3.0. The foreground subject will be identified and kept, while the
// background is replaced based on the prompt and chosen style. Supported image
// formats include JPEG, PNG, and WebP. Images links are available for a limited
// period of time; if you would like to keep the image, you must download it.
func (r *IdeogramV3Service) ReplaceBackground(ctx context.Context, body IdeogramV3ReplaceBackgroundParams, opts ...option.RequestOption) (res *ImageGenerationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/ideogram-v3/replace-background"
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

func (u ColorPaletteUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[ColorPaletteUnionParam](u.OfColorPaletteWithPresetName, u.OfColorPaletteWithMembers)
}
func (u *ColorPaletteUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
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

func (r ColorPaletteColorPaletteWithPresetNameParam) MarshalJSON() (data []byte, err error) {
	type shadow ColorPaletteColorPaletteWithPresetNameParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ColorPaletteColorPaletteWithPresetNameParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ColorPaletteColorPaletteWithPresetNameParam](
		"name", "EMBER", "FRESH", "JUNGLE", "MAGIC", "MELON", "MOSAIC", "PASTEL", "ULTRAMARINE",
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

func (r ColorPaletteColorPaletteWithMembersParam) MarshalJSON() (data []byte, err error) {
	type shadow ColorPaletteColorPaletteWithMembersParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ColorPaletteColorPaletteWithMembersParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

func (r ColorPaletteColorPaletteWithMembersMemberParam) MarshalJSON() (data []byte, err error) {
	type shadow ColorPaletteColorPaletteWithMembersMemberParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ColorPaletteColorPaletteWithMembersMemberParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response which contains information about the generated image, including the
// download link. Images links are available for a limited period of time; if you
// would like to keep the image, you must download it.
type ImageGenerationResponse struct {
	// The time the request was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// A list of ImageObjects that contain the generated image(s).
	Data []ImageGenerationResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Created     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
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
	// The resolutions supported for Ideogram 3.0.
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
	Resolution ResolutionIdeogram `json:"resolution,required"`
	// Random seed. Set for reproducible generation.
	Seed int64 `json:"seed,required"`
	// The style type to generate with.
	//
	// Any of "AUTO", "GENERAL", "REALISTIC", "DESIGN".
	StyleType StyleTypeV3 `json:"style_type"`
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
func (r ImageGenerationResponseData) RawJSON() string { return r.JSON.raw }
func (r *ImageGenerationResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The rendering speed to use.
type RenderingSpeed string

const (
	RenderingSpeedTurbo    RenderingSpeed = "TURBO"
	RenderingSpeedBalanced RenderingSpeed = "BALANCED"
	RenderingSpeedDefault  RenderingSpeed = "DEFAULT"
	RenderingSpeedQuality  RenderingSpeed = "QUALITY"
)

// The resolutions supported for Ideogram 3.0.
type ResolutionIdeogram string

const (
	ResolutionIdeogram512x1536  ResolutionIdeogram = "512x1536"
	ResolutionIdeogram576x1408  ResolutionIdeogram = "576x1408"
	ResolutionIdeogram576x1472  ResolutionIdeogram = "576x1472"
	ResolutionIdeogram576x1536  ResolutionIdeogram = "576x1536"
	ResolutionIdeogram640x1344  ResolutionIdeogram = "640x1344"
	ResolutionIdeogram640x1408  ResolutionIdeogram = "640x1408"
	ResolutionIdeogram640x1472  ResolutionIdeogram = "640x1472"
	ResolutionIdeogram640x1536  ResolutionIdeogram = "640x1536"
	ResolutionIdeogram704x1152  ResolutionIdeogram = "704x1152"
	ResolutionIdeogram704x1216  ResolutionIdeogram = "704x1216"
	ResolutionIdeogram704x1280  ResolutionIdeogram = "704x1280"
	ResolutionIdeogram704x1344  ResolutionIdeogram = "704x1344"
	ResolutionIdeogram704x1408  ResolutionIdeogram = "704x1408"
	ResolutionIdeogram704x1472  ResolutionIdeogram = "704x1472"
	ResolutionIdeogram736x1312  ResolutionIdeogram = "736x1312"
	ResolutionIdeogram768x1088  ResolutionIdeogram = "768x1088"
	ResolutionIdeogram768x1216  ResolutionIdeogram = "768x1216"
	ResolutionIdeogram768x1280  ResolutionIdeogram = "768x1280"
	ResolutionIdeogram768x1344  ResolutionIdeogram = "768x1344"
	ResolutionIdeogram800x1280  ResolutionIdeogram = "800x1280"
	ResolutionIdeogram832x960   ResolutionIdeogram = "832x960"
	ResolutionIdeogram832x1024  ResolutionIdeogram = "832x1024"
	ResolutionIdeogram832x1088  ResolutionIdeogram = "832x1088"
	ResolutionIdeogram832x1152  ResolutionIdeogram = "832x1152"
	ResolutionIdeogram832x1216  ResolutionIdeogram = "832x1216"
	ResolutionIdeogram832x1248  ResolutionIdeogram = "832x1248"
	ResolutionIdeogram864x1152  ResolutionIdeogram = "864x1152"
	ResolutionIdeogram896x960   ResolutionIdeogram = "896x960"
	ResolutionIdeogram896x1024  ResolutionIdeogram = "896x1024"
	ResolutionIdeogram896x1088  ResolutionIdeogram = "896x1088"
	ResolutionIdeogram896x1120  ResolutionIdeogram = "896x1120"
	ResolutionIdeogram896x1152  ResolutionIdeogram = "896x1152"
	ResolutionIdeogram960x832   ResolutionIdeogram = "960x832"
	ResolutionIdeogram960x896   ResolutionIdeogram = "960x896"
	ResolutionIdeogram960x1024  ResolutionIdeogram = "960x1024"
	ResolutionIdeogram960x1088  ResolutionIdeogram = "960x1088"
	ResolutionIdeogram1024x832  ResolutionIdeogram = "1024x832"
	ResolutionIdeogram1024x896  ResolutionIdeogram = "1024x896"
	ResolutionIdeogram1024x960  ResolutionIdeogram = "1024x960"
	ResolutionIdeogram1024x1024 ResolutionIdeogram = "1024x1024"
	ResolutionIdeogram1088x768  ResolutionIdeogram = "1088x768"
	ResolutionIdeogram1088x832  ResolutionIdeogram = "1088x832"
	ResolutionIdeogram1088x896  ResolutionIdeogram = "1088x896"
	ResolutionIdeogram1088x960  ResolutionIdeogram = "1088x960"
	ResolutionIdeogram1120x896  ResolutionIdeogram = "1120x896"
	ResolutionIdeogram1152x704  ResolutionIdeogram = "1152x704"
	ResolutionIdeogram1152x832  ResolutionIdeogram = "1152x832"
	ResolutionIdeogram1152x864  ResolutionIdeogram = "1152x864"
	ResolutionIdeogram1152x896  ResolutionIdeogram = "1152x896"
	ResolutionIdeogram1216x704  ResolutionIdeogram = "1216x704"
	ResolutionIdeogram1216x768  ResolutionIdeogram = "1216x768"
	ResolutionIdeogram1216x832  ResolutionIdeogram = "1216x832"
	ResolutionIdeogram1248x832  ResolutionIdeogram = "1248x832"
	ResolutionIdeogram1280x704  ResolutionIdeogram = "1280x704"
	ResolutionIdeogram1280x768  ResolutionIdeogram = "1280x768"
	ResolutionIdeogram1280x800  ResolutionIdeogram = "1280x800"
	ResolutionIdeogram1312x736  ResolutionIdeogram = "1312x736"
	ResolutionIdeogram1344x640  ResolutionIdeogram = "1344x640"
	ResolutionIdeogram1344x704  ResolutionIdeogram = "1344x704"
	ResolutionIdeogram1344x768  ResolutionIdeogram = "1344x768"
	ResolutionIdeogram1408x576  ResolutionIdeogram = "1408x576"
	ResolutionIdeogram1408x640  ResolutionIdeogram = "1408x640"
	ResolutionIdeogram1408x704  ResolutionIdeogram = "1408x704"
	ResolutionIdeogram1472x576  ResolutionIdeogram = "1472x576"
	ResolutionIdeogram1472x640  ResolutionIdeogram = "1472x640"
	ResolutionIdeogram1472x704  ResolutionIdeogram = "1472x704"
	ResolutionIdeogram1536x512  ResolutionIdeogram = "1536x512"
	ResolutionIdeogram1536x576  ResolutionIdeogram = "1536x576"
	ResolutionIdeogram1536x640  ResolutionIdeogram = "1536x640"
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
	// Any of "TURBO", "BALANCED", "DEFAULT", "QUALITY".
	RenderingSpeed RenderingSpeed `json:"rendering_speed,omitzero"`
	// A list of 8 character hexadecimal codes representing the style of the image.
	// Cannot be used in conjunction with style_reference_images or style_type.
	StyleCodes []string `json:"style_codes,omitzero"`
	// A set of images to use as style references (maximum total size 10MB across all
	// style references). The images should be in JPEG, PNG or WebP format.
	StyleReferenceImages []io.Reader `json:"style_reference_images,omitzero" format:"binary"`
	paramObj
}

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
	// Any of "TURBO", "BALANCED", "DEFAULT", "QUALITY".
	RenderingSpeed RenderingSpeed `json:"rendering_speed,omitzero"`
	// The resolutions supported for Ideogram 3.0.
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
	Resolution ResolutionIdeogram `json:"resolution,omitzero"`
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
	// The resolutions supported for Ideogram 3.0.
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
	Resolution ResolutionIdeogram `json:"resolution,omitzero,required"`
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
	// Any of "TURBO", "BALANCED", "DEFAULT", "QUALITY".
	RenderingSpeed RenderingSpeed `json:"rendering_speed,omitzero"`
	// A list of 8 character hexadecimal codes representing the style of the image.
	// Cannot be used in conjunction with style_reference_images or style_type.
	StyleCodes []string `json:"style_codes,omitzero"`
	// A set of images to use as style references (maximum total size 10MB across all
	// style references). The images should be in JPEG, PNG or WebP format.
	StyleReferenceImages []io.Reader `json:"style_reference_images,omitzero" format:"binary"`
	paramObj
}

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
	// Any of "TURBO", "BALANCED", "DEFAULT", "QUALITY".
	RenderingSpeed RenderingSpeed `json:"rendering_speed,omitzero"`
	// The resolutions supported for Ideogram 3.0.
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
	Resolution ResolutionIdeogram `json:"resolution,omitzero"`
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

type IdeogramV3ReplaceBackgroundParams struct {
	// The image whose background is being replaced (max size 10MB); only JPEG, WebP
	// and PNG formats are supported at this time.
	Image io.Reader `json:"image,omitzero,required" format:"binary"`
	// The prompt describing the desired new background.
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
	// Any of "TURBO", "BALANCED", "DEFAULT", "QUALITY".
	RenderingSpeed RenderingSpeed `json:"rendering_speed,omitzero"`
	// A list of 8 character hexadecimal codes representing the style of the image.
	// Cannot be used in conjunction with style_reference_images or style_type.
	StyleCodes []string `json:"style_codes,omitzero"`
	// A set of images to use as style references (maximum total size 10MB across all
	// style references). The images should be in JPEG, PNG or WebP format.
	StyleReferenceImages []io.Reader `json:"style_reference_images,omitzero" format:"binary"`
	paramObj
}

func (r IdeogramV3ReplaceBackgroundParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
