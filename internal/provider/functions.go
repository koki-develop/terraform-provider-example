package provider

import (
	"context"
	"fmt"

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
		Parameters: []function.Parameter{
			function.StringParameter{Name: "name"},
		},
		Return: function.StringReturn{},
	}
}

func (f *HelloFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var name string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &name))
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, fmt.Sprintf("Hello, %s!", name)))
}
