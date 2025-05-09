// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/NeuralNetLab/ideogram-sdk-go"
	"github.com/NeuralNetLab/ideogram-sdk-go/internal/testutil"
	"github.com/NeuralNetLab/ideogram-sdk-go/option"
)

func TestManageAPISubscriptionGet(t *testing.T) {
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
	_, err := client.Manage.API.Subscription.Get(context.TODO())
	if err != nil {
		var apierr *ideogramsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestManageAPISubscriptionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Manage.API.Subscription.Update(context.TODO(), ideogramsdk.ManageAPISubscriptionUpdateParams{
		RechargeSettings: ideogramsdk.RechargeSettingsSubscriptionParam{
			MinimumBalanceThreshold: ideogramsdk.PriceParam{
				Amount:       1050,
				CurrencyCode: "USD",
			},
			TopUpBalance: ideogramsdk.PriceParam{
				Amount:       1050,
				CurrencyCode: "USD",
			},
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
