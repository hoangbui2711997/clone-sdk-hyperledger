package integration

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

// tc1
// create nft with info invalid
func TestCreateNFTWithInvalidInfo(t *testing.T) {
	c, err := channel.New(getContext("mychannel", "admin", "Org1"))

	t.Run("create nft with invalid params", func(t *testing.T) {
		_, err = c.Execute(channel.Request{
			ChaincodeID: "token_erc721",
			//Fcn:         "Initialize",
			Fcn: "MintWithTokenURI",
			//Fcn: "ClientAccountBalance",
			Args: [][]byte{
				[]byte("101"),
			},
			TransientMap:    nil,
			InvocationChain: nil,
			IsInit:          false,
		})

		assert.NotNil(t, err)
	})

	t.Run("create nft with invalid without params", func(t *testing.T) {
		resp, err := c.Execute(channel.Request{
			ChaincodeID: "token_erc721",
			//Fcn:         "Initialize",
			Fcn: "MintWithTokenURI",
			//Fcn: "ClientAccountBalance",
			Args: [][]byte{},
		})

		_ = resp

		assert.NotNil(t, err)
	})
}

var (
	tokenID = strconv.FormatInt(time.Now().Unix(), 10)
)

// tc2
// create nft with info valid
func TestCreateNFTWithValidInfo(t *testing.T) {
	c, err := channel.New(getContext("mychannel", "admin", "Org1"))

	t.Run("create nft with valid params", func(t *testing.T) {
		_, err = c.Execute(channel.Request{
			ChaincodeID: "token_erc721",
			Fcn:         "MintWithTokenURI",
			Args: [][]byte{
				[]byte(tokenID),
				[]byte(fmt.Sprintf("https://example.com/nft%s.json", tokenID)),
			},
			TransientMap:    nil,
			InvocationChain: nil,
			IsInit:          false,
		})

		assert.Equal(t, nil, err)

		// get info nft
		t.Run("get info nft", func(t *testing.T) {
			resp, err := getInfoToken(c, tokenID)
			assert.Equal(t, nil, err)
			assert.NotEmpty(t, resp)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, string(resp.Payload))
		})
	})
}

// TC3
// Test create nft with id existed
func TestCreateNFTWithIDExisted(t *testing.T) {
	c, _ := channel.New(getContext("mychannel", "admin", "Org1"))
	tokenIDExist := "101"

	t.Run("test create nft with id exited", func(t *testing.T) {
		// get info nft
		t.Run("get info nft", func(t *testing.T) {
			resp, err := getInfoToken(c, tokenIDExist)
			assert.Equal(t, nil, err)
			assert.NotEmpty(t, resp)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, string(resp.Payload))

			// create nft with id existed
			t.Run("create nft with id existed", func(t *testing.T) {
				_, err = c.Execute(channel.Request{
					ChaincodeID: "token_erc721",
					Fcn:         "MintWithTokenURI",
					Args: [][]byte{
						[]byte(tokenIDExist),
						[]byte(fmt.Sprintf("https://example.com/nft%s.json", tokenIDExist)),
					},
				})

				assert.NotNil(t, err)
			})
		})
	})
}

// TC4
// Test case transfer nft
func TestTransferNFT(t *testing.T) {
	c, _ := channel.New(getContext("mychannel", "admin", "Org1"))
	TestCreateNFTWithValidInfo(t)

	t.Run("TestTransferNFTWithInvalidReceipient", func(t *testing.T) {
		// Check owner of token
		t.Run("Check owner of token", func(t *testing.T) {
			resp, err := getOwnerOfToken(c, tokenID)
			assert.Nil(t, err)
			logrus.Infoln(string(resp.Payload))
			minter := GetClientIDAcc("mychannel", "admin", "org1")
			assert.Equal(t, string(resp.Payload), minter)

			// Transfer token
			t.Run("Transfer token", func(t *testing.T) {
				minter := GetClientIDAcc("mychannel", "admin", "org1")
				recipient := GetClientIDAcc("mychannel", "admin", "org2")

				resp, err = c.Execute(channel.Request{
					ChaincodeID: "token_erc721",
					//Fcn:         "Initialize",
					Fcn: "TransferFrom",
					//Fcn: "ClientAccountBalance",
					Args: [][]byte{
						[]byte(minter),
						[]byte(recipient),
						[]byte(tokenID),
					},
				})

				_ = resp
				assert.Nil(t, err)

				resp, err = getOwnerOfToken(c, tokenID)
				assert.Nil(t, err)
				assert.Equal(t, string(resp.Payload), recipient)
			})
		})
	})
}

