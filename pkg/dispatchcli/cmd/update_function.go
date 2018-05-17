///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package cmd

import (
	"context"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// CallUpdateFunction makes the API call to update a function
func CallUpdateFunction(input interface{}) error {
	function := input.(*v1.Function)

	_, err := functionManagerClient().UpdateFunction(context.TODO(), function)
	if err != nil {
		return formatAPIError(err, function)
	}

	return nil
}
