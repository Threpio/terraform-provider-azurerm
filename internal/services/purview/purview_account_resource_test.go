package purview_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/response"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/purview/sdk/2020-12-01-preview/account"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

type PurviewAccountResource struct{}

func TestAccPurviewAccount_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_purview_account", "test")
	r := PurviewAccountResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccPurviewAccount_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_purview_account", "test")
	r := PurviewAccountResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

func TestAccPurviewAccount_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_purview_account", "test")
	r := PurviewAccountResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}
func TestAccPurviewAccount_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_purview_account", "test")
	r := PurviewAccountResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r PurviewAccountResource) Exists(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := account.ParseAccountID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.Purview.AccountsClient.Get(ctx, *id)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return utils.Bool(false), nil
		}
		return nil, fmt.Errorf("retrieving Purview Account %q (Resource Group %q): %+v", id.Name, id.ResourceGroup, err)
	}

	return utils.Bool(true), nil
}

func (r PurviewAccountResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-purview-%d"
  location = "%s"
}


resource "azurerm_purview_account" "test" {
  name                = "acctestpa%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku_name            = "Standard_1"

  identity {
    type = "SystemAssigned"
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}

func (r PurviewAccountResource) requiresImport(data acceptance.TestData) string {
	template := r.basic(data)
	return fmt.Sprintf(`
%s

resource "azurerm_purview_account" "import" {
  name                = azurerm_purview_account.test.name
  resource_group_name = azurerm_purview_account.test.resource_group_name
  location            = azurerm_purview_account.test.location
  sku_name            = azurerm_purview_account.test.sku_name

  identity {
    type = "SystemAssigned"
  }
}
`, template)
}

func (r PurviewAccountResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-purview-%d"
  location = "%s"
}


resource "azurerm_purview_account" "test" {
  name                = "acctestpa%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku_name            = "Standard_1"

  identity {
    type = "SystemAssigned"
  }

  tags {
    hello = "World"
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}
