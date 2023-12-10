package provider

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
    "github.com/hashicorp/terraform-plugin-framework/provider"
    "github.com/hashicorp/terraform-plugin-framework/provider/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource"
		"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
    _ provider.Provider = &x25519Provider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
    return func() provider.Provider {
        return &x25519Provider{
            version: version,
        }
    }
}

// x25519Provider is the provider implementation.
type x25519Provider struct {
    // version is set to the provider version on release, "dev" when the
    // provider is built and ran locally, and "test" when running acceptance
    // testing.
    version string
}

// Metadata returns the provider type name.
func (p *x25519Provider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
    resp.TypeName = "x25519"
    resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *x25519Provider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
    resp.Schema = schema.Schema{}
}

// Configure prepares a x25519 API client for data sources and resources.
func (p *x25519Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring x25519 client")
}

// DataSources defines the data sources implemented in the provider.
func (p *x25519Provider) DataSources(_ context.Context) []func() datasource.DataSource {
  return []func() datasource.DataSource {
    NewX25519DataSource,
  }
}

// Resources defines the resources implemented in the provider.
func (p *x25519Provider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewX25519Resource,
	}
}
