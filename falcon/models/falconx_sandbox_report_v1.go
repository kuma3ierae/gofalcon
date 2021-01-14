// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FalconxSandboxReportV1 falconx sandbox report v1
//
// swagger:model falconx.SandboxReportV1
type FalconxSandboxReportV1 struct {

	// architecture
	Architecture string `json:"architecture,omitempty"`

	// classification
	Classification []string `json:"classification"`

	// classification tags
	ClassificationTags []string `json:"classification_tags"`

	// contacted hosts
	ContactedHosts []*FalconxContactedHost `json:"contacted_hosts"`

	// dns requests
	DNSRequests []*FalconxDNSRequest `json:"dns_requests"`

	// environment description
	EnvironmentDescription string `json:"environment_description,omitempty"`

	// environment id
	EnvironmentID int32 `json:"environment_id,omitempty"`

	// error message
	ErrorMessage string `json:"error_message,omitempty"`

	// error origin
	ErrorOrigin string `json:"error_origin,omitempty"`

	// error type
	ErrorType string `json:"error_type,omitempty"`

	// extracted files
	ExtractedFiles []*FalconxExtractedFile `json:"extracted_files"`

	// extracted interesting strings
	ExtractedInterestingStrings []*FalconxExtractedInterestingString `json:"extracted_interesting_strings"`

	// file imports
	FileImports []*FalconxFileImport `json:"file_imports"`

	// file metadata
	FileMetadata *FalconxFileMetadata `json:"file_metadata,omitempty"`

	// file size
	FileSize int32 `json:"file_size,omitempty"`

	// file type
	FileType string `json:"file_type,omitempty"`

	// file type short
	FileTypeShort []string `json:"file_type_short"`

	// http requests
	HTTPRequests []*FalconxHTTPRequest `json:"http_requests"`

	// incidents
	Incidents []*FalconxIncident `json:"incidents"`

	// ioc report broad artifact id
	IocReportBroadArtifactID string `json:"ioc_report_broad_artifact_id,omitempty"`

	// ioc report strict artifact id
	IocReportStrictArtifactID string `json:"ioc_report_strict_artifact_id,omitempty"`

	// memory forensics
	MemoryForensics []*FalconxMemoryForensic `json:"memory_forensics"`

	// memory strings artifact id
	MemoryStringsArtifactID string `json:"memory_strings_artifact_id,omitempty"`

	// mitre attacks
	MitreAttacks []*FalconxMITREAttack `json:"mitre_attacks"`

	// packer
	Packer string `json:"packer,omitempty"`

	// pcap report artifact id
	PcapReportArtifactID string `json:"pcap_report_artifact_id,omitempty"`

	// processes
	Processes []*FalconxProcess `json:"processes"`

	// sample flags
	SampleFlags []string `json:"sample_flags"`

	// screenshots artifact ids
	ScreenshotsArtifactIds []string `json:"screenshots_artifact_ids"`

	// sha256
	Sha256 string `json:"sha256,omitempty"`

	// signatures
	Signatures []*FalconxSignature `json:"signatures"`

	// submission type
	SubmissionType string `json:"submission_type,omitempty"`

	// submit name
	SubmitName string `json:"submit_name,omitempty"`

	// submit url
	SubmitURL string `json:"submit_url,omitempty"`

	// suricata alerts
	SuricataAlerts []*FalconxSuricataAlert `json:"suricata_alerts"`

	// target url
	TargetURL string `json:"target_url,omitempty"`

	// threat score
	ThreatScore int32 `json:"threat_score,omitempty"`

	// verdict
	Verdict string `json:"verdict,omitempty"`

	// version info
	VersionInfo []*FalconxVersionInfo `json:"version_info"`

	// windows version bitness
	WindowsVersionBitness int32 `json:"windows_version_bitness,omitempty"`

	// windows version edition
	WindowsVersionEdition string `json:"windows_version_edition,omitempty"`

	// windows version name
	WindowsVersionName string `json:"windows_version_name,omitempty"`

	// windows version service pack
	WindowsVersionServicePack string `json:"windows_version_service_pack,omitempty"`

	// windows version version
	WindowsVersionVersion string `json:"windows_version_version,omitempty"`
}

