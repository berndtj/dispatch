///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package client

import (
	"context"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"

	"github.com/vmware/dispatch/pkg/api/v1"
	swaggerclient "github.com/vmware/dispatch/pkg/event-manager/gen/client"
	"github.com/vmware/dispatch/pkg/event-manager/gen/client/drivers"
	"github.com/vmware/dispatch/pkg/event-manager/gen/client/events"
	"github.com/vmware/dispatch/pkg/event-manager/gen/client/subscriptions"
)

// EventsClient defines the event client interface
type EventsClient interface {
	// Emit an event
	EmitEvent(ctx context.Context, emission *v1.Emission) (*v1.Emission, error)

	// Subscriptions
	CreateSubscription(ctx context.Context, subscription *v1.Subscription) (*v1.Subscription, error)
	DeleteSubscription(ctx context.Context, subscriptionName string) (*v1.Subscription, error)
	GetSubscription(ctx context.Context, subscriptionName string) (*v1.Subscription, error)
	ListSubscriptions(ctx context.Context) ([]v1.Subscription, error)
	UpdateSubscription(ctx context.Context, subscription *v1.Subscription) (*v1.Subscription, error)

	// Event Drivers
	CreateEventDriver(ctx context.Context, eventDriver *v1.EventDriver) (*v1.EventDriver, error)
	DeleteEventDriver(ctx context.Context, eventDriverName string) (*v1.EventDriver, error)
	GetEventDriver(ctx context.Context, eventDriverName string) (*v1.EventDriver, error)
	ListEventDrivers(ctx context.Context) ([]v1.EventDriver, error)
	UpdateEventDriver(ctx context.Context, eventDriver *v1.EventDriver) (*v1.EventDriver, error)

	// Event Driver Types
	CreateEventDriverType(ctx context.Context, eventDriverType *v1.EventDriverType) (*v1.EventDriverType, error)
	DeleteEventDriverType(ctx context.Context, eventDriverTypeName string) (*v1.EventDriverType, error)
	GetEventDriverType(ctx context.Context, eventDriverTypeName string) (*v1.EventDriverType, error)
	ListEventDriversType(ctx context.Context) ([]v1.EventDriverType, error)
}

// DefaultEventsClient defines the default client for events API
type DefaultEventsClient struct {
	client *swaggerclient.EventManager
	auth   runtime.ClientAuthInfoWriter
}

// NewEventsClient is used to create a new subscriptions client
func NewEventsClient(host string, auth runtime.ClientAuthInfoWriter) *DefaultEventsClient {
	transport := DefaultHTTPClient(host, swaggerclient.DefaultBasePath)
	return &DefaultEventsClient{
		client: swaggerclient.New(transport, strfmt.Default),
		auth:   auth,
	}
}

// EmitEvent emits an event
func (c *DefaultEventsClient) EmitEvent(ctx context.Context, event *v1.CloudEvent) error {
	emission := &v1.Emission{
		Event: event,
	}
	params := events.EmitEventParams{
		Body:    emission,
		Context: ctx,
	}
	_, err := c.client.Events.EmitEvent(&params, c.auth)
	if err != nil {
		return errors.Wrapf(err, "error when emitting an event %s", event.EventType)
	}
	return nil
}

// CreateSubscription creates and adds a new subscription
func (c *DefaultEventsClient) CreateSubscription(ctx context.Context, subscription *v1.Subscription) (*v1.Subscription, error) {
	params := subscriptions.AddSubscriptionParams{
		Context: ctx,
		Body:    subscription,
	}
	response, err := c.client.Subscriptions.AddSubscription(&params, c.auth)
	if err != nil {
		return nil, errors.Wrap(err, "error when creating a subscription")
	}
	return response.Payload, nil
}

// DeleteSubscription deletes a subscription
func (c *DefaultEventsClient) DeleteSubscription(ctx context.Context, subscriptionName string) (*v1.Subscription, error) {
	params := subscriptions.DeleteSubscriptionParams{
		Context:          ctx,
		SubscriptionName: subscriptionName,
	}
	response, err := c.client.Subscriptions.DeleteSubscription(&params, c.auth)
	if err != nil {
		return nil, errors.Wrapf(err, "error when deleting the subscription %s", subscriptionName)
	}
	return response.Payload, nil
}

// GetSubscription gets a subscription by name
func (c *DefaultEventsClient) GetSubscription(ctx context.Context, subscriptionName string) (*v1.Subscription, error) {
	params := subscriptions.GetSubscriptionParams{
		Context:          ctx,
		SubscriptionName: subscriptionName,
	}
	response, err := c.client.Subscriptions.GetSubscription(&params, c.auth)
	if err != nil {
		return nil, errors.Wrapf(err, "error when retrieving the subscription %s", subscriptionName)
	}
	return response.Payload, nil
}

// ListSubscriptions lists all subscriptions
func (c *DefaultEventsClient) ListSubscriptions(ctx context.Context) ([]v1.Subscription, error) {
	params := subscriptions.GetSubscriptionsParams{
		Context: ctx,
	}
	response, err := c.client.Subscriptions.GetSubscriptions(&params, c.auth)
	if err != nil {
		return nil, errors.Wrap(err, "error when retrieving the subscriptions")
	}
	var subscriptions []v1.Subscription
	for _, f := range response.Payload {
		subscriptions = append(subscriptions, *f)
	}
	return subscriptions, nil
}

