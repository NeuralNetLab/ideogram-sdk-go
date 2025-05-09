// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/NeuralNetLab/ideogram-sdk-go"
	"github.com/NeuralNetLab/ideogram-sdk-go/internal/testutil"
	"github.com/NeuralNetLab/ideogram-sdk-go/option"
)

func TestIdeogramV3EditWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ideogramsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.IdeogramV3.Edit(context.TODO(), ideogramsdk.IdeogramV3EditParams{
		Image:  io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		Mask:   io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		Prompt: "A photo of a cat.",
		ColorPalette: ideogramsdk.ColorPaletteUnionParam{
			OfColorPaletteWithPresetName: &ideogramsdk.ColorPaletteColorPaletteWithPresetNameParam{
				Name: "PASTEL",
			},
		},
		MagicPrompt:          ideogramsdk.MagicPromptOptionOn,
		NumImages:            ideogramsdk.Int(1),
		RenderingSpeed:       ideogramsdk.RenderingSpeedTurbo,
		Seed:                 ideogramsdk.Int(12345),
		StyleCodes:           []string{"AAFF5733", "0133FF57", "DE3357FF"},
		StyleReferenceImages: []io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents")))},
	})
	if err != nil {
		var apierr *ideogramsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIdeogramV3GenerateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ideogramsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.IdeogramV3.Generate(context.TODO(), ideogramsdk.IdeogramV3GenerateParams{
		Prompt:      "A photo of a cat",
		AspectRatio: ideogramsdk.AspectRatio1x3,
		ColorPalette: ideogramsdk.ColorPaletteUnionParam{
			OfColorPaletteWithPresetName: &ideogramsdk.ColorPaletteColorPaletteWithPresetNameParam{
				Name: "PASTEL",
			},
		},
		MagicPrompt:          ideogramsdk.MagicPromptOptionOn,
		NegativePrompt:       ideogramsdk.String("brush strokes, painting"),
		NumImages:            ideogramsdk.Int(1),
		RenderingSpeed:       ideogramsdk.RenderingSpeedTurbo,
		Resolution:           ideogramsdk.ResolutionIdeogram1280x800,
		Seed:                 ideogramsdk.Int(12345),
		StyleCodes:           []string{"AAFF5733", "0133FF57", "DE3357FF"},
		StyleReferenceImages: []io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents")))},
		StyleType:            ideogramsdk.StyleTypeV3General,
	})
	if err != nil {
		var apierr *ideogramsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIdeogramV3ReframeWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ideogramsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.IdeogramV3.Reframe(context.TODO(), ideogramsdk.IdeogramV3ReframeParams{
		Image:      io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		Resolution: ideogramsdk.ResolutionIdeogram1280x800,
		ColorPalette: ideogramsdk.ColorPaletteUnionParam{
			OfColorPaletteWithPresetName: &ideogramsdk.ColorPaletteColorPaletteWithPresetNameParam{
				Name: "PASTEL",
			},
		},
		NumImages:            ideogramsdk.Int(1),
		RenderingSpeed:       ideogramsdk.RenderingSpeedTurbo,
		Seed:                 ideogramsdk.Int(12345),
		StyleCodes:           []string{"AAFF5733", "0133FF57", "DE3357FF"},
		StyleReferenceImages: []io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents")))},
	})
	if err != nil {
		var apierr *ideogramsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIdeogramV3RemixWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ideogramsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.IdeogramV3.Remix(context.TODO(), ideogramsdk.IdeogramV3RemixParams{
		Image:       io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		Prompt:      "A photo of a cat",
		AspectRatio: ideogramsdk.AspectRatio1x3,
		ColorPalette: ideogramsdk.ColorPaletteUnionParam{
			OfColorPaletteWithPresetName: &ideogramsdk.ColorPaletteColorPaletteWithPresetNameParam{
				Name: "PASTEL",
			},
		},
		ImageWeight:          ideogramsdk.Int(50),
		MagicPrompt:          ideogramsdk.MagicPromptOptionOn,
		NegativePrompt:       ideogramsdk.String("brush strokes, painting"),
		NumImages:            ideogramsdk.Int(1),
		RenderingSpeed:       ideogramsdk.RenderingSpeedTurbo,
		Resolution:           ideogramsdk.ResolutionIdeogram1280x800,
		Seed:                 ideogramsdk.Int(12345),
		StyleCodes:           []string{"AAFF5733", "0133FF57", "DE3357FF"},
		StyleReferenceImages: []io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents")))},
		StyleType:            ideogramsdk.StyleTypeV3General,
	})
	if err != nil {
		var apierr *ideogramsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIdeogramV3ReplaceBackgroundWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ideogramsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.IdeogramV3.ReplaceBackground(context.TODO(), ideogramsdk.IdeogramV3ReplaceBackgroundParams{
		Image:  io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		Prompt: "A vibrant cityscape at night.",
		ColorPalette: ideogramsdk.ColorPaletteUnionParam{
			OfColorPaletteWithPresetName: &ideogramsdk.ColorPaletteColorPaletteWithPresetNameParam{
				Name: "PASTEL",
			},
		},
		MagicPrompt:          ideogramsdk.MagicPromptOptionOn,
		NumImages:            ideogramsdk.Int(1),
		RenderingSpeed:       ideogramsdk.RenderingSpeedTurbo,
		Seed:                 ideogramsdk.Int(12345),
		StyleCodes:           []string{"AAFF5733", "0133FF57", "DE3357FF"},
		StyleReferenceImages: []io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents")))},
	})
	if err != nil {
		var apierr *ideogramsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
