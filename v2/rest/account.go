package rest

import (
	"path"

	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/common"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/summary"
)

// AccountService manages Account endpoint
type AccountService struct {
	requestFactory
	Synchronous
}

// Provides an overview of the different fee rates for the account
// See https://docs.bitfinex.com/reference/rest-auth-summary for more info
func (s *AccountService) Summary() (_s *summary.Summary, err error) {
	req, err := s.requestFactory.NewAuthenticatedRequest(common.PermissionRead, path.Join("summary"))
	if err != nil {
		return nil, err
	}
	raw, err := s.Request(req)
	if err != nil {
		return nil, err
	}

	return summary.FromRaw(raw)
}
