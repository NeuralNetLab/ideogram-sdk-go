// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk_test

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/NeuralNetLab/ideogram-sdk-go"
	"github.com/NeuralNetLab/ideogram-sdk-go/internal/testutil"
	"github.com/NeuralNetLab/ideogram-sdk-go/option"
)

func TestUsage(t *testing.T) {
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
	describe, err := client.Describe.New(context.TODO(), ideogramsdk.DescribeNewParams{
		ImageFile: io.Reader(bytes.NewBuffer([]byte("some file contents"))),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", describe.Descriptions)
}
