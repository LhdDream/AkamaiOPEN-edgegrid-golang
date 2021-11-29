package appsec

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	// The AttackGroup interface supports retrieving and updating attack groups along with their
	// associated actions, conditions, and exceptions.
	//
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#attackgroup
	AttackGroup interface {
		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getattackgroups
		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getattackgroupconditionexception
		GetAttackGroups(ctx context.Context, params GetAttackGroupsRequest) (*GetAttackGroupsResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getattackgroup
		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getattackgroupconditionexception
		GetAttackGroup(ctx context.Context, params GetAttackGroupRequest) (*GetAttackGroupResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#putattackgroup
		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#putattackgroupconditionexception
		UpdateAttackGroup(ctx context.Context, params UpdateAttackGroupRequest) (*UpdateAttackGroupResponse, error)
	}

	// AttackGroupConditionException describes an attack group's condition and exception information.
	AttackGroupConditionException struct {
		AdvancedExceptionsList *AttackGroupAdvancedExceptions `json:"advancedExceptions,omitempty"`
		Exception              *AttackGroupException          `json:"exception,omitempty"`
	}

	// AttackGroupAdvancedExceptions describes an attack group's advanced exception information.
	AttackGroupAdvancedExceptions struct {
		ConditionOperator                       string                                                      `json:"conditionOperator,omitempty"`
		Conditions                              *AttackGroupConditions                                      `json:"conditions,omitempty"`
		HeaderCookieOrParamValues               *AttackGroupHeaderCookieOrParamValuesAdvanced               `json:"headerCookieOrParamValues,omitempty"`
		SpecificHeaderCookieOrParamNameValue    *AttackGroupSpecificHeaderCookieOrParamNameValAdvanced      `json:"specificHeaderCookieOrParamNameValue,omitempty"`
		SpecificHeaderCookieParamXMLOrJSONNames *AttackGroupSpecificHeaderCookieParamXMLOrJSONNamesAdvanced `json:"specificHeaderCookieParamXmlOrJsonNames,omitempty"`
	}

	// AttackGroupConditions describes an attack group's condition information.
	AttackGroupConditions []struct {
		Type          string   `json:"type,omitempty"`
		Extensions    []string `json:"extensions,omitempty"`
		Filenames     []string `json:"filenames,omitempty"`
		Hosts         []string `json:"hosts,omitempty"`
		Ips           []string `json:"ips,omitempty"`
		Methods       []string `json:"methods,omitempty"`
		Paths         []string `json:"paths,omitempty"`
		Header        string   `json:"header,omitempty"`
		CaseSensitive bool     `json:"caseSensitive,omitempty"`
		Name          string   `json:"name,omitempty"`
		NameCase      bool     `json:"nameCase,omitempty"`
		PositiveMatch bool     `json:"positiveMatch"`
		Value         string   `json:"value,omitempty"`
		Wildcard      bool     `json:"wildcard,omitempty"`
		ValueCase     bool     `json:"valueCase,omitempty"`
		ValueWildcard bool     `json:"valueWildcard,omitempty"`
		UseHeaders    bool     `json:"useHeaders,omitempty"`
	}

	// AttackGroupAdvancedCriteria describes the hostname and path criteria used to limit the scope of an exception.
	AttackGroupAdvancedCriteria []struct {
		Hostnames []string `json:"hostnames,omitempty"`
		Names     []string `json:"names,omitempty"`
		Paths     []string `json:"paths,omitempty"`
		Values    []string `json:"values,omitempty"`
	}

	// AttackGroupSpecificHeaderCookieOrParamNameValAdvanced describes the excepted name-value pairs in a request.
	AttackGroupSpecificHeaderCookieOrParamNameValAdvanced []struct {
		Criteria    *AttackGroupAdvancedCriteria `json:"criteria,omitempty"`
		NamesValues []struct {
			Names  []string `json:"names"`
			Values []string `json:"values"`
		} `json:"namesValues"`
		Selector      string `json:"selector"`
		ValueWildcard bool   `json:"valueWildcard"`
		Wildcard      bool   `json:"wildcard"`
	}

	// AttackGroupSpecificHeaderCookieParamXMLOrJSONNamesAdvanced describes the advanced exception members that allow you to conditionally exclude requests from inspection.
	AttackGroupSpecificHeaderCookieParamXMLOrJSONNamesAdvanced []struct {
		Criteria *AttackGroupAdvancedCriteria `json:"criteria,omitempty"`
		Names    []string                     `json:"names,omitempty"`
		Selector string                       `json:"selector,omitempty"`
		Wildcard bool                         `json:"wildcard,omitempty"`
	}

	// AttackGroupHeaderCookieOrParamValuesAdvanced describes the list of excepted values in headers, cookies, or query parameters.
	AttackGroupHeaderCookieOrParamValuesAdvanced []struct {
		Criteria      *AttackGroupAdvancedCriteria `json:"criteria,omitempty"`
		ValueWildcard bool                         `json:"valueWildcard"`
		Values        []string                     `json:"values,omitempty"`
	}

	// AttackGroupException is used to describe an exception that can be used to conditionally exclude requests from inspection.
	AttackGroupException struct {
		SpecificHeaderCookieParamXMLOrJSONNames *AttackGroupSpecificHeaderCookieParamXMLOrJSONNames `json:"specificHeaderCookieParamXmlOrJsonNames,omitempty"`
	}

	// AttackGroupSpecificHeaderCookieParamXMLOrJSONNames describes the advanced exception members that can be used to conditionally exclude requests from inspection.
	AttackGroupSpecificHeaderCookieParamXMLOrJSONNames []struct {
		Names    []string `json:"names,omitempty"`
		Selector string   `json:"selector,omitempty"`
		Wildcard bool     `json:"wildcard,omitempty"`
	}

	// GetAttackGroupsRequest is used to retrieve a list of attack groups with their associated actions.
	GetAttackGroupsRequest struct {
		ConfigID int    `json:"-"`
		Version  int    `json:"-"`
		PolicyID string `json:"-"`
		Group    string `json:"group,omitempty"`
	}

	// GetAttackGroupsResponse is returned from a call to GetAttackGroups.
	GetAttackGroupsResponse struct {
		AttackGroups []struct {
			Group              string                         `json:"group,omitempty"`
			Action             string                         `json:"action,omitempty"`
			ConditionException *AttackGroupConditionException `json:"conditionException,omitempty"`
		} `json:"attackGroupActions,omitempty"`
	}

	// GetAttackGroupRequest is used to retrieve a list of attack groups with their associated actions.
	GetAttackGroupRequest struct {
		ConfigID int    `json:"-"`
		Version  int    `json:"-"`
		PolicyID string `json:"-"`
		Group    string `json:"group"`
	}

	// GetAttackGroupResponse is returned from a call to GetAttackGroup.
	GetAttackGroupResponse struct {
		Action             string                         `json:"action,omitempty"`
		ConditionException *AttackGroupConditionException `json:"conditionException,omitempty"`
	}

	// UpdateAttackGroupRequest is used to modify what action to take when an attack group’s rule triggers.
	UpdateAttackGroupRequest struct {
		ConfigID       int             `json:"-"`
		Version        int             `json:"-"`
		PolicyID       string          `json:"-"`
		Group          string          `json:"-"`
		Action         string          `json:"action"`
		JsonPayloadRaw json.RawMessage `json:"conditionException,omitempty"`
	}

	// UpdateAttackGroupResponse is returned from a call to UpdateAttackGroup.
	UpdateAttackGroupResponse struct {
		Action             string                         `json:"action,omitempty"`
		ConditionException *AttackGroupConditionException `json:"conditionException,omitempty"`
	}
)

// IsEmptyConditionException checks whether an attack group's ConditionException field is empty.
func (r GetAttackGroupResponse) IsEmptyConditionException() bool {
	return r.ConditionException == nil
}

// Validate validates a GetAttackGroupConditionExceptionRequest.
func (v GetAttackGroupRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
		"PolicyID": validation.Validate(v.PolicyID, validation.Required),
	}.Filter()
}

