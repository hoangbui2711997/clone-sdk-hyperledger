package usercontroller

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/util/pathvar"
	"github.com/sirupsen/logrus"
)

var (
	myConfigPath          = "/Users/jeff/go/src/fabric-sdk-go/pkg/core/config/testdata/config_local_test.yaml"
	entityMatcherOverride = "/Users/jeff/go/src/fabric-sdk-go/test/fixtures/config/overrides/local_entity_matchers.yaml"
)

func getContext(channelID, admin, org string) context.ChannelProvider {
	configProvider := config.FromFile(pathvar.Subst(myConfigPath))
	backendConfig := func() ([]core.ConfigBackend, error) {
		return appendLocalEntityMappingBackend(configProvider, entityMatcherOverride)
	}

	sdk, err := fabsdk.New(backendConfig)
	if err != nil {
		logrus.Errorln(sdk, err)
		return nil
	}

	//prepare contexts
	org1AdminChannelContext := sdk.ChannelContext(channelID, fabsdk.WithUser(admin), fabsdk.WithOrg(org))

	return org1AdminChannelContext
}

func appendLocalEntityMappingBackend(configProvider core.ConfigProvider, entityMatcherOverridePath string) ([]core.ConfigBackend, error) {
	currentBackends, err := extractBackend(configProvider)
	if err != nil {
		return nil, err
	}

	//Entity matcher config backend
	configProvider = config.FromFile(pathvar.Subst(entityMatcherOverridePath))
	matcherBackends, err := configProvider()
	if err != nil {
		return nil, err
	}

	//backends should fal back in this order - matcherBackends, localBackends, currentBackends
	localBackends := append([]core.ConfigBackend{}, matcherBackends...)
	localBackends = append(localBackends, currentBackends...)

	return localBackends, nil
}

func extractBackend(configProvider core.ConfigProvider) ([]core.ConfigBackend, error) {
	if configProvider == nil {
		return []core.ConfigBackend{}, nil
	}
	return configProvider()
}

func GetClientIDAccErc20(channelName, userName, orgName string) string {
	c, err := channel.New(getContext(channelName, userName, orgName))
	if err != nil {
		logrus.Errorln(err)
		return ""
	}
	//logrus.Infoln(c, err)

	resp, err := c.Query(channel.Request{
		ChaincodeID: "token_erc20",
		//Fcn:         "Initialize",
		//Fcn: "MintWithTokenURI",
		Fcn:  "ClientAccountID",
		Args: [][]byte{
			//[]byte("102"),
			//[]byte("https://example.com/nft101.json"),
		},
		TransientMap:    nil,
		InvocationChain: nil,
		IsInit:          false,
	})

	if err != nil {
		logrus.Errorln(err)
		return ""
	}

	return string(resp.Payload)
}

func GetInfoToken(c *channel.Client, tokenID string) (channel.Response, error) {
	return c.Query(channel.Request{
		ChaincodeID: "token_erc721",
		Fcn:         "TokenURI",
		Args: [][]byte{
			[]byte(tokenID),
		},
	})
}

func GetClientIDAcc(channelName, userName, orgName string) string {
	c, err := channel.New(getContext(channelName, userName, orgName))
	if err != nil {
		logrus.Errorln(err)
		return ""
	}

	resp, err := c.Query(channel.Request{
		ChaincodeID: "token_erc721",
		Fcn:         "ClientAccountID",
	})

	if err != nil {
		logrus.Errorln(err)
		return ""
	}

	return string(resp.Payload)
}

func GetOwnerOfToken(c *channel.Client, tokenID string) (channel.Response, error) {
	return c.Query(channel.Request{
		ChaincodeID: "token_erc721",
		Fcn:         "OwnerOf",
		Args: [][]byte{
			[]byte(tokenID),
		},
	})
}
