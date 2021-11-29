package appsec

import (
	"context"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	// The ReputationAnalysis interface supports retrieving and modifying the reputation analysis
	// settings for a configuration and policy.
	//
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#reputationanalysis
	ReputationAnalysis interface {
		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getreputationanalysis
		GetReputationAnalysis(ctx context.Context, params GetReputationAnalysisRequest) (*GetReputationAnalysisResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#putreputationanalysis
		UpdateReputationAnalysis(ctx context.Context, params UpdateReputationAnalysisRequest) (*UpdateReputationAnalysisResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#putreputationanalysis
		RemoveReputationAnalysis(ctx context.Context, params RemoveReputationAnalysisRequest) (*RemoveReputationAnalysisResponse, error)
	}

	// GetReputationAnalysisRequest is used to retrieve the reputation analysis settings for a security policy.
	GetReputationAnalysisRequest struct {
		ConfigID int    `json:"configId"`
		Version  int    `json:"version"`
		PolicyID string `json:"policyId"`
	}

	// GetReputationAnalysisResponse is returned from a call to GetReputationAnalysis.
	GetReputationAnalysisResponse struct {
		ConfigID                           int    `json:"-"`
		Version                            int    `json:"-"`
		PolicyID                           string `json:"-"`
		ForwardToHTTPHeader                bool   `json:"forwardToHTTPHeader"`
		ForwardSharedIPToHTTPHeaderAndSIEM bool   `json:"forwardSharedIPToHTTPHeaderAndSIEM"`
	}

	// UpdateReputationAnalysisRequest is used to modify the reputation analysis settings for a security poliyc.
	UpdateReputationAnalysisRequest struct {
		ConfigID                           int    `json:"-"`
		Version                            int    `json:"-"`
		PolicyID                           string `json:"-"`
		ForwardToHTTPHeader                bool   `json:"forwardToHTTPHeader"`
		ForwardSharedIPToHTTPHeaderAndSIEM bool   `json:"forwardSharedIPToHTTPHeaderAndSIEM"`
	}

	// UpdateReputationAnalysisResponse is returned from a call to UpdateReputationAnalysis.
	UpdateReputationAnalysisResponse struct {
		ForwardToHTTPHeader                bool `json:"forwardToHTTPHeader"`
		ForwardSharedIPToHTTPHeaderAndSIEM bool `json:"forwardSharedIPToHTTPHeaderAndSIEM"`
	}

	// RemoveReputationAnalysisRequest is used to remove the reputation analysis settings for a security policy.
	RemoveReputationAnalysisRequest struct {
		ConfigID                           int    `json:"-"`
		Version                            int    `json:"-"`
		PolicyID                           string `json:"-"`
		ForwardToHTTPHeader                bool   `json:"forwardToHTTPHeader"`
		ForwardSharedIPToHTTPHeaderAndSIEM bool   `json:"forwardSharedIPToHTTPHeaderAndSIEM"`
	}

	// RemoveReputationAnalysisResponse is returned from a call to RemoveReputationAnalysis.
	RemoveReputationAnalysisResponse struct {
		ForwardToHTTPHeader                bool `json:"forwardToHTTPHeader"`
		ForwardSharedIPToHTTPHeaderAndSIEM bool `json:"forwardSharedIPToHTTPHeaderAndSIEM"`
	}
)

// Validate validates a GetReputationAnalysisRequest.
func (v GetReputationAnalysisRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
		"PolicyID": validation.Validate(v.PolicyID, validation.Required),
	}.Filter()
}

// Validate validates an UpdateReputationAnalysisRequest.
func (v UpdateReputationAnalysisRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
		"PolicyID": validation.Validate(v.PolicyID, validation.Required),
	}.Filter()
}

// Validate validates a RemoveReputationAnalysisRequest.
func (v RemoveReputationAnalysisRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
		"PolicyID": validation.Validate(v.PolicyID, validation.Required),
	}.Filter()
}

func (p *appsec) GetReputationAnalysis(ctx context.Context, params GetReputationAnalysisRequest) (*GetReputationAnalysisResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("GetReputationAnalysis")

	var rval GetReputationAnalysisResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/security-policies/%s/reputation-analysis",
		params.ConfigID,
		params.Version,
		params.PolicyID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GetReputationAnalysis request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("GetReputationAnalysis request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil

}

func (p *appsec) UpdateReputationAnalysis(ctx context.Context, params UpdateReputationAnalysisRequest) (*UpdateReputationAnalysisResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("UpdateReputationAnalysis")

	putURL := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/security-policies/%s/reputation-analysis",
		params.ConfigID,
		params.Version,
		params.PolicyID,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, putURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create UpdateReputationAnalysis request: %w", err)
	}

	var rval UpdateReputationAnalysisResponse
	resp, err := p.Exec(req, &rval, params)
	if err != nil {
		return nil, fmt.Errorf("UpdateReputationAnalysis request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, p.Error(resp)
	}

	return &rval, nil
}

func (p *appsec) RemoveReputationAnalysis(ctx context.Context, params RemoveReputationAnalysisRequest) (*RemoveReputationAnalysisResponse, error) {

	var rval RemoveReputationAnalysisResponse

	logger := p.Log(ctx)
	logger.Debug("RemoveReputationAnalysis")

	putURL := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/security-policies/%s/reputation-analysis",
		params.ConfigID,
		params.Version,
		params.PolicyID,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, putURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create RemoveReputationAnalysis request: %w", err)
	}

	resp, err := p.Exec(req, &rval, params)
	if err != nil {
		return nil, fmt.Errorf("RemoveReputationAnalysis request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, p.Error(resp)
	}

	return &rval, nil
}