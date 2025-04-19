// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/ideogram-sdk-go"
	"github.com/stainless-sdks/ideogram-sdk-go/internal/testutil"
	"github.com/stainless-sdks/ideogram-sdk-go/option"
)

func TestInternalTestingNewWithOptionalParams(t *testing.T) {
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
	_, err := client.InternalTesting.New(context.TODO(), ideogramsdk.InternalTestingNewParams{
		RequiredDateTypeField: time.Now(),
		AnotherImageFile:      io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		DateTimeTypeField:     ideogramsdk.Time(time.Now()),
		DateTypeField:         ideogramsdk.Time(time.Now()),
		EnumTypeField:         ideogramsdk.InternalTestingNewParamsEnumTypeFieldEin,
		ImageFile:             io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		NestedObject: ideogramsdk.NestedObjectParam{
			PropOne: ideogramsdk.String("prop_one"),
			PropTwo: ideogramsdk.String("prop_two"),
		},
		NestedObjectRequiredFields: ideogramsdk.InternalTestingNewParamsNestedObjectRequiredFields{
			PropOne: "prop_one",
			PropTwo: "prop_two",
		},
		RepeatedComplexField: []ideogramsdk.NestedObjectParam{{
			PropOne: ideogramsdk.String("prop_one"),
			PropTwo: ideogramsdk.String("prop_two"),
		}},
		RepeatedPrimitiveField: []string{"string"},
		SomeText:               ideogramsdk.String("some_text"),
		XPosition:              ideogramsdk.Int(0),
		XTestHeader:            ideogramsdk.String("X-Test-Header"),
		XTestHeader2:           ideogramsdk.String("X-Test-Header-2"),
	})
	if err != nil {
		var apierr *ideogramsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
