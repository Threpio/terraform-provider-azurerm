package iothub

//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Enrichment -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/IotHubs/hub1/Enrichments/enrichment1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=IotHub -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/IotHubs/hub1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ConsumerGroup -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/IotHubs/hub1/eventHubEndpoints/eventHubEndpoint1/ConsumerGroups/consumerGroup1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=FallbackRoute -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/IotHubs/hub1/FallbackRoute/default
