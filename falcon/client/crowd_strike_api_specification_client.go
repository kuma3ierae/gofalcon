// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/crowdstrike/gofalcon/falcon/client/alerts"
	"github.com/crowdstrike/gofalcon/falcon/client/cloud_connect_aws"
	"github.com/crowdstrike/gofalcon/falcon/client/cspm_registration"
	"github.com/crowdstrike/gofalcon/falcon/client/custom_ioa"
	"github.com/crowdstrike/gofalcon/falcon/client/d4c_registration"
	"github.com/crowdstrike/gofalcon/falcon/client/detects"
	"github.com/crowdstrike/gofalcon/falcon/client/device_control_policies"
	"github.com/crowdstrike/gofalcon/falcon/client/discover"
	"github.com/crowdstrike/gofalcon/falcon/client/event_streams"
	"github.com/crowdstrike/gofalcon/falcon/client/falcon_complete_dashboard"
	"github.com/crowdstrike/gofalcon/falcon/client/falcon_container"
	"github.com/crowdstrike/gofalcon/falcon/client/falcon_container_cli"
	"github.com/crowdstrike/gofalcon/falcon/client/falconx_sandbox"
	"github.com/crowdstrike/gofalcon/falcon/client/filevantage"
	"github.com/crowdstrike/gofalcon/falcon/client/firewall_management"
	"github.com/crowdstrike/gofalcon/falcon/client/firewall_policies"
	"github.com/crowdstrike/gofalcon/falcon/client/host_group"
	"github.com/crowdstrike/gofalcon/falcon/client/hosts"
	"github.com/crowdstrike/gofalcon/falcon/client/identity_protection"
	"github.com/crowdstrike/gofalcon/falcon/client/incidents"
	"github.com/crowdstrike/gofalcon/falcon/client/incoming_webhook_requests"
	"github.com/crowdstrike/gofalcon/falcon/client/installation_tokens"
	"github.com/crowdstrike/gofalcon/falcon/client/intel"
	"github.com/crowdstrike/gofalcon/falcon/client/ioa_exclusions"
	"github.com/crowdstrike/gofalcon/falcon/client/ioc"
	"github.com/crowdstrike/gofalcon/falcon/client/iocs"
	"github.com/crowdstrike/gofalcon/falcon/client/kubernetes_protection"
	"github.com/crowdstrike/gofalcon/falcon/client/malquery"
	"github.com/crowdstrike/gofalcon/falcon/client/message_center"
	"github.com/crowdstrike/gofalcon/falcon/client/ml_exclusions"
	"github.com/crowdstrike/gofalcon/falcon/client/mobile_enrollment"
	"github.com/crowdstrike/gofalcon/falcon/client/mssp"
	"github.com/crowdstrike/gofalcon/falcon/client/oauth2"
	"github.com/crowdstrike/gofalcon/falcon/client/ods"
	"github.com/crowdstrike/gofalcon/falcon/client/operations"
	"github.com/crowdstrike/gofalcon/falcon/client/overwatch_dashboard"
	"github.com/crowdstrike/gofalcon/falcon/client/prevention_policies"
	"github.com/crowdstrike/gofalcon/falcon/client/quarantine"
	"github.com/crowdstrike/gofalcon/falcon/client/quick_scan"
	"github.com/crowdstrike/gofalcon/falcon/client/real_time_response"
	"github.com/crowdstrike/gofalcon/falcon/client/real_time_response_admin"
	"github.com/crowdstrike/gofalcon/falcon/client/recon"
	"github.com/crowdstrike/gofalcon/falcon/client/report_executions"
	"github.com/crowdstrike/gofalcon/falcon/client/response_policies"
	"github.com/crowdstrike/gofalcon/falcon/client/sample_uploads"
	"github.com/crowdstrike/gofalcon/falcon/client/scheduled_reports"
	"github.com/crowdstrike/gofalcon/falcon/client/sensor_download"
	"github.com/crowdstrike/gofalcon/falcon/client/sensor_update_policies"
	"github.com/crowdstrike/gofalcon/falcon/client/sensor_visibility_exclusions"
	"github.com/crowdstrike/gofalcon/falcon/client/spotlight_evaluation_logic"
	"github.com/crowdstrike/gofalcon/falcon/client/spotlight_vulnerabilities"
	"github.com/crowdstrike/gofalcon/falcon/client/store"
	"github.com/crowdstrike/gofalcon/falcon/client/tailored_intelligence"
	"github.com/crowdstrike/gofalcon/falcon/client/user_management"
	"github.com/crowdstrike/gofalcon/falcon/client/zero_trust_assessment"
)

