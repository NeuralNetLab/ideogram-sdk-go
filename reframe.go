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
type ResolutionImageGeneration string

const (
	ResolutionImageGenerationResolution512_1536  ResolutionImageGeneration = "RESOLUTION_512_1536"
	ResolutionImageGenerationResolution576_1408  ResolutionImageGeneration = "RESOLUTION_576_1408"
	ResolutionImageGenerationResolution576_1472  ResolutionImageGeneration = "RESOLUTION_576_1472"
	ResolutionImageGenerationResolution576_1536  ResolutionImageGeneration = "RESOLUTION_576_1536"
	ResolutionImageGenerationResolution640_1024  ResolutionImageGeneration = "RESOLUTION_640_1024"
	ResolutionImageGenerationResolution640_1344  ResolutionImageGeneration = "RESOLUTION_640_1344"
	ResolutionImageGenerationResolution640_1408  ResolutionImageGeneration = "RESOLUTION_640_1408"
	ResolutionImageGenerationResolution640_1472  ResolutionImageGeneration = "RESOLUTION_640_1472"
	ResolutionImageGenerationResolution640_1536  ResolutionImageGeneration = "RESOLUTION_640_1536"
	ResolutionImageGenerationResolution704_1152  ResolutionImageGeneration = "RESOLUTION_704_1152"
	ResolutionImageGenerationResolution704_1216  ResolutionImageGeneration = "RESOLUTION_704_1216"
	ResolutionImageGenerationResolution704_1280  ResolutionImageGeneration = "RESOLUTION_704_1280"
	ResolutionImageGenerationResolution704_1344  ResolutionImageGeneration = "RESOLUTION_704_1344"
	ResolutionImageGenerationResolution704_1408  ResolutionImageGeneration = "RESOLUTION_704_1408"
	ResolutionImageGenerationResolution704_1472  ResolutionImageGeneration = "RESOLUTION_704_1472"
	ResolutionImageGenerationResolution720_1280  ResolutionImageGeneration = "RESOLUTION_720_1280"
	ResolutionImageGenerationResolution736_1312  ResolutionImageGeneration = "RESOLUTION_736_1312"
	ResolutionImageGenerationResolution768_1024  ResolutionImageGeneration = "RESOLUTION_768_1024"
	ResolutionImageGenerationResolution768_1088  ResolutionImageGeneration = "RESOLUTION_768_1088"
	ResolutionImageGenerationResolution768_1152  ResolutionImageGeneration = "RESOLUTION_768_1152"
	ResolutionImageGenerationResolution768_1216  ResolutionImageGeneration = "RESOLUTION_768_1216"
	ResolutionImageGenerationResolution768_1232  ResolutionImageGeneration = "RESOLUTION_768_1232"
	ResolutionImageGenerationResolution768_1280  ResolutionImageGeneration = "RESOLUTION_768_1280"
	ResolutionImageGenerationResolution768_1344  ResolutionImageGeneration = "RESOLUTION_768_1344"
	ResolutionImageGenerationResolution832_960   ResolutionImageGeneration = "RESOLUTION_832_960"
	ResolutionImageGenerationResolution832_1024  ResolutionImageGeneration = "RESOLUTION_832_1024"
	ResolutionImageGenerationResolution832_1088  ResolutionImageGeneration = "RESOLUTION_832_1088"
	ResolutionImageGenerationResolution832_1152  ResolutionImageGeneration = "RESOLUTION_832_1152"
	ResolutionImageGenerationResolution832_1216  ResolutionImageGeneration = "RESOLUTION_832_1216"
	ResolutionImageGenerationResolution832_1248  ResolutionImageGeneration = "RESOLUTION_832_1248"
	ResolutionImageGenerationResolution864_1152  ResolutionImageGeneration = "RESOLUTION_864_1152"
	ResolutionImageGenerationResolution896_960   ResolutionImageGeneration = "RESOLUTION_896_960"
	ResolutionImageGenerationResolution896_1024  ResolutionImageGeneration = "RESOLUTION_896_1024"
	ResolutionImageGenerationResolution896_1088  ResolutionImageGeneration = "RESOLUTION_896_1088"
	ResolutionImageGenerationResolution896_1120  ResolutionImageGeneration = "RESOLUTION_896_1120"
	ResolutionImageGenerationResolution896_1152  ResolutionImageGeneration = "RESOLUTION_896_1152"
	ResolutionImageGenerationResolution960_832   ResolutionImageGeneration = "RESOLUTION_960_832"
	ResolutionImageGenerationResolution960_896   ResolutionImageGeneration = "RESOLUTION_960_896"
	ResolutionImageGenerationResolution960_1024  ResolutionImageGeneration = "RESOLUTION_960_1024"
	ResolutionImageGenerationResolution960_1088  ResolutionImageGeneration = "RESOLUTION_960_1088"
	ResolutionImageGenerationResolution1024_640  ResolutionImageGeneration = "RESOLUTION_1024_640"
	ResolutionImageGenerationResolution1024_768  ResolutionImageGeneration = "RESOLUTION_1024_768"
	ResolutionImageGenerationResolution1024_832  ResolutionImageGeneration = "RESOLUTION_1024_832"
	ResolutionImageGenerationResolution1024_896  ResolutionImageGeneration = "RESOLUTION_1024_896"
	ResolutionImageGenerationResolution1024_960  ResolutionImageGeneration = "RESOLUTION_1024_960"
	ResolutionImageGenerationResolution1024_1024 ResolutionImageGeneration = "RESOLUTION_1024_1024"
	ResolutionImageGenerationResolution1088_768  ResolutionImageGeneration = "RESOLUTION_1088_768"
	ResolutionImageGenerationResolution1088_832  ResolutionImageGeneration = "RESOLUTION_1088_832"
	ResolutionImageGenerationResolution1088_896  ResolutionImageGeneration = "RESOLUTION_1088_896"
	ResolutionImageGenerationResolution1088_960  ResolutionImageGeneration = "RESOLUTION_1088_960"
	ResolutionImageGenerationResolution1120_896  ResolutionImageGeneration = "RESOLUTION_1120_896"
	ResolutionImageGenerationResolution1152_704  ResolutionImageGeneration = "RESOLUTION_1152_704"
	ResolutionImageGenerationResolution1152_768  ResolutionImageGeneration = "RESOLUTION_1152_768"
	ResolutionImageGenerationResolution1152_832  ResolutionImageGeneration = "RESOLUTION_1152_832"
	ResolutionImageGenerationResolution1152_864  ResolutionImageGeneration = "RESOLUTION_1152_864"
	ResolutionImageGenerationResolution1152_896  ResolutionImageGeneration = "RESOLUTION_1152_896"
	ResolutionImageGenerationResolution1216_704  ResolutionImageGeneration = "RESOLUTION_1216_704"
	ResolutionImageGenerationResolution1216_768  ResolutionImageGeneration = "RESOLUTION_1216_768"
	ResolutionImageGenerationResolution1216_832  ResolutionImageGeneration = "RESOLUTION_1216_832"
	ResolutionImageGenerationResolution1232_768  ResolutionImageGeneration = "RESOLUTION_1232_768"
	ResolutionImageGenerationResolution1248_832  ResolutionImageGeneration = "RESOLUTION_1248_832"
	ResolutionImageGenerationResolution1280_704  ResolutionImageGeneration = "RESOLUTION_1280_704"
	ResolutionImageGenerationResolution1280_720  ResolutionImageGeneration = "RESOLUTION_1280_720"
	ResolutionImageGenerationResolution1280_768  ResolutionImageGeneration = "RESOLUTION_1280_768"
	ResolutionImageGenerationResolution1280_800  ResolutionImageGeneration = "RESOLUTION_1280_800"
	ResolutionImageGenerationResolution1312_736  ResolutionImageGeneration = "RESOLUTION_1312_736"
	ResolutionImageGenerationResolution1344_640  ResolutionImageGeneration = "RESOLUTION_1344_640"
	ResolutionImageGenerationResolution1344_704  ResolutionImageGeneration = "RESOLUTION_1344_704"
	ResolutionImageGenerationResolution1344_768  ResolutionImageGeneration = "RESOLUTION_1344_768"
	ResolutionImageGenerationResolution1408_576  ResolutionImageGeneration = "RESOLUTION_1408_576"
	ResolutionImageGenerationResolution1408_640  ResolutionImageGeneration = "RESOLUTION_1408_640"
	ResolutionImageGenerationResolution1408_704  ResolutionImageGeneration = "RESOLUTION_1408_704"
	ResolutionImageGenerationResolution1472_576  ResolutionImageGeneration = "RESOLUTION_1472_576"
	ResolutionImageGenerationResolution1472_640  ResolutionImageGeneration = "RESOLUTION_1472_640"
	ResolutionImageGenerationResolution1472_704  ResolutionImageGeneration = "RESOLUTION_1472_704"
	ResolutionImageGenerationResolution1536_512  ResolutionImageGeneration = "RESOLUTION_1536_512"
	ResolutionImageGenerationResolution1536_576  ResolutionImageGeneration = "RESOLUTION_1536_576"
	ResolutionImageGenerationResolution1536_640  ResolutionImageGeneration = "RESOLUTION_1536_640"
)

type ReframeNewParams struct {
	// The image being reframed (max size 10MB); only JPEG, WebP and PNG formats are
	// supported at this time.
	ImageFile io.Reader `json:"image_file,omitzero,required" format:"binary"`
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
	Resolution ResolutionImageGeneration `json:"resolution,omitzero,required"`
	// The number of images to generate.
	NumImages param.Opt[int64] `json:"num_images,omitzero"`
	// Random seed. Set for reproducible generation.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// The style type to generate with; this is only applicable for models V_2 and
	// above and should not be specified for model versions V_1.
	//
	// Any of "AUTO", "GENERAL", "REALISTIC", "DESIGN", "RENDER_3D", "ANIME", "CUSTOM".
	StyleType StyleTypeV2AndAbove `json:"style_type,omitzero"`
	paramObj
}

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
