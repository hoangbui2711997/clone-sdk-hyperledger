// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/hyperledger/fabric-sdk-go/test/integration

replace github.com/hyperledger/fabric-sdk-go => ../../

require (
	github.com/golang/protobuf v1.3.3
	github.com/hyperledger/fabric-config v0.0.5
	github.com/hyperledger/fabric-protos-go v0.0.0-20211118165945-23d738fc3553
	github.com/hyperledger/fabric-sdk-go v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.8.1
	github.com/prometheus/common v0.7.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.29.1
)

go 1.14
