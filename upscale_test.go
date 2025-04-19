// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/ideogram-sdk-go"
	"github.com/stainless-sdks/ideogram-sdk-go/internal/testutil"
	"github.com/stainless-sdks/ideogram-sdk-go/option"
)

func TestUpscaleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Upscale.New(context.TODO(), ideogramsdk.UpscaleNewParams{
		ImageFile: io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		ImageRequest: ideogramsdk.UpscaleNewParamsImageRequest{
			Detail:             ideogramsdk.Int(50),
			MagicPromptOption:  ideogramsdk.MagicPromptOptionOn,
			MagicPromptVersion: ideogramsdk.MagicPromptVersionEnumV0,
			NumImages:          ideogramsdk.Int(1),
			Prompt:             ideogramsdk.String("A serene tropical beach scene. Dominating the foreground are tall palm trees with lush green leaves, standing tall against a backdrop of a sandy beach. The beach leads to the azure waters of the sea, which gently kisses the shoreline. In the distance, there is an island or landmass with a silhouette of what appears to be a lighthouse or tower. The sky above is painted with fluffy white clouds, some of which are tinged with hues of pink and orange, suggesting either a sunrise or sunset."),
			Resemblance:        ideogramsdk.Int(50),
			Seed:               ideogramsdk.Int(12345),
		},
	})
	if err != nil {
		var apierr *ideogramsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
