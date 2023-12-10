package provider

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"golang.org/x/crypto/curve25519"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &x25519Resource{}
)

// x25519ResourceModel maps the resource schema data.
type x25519ResourceModel struct {
	ID         types.String `tfsdk:"id"`
	PrivateKey types.String `tfsdk:"private_key"`
	PublicKey  types.String `tfsdk:"public_key"`
}

// NewX25519Resource is a helper function to simplify the provider implementation.
func NewX25519Resource() resource.Resource {
	return &x25519Resource{}
}

// x25519Resource is the resource implementation.
type x25519Resource struct{}

// Metadata returns the resource type name.
func (r *x25519Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_private_key"
}

// Schema defines the schema for the resource.
func (r *x25519Resource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "Unique identifier for this resource: " +
					"hexadecimal representation of the SHA1 checksum of the resource.",
			},
			"private_key": schema.StringAttribute{
				Computed:            true,
				Sensitive:           true,
				MarkdownDescription: "Base64 encoded private key data [(RFC 7748)](https://datatracker.ietf.org/doc/html/rfc7748#section-5) format.",
			},
			"public_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Base64 encoded public key data [(RFC 7748)](https://datatracker.ietf.org/doc/html/rfc7748#section-5) format.",
			},
		},
		MarkdownDescription: "Generates a base64-encoded private key for Curve25519. " +
			"[(RFC 7748)](https://datatracker.ietf.org/doc/html/rfc7748#section-5).\n\n" +
			"This resource is primarily intended for easily bootstrapping throwaway development environments.",
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *x25519Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan x25519ResourceModel

	tflog.Info(ctx, "Getting the plan")
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate a random 32-byte private key
	tflog.Info(ctx, "Generating the private key")
	privateKey := make([]byte, curve25519.ScalarSize)
	_, err := rand.Read(privateKey)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error generating the random key",
			err.Error(),
		)
		return
	}

	// Clamp the private key
	tflog.Info(ctx, "Clamping the private key")
	privateKey[0] &= 248
	privateKey[31] = (privateKey[31] & 127) | 64

	// Derive the public key from the private key
	tflog.Info(ctx, "Deriving the public key")
	publicKey, err := curve25519.X25519(privateKey, curve25519.Basepoint)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deriving public key",
			err.Error(),
		)
		return
	}

	// Encode private and public keys to base64
	tflog.Info(ctx, "Encoding the keys")
	encodedPrivateKey := base64.StdEncoding.EncodeToString(privateKey)
	encodedPublicKey := base64.StdEncoding.EncodeToString(publicKey)

	// Set the keys in the plan
	tflog.Info(ctx, "Setting the plan values")
	plan.PrivateKey = types.StringValue(encodedPrivateKey)
	plan.PrivateKey = types.StringValue(encodedPrivateKey)
	plan.PublicKey = types.StringValue(encodedPublicKey)

	// Set the ID
	tflog.Info(ctx, "Setting the resource ID")
	id := hashForState(string(publicKey))
	plan.ID = types.StringValue(id)

	// Set state to fully populated data
	tflog.Info(ctx, "Setting the state")
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *x25519Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *x25519Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *x25519Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
