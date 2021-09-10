package iothub

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-provider-azurerm/internal/services/iothub/parse"

	"github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2020-03-01/devices"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/azure"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/locks"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/iothub/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

func resourceIotHubFallbackRoute() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceIotHubFallbackRouteCreateUpdate,
		Read:   resourceIotHubFallbackRouteRead,
		Update: resourceIotHubFallbackRouteCreateUpdate,
		Delete: resourceIotHubFallbackRouteDelete,
		// TODO: replace this with an importer which validates the ID during import
		Importer: pluginsdk.DefaultImporter(),

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(30 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(30 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"resource_group_name": azure.SchemaResourceGroupName(),

			"iothub_name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.IoTHubName,
			},

			"condition": {
				// The condition is a string value representing device-to-cloud message routes query expression
				// https://docs.microsoft.com/en-us/azure/iot-hub/iot-hub-devguide-query-language#device-to-cloud-message-routes-query-expressions
				Type:     pluginsdk.TypeString,
				Optional: true,
				Default:  "true",
			},

			"endpoint_names": {
				Type:     pluginsdk.TypeList,
				Required: true,
				// Currently only one endpoint is allowed. With that comment from Microsoft, we'll leave this open to enhancement when they add multiple endpoint support.
				MaxItems: 1,
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validate.IoTHubEndpointName,
				},
			},

			"enabled": {
				Type:     pluginsdk.TypeBool,
				Required: true,
			},
		},
	}
}

func resourceIotHubFallbackRouteCreateUpdate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).IoTHub.ResourceClient
	subscriptionId := meta.(*clients.Client).IoTHub.ResourceClient.SubscriptionID
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	iothubId := parse.NewIotHubID(subscriptionId, d.Get("resource_group_name").(string), d.Get("iothub_name").(string))

	id := parse.NewFallbackRouteID(iothubId.SubscriptionId, iothubId.ResourceGroup, iothubId.Name, "default")
	locks.ByName(id.IotHubName, IothubResourceName)
	defer locks.UnlockByName(id.IotHubName, IothubResourceName)

	iothub, err := client.Get(ctx, id.ResourceGroup, id.IotHubName)
	if err != nil {
		if utils.ResponseWasNotFound(iothub.Response) {
			return fmt.Errorf("checking for presence of existing %s: %+v", id.String(), err)
		}

		return fmt.Errorf("loading %s: %+v", id.String(), err)
	}

	// NOTE: this resource intentionally doesn't support Requires Import
	//       since a fallback route is created by default

	routing := iothub.Properties.Routing

	if routing == nil {
		routing = &devices.RoutingProperties{}
	}

	routing.FallbackRoute = &devices.FallbackRouteProperties{
		Source:        utils.String(string(devices.RoutingSourceDeviceMessages)),
		Condition:     utils.String(d.Get("condition").(string)),
		EndpointNames: utils.ExpandStringSlice(d.Get("endpoint_names").([]interface{})),
		IsEnabled:     utils.Bool(d.Get("enabled").(bool)),
	}

	future, err := client.CreateOrUpdate(ctx, id.ResourceGroup, id.IotHubName, iothub, "")
	if err != nil {
		return fmt.Errorf("creating/updating %s: %+v", id.String(), err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for the completion of the creating/updating of %s: %+v", id.String(), err)
	}

	d.SetId(id.ID())

	return resourceIotHubFallbackRouteRead(d, meta)
}

func resourceIotHubFallbackRouteRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).IoTHub.ResourceClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.FallbackRouteID(d.Id())
	if err != nil {
		return err
	}

	iothub, err := client.Get(ctx, id.ResourceGroup, id.IotHubName)
	if err != nil {
		return fmt.Errorf("loading %s: %+v", id.String(), err)
	}

	d.Set("iothub_name", id.IotHubName)
	d.Set("resource_group_name", id.ResourceGroup)

	if props := iothub.Properties; props != nil {
		if routing := props.Routing; routing != nil {
			if fallbackRoute := routing.FallbackRoute; fallbackRoute != nil {
				d.Set("condition", fallbackRoute.Condition)
				d.Set("enabled", fallbackRoute.IsEnabled)
				d.Set("endpoint_names", fallbackRoute.EndpointNames)
			}
		}
	}

	return nil
}

func resourceIotHubFallbackRouteDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).IoTHub.ResourceClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.FallbackRouteID(d.Id())
	if err != nil {
		return err
	}

	locks.ByName(id.IotHubName, IothubResourceName)
	defer locks.UnlockByName(id.IotHubName, IothubResourceName)

	iothub, err := client.Get(ctx, id.ResourceGroup, id.IotHubName)
	if err != nil {
		if utils.ResponseWasNotFound(iothub.Response) {
			return fmt.Errorf("IotHub %s was not found", id.String())
		}

		return fmt.Errorf("loading %s: %+v", id.String(), err)
	}

	if iothub.Properties == nil || iothub.Properties.Routing == nil || iothub.Properties.Routing.FallbackRoute == nil {
		return nil
	}

	iothub.Properties.Routing.FallbackRoute = nil
	future, err := client.CreateOrUpdate(ctx, id.ResourceGroup, id.IotHubName, iothub, "")
	if err != nil {
		return fmt.Errorf("updating %s with Fallback Route: %+v", id.String(), err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for %s to finish updating Fallback Route: %+v", id.String(), err)
	}

	return nil
}
