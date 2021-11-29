package appsec

import (
	"context"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	// The VersionNotes interface supports retrieving and modifying the version notes for a configuration and version.
	//
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#versionnotesgroup
	VersionNotes interface {
		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getversionnotes
		GetVersionNotes(ctx context.Context, params GetVersionNotesRequest) (*GetVersionNotesResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#putversionnotes
		UpdateVersionNotes(ctx context.Context, params UpdateVersionNotesRequest) (*UpdateVersionNotesResponse, error)
	}

	// GetVersionNotesRequest is used to retrieve the version notes for a configuration version.
	GetVersionNotesRequest struct {
		ConfigID int `json:"-"`
		Version  int `json:"-"`
	}

	// GetVersionNotesResponse is returned from a call to GetVersionNotes.
	GetVersionNotesResponse struct {
		Notes string `json:"notes"`
	}

	// UpdateVersionNotesRequest is used to modify the version notes for a configuration version.
	UpdateVersionNotesRequest struct {
		ConfigID int `json:"-"`
		Version  int `json:"-"`

		Notes string `json:"notes"`
	}

	// UpdateVersionNotesResponse is returned from a call to UpdateVersionNotes.
	UpdateVersionNotesResponse struct {
		Notes string `json:"notes"`
	}
)

// Validate validates a GetVersionNotesRequest.
func (v GetVersionNotesRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
	}.Filter()
}

// Validate validates an UpdateVersionNotesRequest.
func (v UpdateVersionNotesRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
	}.Filter()
}

func (p *appsec) GetVersionNotes(ctx context.Context, params GetVersionNotesRequest) (*GetVersionNotesResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("GetVersionNotes")

	var rval GetVersionNotesResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/version-notes",
		params.ConfigID,
		params.Version,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GetVersionNotes request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("GetVersionNotes request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil

}

func (p *appsec) UpdateVersionNotes(ctx context.Context, params UpdateVersionNotesRequest) (*UpdateVersionNotesResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("UpdateVersionNotes")

	putURL := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/version-notes",
		params.ConfigID,
		params.Version,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, putURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create UpdateVersionNotes request: %w", err)
	}

	var rval UpdateVersionNotesResponse
	resp, err := p.Exec(req, &rval, params)
	if err != nil {
		return nil, fmt.Errorf("UpdateVersionNotes request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, p.Error(resp)
	}

	return &rval, nil
}