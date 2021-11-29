package appsec

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	// The CustomDeny interface supports creating, retrievinfg, modifying and removing custom deny actions
	// for a configuration.
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#customdeny
	//
	CustomDeny interface {
		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getcustomdeny
		GetCustomDenyList(ctx context.Context, params GetCustomDenyListRequest) (*GetCustomDenyListResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getcustomdenyaction
		GetCustomDeny(ctx context.Context, params GetCustomDenyRequest) (*GetCustomDenyResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#postcustomdeny
		CreateCustomDeny(ctx context.Context, params CreateCustomDenyRequest) (*CreateCustomDenyResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#putcustomdenyaction
		UpdateCustomDeny(ctx context.Context, params UpdateCustomDenyRequest) (*UpdateCustomDenyResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#deletecustomdenyaction
		RemoveCustomDeny(ctx context.Context, params RemoveCustomDenyRequest) (*RemoveCustomDenyResponse, error)
	}

	customDenyID string

	// GetCustomDenyListRequest is used to retrieve the custom deny actions for a configuration.
	GetCustomDenyListRequest struct {
		ConfigID int    `json:"configId"`
		Version  int    `json:"version"`
		ID       string `json:"id,omitempty"`
	}

	// GetCustomDenyListResponse is returned from a call to GetCustomDenyList.
	GetCustomDenyListResponse struct {
		CustomDenyList []struct {
			Description string       `json:"description,omitempty"`
			Name        string       `json:"name"`
			ID          customDenyID `json:"id"`
			Parameters  []struct {
				DisplayName string `json:"-"`
				Name        string `json:"name"`
				Value       string `json:"value"`
			} `json:"parameters"`
		} `json:"customDenyList"`
	}

	// GetCustomDenyRequest is used to retrieve a specific custom deny action.
	GetCustomDenyRequest struct {
		ConfigID int    `json:"configId"`
		Version  int    `json:"version"`
		ID       string `json:"id,omitempty"`
	}

	// GetCustomDenyResponse is returned from a call to GetCustomDeny.
	GetCustomDenyResponse struct {
		Description string       `json:"description,omitempty"`
		Name        string       `json:"name"`
		ID          customDenyID `json:"-"`
		Parameters  []struct {
			DisplayName string `json:"-"`
			Name        string `json:"name"`
			Value       string `json:"value"`
		} `json:"parameters"`
	}

	// CreateCustomDenyRequest is used to create a new custom deny action for a specific configuration.
	CreateCustomDenyRequest struct {
		ConfigID       int             `json:"-"`
		Version        int             `json:"-"`
		JsonPayloadRaw json.RawMessage `json:"-"`
	}

	// CreateCustomDenyResponse is returned from a call to CreateCustomDeny.
	CreateCustomDenyResponse struct {
		Description string       `json:"description,omitempty"`
		Name        string       `json:"name"`
		ID          customDenyID `json:"id"`
		Parameters  []struct {
			DisplayName string `json:"-"`
			Name        string `json:"name"`
			Value       string `json:"value"`
		} `json:"parameters"`
	}

	// UpdateCustomDenyRequest is used to details for a specific custom deny action.
	UpdateCustomDenyRequest struct {
		ConfigID       int             `json:"-"`
		Version        int             `json:"-"`
		ID             string          `json:"id"`
		JsonPayloadRaw json.RawMessage `json:"-"`
	}

	// UpdateCustomDenyResponse is returned from a call to UpdateCustomDeny.
	UpdateCustomDenyResponse struct {
		Description string       `json:"description,omitempty"`
		Name        string       `json:"name"`
		ID          customDenyID `json:"-"`
		Parameters  []struct {
			DisplayName string `json:"-"`
			Name        string `json:"name"`
			Value       string `json:"value"`
		} `json:"parameters"`
	}

	// RemoveCustomDenyRequest is used to remove an existing custom deny action.
	RemoveCustomDenyRequest struct {
		ConfigID int    `json:"-"`
		Version  int    `json:"-"`
		ID       string `json:"id,omitempty"`
	}

	// RemoveCustomDenyResponse is returned from a call to RemoveCustomDeny.
	RemoveCustomDenyResponse struct {
		Empty string `json:"-"`
	}
)

// UnmarshalJSON reads a customDenyID struct from its data argument.
func (c *customDenyID) UnmarshalJSON(data []byte) error {
	var nums interface{}
	err := json.Unmarshal(data, &nums)
	if err != nil {
		return err
	}

	items := reflect.ValueOf(nums)
	switch items.Kind() {
	case reflect.String:
		*c = customDenyID(nums.(string))
	case reflect.Int:

		*c = customDenyID(strconv.Itoa(nums.(int)))

	}
	return nil
}

// Validate validates a GetCustomDenyRequest.
func (v GetCustomDenyRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
		"ID":       validation.Validate(v.ID, validation.Required),
	}.Filter()
}

// Validate validates a GetCustomDenysRequest.
func (v GetCustomDenyListRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
	}.Filter()
}

// Validate validates a CreateCustomDenyRequest.
func (v CreateCustomDenyRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
	}.Filter()
}

