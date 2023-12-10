package provider

import (
  "context"
	"encoding/base64"

  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"golang.org/x/crypto/curve25519"
)

// Ensure the implementation satisfies the expected interfaces.
var (
  _ datasource.DataSource = &x25519DataSource{}
)

// x25519Model maps x25519 schema data.
type x25519DataSourceModel struct {
	ID      						types.String   				`tfsdk:"id"`
  PrivateKey					types.String					`tfsdk:"private_key"`
  PublicKey      			types.String					`tfsdk:"public_key"`
}

// x25519DataSource is the data source implementation.
type x25519DataSource struct{}

// Schema defines the schema for the data source.
func (d *x25519DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Computed: true,
      },
			"public_key": schema.StringAttribute{
				Computed: true,
			},
			"private_key": schema.StringAttribute{
				Required:  true,
			},
    },
  }
}

// NewX25519DataSource is a helper function to simplify the provider implementation.
func NewX25519DataSource() datasource.DataSource {
  return &x25519DataSource{}
}

// Metadata returns the data source type name.
func (d *x25519DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_public_key"
}

// Read refreshes the Terraform state with the latest data.
func (d *x25519DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state x25519DataSourceModel

	// Get the private key from Terraform state
	var privateKeyStr types.String
	tflog.Info(ctx, "Getting private_key from the state")
	if err := req.Config.GetAttribute(ctx, path.Root("private_key"), &privateKeyStr); err != nil {
		resp.Diagnostics.AddError(
			"Error getting private key attribute",
			"",
		)
		return
	}

	// Decode the private key form base64 to bytes
	tflog.Info(ctx, "Decoding private_key")
	privateKey, err := base64.StdEncoding.DecodeString(privateKeyStr.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error decoding private key",
			err.Error(),
		)
		return
	}

	// Ensure privateKey is exactly 32 bytes long
	tflog.Info(ctx, "Veryfing private_key")
	if len(privateKey) != 32 {
		resp.Diagnostics.AddError("Private key must be 32 bytes long", "")
		return
	}

	// Use curve25519 to derive the public key
	tflog.Info(ctx, "Deriving public key")
	publicKey, err := curve25519.X25519(privateKey, curve25519.Basepoint)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deriving public key",
			err.Error(),
		)
		return
	}

	// Encode the public key to base64
	tflog.Info(ctx, "Encoding public key")
	encodedPublicKey := base64.StdEncoding.EncodeToString(publicKey)

	// Set the keys in the state
	tflog.Info(ctx, "Setting the state keys")
	state.PrivateKey = privateKeyStr
	state.PublicKey = types.StringValue(encodedPublicKey)

	// Set the ID
	tflog.Info(ctx, "Setting the data source ID")
	id := hashForState(string(publicKey))
	state.ID = types.StringValue(id)

	// Set the state in the response
	tflog.Info(ctx, "Setting the state")
  diags := resp.State.Set(ctx, &state)
  resp.Diagnostics.Append(diags...)
  if resp.Diagnostics.HasError() {
    return
  }
}

