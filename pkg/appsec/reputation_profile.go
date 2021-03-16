package appsec

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ReputationProfile represents a collection of ReputationProfile
//
// See: ReputationProfile.GetReputationProfile()
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html

type (
	// ReputationProfile  contains operations available on ReputationProfile  resource
	// See: // appsec v1
	//
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getreputationprofile
	ReputationProfile interface {
		GetReputationProfiles(ctx context.Context, params GetReputationProfilesRequest) (*GetReputationProfilesResponse, error)
		GetReputationProfile(ctx context.Context, params GetReputationProfileRequest) (*GetReputationProfileResponse, error)
		CreateReputationProfile(ctx context.Context, params CreateReputationProfileRequest) (*CreateReputationProfileResponse, error)
		UpdateReputationProfile(ctx context.Context, params UpdateReputationProfileRequest) (*UpdateReputationProfileResponse, error)
		RemoveReputationProfile(ctx context.Context, params RemoveReputationProfileRequest) (*RemoveReputationProfileResponse, error)
	}

	atomicConditionsName []string

	GetReputationProfilesResponse struct {
		ReputationProfiles []struct {
			Condition        *ReputationProfileCondition `json:"condition,omitempty"`
			Context          string                      `json:"context,omitempty"`
			ContextReadable  string                      `json:"-"`
			Enabled          bool                        `json:"-"`
			ID               int                         `json:"id,omitempty"`
			Name             string                      `json:"name,omitempty"`
			SharedIPHandling string                      `json:"sharedIpHandling,omitempty"`
			Threshold        int                         `json:"threshold,omitempty"`
		} `json:"reputationProfiles,omitempty"`
	}

	ReputationProfileCondition struct {
		AtomicConditions []struct {
			CheckIps      *json.RawMessage `json:"checkIps,omitempty"`
			ClassName     string           `json:"className,omitempty"`
			Index         int              `json:"index,omitempty"`
			PositiveMatch *json.RawMessage `json:"positiveMatch,omitempty"`
			Value         []string         `json:"value,omitempty"`
			Name          *json.RawMessage `json:"name,omitempty"`
			NameCase      bool             `json:"nameCase,omitempty"`
			NameWildcard  *json.RawMessage `json:"nameWildcard,omitempty"`
			ValueCase     bool             `json:"valueCase,omitempty"`
			ValueWildcard *json.RawMessage `json:"valueWildcard,omitempty"`
			Host          []string         `json:"host,omitempty"`
		} `json:"atomicConditions,omitempty"`
		PositiveMatch *json.RawMessage `json:"positiveMatch,omitempty"`
	}

	GetReputationProfileResponse struct {
		Condition        *GetReputationProfileResponseCondition `json:"condition,omitempty"`
		Context          string                                 `json:"context,omitempty"`
		ContextReadable  string                                 `json:"-"`
		Enabled          bool                                   `json:"-"`
		ID               int                                    `json:"-"`
		Name             string                                 `json:"name,omitempty"`
		SharedIPHandling string                                 `json:"sharedIpHandling,omitempty"`
		Threshold        int                                    `json:"threshold,omitempty"`
	}

	GetReputationProfileResponseCondition struct {
		AtomicConditions []struct {
			CheckIps      *json.RawMessage `json:"checkIps,omitempty"`
			ClassName     string           `json:"className,omitempty"`
			Index         int              `json:"index,omitempty"`
			PositiveMatch json.RawMessage  `json:"positiveMatch,omitempty"`
			Value         []string         `json:"value,omitempty"`
			Name          *json.RawMessage `json:"name,omitempty"`
			NameCase      bool             `json:"nameCase,omitempty"`
			NameWildcard  *json.RawMessage `json:"nameWildcard,omitempty"`
			ValueCase     bool             `json:"valueCase,omitempty"`
			ValueWildcard *json.RawMessage `json:"valueWildcard,omitempty"`
			Host          []string         `json:"host,omitempty"`
		} `json:"atomicConditions,omitempty"`
		PositiveMatch *json.RawMessage `json:"positiveMatch,omitempty"`
	}

	CreateReputationProfileResponse struct {
		ID               int    `json:"id"`
		Name             string `json:"name"`
		Context          string `json:"context"`
		Description      string `json:"description"`
		Threshold        int    `json:"threshold"`
		SharedIPHandling string `json:"sharedIpHandling"`
		Condition        struct {
			AtomicConditions []struct {
				CheckIps      string               `json:"checkIps,omitempty"`
				ClassName     string               `json:"className"`
				Index         int                  `json:"index"`
				PositiveMatch bool                 `json:"positiveMatch"`
				Value         []string             `json:"value,omitempty"`
				Name          atomicConditionsName `json:"name,omitempty"`
				NameCase      bool                 `json:"nameCase,omitempty"`
				NameWildcard  bool                 `json:"nameWildcard,omitempty"`
				ValueCase     bool                 `json:"valueCase,omitempty"`
				ValueWildcard bool                 `json:"valueWildcard,omitempty"`
				Host          []string             `json:"host,omitempty"`
			} `json:"atomicConditions"`
			PositiveMatch *json.RawMessage `json:"positiveMatch,omitempty"`
		} `json:"condition"`
		Enabled bool `json:"enabled"`
	}

	UpdateReputationProfileResponse struct {
		ID                    int    `json:"id"`
		PolicyID              int    `json:"policyId"`
		ConfigID              int    `json:"configId"`
		ConfigVersion         int    `json:"configVersion"`
		MatchType             string `json:"matchType"`
		Type                  string `json:"type"`
		Name                  string `json:"name"`
		Description           string `json:"description"`
		AverageThreshold      int    `json:"averageThreshold"`
		BurstThreshold        int    `json:"burstThreshold"`
		ClientIdentifier      string `json:"clientIdentifier"`
		UseXForwardForHeaders bool   `json:"useXForwardForHeaders"`
		RequestType           string `json:"requestType"`
		SameActionOnIpv6      bool   `json:"sameActionOnIpv6"`
		Path                  struct {
			PositiveMatch bool     `json:"positiveMatch"`
			Values        []string `json:"values"`
		} `json:"path"`
		PathMatchType        string `json:"pathMatchType"`
		PathURIPositiveMatch bool   `json:"pathUriPositiveMatch"`
		FileExtensions       struct {
			PositiveMatch bool     `json:"positiveMatch"`
			Values        []string `json:"values"`
		} `json:"fileExtensions"`
		Hostnames              []string `json:"hostNames"`
		AdditionalMatchOptions []struct {
			PositiveMatch bool     `json:"positiveMatch"`
			Type          string   `json:"type"`
			Values        []string `json:"values"`
		} `json:"additionalMatchOptions"`
		QueryParameters []struct {
			Name          string   `json:"name"`
			Values        []string `json:"values"`
			PositiveMatch bool     `json:"positiveMatch"`
			ValueInRange  bool     `json:"valueInRange"`
		} `json:"queryParameters"`
		CreateDate string `json:"createDate"`
		UpdateDate string `json:"updateDate"`
		Used       bool   `json:"used"`
	}

	RemoveReputationProfileResponse struct {
		ID                    int    `json:"id"`
		PolicyID              int    `json:"policyId"`
		ConfigID              int    `json:"configId"`
		ConfigVersion         int    `json:"configVersion"`
		MatchType             string `json:"matchType"`
		Type                  string `json:"type"`
		Name                  string `json:"name"`
		Description           string `json:"description"`
		AverageThreshold      int    `json:"averageThreshold"`
		BurstThreshold        int    `json:"burstThreshold"`
		ClientIdentifier      string `json:"clientIdentifier"`
		UseXForwardForHeaders bool   `json:"useXForwardForHeaders"`
		RequestType           string `json:"requestType"`
		SameActionOnIpv6      bool   `json:"sameActionOnIpv6"`
		Path                  struct {
			PositiveMatch bool     `json:"positiveMatch"`
			Values        []string `json:"values"`
		} `json:"path"`
		PathMatchType        string `json:"pathMatchType"`
		PathURIPositiveMatch bool   `json:"pathUriPositiveMatch"`
		FileExtensions       struct {
			PositiveMatch bool     `json:"positiveMatch"`
			Values        []string `json:"values"`
		} `json:"fileExtensions"`
		Hostnames              []string `json:"hostNames"`
		AdditionalMatchOptions []struct {
			PositiveMatch bool     `json:"positiveMatch"`
			Type          string   `json:"type"`
			Values        []string `json:"values"`
		} `json:"additionalMatchOptions"`
		QueryParameters []struct {
			Name          string   `json:"name"`
			Values        []string `json:"values"`
			PositiveMatch bool     `json:"positiveMatch"`
			ValueInRange  bool     `json:"valueInRange"`
		} `json:"queryParameters"`
		CreateDate string `json:"createDate"`
		UpdateDate string `json:"updateDate"`
		Used       bool   `json:"used"`
	}

	GetReputationProfilesRequest struct {
		ConfigID            int `json:"configId"`
		ConfigVersion       int `json:"configVersion"`
		ReputationProfileId int `json:"-"`
	}

	GetReputationProfileRequest struct {
		ConfigID            int `json:"configId"`
		ConfigVersion       int `json:"configVersion"`
		ReputationProfileId int `json:"-"`
	}

	CreateReputationProfileRequest struct {
		ConfigID       int             `json:"-"`
		ConfigVersion  int             `json:"-"`
		JsonPayloadRaw json.RawMessage `json:"-"`
	}

	UpdateReputationProfileRequest struct {
		ConfigID            int             `json:"-"`
		ConfigVersion       int             `json:"-"`
		ReputationProfileId int             `json:"-"`
		JsonPayloadRaw      json.RawMessage `json:"-"`
	}

	RemoveReputationProfileRequest struct {
		ConfigID            int `json:"configId"`
		ConfigVersion       int `json:"configVersion"`
		ReputationProfileId int `json:"-"`
	}
)

