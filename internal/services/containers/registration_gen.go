package containers

// @tombuildsstuff: this file is normally autogenerated and will be replaced in the next regeneration
// it's just present to allow the `containers` service to become auto-generated

import "github.com/hashicorp/terraform-provider-azurerm/internal/sdk"

var _ sdk.TypedServiceRegistration = autoRegistration{}

type autoRegistration struct{}

func (a autoRegistration) Name() string {
	return "TODO"
}

func (a autoRegistration) DataSources() []sdk.DataSource {
	return []sdk.DataSource{}
}

func (a autoRegistration) Resources() []sdk.Resource {
	return []sdk.Resource{}
}

func (a autoRegistration) WebsiteCategories() []string {
	return []string{}
}