// Validate validates this falconx sandbox report v1
func (m *FalconxSandboxReportV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactedHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtractedFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtractedInterestingStrings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileImports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncidents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryForensics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMitreAttacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcesses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuricataAlerts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxSandboxReportV1) validateContactedHosts(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactedHosts) { // not required
		return nil
	}

	for i := 0; i < len(m.ContactedHosts); i++ {
		if swag.IsZero(m.ContactedHosts[i]) { // not required
			continue
		}

		if m.ContactedHosts[i] != nil {
			if err := m.ContactedHosts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contacted_hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateDNSRequests(formats strfmt.Registry) error {

	if swag.IsZero(m.DNSRequests) { // not required
		return nil
	}

	for i := 0; i < len(m.DNSRequests); i++ {
		if swag.IsZero(m.DNSRequests[i]) { // not required
			continue
		}

		if m.DNSRequests[i] != nil {
			if err := m.DNSRequests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dns_requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateExtractedFiles(formats strfmt.Registry) error {

	if swag.IsZero(m.ExtractedFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.ExtractedFiles); i++ {
		if swag.IsZero(m.ExtractedFiles[i]) { // not required
			continue
		}

		if m.ExtractedFiles[i] != nil {
			if err := m.ExtractedFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extracted_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateExtractedInterestingStrings(formats strfmt.Registry) error {

	if swag.IsZero(m.ExtractedInterestingStrings) { // not required
		return nil
	}

	for i := 0; i < len(m.ExtractedInterestingStrings); i++ {
		if swag.IsZero(m.ExtractedInterestingStrings[i]) { // not required
			continue
		}

		if m.ExtractedInterestingStrings[i] != nil {
			if err := m.ExtractedInterestingStrings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extracted_interesting_strings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateFileImports(formats strfmt.Registry) error {

	if swag.IsZero(m.FileImports) { // not required
		return nil
	}

	for i := 0; i < len(m.FileImports); i++ {
		if swag.IsZero(m.FileImports[i]) { // not required
			continue
		}

		if m.FileImports[i] != nil {
			if err := m.FileImports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("file_imports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateFileMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.FileMetadata) { // not required
		return nil
	}

	if m.FileMetadata != nil {
		if err := m.FileMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file_metadata")
			}
			return err
		}
	}

	return nil
}

func (m *FalconxSandboxReportV1) validateHTTPRequests(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPRequests) { // not required
		return nil
	}

	for i := 0; i < len(m.HTTPRequests); i++ {
		if swag.IsZero(m.HTTPRequests[i]) { // not required
			continue
		}

		if m.HTTPRequests[i] != nil {
			if err := m.HTTPRequests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("http_requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateIncidents(formats strfmt.Registry) error {

	if swag.IsZero(m.Incidents) { // not required
		return nil
	}

	for i := 0; i < len(m.Incidents); i++ {
		if swag.IsZero(m.Incidents[i]) { // not required
			continue
		}

		if m.Incidents[i] != nil {
			if err := m.Incidents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incidents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateMemoryForensics(formats strfmt.Registry) error {

	if swag.IsZero(m.MemoryForensics) { // not required
		return nil
	}

	for i := 0; i < len(m.MemoryForensics); i++ {
		if swag.IsZero(m.MemoryForensics[i]) { // not required
			continue
		}

		if m.MemoryForensics[i] != nil {
			if err := m.MemoryForensics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("memory_forensics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateMitreAttacks(formats strfmt.Registry) error {

	if swag.IsZero(m.MitreAttacks) { // not required
		return nil
	}

	for i := 0; i < len(m.MitreAttacks); i++ {
		if swag.IsZero(m.MitreAttacks[i]) { // not required
			continue
		}

		if m.MitreAttacks[i] != nil {
			if err := m.MitreAttacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mitre_attacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateProcesses(formats strfmt.Registry) error {

	if swag.IsZero(m.Processes) { // not required
		return nil
	}

	for i := 0; i < len(m.Processes); i++ {
		if swag.IsZero(m.Processes[i]) { // not required
			continue
		}

		if m.Processes[i] != nil {
			if err := m.Processes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateSignatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Signatures) { // not required
		return nil
	}

	for i := 0; i < len(m.Signatures); i++ {
		if swag.IsZero(m.Signatures[i]) { // not required
			continue
		}

		if m.Signatures[i] != nil {
			if err := m.Signatures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signatures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateSuricataAlerts(formats strfmt.Registry) error {

	if swag.IsZero(m.SuricataAlerts) { // not required
		return nil
	}

	for i := 0; i < len(m.SuricataAlerts); i++ {
		if swag.IsZero(m.SuricataAlerts[i]) { // not required
			continue
		}

		if m.SuricataAlerts[i] != nil {
			if err := m.SuricataAlerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("suricata_alerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateVersionInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.VersionInfo); i++ {
		if swag.IsZero(m.VersionInfo[i]) { // not required
			continue
		}

		if m.VersionInfo[i] != nil {
			if err := m.VersionInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("version_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FalconxSandboxReportV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxSandboxReportV1) UnmarshalBinary(b []byte) error {
	var res FalconxSandboxReportV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