func (c *atomicConditionsName) UnmarshalJSON(data []byte) error {
	var nums interface{}
	err := json.Unmarshal(data, &nums)
	if err != nil {
		return err
	}

	items := reflect.ValueOf(nums)
	switch items.Kind() {
	case reflect.String:
		*c = append(*c, items.String())

	case reflect.Slice:
		*c = make(atomicConditionsName, 0, items.Len())
		for i := 0; i < items.Len(); i++ {
			item := items.Index(i)
			switch item.Kind() {
			case reflect.String:
				*c = append(*c, item.String())
			case reflect.Interface:
				*c = append(*c, item.Interface().(string))
			}
		}
	}
	return nil
}

// Validate validates GetReputationProfileRequest
func (v GetReputationProfileRequest) Validate() error {
	return validation.Errors{
		"ConfigID":      validation.Validate(v.ConfigID, validation.Required),
		"ConfigVersion": validation.Validate(v.ConfigVersion, validation.Required),
		"RatePolicyID":  validation.Validate(v.ReputationProfileId, validation.Required),
	}.Filter()
}

// Validate validates GetReputationProfilesRequest
func (v GetReputationProfilesRequest) Validate() error {
	return validation.Errors{
		"ConfigID":      validation.Validate(v.ConfigID, validation.Required),
		"ConfigVersion": validation.Validate(v.ConfigVersion, validation.Required),
	}.Filter()
}

