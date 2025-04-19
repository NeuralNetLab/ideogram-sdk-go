// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/stainless-sdks/ideogram-sdk-go/internal/apiform"
	"github.com/stainless-sdks/ideogram-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/ideogram-sdk-go/option"
	"github.com/stainless-sdks/ideogram-sdk-go/packages/param"
)

// ReframeService contains methods and other services that help with interacting
// with the ideogram-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReframeService] method instead.
type ReframeService struct {
	Options []option.RequestOption
}

// NewReframeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewReframeService(opts ...option.RequestOption) (r ReframeService) {
	r = ReframeService{}
	r.Options = opts
	return
}

// Reframe a square image to a chosen resolution. The supported image formats
// include JPEG, PNG, and WebP.
//
// Image links are available for a limited period of time; if you would like to
// keep the image, you must download it.
func (r *ReframeService) New(ctx context.Context, body ReframeNewParams, opts ...option.RequestOption) (res *GenerateImage, err error) {
	opts = append(r.Options[:], opts...)
	path := "reframe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// (For model_version for 2.0 only, cannot be used in conjunction with
// aspect_ratio) The resolution to use for image generation, represented in width x
// height. If not specified, defaults to using aspect_ratio.
type ImageGenerationResolution string

const (
	ImageGenerationResolutionResolution512_1536  ImageGenerationResolution = "RESOLUTION_512_1536"
	ImageGenerationResolutionResolution576_1408  ImageGenerationResolution = "RESOLUTION_576_1408"
	ImageGenerationResolutionResolution576_1472  ImageGenerationResolution = "RESOLUTION_576_1472"
	ImageGenerationResolutionResolution576_1536  ImageGenerationResolution = "RESOLUTION_576_1536"
	ImageGenerationResolutionResolution640_1024  ImageGenerationResolution = "RESOLUTION_640_1024"
	ImageGenerationResolutionResolution640_1344  ImageGenerationResolution = "RESOLUTION_640_1344"
	ImageGenerationResolutionResolution640_1408  ImageGenerationResolution = "RESOLUTION_640_1408"
	ImageGenerationResolutionResolution640_1472  ImageGenerationResolution = "RESOLUTION_640_1472"
	ImageGenerationResolutionResolution640_1536  ImageGenerationResolution = "RESOLUTION_640_1536"
	ImageGenerationResolutionResolution704_1152  ImageGenerationResolution = "RESOLUTION_704_1152"
	ImageGenerationResolutionResolution704_1216  ImageGenerationResolution = "RESOLUTION_704_1216"
	ImageGenerationResolutionResolution704_1280  ImageGenerationResolution = "RESOLUTION_704_1280"
	ImageGenerationResolutionResolution704_1344  ImageGenerationResolution = "RESOLUTION_704_1344"
	ImageGenerationResolutionResolution704_1408  ImageGenerationResolution = "RESOLUTION_704_1408"
	ImageGenerationResolutionResolution704_1472  ImageGenerationResolution = "RESOLUTION_704_1472"
	ImageGenerationResolutionResolution720_1280  ImageGenerationResolution = "RESOLUTION_720_1280"
	ImageGenerationResolutionResolution736_1312  ImageGenerationResolution = "RESOLUTION_736_1312"
	ImageGenerationResolutionResolution768_1024  ImageGenerationResolution = "RESOLUTION_768_1024"
	ImageGenerationResolutionResolution768_1088  ImageGenerationResolution = "RESOLUTION_768_1088"
	ImageGenerationResolutionResolution768_1152  ImageGenerationResolution = "RESOLUTION_768_1152"
	ImageGenerationResolutionResolution768_1216  ImageGenerationResolution = "RESOLUTION_768_1216"
	ImageGenerationResolutionResolution768_1232  ImageGenerationResolution = "RESOLUTION_768_1232"
	ImageGenerationResolutionResolution768_1280  ImageGenerationResolution = "RESOLUTION_768_1280"
	ImageGenerationResolutionResolution768_1344  ImageGenerationResolution = "RESOLUTION_768_1344"
	ImageGenerationResolutionResolution832_960   ImageGenerationResolution = "RESOLUTION_832_960"
	ImageGenerationResolutionResolution832_1024  ImageGenerationResolution = "RESOLUTION_832_1024"
	ImageGenerationResolutionResolution832_1088  ImageGenerationResolution = "RESOLUTION_832_1088"
	ImageGenerationResolutionResolution832_1152  ImageGenerationResolution = "RESOLUTION_832_1152"
	ImageGenerationResolutionResolution832_1216  ImageGenerationResolution = "RESOLUTION_832_1216"
	ImageGenerationResolutionResolution832_1248  ImageGenerationResolution = "RESOLUTION_832_1248"
	ImageGenerationResolutionResolution864_1152  ImageGenerationResolution = "RESOLUTION_864_1152"
	ImageGenerationResolutionResolution896_960   ImageGenerationResolution = "RESOLUTION_896_960"
	ImageGenerationResolutionResolution896_1024  ImageGenerationResolution = "RESOLUTION_896_1024"
	ImageGenerationResolutionResolution896_1088  ImageGenerationResolution = "RESOLUTION_896_1088"
	ImageGenerationResolutionResolution896_1120  ImageGenerationResolution = "RESOLUTION_896_1120"
	ImageGenerationResolutionResolution896_1152  ImageGenerationResolution = "RESOLUTION_896_1152"
	ImageGenerationResolutionResolution960_832   ImageGenerationResolution = "RESOLUTION_960_832"
	ImageGenerationResolutionResolution960_896   ImageGenerationResolution = "RESOLUTION_960_896"
	ImageGenerationResolutionResolution960_1024  ImageGenerationResolution = "RESOLUTION_960_1024"
	ImageGenerationResolutionResolution960_1088  ImageGenerationResolution = "RESOLUTION_960_1088"
	ImageGenerationResolutionResolution1024_640  ImageGenerationResolution = "RESOLUTION_1024_640"
	ImageGenerationResolutionResolution1024_768  ImageGenerationResolution = "RESOLUTION_1024_768"
	ImageGenerationResolutionResolution1024_832  ImageGenerationResolution = "RESOLUTION_1024_832"
	ImageGenerationResolutionResolution1024_896  ImageGenerationResolution = "RESOLUTION_1024_896"
	ImageGenerationResolutionResolution1024_960  ImageGenerationResolution = "RESOLUTION_1024_960"
	ImageGenerationResolutionResolution1024_1024 ImageGenerationResolution = "RESOLUTION_1024_1024"
	ImageGenerationResolutionResolution1088_768  ImageGenerationResolution = "RESOLUTION_1088_768"
	ImageGenerationResolutionResolution1088_832  ImageGenerationResolution = "RESOLUTION_1088_832"
	ImageGenerationResolutionResolution1088_896  ImageGenerationResolution = "RESOLUTION_1088_896"
	ImageGenerationResolutionResolution1088_960  ImageGenerationResolution = "RESOLUTION_1088_960"
	ImageGenerationResolutionResolution1120_896  ImageGenerationResolution = "RESOLUTION_1120_896"
	ImageGenerationResolutionResolution1152_704  ImageGenerationResolution = "RESOLUTION_1152_704"
	ImageGenerationResolutionResolution1152_768  ImageGenerationResolution = "RESOLUTION_1152_768"
	ImageGenerationResolutionResolution1152_832  ImageGenerationResolution = "RESOLUTION_1152_832"
	ImageGenerationResolutionResolution1152_864  ImageGenerationResolution = "RESOLUTION_1152_864"
	ImageGenerationResolutionResolution1152_896  ImageGenerationResolution = "RESOLUTION_1152_896"
	ImageGenerationResolutionResolution1216_704  ImageGenerationResolution = "RESOLUTION_1216_704"
	ImageGenerationResolutionResolution1216_768  ImageGenerationResolution = "RESOLUTION_1216_768"
	ImageGenerationResolutionResolution1216_832  ImageGenerationResolution = "RESOLUTION_1216_832"
	ImageGenerationResolutionResolution1232_768  ImageGenerationResolution = "RESOLUTION_1232_768"
	ImageGenerationResolutionResolution1248_832  ImageGenerationResolution = "RESOLUTION_1248_832"
	ImageGenerationResolutionResolution1280_704  ImageGenerationResolution = "RESOLUTION_1280_704"
	ImageGenerationResolutionResolution1280_720  ImageGenerationResolution = "RESOLUTION_1280_720"
	ImageGenerationResolutionResolution1280_768  ImageGenerationResolution = "RESOLUTION_1280_768"
	ImageGenerationResolutionResolution1280_800  ImageGenerationResolution = "RESOLUTION_1280_800"
	ImageGenerationResolutionResolution1312_736  ImageGenerationResolution = "RESOLUTION_1312_736"
	ImageGenerationResolutionResolution1344_640  ImageGenerationResolution = "RESOLUTION_1344_640"
	ImageGenerationResolutionResolution1344_704  ImageGenerationResolution = "RESOLUTION_1344_704"
	ImageGenerationResolutionResolution1344_768  ImageGenerationResolution = "RESOLUTION_1344_768"
	ImageGenerationResolutionResolution1408_576  ImageGenerationResolution = "RESOLUTION_1408_576"
	ImageGenerationResolutionResolution1408_640  ImageGenerationResolution = "RESOLUTION_1408_640"
	ImageGenerationResolutionResolution1408_704  ImageGenerationResolution = "RESOLUTION_1408_704"
	ImageGenerationResolutionResolution1472_576  ImageGenerationResolution = "RESOLUTION_1472_576"
	ImageGenerationResolutionResolution1472_640  ImageGenerationResolution = "RESOLUTION_1472_640"
	ImageGenerationResolutionResolution1472_704  ImageGenerationResolution = "RESOLUTION_1472_704"
	ImageGenerationResolutionResolution1536_512  ImageGenerationResolution = "RESOLUTION_1536_512"
	ImageGenerationResolutionResolution1536_576  ImageGenerationResolution = "RESOLUTION_1536_576"
	ImageGenerationResolutionResolution1536_640  ImageGenerationResolution = "RESOLUTION_1536_640"
)

type ReframeNewParams struct {
	// The image being reframed (max size 10MB); only JPEG, WebP and PNG formats are
	// supported at this time.
	ImageFile io.Reader `json:"image_file,required" format:"binary"`
	// The model used to generate an image or edit one. /generate and /remix supports
	// all model types, however, /edit is only supported for V_2 and V_2_TURBO.
	//
	// Any of "V_1", "V_1_TURBO", "V_2", "V_2_TURBO", "V_2A", "V_2A_TURBO", "V_3".
	Model ModelEnum `json:"model,omitzero,required"`
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
	Resolution ImageGenerationResolution `json:"resolution,omitzero,required"`
	// The number of images to generate.
	NumImages param.Opt[int64] `json:"num_images,omitzero"`
	// Random seed. Set for reproducible generation.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// The style type to generate with; this is only applicable for models V_2 and
	// above and should not be specified for model versions V_1.
	//
	// Any of "AUTO", "GENERAL", "REALISTIC", "DESIGN", "RENDER_3D", "ANIME".
	StyleType StyleTypeV2AndAbove `json:"style_type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ReframeNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ReframeNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
