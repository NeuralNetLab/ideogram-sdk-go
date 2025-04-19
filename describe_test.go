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

func TestDescribeNew(t *testing.T) {
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
	_, err := client.Describe.New(context.TODO(), ideogramsdk.DescribeNewParams{
		ImageFile: io.Reader(bytes.NewBuffer([]byte("some file contents"))),
	})
	if err != nil {
		var apierr *ideogramsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
