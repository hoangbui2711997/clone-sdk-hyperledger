package integration

import (
	"encoding/json"
	"github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwsetutil"
	"github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/protoutil"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	mspprovide "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	mspimpl "github.com/hyperledger/fabric-sdk-go/pkg/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/util/pathvar"
	"github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

var (
	myConfigPath          = "/Users/jeff/go/src/fabric-sdk-go/pkg/core/config/testdata/config_local_test.yaml"
	entityMatcherOverride = "/Users/jeff/go/src/fabric-sdk-go/test/fixtures/config/overrides/local_entity_matchers.yaml"
)

func TestCreateID(t *testing.T) {
	configProvider := config.FromFile(pathvar.Subst(myConfigPath))
	backendConfig := func() ([]core.ConfigBackend, error) {
		return appendLocalEntityMappingBackend(configProvider, entityMatcherOverride)
	}

	sdk, _ := fabsdk.New(backendConfig)
	context := sdk.Context(fabsdk.WithOrg("Org1"), fabsdk.WithUser("admin"))
	c, _ := msp.New(context)

	identity, err := c.CreateIdentity(&msp.IdentityRequest{
		ID:          "hoangbm",
		Affiliation: "org1",
		Attributes:  nil,
		Type:        "client",
		Secret:      "hoangbmw",
		CAName:      "ca-org1",
	})

	id, err := c.GetIdentity("hoangbm")

	//c.Enroll("hoangbm", msp.WithSecret(identity.Secret))

	logrus.Infoln(id, identity, err)

	//TestGetID(t)
}

func TestCreateUserStoreInFile(t *testing.T) {
	//store, err := mspimpl.NewCertFileUserStore("/Users/jeff/go/src/fabric-sdk-go")
	store, err := mspimpl.NewCertFileUserStore("/tmp/state-store")
	if err != nil {
		logrus.Errorln(err)
	}

	pathCertO1 := "/Users/jeff/go/src/fabric-samples/test-network/organizations/fabric-ca/org1/tls-cert.pem"
	pathCertO2 := "/Users/jeff/go/src/fabric-samples/test-network/organizations/fabric-ca/org2/tls-cert.pem"

	bCert1, err := os.ReadFile(pathCertO1)
	if err != nil {
		logrus.Errorln(err)
	}
	bCert2, err := os.ReadFile(pathCertO2)
	if err != nil {
		logrus.Errorln(err)
	}

	user1 := &mspprovide.UserData{
		MSPID:                 "Org1MSP",
		ID:                    "hoangbm",
		EnrollmentCertificate: bCert1,
	}
	user2 := &mspprovide.UserData{
		MSPID:                 "Org2MSP",
		ID:                    "hoangbm",
		EnrollmentCertificate: bCert2,
	}

	_, _ = user1, user2

	if err := store.Store(user1); err != nil {
		t.Fatalf("Store %s failed [%s]", user1.ID, err)
	}
	if err := store.Store(user2); err != nil {
		t.Fatalf("Store %s failed [%s]", user2.ID, err)
	}

	//configProvider := config.FromFile(pathvar.Subst(myConfigPath))
	//backendConfig := func() ([]core.ConfigBackend, error) {
	//	return appendLocalEntityMappingBackend(configProvider, entityMatcherOverride)
	//}
	//
	//backends, err := backendConfig()
	//if err != nil {
	//	logrus.Errorln(err)
	//}
	//
	//idConfig, err := mspimpl.ConfigFromBackend(backends...)
	//if err != nil {
	//	logrus.Errorln(err)
	//}
	//idConfig.CredentialStorePath()
	//logrus.Infoln(idConfig.CredentialStorePath())
	//stateStore, err := kvs.New(&kvs.FileKeyValueStoreOptions{Path: idConfig.CredentialStorePath()})
	//userStore, err := mspimpl.NewCertFileUserStore1(stateStore)
	//
	////if err := userStore.Store(user1); err != nil {
	////	t.Fatalf("Store %s failed [%s]", user1.ID, err)
	////}
	////if err := userStore.Store(user2); err != nil {
	////	t.Fatalf("Store %s failed [%s]", user2.ID, err)
	////}
	//
	//load, err := userStore.Load(mspprovide.IdentityIdentifier{
	//	MSPID: "Org1",
	//	ID:    "user1",
	//})
	//
	//logrus.Infoln(load, err)
	//println(string(load.EnrollmentCertificate))
	//
	//_ = userStore
	//cryptosuite.ConfigFromBackend(backends...)

	//load, err := store.Load(mspprovide.IdentityIdentifier{
	//	MSPID: "Org1",
	//	ID:    "hoangbm",
	//})
	//
	//logrus.Infoln(load, err)
}

func TestRegisterClient(t *testing.T) {
	configProvider := config.FromFile(pathvar.Subst(myConfigPath))
	backendConfig := func() ([]core.ConfigBackend, error) {
		backends, err := appendLocalEntityMappingBackend(configProvider, entityMatcherOverride)
		backends = append(backends)
		return backends, err
	}

	sdk, _ := fabsdk.New(backendConfig)
	context := sdk.Context(fabsdk.WithOrg("org1"), fabsdk.WithUser("admin"))
	c, err := msp.New(context)
	logrus.Infoln(c, err)

	register, err := c.Register(&msp.RegistrationRequest{
		Name: "hoangbm",
		Type: "client",
		//MaxEnrollments: 0,
		Affiliation: "org1",
		Attributes:  nil,
		CAName:      "ca-org1",
		Secret:      "hoangbmpw",
	})

	logrus.Infoln(register, err)

	c.GetIdentity("hoangbm")
}

func TestGetID(t *testing.T) {
	configProvider := config.FromFile(pathvar.Subst(myConfigPath))
	backendConfig := func() ([]core.ConfigBackend, error) {
		backends, err := appendLocalEntityMappingBackend(configProvider, entityMatcherOverride)
		backends = append(backends)
		return backends, err
	}

	sdk, _ := fabsdk.New(backendConfig)
	context := sdk.Context(fabsdk.WithOrg("Org2"), fabsdk.WithUser("admin"))
	c, err := msp.New(context)
	logrus.Infoln(c, err)

	identity, err := c.GetIdentity("hoangbm")

	//c.Enroll("hoangbm", msp.WithSecret(identity.Secret))

	logrus.Infoln(identity, err)
}

// work
func TestQueryGetAllIDs(t *testing.T) {
	configProvider := config.FromFile(pathvar.Subst(myConfigPath))
	backendConfig := func() ([]core.ConfigBackend, error) {
		return appendLocalEntityMappingBackend(configProvider, entityMatcherOverride)
	}

	sdk, err := fabsdk.New(backendConfig)
	logrus.Infoln(sdk, err)

	//prepare contexts
	context := sdk.Context()
	c, err := msp.New(context, msp.WithOrg("org1"))
	logrus.Infoln(c, err)

	identities, err := c.GetAllIdentities()
	logrus.Infoln(identities, err)
	for _, identityResponse := range identities {
		bData, err := json.Marshal(identityResponse)
		logrus.Infoln(string(bData), err)
	}
}

// work
func TestQueryChainDemo(t *testing.T) {
	configProvider := config.FromFile(pathvar.Subst(myConfigPath))
	backendConfig := func() ([]core.ConfigBackend, error) {
		return appendLocalEntityMappingBackend(configProvider, entityMatcherOverride)
	}

	//configBackend := fetchConfigBackend(configPath, entityMatcherOverride)
	sdk, err := fabsdk.New(backendConfig)
	logrus.Infoln(sdk, err)

	org1AdminChannelContext := sdk.ChannelContext("mychannel", fabsdk.WithUser("admin"), fabsdk.WithOrg("Org1"))
	client, err := ledger.New(org1AdminChannelContext)
	if err != nil {
		t.Fatalf("Failed to create new resource management client: %s", err)
	}
	//
	block, err := client.QueryBlock(1)
	if err != nil {
		logrus.Errorln(err)
	}

	info, err := client.QueryInfo()
	logrus.Infoln(info, err)
	//client.QueryTransaction("")
	log.Infoln(block, err)

	for i := uint64(0); i < 100; i++ {
		block, err = client.QueryBlock(i)
		if err != nil {
			logrus.Errorln(err)
		}
		//log.Infoln(block, err)
		if block == nil {
			continue
		}

		for _, datum := range block.Data.Data {
			envelop := protoutil.UnmarshalEnvelopeOrPanic(datum)
			payloadActions, _ := protoutil.GetActionFromEnvelopeMsg(envelop)
			if payloadActions != nil {
				txRwSet := &rwsetutil.TxRwSet{}
				txRwSet.FromProtoBytes(payloadActions.Results)
				if len(txRwSet.NsRwSets) > 0 {
					for _, set := range txRwSet.NsRwSets {
						logrus.Infoln(set.NameSpace)
						logrus.Infoln(set.KvRwSet.String())
						//logrus.Infoln(set.CollHashedRwSets)
					}
				}
			}
		}
	}

	tx, err := client.QueryTransaction("88a23c1a6ee007db5d1c666fea6cb9d1f33fc6f91bfb9c9eb7aa324fa24f6a23")
	payloadOrPanic := protoutil.UnmarshalPayloadOrPanic(tx.TransactionEnvelope.Payload)
	channelHeader := protoutil.UnmarshalChannelHeaderOrPanic(payloadOrPanic.Header.ChannelHeader)
	logrus.Infoln(tx, channelHeader, channelHeader, err)
}

func TestQueryByTxHash(t *testing.T) {
	configProvider := config.FromFile(pathvar.Subst(myConfigPath))
	backendConfig := func() ([]core.ConfigBackend, error) {
		return appendLocalEntityMappingBackend(configProvider, entityMatcherOverride)
	}

	//configBackend := fetchConfigBackend(configPath, entityMatcherOverride)
	sdk, err := fabsdk.New(backendConfig)
	logrus.Infoln(sdk, err)

	org1AdminChannelContext := sdk.ChannelContext("mychannel", fabsdk.WithUser("hoangbm"), fabsdk.WithOrg("Org1"))
	client, err := ledger.New(org1AdminChannelContext)
	if err != nil {
		t.Fatalf("Failed to create new resource management client: %s", err)
	}

	tx, err := client.QueryTransaction("88a23c1a6ee007db5d1c666fea6cb9d1f33fc6f91bfb9c9eb7aa324fa24f6a23")
	payloadActions, _ := protoutil.GetActionFromEnvelopeMsg(tx.TransactionEnvelope)
	if payloadActions != nil {
		txRwSet := &rwsetutil.TxRwSet{}
		txRwSet.FromProtoBytes(payloadActions.Results)
		if len(txRwSet.NsRwSets) > 0 {
			for _, set := range txRwSet.NsRwSets {
				logrus.Infoln(set.NameSpace)
				logrus.Infoln(set.KvRwSet.String())
				//logrus.Infoln(set.CollHashedRwSets)
			}
		}
	}
}

func GetClientIDAcc(channelName, userName, orgName string) string {
	c, err := channel.New(getContext(channelName, userName, orgName))
	if err != nil {
		logrus.Errorln(err)
		return ""
	}
	//logrus.Infoln(c, err)

	resp, err := c.Query(channel.Request{
		ChaincodeID: "token_erc721",
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

func TestQueryClientID(t *testing.T) {
	logrus.Infoln(GetClientIDAcc("mychannel", "Admin", "org1"))
	logrus.Infoln(GetClientIDAcc("mychannel", "admin", "org2"))
}

func TestQueryClientBalance(t *testing.T) {
	c, err := channel.New(getContext("mychannel", "Admin", "org1"))
	logrus.Infoln(c, err)

	resp, err := c.Query(channel.Request{
		ChaincodeID: "token_erc721",
		//Fcn:         "Initialize",
		//Fcn: "MintWithTokenURI",
		Fcn:  "ClientAccountBalance",
		Args: [][]byte{
			//[]byte("102"),
			//[]byte("https://example.com/nft101.json"),
		},
		TransientMap:    nil,
		InvocationChain: nil,
		IsInit:          false,
	})

	logrus.Infoln(string(resp.Payload))
	//_ = resp
	//for _, resp := range resp.Responses {
	//	logrus.Infoln(
	//		resp.Response.Message,
	//		resp.Response.Status,
	//		string(resp.Response.Payload),
	//	)
	//}
}

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
