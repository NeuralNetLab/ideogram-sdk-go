// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ideogramsdk

import (
	"github.com/stainless-sdks/ideogram-sdk-go/option"
)

// ManageService contains methods and other services that help with interacting
// with the ideogram-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewManageService] method instead.
type ManageService struct {
	Options []option.RequestOption
	API     ManageAPIService
}

// NewManageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewManageService(opts ...option.RequestOption) (r ManageService) {
	r = ManageService{}
	r.Options = opts
	r.API = NewManageAPIService(opts...)
	return
}
