package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = (*HelloFunction)(nil)
)

type HelloFunction struct{}

func NewHelloFunction() function.Function {
	return &HelloFunction{}
}

func (f *HelloFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "hello"
}

func (f *HelloFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Return: function.StringReturn{},
	}
}

func (f *HelloFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, "Hello, world!"))
}
