package shim

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
	p "github.com/itoam/terraform-provider-openstatus/internal/provider"
)

func NewProvider() provider.Provider {
	return p.New()()
}