// Validate validates CreateReputationProfileRequest
func (v CreateReputationProfileRequest) Validate() error {
	return validation.Errors{
		"ConfigID":      validation.Validate(v.ConfigID, validation.Required),
		"ConfigVersion": validation.Validate(v.ConfigVersion, validation.Required),
	}.Filter()
}

// Validate validates UpdateReputationProfileRequest
func (v UpdateReputationProfileRequest) Validate() error {
	return validation.Errors{
		"ConfigID":            validation.Validate(v.ConfigID, validation.Required),
		"ConfigVersion":       validation.Validate(v.ConfigVersion, validation.Required),
		"ReputationProfileId": validation.Validate(v.ReputationProfileId, validation.Required),
	}.Filter()
}

// Validate validates RemoveReputationProfileRequest
func (v RemoveReputationProfileRequest) Validate() error {
	return validation.Errors{
		"ConfigID":            validation.Validate(v.ConfigID, validation.Required),
		"ConfigVersion":       validation.Validate(v.ConfigVersion, validation.Required),
		"ReputationProfileId": validation.Validate(v.ReputationProfileId, validation.Required),
	}.Filter()
}

func (p *appsec) GetReputationProfile(ctx context.Context, params GetReputationProfileRequest) (*GetReputationProfileResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("GetReputationProfile")

	var rval GetReputationProfileResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/reputation-profiles/%d",
		params.ConfigID,
		params.ConfigVersion,
		params.ReputationProfileId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create getreputationprofile request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("getproperties request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil

}

func (p *appsec) GetReputationProfiles(ctx context.Context, params GetReputationProfilesRequest) (*GetReputationProfilesResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("GetReputationProfiles")

	var rval GetReputationProfilesResponse
	var rvalfiltered GetReputationProfilesResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/reputation-profiles",
		params.ConfigID,
		params.ConfigVersion,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create getreputationprofiles request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("getreputationprofiles request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	if params.ReputationProfileId != 0 {
		for _, val := range rval.ReputationProfiles {
			if val.ID == params.ReputationProfileId {
				rvalfiltered.ReputationProfiles = append(rvalfiltered.ReputationProfiles, val)
			}
		}

	} else {
		rvalfiltered = rval
	}

	return &rvalfiltered, nil

}

// Update will update a ReputationProfile.
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#putreputationprofile

func (p *appsec) UpdateReputationProfile(ctx context.Context, params UpdateReputationProfileRequest) (*UpdateReputationProfileResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("UpdateReputationProfile")

	putURL := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/reputation-profiles/%d",
		params.ConfigID,
		params.ConfigVersion,
		params.ReputationProfileId,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, putURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create create ReputationProfilerequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	var rval UpdateReputationProfileResponse
	resp, err := p.Exec(req, &rval, params.JsonPayloadRaw)
	if err != nil {
		return nil, fmt.Errorf("create ReputationProfile request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, p.Error(resp)
	}

	return &rval, nil
}

// Create will create a new reputationprofile.
//
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#postreputationprofile
func (p *appsec) CreateReputationProfile(ctx context.Context, params CreateReputationProfileRequest) (*CreateReputationProfileResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("CreateReputationProfile")

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/reputation-profiles",
		params.ConfigID,
		params.ConfigVersion,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create create reputationprofile request: %w", err)
	}

	var rval CreateReputationProfileResponse
	req.Header.Set("Content-Type", "application/json")
	resp, err := p.Exec(req, &rval, params.JsonPayloadRaw)
	if err != nil {
		return nil, fmt.Errorf("create reputationprofilerequest failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, p.Error(resp)
	}

	return &rval, nil

}

// Delete will delete a ReputationProfile
//
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#deletereputationprofile

func (p *appsec) RemoveReputationProfile(ctx context.Context, params RemoveReputationProfileRequest) (*RemoveReputationProfileResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	var rval RemoveReputationProfileResponse

	logger := p.Log(ctx)
	logger.Debug("RemoveReputationProfile")

	uri, err := url.Parse(fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/reputation-profiles/%d",
		params.ConfigID,
		params.ConfigVersion,
		params.ReputationProfileId),
	)
	if err != nil {
		return nil, fmt.Errorf("failed parse url: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, uri.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create delreputationprofile request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("delreputationprofile request failed: %w", err)
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil
}