// UpdateSubscription updates a specific subscription
func (c *DefaultEventsClient) UpdateSubscription(ctx context.Context, subscription *v1.Subscription) (*v1.Subscription, error) {
	params := subscriptions.UpdateSubscriptionParams{
		Context:          ctx,
		Body:             subscription,
		SubscriptionName: *subscription.Name,
	}
	response, err := c.client.Subscriptions.UpdateSubscription(&params, c.auth)
	if err != nil {
		return nil, errors.Wrapf(err, "error when updating the subscription %s", *subscription.Name)
	}
	return response.Payload, nil
}

// CreateEventDriver creates and adds a new event driver
func (c *DefaultEventsClient) CreateEventDriver(ctx context.Context, driver *v1.EventDriver) (*v1.EventDriver, error) {
	params := drivers.AddDriverParams{
		Context: ctx,
		Body:    driver,
	}
	response, err := c.client.Drivers.AddDriver(&params, c.auth)
	if err != nil {
		return nil, errors.Wrap(err, "error when creating a driver")
	}
	return response.Payload, nil
}

// DeleteEventDriver deletes a driver
func (c *DefaultEventsClient) DeleteEventDriver(ctx context.Context, driverName string) (*v1.EventDriver, error) {
	params := drivers.DeleteDriverParams{
		Context:    ctx,
		DriverName: driverName,
	}
	response, err := c.client.Drivers.DeleteDriver(&params, c.auth)
	if err != nil {
		return nil, errors.Wrapf(err, "error when deleting the driver %s", driverName)
	}
	return response.Payload, nil
}

// GetEventDriver gets a driver by name
func (c *DefaultEventsClient) GetEventDriver(ctx context.Context, driverName string) (*v1.EventDriver, error) {
	params := drivers.GetDriverParams{
		Context:    ctx,
		DriverName: driverName,
	}
	response, err := c.client.Drivers.GetDriver(&params, c.auth)
	if err != nil {
		return nil, errors.Wrapf(err, "error when retrieving the driver %s", driverName)
	}
	return response.Payload, nil
}

// ListEventDrivers lists all drivers
func (c *DefaultEventsClient) ListEventDrivers(ctx context.Context) ([]v1.EventDriver, error) {
	params := drivers.GetDriversParams{
		Context: ctx,
	}
	response, err := c.client.Drivers.GetDrivers(&params, c.auth)
	if err != nil {
		return nil, errors.Wrap(err, "error when retrieving the drivers")
	}
	var drivers []v1.EventDriver
	for _, f := range response.Payload {
		drivers = append(drivers, *f)
	}
	return drivers, nil
}

// UpdateEventDriver updates a specific driver
func (c *DefaultEventsClient) UpdateEventDriver(ctx context.Context, driver *v1.EventDriver) (*v1.EventDriver, error) {
	params := drivers.UpdateDriverParams{
		Context:    ctx,
		Body:       driver,
		DriverName: *driver.Name,
	}
	response, err := c.client.Drivers.UpdateDriver(&params, c.auth)
	if err != nil {
		return nil, errors.Wrapf(err, "error when updating the driver %s", *driver.Name)
	}
	return response.Payload, nil
}

// CreateEventDriverType creates and adds a new subscription
func (c *DefaultEventsClient) CreateEventDriverType(ctx context.Context, driverType *v1.EventDriverType) (*v1.EventDriverType, error) {
	params := drivers.AddDriverTypeParams{
		Context: ctx,
		Body:    driverType,
	}
	response, err := c.client.Drivers.AddDriverType(&params, c.auth)
	if err != nil {
		return nil, errors.Wrap(err, "error when creating a driver type")
	}
	return response.Payload, nil
}

// DeleteEventDriverType deletes a driver
func (c *DefaultEventsClient) DeleteEventDriverType(ctx context.Context, driverTypeName string) (*v1.EventDriverType, error) {
	params := drivers.DeleteDriverTypeParams{
		Context:        ctx,
		DriverTypeName: driverTypeName,
	}
	response, err := c.client.Drivers.DeleteDriverType(&params, c.auth)
	if err != nil {
		return nil, errors.Wrapf(err, "error when deleting the driver type %s", driverTypeName)
	}
	return response.Payload, nil
}

// GetEventDriverType gets a driver by name
func (c *DefaultEventsClient) GetEventDriverType(ctx context.Context, driverTypeName string) (*v1.EventDriverType, error) {
	params := drivers.GetDriverTypeParams{
		Context:        ctx,
		DriverTypeName: driverTypeName,
	}
	response, err := c.client.Drivers.GetDriverType(&params, c.auth)
	if err != nil {
		return nil, errors.Wrapf(err, "error when retrieving the driver type %s", driverTypeName)
	}
	return response.Payload, nil
}

// ListEventDriverTypes lists all drivers
func (c *DefaultEventsClient) ListEventDriverTypes(ctx context.Context) ([]v1.EventDriverType, error) {
	params := drivers.GetDriverTypesParams{
		Context: ctx,
	}
	response, err := c.client.Drivers.GetDriverTypes(&params, c.auth)
	if err != nil {
		return nil, errors.Wrap(err, "error when retrieving the driver types")
	}
	var drivers []v1.EventDriverType
	for _, f := range response.Payload {
		drivers = append(drivers, *f)
	}
	return drivers, nil
}