// Default crowd strike API specification HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.crowdstrike.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new crowd strike API specification HTTP client.
func NewHTTPClient(formats strfmt.Registry) *CrowdStrikeAPISpecification {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new crowd strike API specification HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *CrowdStrikeAPISpecification {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new crowd strike API specification client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *CrowdStrikeAPISpecification {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(CrowdStrikeAPISpecification)
	cli.Transport = transport
	cli.Alerts = alerts.New(transport, formats)
	cli.CloudConnectAws = cloud_connect_aws.New(transport, formats)
	cli.CspmRegistration = cspm_registration.New(transport, formats)
	cli.CustomIoa = custom_ioa.New(transport, formats)
	cli.D4cRegistration = d4c_registration.New(transport, formats)
	cli.Detects = detects.New(transport, formats)
	cli.DeviceControlPolicies = device_control_policies.New(transport, formats)
	cli.Discover = discover.New(transport, formats)
	cli.EventStreams = event_streams.New(transport, formats)
	cli.FalconCompleteDashboard = falcon_complete_dashboard.New(transport, formats)
	cli.FalconContainer = falcon_container.New(transport, formats)
	cli.FalconContainerCli = falcon_container_cli.New(transport, formats)
	cli.FalconxSandbox = falconx_sandbox.New(transport, formats)
	cli.Filevantage = filevantage.New(transport, formats)
	cli.FirewallManagement = firewall_management.New(transport, formats)
	cli.FirewallPolicies = firewall_policies.New(transport, formats)
	cli.HostGroup = host_group.New(transport, formats)
	cli.Hosts = hosts.New(transport, formats)
	cli.IdentityProtection = identity_protection.New(transport, formats)
	cli.Incidents = incidents.New(transport, formats)
	cli.IncomingWebhookRequests = incoming_webhook_requests.New(transport, formats)
	cli.InstallationTokens = installation_tokens.New(transport, formats)
	cli.Intel = intel.New(transport, formats)
	cli.IoaExclusions = ioa_exclusions.New(transport, formats)
	cli.Ioc = ioc.New(transport, formats)
	cli.Iocs = iocs.New(transport, formats)
	cli.KubernetesProtection = kubernetes_protection.New(transport, formats)
	cli.Malquery = malquery.New(transport, formats)
	cli.MessageCenter = message_center.New(transport, formats)
	cli.MlExclusions = ml_exclusions.New(transport, formats)
	cli.MobileEnrollment = mobile_enrollment.New(transport, formats)
	cli.Mssp = mssp.New(transport, formats)
	cli.Oauth2 = oauth2.New(transport, formats)
	cli.Ods = ods.New(transport, formats)
	cli.Operations = operations.New(transport, formats)
	cli.OverwatchDashboard = overwatch_dashboard.New(transport, formats)
	cli.PreventionPolicies = prevention_policies.New(transport, formats)
	cli.Quarantine = quarantine.New(transport, formats)
	cli.QuickScan = quick_scan.New(transport, formats)
	cli.RealTimeResponse = real_time_response.New(transport, formats)
	cli.RealTimeResponseAdmin = real_time_response_admin.New(transport, formats)
	cli.Recon = recon.New(transport, formats)
	cli.ReportExecutions = report_executions.New(transport, formats)
	cli.ResponsePolicies = response_policies.New(transport, formats)
	cli.SampleUploads = sample_uploads.New(transport, formats)
	cli.ScheduledReports = scheduled_reports.New(transport, formats)
	cli.SensorDownload = sensor_download.New(transport, formats)
	cli.SensorUpdatePolicies = sensor_update_policies.New(transport, formats)
	cli.SensorVisibilityExclusions = sensor_visibility_exclusions.New(transport, formats)
	cli.SpotlightEvaluationLogic = spotlight_evaluation_logic.New(transport, formats)
	cli.SpotlightVulnerabilities = spotlight_vulnerabilities.New(transport, formats)
	cli.Store = store.New(transport, formats)
	cli.TailoredIntelligence = tailored_intelligence.New(transport, formats)
	cli.UserManagement = user_management.New(transport, formats)
	cli.ZeroTrustAssessment = zero_trust_assessment.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// CrowdStrikeAPISpecification is a client for crowd strike API specification
type CrowdStrikeAPISpecification struct {
	Alerts alerts.ClientService

	CloudConnectAws cloud_connect_aws.ClientService

	CspmRegistration cspm_registration.ClientService

	CustomIoa custom_ioa.ClientService

	D4cRegistration d4c_registration.ClientService

	Detects detects.ClientService

	DeviceControlPolicies device_control_policies.ClientService

	Discover discover.ClientService

	EventStreams event_streams.ClientService

	FalconCompleteDashboard falcon_complete_dashboard.ClientService

	FalconContainer falcon_container.ClientService

	FalconContainerCli falcon_container_cli.ClientService

	FalconxSandbox falconx_sandbox.ClientService

	Filevantage filevantage.ClientService

	FirewallManagement firewall_management.ClientService

	FirewallPolicies firewall_policies.ClientService

	HostGroup host_group.ClientService

	Hosts hosts.ClientService

	IdentityProtection identity_protection.ClientService

	Incidents incidents.ClientService

	IncomingWebhookRequests incoming_webhook_requests.ClientService

	InstallationTokens installation_tokens.ClientService

	Intel intel.ClientService

	IoaExclusions ioa_exclusions.ClientService

	Ioc ioc.ClientService

	Iocs iocs.ClientService

	KubernetesProtection kubernetes_protection.ClientService

	Malquery malquery.ClientService

	MessageCenter message_center.ClientService

	MlExclusions ml_exclusions.ClientService

	MobileEnrollment mobile_enrollment.ClientService

	Mssp mssp.ClientService

	Oauth2 oauth2.ClientService

	Ods ods.ClientService

	Operations operations.ClientService

	OverwatchDashboard overwatch_dashboard.ClientService

	PreventionPolicies prevention_policies.ClientService

	Quarantine quarantine.ClientService

	QuickScan quick_scan.ClientService

	RealTimeResponse real_time_response.ClientService

	RealTimeResponseAdmin real_time_response_admin.ClientService

	Recon recon.ClientService

	ReportExecutions report_executions.ClientService

	ResponsePolicies response_policies.ClientService

	SampleUploads sample_uploads.ClientService

	ScheduledReports scheduled_reports.ClientService

	SensorDownload sensor_download.ClientService

	SensorUpdatePolicies sensor_update_policies.ClientService

	SensorVisibilityExclusions sensor_visibility_exclusions.ClientService

	SpotlightEvaluationLogic spotlight_evaluation_logic.ClientService

	SpotlightVulnerabilities spotlight_vulnerabilities.ClientService

	Store store.ClientService

	TailoredIntelligence tailored_intelligence.ClientService

	UserManagement user_management.ClientService

	ZeroTrustAssessment zero_trust_assessment.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *CrowdStrikeAPISpecification) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Alerts.SetTransport(transport)
	c.CloudConnectAws.SetTransport(transport)
	c.CspmRegistration.SetTransport(transport)
	c.CustomIoa.SetTransport(transport)
	c.D4cRegistration.SetTransport(transport)
	c.Detects.SetTransport(transport)
	c.DeviceControlPolicies.SetTransport(transport)
	c.Discover.SetTransport(transport)
	c.EventStreams.SetTransport(transport)
	c.FalconCompleteDashboard.SetTransport(transport)
	c.FalconContainer.SetTransport(transport)
	c.FalconContainerCli.SetTransport(transport)
	c.FalconxSandbox.SetTransport(transport)
	c.Filevantage.SetTransport(transport)
	c.FirewallManagement.SetTransport(transport)
	c.FirewallPolicies.SetTransport(transport)
	c.HostGroup.SetTransport(transport)
	c.Hosts.SetTransport(transport)
	c.IdentityProtection.SetTransport(transport)
	c.Incidents.SetTransport(transport)
	c.IncomingWebhookRequests.SetTransport(transport)
	c.InstallationTokens.SetTransport(transport)
	c.Intel.SetTransport(transport)
	c.IoaExclusions.SetTransport(transport)
	c.Ioc.SetTransport(transport)
	c.Iocs.SetTransport(transport)
	c.KubernetesProtection.SetTransport(transport)
	c.Malquery.SetTransport(transport)
	c.MessageCenter.SetTransport(transport)
	c.MlExclusions.SetTransport(transport)
	c.MobileEnrollment.SetTransport(transport)
	c.Mssp.SetTransport(transport)
	c.Oauth2.SetTransport(transport)
	c.Ods.SetTransport(transport)
	c.Operations.SetTransport(transport)
	c.OverwatchDashboard.SetTransport(transport)
	c.PreventionPolicies.SetTransport(transport)
	c.Quarantine.SetTransport(transport)
	c.QuickScan.SetTransport(transport)
	c.RealTimeResponse.SetTransport(transport)
	c.RealTimeResponseAdmin.SetTransport(transport)
	c.Recon.SetTransport(transport)
	c.ReportExecutions.SetTransport(transport)
	c.ResponsePolicies.SetTransport(transport)
	c.SampleUploads.SetTransport(transport)
	c.ScheduledReports.SetTransport(transport)
	c.SensorDownload.SetTransport(transport)
	c.SensorUpdatePolicies.SetTransport(transport)
	c.SensorVisibilityExclusions.SetTransport(transport)
	c.SpotlightEvaluationLogic.SetTransport(transport)
	c.SpotlightVulnerabilities.SetTransport(transport)
	c.Store.SetTransport(transport)
	c.TailoredIntelligence.SetTransport(transport)
	c.UserManagement.SetTransport(transport)
	c.ZeroTrustAssessment.SetTransport(transport)
}
