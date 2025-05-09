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

func TestReframeNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Reframe.New(context.TODO(), ideogramsdk.ReframeNewParams{
		ImageFile:  io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		Model:      ideogramsdk.ModelEnumV2Turbo,
		Resolution: ideogramsdk.ResolutionImageGenerationResolution1024_1024,
		NumImages:  ideogramsdk.Int(1),
		Seed:       ideogramsdk.Int(12345),
		StyleType:  ideogramsdk.StyleTypeV2AndAboveRealistic,
	})
	if err != nil {
		var apierr *ideogramsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
