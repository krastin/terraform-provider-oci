// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_streaming "github.com/oracle/oci-go-sdk/streaming"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_streaming.StreamAdminClient", &OracleClient{initClientFn: initStreamingStreamAdminClient})
}

func initStreamingStreamAdminClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_streaming.NewStreamAdminClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) streamAdminClient() *oci_streaming.StreamAdminClient {
	return m.GetClient("oci_streaming.StreamAdminClient").(*oci_streaming.StreamAdminClient)
}
