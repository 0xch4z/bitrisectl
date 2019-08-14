// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/activity"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/addons"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/android_keystore_file"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/app_setup"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/application"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/avatar_candidate"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/build_artifact"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/build_certificate"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/build_request"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/builds"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/generic_project_file"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/organizations"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/outgoing_webhook"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/provisioning_profile"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/test_devices"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/user"
	"github.com/charliekenney23/bitrisectl/var/work/bitrise/client/webhook_delivery_item"
)

// Default bitrise HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.bitrise.io"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v0.1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new bitrise HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Bitrise {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new bitrise HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Bitrise {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new bitrise client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Bitrise {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Bitrise)
	cli.Transport = transport

	cli.Activity = activity.New(transport, formats)

	cli.Addons = addons.New(transport, formats)

	cli.AndroidKeystoreFile = android_keystore_file.New(transport, formats)

	cli.AppSetup = app_setup.New(transport, formats)

	cli.Application = application.New(transport, formats)

	cli.AvatarCandidate = avatar_candidate.New(transport, formats)

	cli.BuildArtifact = build_artifact.New(transport, formats)

	cli.BuildCertificate = build_certificate.New(transport, formats)

	cli.BuildRequest = build_request.New(transport, formats)

	cli.Builds = builds.New(transport, formats)

	cli.GenericProjectFile = generic_project_file.New(transport, formats)

	cli.Organizations = organizations.New(transport, formats)

	cli.OutgoingWebhook = outgoing_webhook.New(transport, formats)

	cli.ProvisioningProfile = provisioning_profile.New(transport, formats)

	cli.TestDevices = test_devices.New(transport, formats)

	cli.User = user.New(transport, formats)

	cli.WebhookDeliveryItem = webhook_delivery_item.New(transport, formats)

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

// Bitrise is a client for bitrise
type Bitrise struct {
	Activity *activity.Client

	Addons *addons.Client

	AndroidKeystoreFile *android_keystore_file.Client

	AppSetup *app_setup.Client

	Application *application.Client

	AvatarCandidate *avatar_candidate.Client

	BuildArtifact *build_artifact.Client

	BuildCertificate *build_certificate.Client

	BuildRequest *build_request.Client

	Builds *builds.Client

	GenericProjectFile *generic_project_file.Client

	Organizations *organizations.Client

	OutgoingWebhook *outgoing_webhook.Client

	ProvisioningProfile *provisioning_profile.Client

	TestDevices *test_devices.Client

	User *user.Client

	WebhookDeliveryItem *webhook_delivery_item.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Bitrise) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Activity.SetTransport(transport)

	c.Addons.SetTransport(transport)

	c.AndroidKeystoreFile.SetTransport(transport)

	c.AppSetup.SetTransport(transport)

	c.Application.SetTransport(transport)

	c.AvatarCandidate.SetTransport(transport)

	c.BuildArtifact.SetTransport(transport)

	c.BuildCertificate.SetTransport(transport)

	c.BuildRequest.SetTransport(transport)

	c.Builds.SetTransport(transport)

	c.GenericProjectFile.SetTransport(transport)

	c.Organizations.SetTransport(transport)

	c.OutgoingWebhook.SetTransport(transport)

	c.ProvisioningProfile.SetTransport(transport)

	c.TestDevices.SetTransport(transport)

	c.User.SetTransport(transport)

	c.WebhookDeliveryItem.SetTransport(transport)

}