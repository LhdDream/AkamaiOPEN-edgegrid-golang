package appsec

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	// The MatchTarget interface supports creating, retrieving, updating and removing match targets.
	//
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#matchtarget
	MatchTarget interface {
		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getmatchtargets
		GetMatchTargets(ctx context.Context, params GetMatchTargetsRequest) (*GetMatchTargetsResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getmatchtargetid
		GetMatchTarget(ctx context.Context, params GetMatchTargetRequest) (*GetMatchTargetResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getmatchtargetid
		CreateMatchTarget(ctx context.Context, params CreateMatchTargetRequest) (*CreateMatchTargetResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#putmatchtargetid
		UpdateMatchTarget(ctx context.Context, params UpdateMatchTargetRequest) (*UpdateMatchTargetResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#deletematchtargetid
		RemoveMatchTarget(ctx context.Context, params RemoveMatchTargetRequest) (*RemoveMatchTargetResponse, error)
	}

	// GetMatchTargetsRequest is used to retrieve the match targets for a configuration.
	GetMatchTargetsRequest struct {
		ConfigID      int `json:"configId"`
		ConfigVersion int `json:"configVersion"`
		TargetID      int `json:"targetId"`
	}

	// GetMatchTargetsResponse is returned from a call to GetMatchTargets.
	GetMatchTargetsResponse struct {
		MatchTargets struct {
			APITargets []struct {
				Type string `json:"type,omitempty"`
				Apis []struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"apis"`
				Sequence      int `json:"sequence"`
				TargetID      int `json:"targetId"`
				ConfigID      int `json:"configId,omitempty"`
				ConfigVersion int `json:"configVersion,omitempty"`

				SecurityPolicy struct {
					PolicyID string `json:"policyId,omitempty"`
				} `json:"securityPolicy,omitempty"`

				BypassNetworkLists []struct {
					Name string `json:"name,omitempty"`
					ID   string `json:"id,omitempty"`
				} `json:"bypassNetworkLists,omitempty"`
			} `json:"apiTargets,omitempty"`
			WebsiteTargets []struct {
				ConfigID                     int              `json:"configId,omitempty"`
				ConfigVersion                int              `json:"configVersion,omitempty"`
				DefaultFile                  string           `json:"defaultFile,omitempty"`
				IsNegativeFileExtensionMatch bool             `json:"isNegativeFileExtensionMatch,omitempty"`
				IsNegativePathMatch          *json.RawMessage `json:"isNegativePathMatch,omitempty"`
				Sequence                     int              `json:"-"`
				TargetID                     int              `json:"targetId,omitempty"`
				Type                         string           `json:"type,omitempty"`
				FileExtensions               []string         `json:"fileExtensions,omitempty"`
				FilePaths                    []string         `json:"filePaths,omitempty"`
				Hostnames                    []string         `json:"hostnames,omitempty"`
				SecurityPolicy               struct {
					PolicyID string `json:"policyId,omitempty"`
				} `json:"securityPolicy,omitempty"`
				BypassNetworkLists []struct {
					Name string `json:"name,omitempty"`
					ID   string `json:"id,omitempty"`
				} `json:"bypassNetworkLists,omitempty"`
			} `json:"websiteTargets,omitempty"`
		} `json:"matchTargets,omitempty"`
	}

	// GetMatchTargetRequest is used to retrieve a match target.
	GetMatchTargetRequest struct {
		ConfigID      int `json:"configId"`
		ConfigVersion int `json:"configVersion"`
		TargetID      int `json:"targetId"`
	}

	// GetMatchTargetResponse is returned from a call to GetMatchTarget.
	GetMatchTargetResponse struct {
		Type string `json:"type,omitempty"`
		Apis []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"apis,omitempty"`
		DefaultFile                  string           `json:"defaultFile,omitempty"`
		Hostnames                    []string         `json:"hostnames,omitempty"`
		IsNegativeFileExtensionMatch bool             `json:"isNegativeFileExtensionMatch,omitempty"`
		IsNegativePathMatch          *json.RawMessage `json:"isNegativePathMatch,omitempty"`
		FilePaths                    []string         `json:"filePaths,omitempty"`
		FileExtensions               []string         `json:"fileExtensions,omitempty"`
		SecurityPolicy               struct {
			PolicyID string `json:"policyId,omitempty"`
		} `json:"securityPolicy,omitempty"`
		Sequence           int `json:"-"`
		TargetID           int `json:"targetId"`
		BypassNetworkLists []struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"bypassNetworkLists,omitempty"`
	}

	// CreateMatchTargetRequest is used to create a match target.
	CreateMatchTargetRequest struct {
		Type           string          `json:"type"`
		ConfigID       int             `json:"configId"`
		ConfigVersion  int             `json:"configVersion"`
		JsonPayloadRaw json.RawMessage `json:"-"`
	}

	// CreateMatchTargetResponse is returned from a call to CreateMatchTarget.
	CreateMatchTargetResponse struct {
		MType string `json:"type"`
		Apis  []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"apis,omitempty"`
		DefaultFile                  string           `json:"defaultFile"`
		Hostnames                    []string         `json:"hostnames"`
		IsNegativeFileExtensionMatch bool             `json:"isNegativeFileExtensionMatch"`
		IsNegativePathMatch          *json.RawMessage `json:"isNegativePathMatch,omitempty"`
		FilePaths                    []string         `json:"filePaths"`
		FileExtensions               []string         `json:"fileExtensions"`
		SecurityPolicy               struct {
			PolicyID string `json:"policyId"`
		} `json:"securityPolicy"`
		Sequence           int `json:"-"`
		TargetID           int `json:"targetId"`
		BypassNetworkLists []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"bypassNetworkLists"`
	}

	// UpdateMatchTargetRequest is used to modify an existing match target.
	UpdateMatchTargetRequest struct {
		ConfigID       int             `json:"configId"`
		ConfigVersion  int             `json:"configVersion"`
		JsonPayloadRaw json.RawMessage `json:"-"`
		TargetID       int             `json:"targetId"`
	}

	// UpdateMatchTargetResponse is returned from a call to UpdateMatchTarget.
	UpdateMatchTargetResponse struct {
		Type                         string           `json:"type"`
		ConfigID                     int              `json:"configId"`
		ConfigVersion                int              `json:"configVersion"`
		DefaultFile                  string           `json:"defaultFile"`
		Hostnames                    []string         `json:"hostnames"`
		IsNegativeFileExtensionMatch bool             `json:"isNegativeFileExtensionMatch"`
		IsNegativePathMatch          *json.RawMessage `json:"isNegativePathMatch,omitempty"`
		FilePaths                    []string         `json:"filePaths"`
		FileExtensions               []string         `json:"fileExtensions"`
		SecurityPolicy               struct {
			PolicyID string `json:"policyId"`
		} `json:"securityPolicy"`
		Sequence           int `json:"-"`
		TargetID           int `json:"targetId"`
		BypassNetworkLists []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"bypassNetworkLists"`
	}

	// RemoveMatchTargetRequest is used to remove a match target.
	RemoveMatchTargetRequest struct {
		ConfigID      int `json:"configId"`
		ConfigVersion int `json:"configVersion"`
		TargetID      int `json:"targetId"`
	}

	// RemoveMatchTargetResponse is returned from a call to RemoveMatchTarget.
	RemoveMatchTargetResponse struct {
		Type                         string   `json:"type"`
		ConfigID                     int      `json:"configId"`
		ConfigVersion                int      `json:"configVersion"`
		DefaultFile                  string   `json:"defaultFile"`
		Hostnames                    []string `json:"hostnames"`
		IsNegativeFileExtensionMatch bool     `json:"isNegativeFileExtensionMatch"`
		IsNegativePathMatch          bool     `json:"isNegativePathMatch"`
		FilePaths                    []string `json:"filePaths"`
		FileExtensions               []string `json:"fileExtensions"`
		SecurityPolicy               struct {
			PolicyID string `json:"policyId"`
		} `json:"securityPolicy"`
		Sequence           int `json:"sequence"`
		TargetID           int `json:"targetId"`
		BypassNetworkLists []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"bypassNetworkLists"`
	}

	// BypassNetworkList describes a network list used in the bypass network lists for the specified configuration.
	BypassNetworkList struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	}

	// Hostnames contains one or more hostnames.
	Hostnames struct {
		Hostnames string `json:"hostnames"`
	}

	// AutoGenerated is currently unused.
	AutoGenerated struct {
		Type string `json:"type"`
		Apis []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"apis"`
		BypassNetworkLists []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"bypassNetworkLists"`
		ConfigID       int `json:"configId"`
		ConfigVersion  int `json:"configVersion"`
		SecurityPolicy struct {
			PolicyID string `json:"policyId"`
		} `json:"securityPolicy"`
		Sequence int `json:"-"`
		TargetID int `json:"targetId"`
	}
)

// Validate validates a GetMatchTargetRequest.
func (v GetMatchTargetRequest) Validate() error {
	return validation.Errors{
		"ConfigID":      validation.Validate(v.ConfigID, validation.Required),
		"ConfigVersion": validation.Validate(v.ConfigVersion, validation.Required),
		"TargetID":      validation.Validate(v.TargetID, validation.Required),
	}.Filter()
}

// Validate validates a GetMatchTargetsRequest.
func (v GetMatchTargetsRequest) Validate() error {
	return validation.Errors{
		"ConfigID":      validation.Validate(v.ConfigID, validation.Required),
		"ConfigVersion": validation.Validate(v.ConfigVersion, validation.Required),
	}.Filter()
}

// Validate validates a CreateMatchTargetRequest.
func (v CreateMatchTargetRequest) Validate() error {
	return validation.Errors{
		"ConfigID":      validation.Validate(v.ConfigID, validation.Required),
		"ConfigVersion": validation.Validate(v.ConfigVersion, validation.Required),
	}.Filter()
}

// Validate validates an UpdateMatchTargetRequest.
func (v UpdateMatchTargetRequest) Validate() error {
	return validation.Errors{
		"ConfigID":      validation.Validate(v.ConfigID, validation.Required),
		"ConfigVersion": validation.Validate(v.ConfigVersion, validation.Required),
		"TargetID":      validation.Validate(v.TargetID, validation.Required),
	}.Filter()
}

// Validate validates a RemoveMatchTargetRequest.
func (v RemoveMatchTargetRequest) Validate() error {
	return validation.Errors{
		"ConfigID":      validation.Validate(v.ConfigID, validation.Required),
		"ConfigVersion": validation.Validate(v.ConfigVersion, validation.Required),
		"TargetID":      validation.Validate(v.TargetID, validation.Required),
	}.Filter()
}

func (p *appsec) GetMatchTarget(ctx context.Context, params GetMatchTargetRequest) (*GetMatchTargetResponse, error) {
	logger := p.Log(ctx)
	logger.Debug("GetMatchTarget")

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	var result GetMatchTargetResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/match-targets/%d?includeChildObjectName=true",
		params.ConfigID,
		params.ConfigVersion,
		params.TargetID,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GetMatchTarget request: %w", err)
	}

	resp, err := p.Exec(req, &result)
	if err != nil {
		return nil, fmt.Errorf("GetMatchTarget request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &result, nil

}

func (p *appsec) GetMatchTargets(ctx context.Context, params GetMatchTargetsRequest) (*GetMatchTargetsResponse, error) {
	logger := p.Log(ctx)
	logger.Debug("GetMatchTargets")

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	var result GetMatchTargetsResponse
	var filteredResult GetMatchTargetsResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/match-targets",
		params.ConfigID,
		params.ConfigVersion,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GetMatchTargets request: %w", err)
	}

	resp, err := p.Exec(req, &result)
	if err != nil {
		return nil, fmt.Errorf("GetMatchTargets request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	if params.TargetID != 0 {
		for _, val := range result.MatchTargets.WebsiteTargets {
			if val.TargetID == params.TargetID {
				filteredResult.MatchTargets.WebsiteTargets = append(filteredResult.MatchTargets.WebsiteTargets, val)
			}
		}
		for _, val := range result.MatchTargets.APITargets {
			if val.TargetID == params.TargetID {
				filteredResult.MatchTargets.APITargets = append(filteredResult.MatchTargets.APITargets, val)
			}
		}
	} else {
		filteredResult = result
	}
	return &filteredResult, nil

}

func (p *appsec) UpdateMatchTarget(ctx context.Context, params UpdateMatchTargetRequest) (*UpdateMatchTargetResponse, error) {
	logger := p.Log(ctx)
	logger.Debug("UpdateMatchTarget")

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/match-targets/%d",
		params.ConfigID,
		params.ConfigVersion,
		params.TargetID,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create UpdateMatchTarget request: %w", err)
	}

	var result UpdateMatchTargetResponse
	req.Header.Set("Content-Type", "application/json")
	resp, err := p.Exec(req, &result, params.JsonPayloadRaw)
	if err != nil {
		return nil, fmt.Errorf("UpdateMatchTarget request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, p.Error(resp)
	}

	return &result, nil
}

func (p *appsec) CreateMatchTarget(ctx context.Context, params CreateMatchTargetRequest) (*CreateMatchTargetResponse, error) {
	logger := p.Log(ctx)
	logger.Debug("CreateMatchTarget")

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/match-targets",
		params.ConfigID,
		params.ConfigVersion,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create CreateMatchTarget request: %w", err)
	}

	var result CreateMatchTargetResponse
	req.Header.Set("Content-Type", "application/json")
	resp, err := p.Exec(req, &result, params.JsonPayloadRaw)
	if err != nil {
		return nil, fmt.Errorf("CreateMatchTarget request failed: %w", err)
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &result, nil

}

func (p *appsec) RemoveMatchTarget(ctx context.Context, params RemoveMatchTargetRequest) (*RemoveMatchTargetResponse, error) {
	logger := p.Log(ctx)
	logger.Debug("RemoveMatchTarget")

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	var result RemoveMatchTargetResponse

	uri := fmt.Sprintf("/appsec/v1/configs/%d/versions/%d/match-targets/%d", params.ConfigID, params.ConfigVersion, params.TargetID)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create RemoveMatchTarget request: %w", err)
	}

	resp, errd := p.Exec(req, nil)
	if errd != nil {
		return nil, fmt.Errorf("RemoveMatchTarget request failed: %w", err)
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &result, nil
}
