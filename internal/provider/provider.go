package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
	_ provider.Provider = (*exampleProvider)(nil)
)

type exampleProvider struct {
	version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &exampleProvider{
			version: version,
		}
	}
}

func (p *exampleProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "example"
	resp.Version = p.version
}

func (p *exampleProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *exampleProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
}

func (p *exampleProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func (p *exampleProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
