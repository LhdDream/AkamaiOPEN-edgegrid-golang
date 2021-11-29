package appsec

import (
	"context"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	// The ApiHostnameCoverageOverlapping interface supports listing the configuration versions that
	// contain a hostname also included in the given configuration version.
	//
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#hostnameoverlap
	ApiHostnameCoverageOverlapping interface {
		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#gethostnamecoverageoverlapping
		GetApiHostnameCoverageOverlapping(ctx context.Context, params GetApiHostnameCoverageOverlappingRequest) (*GetApiHostnameCoverageOverlappingResponse, error)
	}

	// GetApiHostnameCoverageOverlappingRequest is used to retrieve the configuration versions that contain a hostname included in the current configuration version.
	GetApiHostnameCoverageOverlappingRequest struct {
		ConfigID int    `json:"-"`
		Version  int    `json:"-"`
		Hostname string `json:"-"`
	}

	// GetApiHostnameCoverageOverlappingResponse is returned from a call to GetApiHostnameCoverageOverlapping.
	GetApiHostnameCoverageOverlappingResponse struct {
		OverLappingList []struct {
			ConfigID      int      `json:"configId"`
			ConfigName    string   `json:"configName"`
			ConfigVersion int      `json:"configVersion"`
			ContractID    string   `json:"contractId"`
			ContractName  string   `json:"contractName"`
			VersionTags   []string `json:"versionTags"`
		} `json:"overLappingList"`
	}
)

// Validate validates a GetApiHostnameCoverageOverlappingRequest.
func (v GetApiHostnameCoverageOverlappingRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
	}.Filter()
}

func (p *appsec) GetApiHostnameCoverageOverlapping(ctx context.Context, params GetApiHostnameCoverageOverlappingRequest) (*GetApiHostnameCoverageOverlappingResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("GetApiHostnameCoverageOverlapping")

	var rval GetApiHostnameCoverageOverlappingResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/hostname-coverage/overlapping?hostname=%s",
		params.ConfigID,
		params.Version,
		params.Hostname,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GetApiHostnameCoverageOverlapping request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("GetApiHostnameCoverageOverlapping request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil

}