// Validate validates a GetAttackGroupConditionExceptionsRequest.
func (v GetAttackGroupsRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
		"PolicyID": validation.Validate(v.PolicyID, validation.Required),
	}.Filter()
}

// Validate validates an UpdateAttackGroupConditionExceptionRequest.
func (v UpdateAttackGroupRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
		"PolicyID": validation.Validate(v.PolicyID, validation.Required),
	}.Filter()
}

func (p *appsec) GetAttackGroup(ctx context.Context, params GetAttackGroupRequest) (*GetAttackGroupResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("GetAttackGroup")

	var rval GetAttackGroupResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/security-policies/%s/attack-groups/%s?includeConditionException=true",
		params.ConfigID,
		params.Version,
		params.PolicyID,
		params.Group)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GetAttackGroup request: %w", err)
	}
	logger.Debugf("BEFORE GetAttackGroup %v", rval)
	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("GetAttackGroup request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}
	logger.Debugf("GetAttackGroup %v", rval)
	return &rval, nil

}

func (p *appsec) GetAttackGroups(ctx context.Context, params GetAttackGroupsRequest) (*GetAttackGroupsResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("GetAttackGroupConditionExceptions")

	var rval GetAttackGroupsResponse
	var rvalfiltered GetAttackGroupsResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/security-policies/%s/attack-groups?includeConditionException=true",
		params.ConfigID,
		params.Version,
		params.PolicyID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GetAttackGroups request: %w", err)
	}
	logger.Debugf("BEFORE GetAttackGroupConditionException %v", rval)
	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("GetAttackGroups request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	if params.Group != "" {
		for k, val := range rval.AttackGroups {
			if val.Group == params.Group {
				rvalfiltered.AttackGroups = append(rvalfiltered.AttackGroups, rval.AttackGroups[k])
			}
		}
	} else {
		rvalfiltered = rval
	}

	return &rvalfiltered, nil

}

func (p *appsec) UpdateAttackGroup(ctx context.Context, params UpdateAttackGroupRequest) (*UpdateAttackGroupResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("UpdateAttackGroup")

	putURL := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/security-policies/%s/attack-groups/%s/action-condition-exception",
		params.ConfigID,
		params.Version,
		params.PolicyID,
		params.Group,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, putURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create UpdateAttackGroup request: %w", err)
	}

	var rval UpdateAttackGroupResponse
	req.Header.Set("Content-Type", "application/json")
	resp, err := p.Exec(req, &rval, params)
	if err != nil {
		return nil, fmt.Errorf("UpdateAttackGroup request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, p.Error(resp)
	}

	return &rval, nil
}