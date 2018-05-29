///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"golang.org/x/net/context"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"

	"github.com/vmware/dispatch/pkg/api/v1"
	"github.com/vmware/dispatch/pkg/dispatchcli/i18n"
	imageclient "github.com/vmware/dispatch/pkg/image-manager/gen/client"
	baseimage "github.com/vmware/dispatch/pkg/image-manager/gen/client/base_image"
)

var (
	createBaseImageLong = i18n.T(`Create base image.`)

	// TODO: add examples
	createBaseImageExample = i18n.T(``)
	public                 = false
	language               = i18n.T(``)
)

// NewCmdCreateBaseImage creates command responsible for base image creation.
func NewCmdCreateBaseImage(out io.Writer, errOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "base-image IMAGE_NAME IMAGE_URL [--public] [--language LANGUAGE]",
		Short:   i18n.T("Create base image"),
		Long:    createBaseImageLong,
		Example: createBaseImageExample,
		Args:    cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			ic := NewBaseImageClient()
			err := createBaseImage(out, errOut, cmd, args, ic)
			CheckErr(err)
		},
	}
	cmd.Flags().StringVar(&language, "language", "", "Specify the runtime language for the image")
	return cmd
}

// BaseImageClient is a client interface for base image operations
type BaseImageClient interface {
	Add(i interface{}) error
}

type defaultBaseImageClient struct {
	client *imageclient.ImageManager
}

func (c *defaultBaseImageClient) Add(i interface{}) error {
	baseImageModel := i.(*v1.BaseImage)

	params := &baseimage.AddBaseImageParams{
		Body:    baseImageModel,
		Context: context.Background(),
	}

	created, err := c.client.BaseImage.AddBaseImage(params, GetAuthInfoWriter())
	if err != nil {
		return formatAPIError(err, params)
	}
	*baseImageModel = *created.Payload
	return nil
}

// NewBaseImageClient is the constructor for a defaultBaseImageClient
func NewBaseImageClient() *defaultBaseImageClient {
	return &defaultBaseImageClient{
		client: imageManagerClient(),
	}
}

func createBaseImage(out, errOut io.Writer, cmd *cobra.Command, args []string, ic BaseImageClient) error {
	baseImage := &v1.BaseImage{
		Name:      &args[0],
		DockerURL: &args[1],
		Language:  swag.String(language),
	}
	err := ic.Add(baseImage)
	if err != nil {
		return err
	}
	if dispatchConfig.JSON {
		encoder := json.NewEncoder(out)
		encoder.SetIndent("", "    ")
		return encoder.Encode(*baseImage)
	}
	fmt.Fprintf(out, "Created base image: %s\n", *baseImage.Name)
	return nil
}