// TC5
// Test case transfer nft with invalid recipient
func TestTransferNFTWithInvalidRecipient(t *testing.T) {
	c, _ := channel.New(getContext("mychannel", "admin", "Org1"))
	TestCreateNFTWithValidInfo(t)

	t.Run("TestTransferNFTWithInvalidRecipient", func(t *testing.T) {
		// Check owner of token
		t.Run("Check owner of token", func(t *testing.T) {
			resp, err := getOwnerOfToken(c, tokenID)
			assert.Nil(t, err)
			logrus.Infoln(string(resp.Payload))
			minter := GetClientIDAcc("mychannel", "admin", "org1")
			assert.Equal(t, string(resp.Payload), minter)

			// Transfer token
			t.Run("Transfer token", func(t *testing.T) {
				minter := GetClientIDAcc("mychannel", "admin", "org1")
				recipient := GetClientIDAcc("mychannel", "not_exist", "org2")

				resp, err = c.Execute(channel.Request{
					ChaincodeID: "token_erc721",
					Fcn:         "TransferFrom",
					Args: [][]byte{
						[]byte(minter),
						[]byte(recipient),
						[]byte(tokenID),
					},
				})

				_ = resp
				assert.NotNil(t, err, errors.New("recipient exist"))

				resp, err = getOwnerOfToken(c, tokenID)
				assert.Nil(t, err)
				assert.Equal(t, string(resp.Payload), minter)
			})
		})
	})
}

// TC6
// Test case transfer nft not owned or not exist
func TestTransferNFTNotOwnedOrNotExist(t *testing.T) {
	c, _ := channel.New(getContext("mychannel", "admin", "Org1"))

	t.Run("TestTransferNFTNotOwned", func(t *testing.T) {
		// Check owner of token
		t.Run("Check owner of token", func(t *testing.T) {
			resp, err := getOwnerOfToken(c, tokenID)
			assert.NotNil(t, err)
			assert.Equal(t, string(resp.Payload), "")

			// Transfer token
			t.Run("Transfer token", func(t *testing.T) {
				minter := GetClientIDAcc("mychannel", "admin", "org1")
				recipient := GetClientIDAcc("mychannel", "admin", "org2")

				resp, err = c.Execute(channel.Request{
					ChaincodeID: "token_erc721",
					Fcn:         "TransferFrom",
					Args: [][]byte{
						[]byte(minter),
						[]byte(recipient),
						[]byte(tokenID),
					},
				})

				_ = resp
				assert.NotNil(t, err, errors.New("recipient exist"))

				resp, err = getOwnerOfToken(c, tokenID)
				assert.NotNil(t, err)
				assert.NotEqual(t, string(resp.Payload), recipient)
			})
		})
	})
}

// TC7
// test case delete nft with valid info
func TestDeleteNFTWithValidInfo(t *testing.T) {
	c, _ := channel.New(getContext("mychannel", "admin", "Org1"))

	// create token
	TestCreateNFTWithValidInfo(t)

	t.Run("TestDeleteNFTWithValidInfo", func(t *testing.T) {
		// Check owner of token
		t.Run("Check owner of token", func(t *testing.T) {
			resp, err := getOwnerOfToken(c, tokenID)
			assert.Nil(t, err)
			minter := GetClientIDAcc("mychannel", "admin", "org1")
			assert.Equal(t, string(resp.Payload), minter)

			// Transfer token
			t.Run("Transfer token", func(t *testing.T) {
				resp, err = c.Execute(channel.Request{
					ChaincodeID: "token_erc721",
					Fcn:         "Burn",
					Args: [][]byte{
						[]byte(tokenID),
					},
				})

				_ = resp
				assert.Nil(t, err)

				resp, err = getOwnerOfToken(c, tokenID)
				assert.NotNil(t, err)
				assert.NotEqual(t, string(resp.Payload), minter)
			})
		})
	})
}

// TC8
// burn nft but not owned
func TestBurnNFTNotOwned(t *testing.T) {
	c, _ := channel.New(getContext("mychannel", "admin", "Org1"))

	// transfer nft
	TestTransferNFT(t)

	t.Run("TestBurnNFTNotOwned", func(t *testing.T) {
		// Check owner of token
		t.Run("Check owner of token", func(t *testing.T) {
			resp, err := getOwnerOfToken(c, tokenID)
			assert.Nil(t, err)
			recipient := GetClientIDAcc("mychannel", "admin", "org2")
			assert.Equal(t, string(resp.Payload), recipient)

			// Transfer token
			t.Run("Transfer token", func(t *testing.T) {
				resp, err = c.Execute(channel.Request{
					ChaincodeID: "token_erc721",
					Fcn:         "Burn",
					Args: [][]byte{
						[]byte(tokenID),
					},
				})

				_ = resp
				assert.NotNil(t, err)
			})
		})
	})
}

// TC9
// burn nft but not exist
func TestBurnNFTNotExist(t *testing.T) {
	c, _ := channel.New(getContext("mychannel", "admin", "Org1"))

	t.Run("TestBurnNFTNotExist", func(t *testing.T) {
		// Check owner of token
		t.Run("Check owner of token", func(t *testing.T) {
			_, err := getOwnerOfToken(c, tokenID)
			assert.NotNil(t, err)

			// Transfer token
			t.Run("Transfer token", func(t *testing.T) {
				resp, err := c.Execute(channel.Request{
					ChaincodeID: "token_erc721",
					Fcn:         "Burn",
					Args: [][]byte{
						[]byte(tokenID),
					},
				})

				_ = resp
				assert.NotNil(t, err)
			})
		})
	})
}

// TC10
