///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package eventmanager

import (
	"reflect"

	ewrapper "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/vmware/dispatch/pkg/entity-store"
	"github.com/vmware/dispatch/pkg/trace"
)

type subscriptionEntityHandler struct {
	store   entitystore.EntityStore
	manager SubscriptionManager
}

func (h *subscriptionEntityHandler) Type() reflect.Type {
	defer trace.Trace("")()

	return reflect.TypeOf(&Subscription{})
}

func (h *subscriptionEntityHandler) Add(obj entitystore.Entity) (err error) {
	defer trace.Tracef("name %s", obj.GetName())()

	sub := obj.(*Subscription)
	defer func() { h.store.UpdateWithError(sub, err) }()

	if err := h.manager.Create(sub); err != nil {
		return ewrapper.Wrap(err, "error activating subscription")
	}

	sub.Status = entitystore.StatusREADY

	log.Infof("subscription %s for topic %s has been activated", sub.Name, sub.Topic)

	return nil
}

func (h *subscriptionEntityHandler) Update(obj entitystore.Entity) error {
	defer trace.Trace("")()
	return h.Add(obj)
}

func (h *subscriptionEntityHandler) Delete(obj entitystore.Entity) error {
	defer trace.Tracef("name '%s'", obj.GetName())()

	sub := obj.(*Subscription)

	// unsubscribe from queue
	err := h.manager.Delete(sub)
	if err != nil {
		return ewrapper.Wrap(err, "error deactivating subscription")
	}

	// hard deletion
	if err := h.store.Delete(sub.OrganizationID, sub.Name, sub); err != nil {
		return ewrapper.Wrap(err, "store error when deleting subscription")
	}
	log.Infof("subscription %s deactivated and deleted from the entity store", sub.Name)
	return nil
}

func (h *subscriptionEntityHandler) Error(obj entitystore.Entity) error {
	defer trace.Tracef("")()

	log.Errorf("handleError func not implemented yet")
	return nil
}
