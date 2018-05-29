///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////
package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
	"github.com/vmware/dispatch/pkg/dispatchcli/cmd/mocks"
)

func TestCmdCreateBaseImage(t *testing.T) {
	var buf bytes.Buffer

	cli := NewCLI(os.Stdin, &buf, &buf)
	cli.SetOutput(&buf)
	cli.SetArgs([]string{"create", "base-image", "--help"})
	err := cli.Execute()
	assert.Nil(t, err)
	assert.True(t, strings.Contains(buf.String(), "Create base image"))
}

func TestAddBaseImage(t *testing.T) {
	var stdout, stderr bytes.Buffer

	cli := NewCLI(os.Stdin, &stdout, &stderr)

	bic := &mocks.BaseImageClient{}

	language = "python3"
	args := []string{"test", "vmware/base-python3:test"}
	dispatchConfig.JSON = true

	bic.On("Add", mock.Anything).Once().Return(nil)
	err := createBaseImage(&stdout, &stderr, cli, args, bic)
	assert.NoError(t, err)

	imgObj := make(map[string]interface{})
	err = json.Unmarshal(stdout.Bytes(), &imgObj)
	assert.NoError(t, err)
	assert.EqualValues(t, "test", imgObj["name"])
	assert.EqualValues(t, "vmware/base-python3:test", imgObj["dockerUrl"])
}