// Validate validates an UpdateCustomDenyRequest.
func (v UpdateCustomDenyRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
		"ID":       validation.Validate(v.ID, validation.Required),
	}.Filter()
}

// Validate validates a RemoveCustomDenyRequest.
func (v RemoveCustomDenyRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
		"ID":       validation.Validate(v.ID, validation.Required),
	}.Filter()
}

func (p *appsec) GetCustomDeny(ctx context.Context, params GetCustomDenyRequest) (*GetCustomDenyResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("GetCustomDeny")

	var rval GetCustomDenyResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/custom-deny/%s",
		params.ConfigID,
		params.Version,
		params.ID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GetCustomDeny request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("GetCustomDeny request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil

}

func (p *appsec) GetCustomDenyList(ctx context.Context, params GetCustomDenyListRequest) (*GetCustomDenyListResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("GetCustomDenyList")

	var rval GetCustomDenyListResponse
	var rvalfiltered GetCustomDenyListResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/custom-deny",
		params.ConfigID,
		params.Version,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GetlustomDenyList request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("GetCustomDenyList request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	if params.ID != "" {
		for _, val := range rval.CustomDenyList {
			if string(val.ID) == params.ID {
				rvalfiltered.CustomDenyList = append(rvalfiltered.CustomDenyList, val)
			}
		}

	} else {
		rvalfiltered = rval
	}
	return &rvalfiltered, nil

}

func (p *appsec) UpdateCustomDeny(ctx context.Context, params UpdateCustomDenyRequest) (*UpdateCustomDenyResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("UpdateCustomDeny")

	putURL := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/custom-deny/%s",
		params.ConfigID,
		params.Version,
		params.ID,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, putURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create UpdateCustomDeny request: %w", err)
	}

	var rval UpdateCustomDenyResponse
	req.Header.Set("Content-Type", "application/json")
	resp, err := p.Exec(req, &rval, params.JsonPayloadRaw)
	if err != nil {
		return nil, fmt.Errorf("UpdateCustomDeny request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, p.Error(resp)
	}

	return &rval, nil
}

func (p *appsec) CreateCustomDeny(ctx context.Context, params CreateCustomDenyRequest) (*CreateCustomDenyResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("CreateCustomDeny")

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/custom-deny",
		params.ConfigID,
		params.Version,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create CreateCustomDeny request: %w", err)
	}

	var rval CreateCustomDenyResponse
	req.Header.Set("Content-Type", "application/json")
	resp, err := p.Exec(req, &rval, params.JsonPayloadRaw)
	if err != nil {
		return nil, fmt.Errorf("CreateCustomDeny request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, p.Error(resp)
	}

	return &rval, nil

}

func (p *appsec) RemoveCustomDeny(ctx context.Context, params RemoveCustomDenyRequest) (*RemoveCustomDenyResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	var rval RemoveCustomDenyResponse

	logger := p.Log(ctx)
	logger.Debug("RemoveCustomDeny")

	uri, err := url.Parse(fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/custom-deny/%s",
		params.ConfigID,
		params.Version,
		params.ID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, uri.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create RemoveCustomDeny request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("RemoveCustomDeny request failed: %w", err)
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil
}