// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DangerousArtifact the dangerous artifact information
//
// swagger:model DangerousArtifact
type DangerousArtifact struct {

	// the count of critical vulnerabilities
	CriticalCnt int64 `json:"critical_cnt" js:"criticalCnt"`

	// the digest of the artifact
	Digest string `json:"digest,omitempty" js:"digest"`

	// the count of high vulnerabilities
	HighCnt int64 `json:"high_cnt" js:"highCnt"`

	// the count of medium vulnerabilities
	MediumCnt int64 `json:"medium_cnt" js:"mediumCnt"`

	// the project id of the artifact
	ProjectID int64 `json:"project_id,omitempty" js:"projectID"`

	// the repository name of the artifact
	RepositoryName string `json:"repository_name,omitempty" js:"repositoryName"`
}

// Validate validates this dangerous artifact
func (m *DangerousArtifact) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dangerous artifact based on context it is used
func (m *DangerousArtifact) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DangerousArtifact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DangerousArtifact) UnmarshalBinary(b []byte) error {
	var res DangerousArtifact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
