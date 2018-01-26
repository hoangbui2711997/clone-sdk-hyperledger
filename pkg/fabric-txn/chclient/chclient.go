/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package chclient enables channel client
package chclient

import (
	"reflect"
	"time"

	"github.com/hyperledger/fabric-sdk-go/api/apiconfig"
	fab "github.com/hyperledger/fabric-sdk-go/api/apifabclient"
	"github.com/hyperledger/fabric-sdk-go/api/apitxn"

	"github.com/hyperledger/fabric-sdk-go/api/apitxn/txnhandler"
	"github.com/hyperledger/fabric-sdk-go/pkg/errors"
	txnHandlerImpl "github.com/hyperledger/fabric-sdk-go/pkg/fabric-txn/txnhandler"
)

const (
	defaultHandlerTimeout = time.Second * 10
)

// ChannelClient enables access to a Fabric network.
type ChannelClient struct {
	client    fab.Resource
	channel   fab.Channel
	discovery fab.DiscoveryService
	selection fab.SelectionService
	eventHub  fab.EventHub
}

// NewChannelClient returns a ChannelClient instance.
func NewChannelClient(client fab.Resource, channel fab.Channel, discovery fab.DiscoveryService, selection fab.SelectionService, eventHub fab.EventHub) (*ChannelClient, error) {

	channelClient := ChannelClient{client: client, channel: channel, discovery: discovery, selection: selection, eventHub: eventHub}

	return &channelClient, nil
}

// Query chaincode using request and optional options provided
func (cc *ChannelClient) Query(request apitxn.Request, options ...apitxn.Option) ([]byte, error) {

	response := cc.InvokeHandler(txnHandlerImpl.NewQueryHandler(), request, cc.addDefaultTimeout(apiconfig.Query, options...)...)

	return response.Payload, response.Error
}

// Execute prepares and executes transaction using request and optional options provided
func (cc *ChannelClient) Execute(request apitxn.Request, options ...apitxn.Option) ([]byte, apitxn.TransactionID, error) {

	response := cc.InvokeHandler(txnHandlerImpl.NewExecuteHandler(), request, cc.addDefaultTimeout(apiconfig.Execute, options...)...)

	return response.Payload, response.TransactionID, response.Error
}

//InvokeHandler invokes handler using request and options provided
func (cc *ChannelClient) InvokeHandler(handler txnhandler.Handler, request apitxn.Request, options ...apitxn.Option) apitxn.Response {
	//TODO: this function going to be exposed through ChannelClient interface
	//Read execute tx options
	txnOpts, err := cc.prepareOptsFromOptions(options...)
	if err != nil {
		return apitxn.Response{Error: err}
	}

	//Prepare context objects for handler
	requestContext, clientContext, err := cc.prepareHandlerContexts(request, txnOpts)
	if err != nil {
		return apitxn.Response{Error: err}
	}

	//Perform action through handler
	go handler.Handle(requestContext, clientContext)

	//notifier in options will handle response if provided
	if txnOpts.Notifier != nil {
		return apitxn.Response{}
	}

	select {
	case response := <-requestContext.Opts.Notifier:
		return response
	case <-time.After(requestContext.Opts.Timeout):
		return apitxn.Response{Error: errors.New("handler timed out while performing operation")}
	}
}

//prepareHandlerContexts prepares context objects for handlers
func (cc *ChannelClient) prepareHandlerContexts(request apitxn.Request, options apitxn.Opts) (*txnhandler.RequestContext, *txnhandler.ClientContext, error) {

	if request.ChaincodeID == "" || request.Fcn == "" {
		return nil, nil, errors.New("ChaincodeID and Fcn are required")
	}

	clientContext := &txnhandler.ClientContext{
		Channel:   cc.channel,
		Selection: cc.selection,
		Discovery: cc.discovery,
		EventHub:  cc.eventHub,
	}

	requestContext := &txnhandler.RequestContext{
		Request:  request,
		Opts:     options,
		Response: apitxn.Response{},
	}

	if requestContext.Opts.Timeout == 0 {
		requestContext.Opts.Timeout = defaultHandlerTimeout
	}

	if requestContext.Opts.Notifier == nil {
		requestContext.Opts.Notifier = make(chan apitxn.Response)
	}

	return requestContext, clientContext, nil

}

//prepareOptsFromOptions Reads apitxn.Opts from apitxn.Option array
func (cc *ChannelClient) prepareOptsFromOptions(options ...apitxn.Option) (apitxn.Opts, error) {
	txnOpts := apitxn.Opts{}
	for _, option := range options {
		err := option(&txnOpts)
		if err != nil {
			return txnOpts, errors.WithMessage(err, "Failed to read opts")
		}
	}
	return txnOpts, nil
}

//addDefaultTimeout adds given default timeout if it is missing in options
func (cc *ChannelClient) addDefaultTimeout(timeOutType apiconfig.TimeoutType, options ...apitxn.Option) []apitxn.Option {
	txnOpts := apitxn.Opts{}
	for _, option := range options {
		option(&txnOpts)
	}

	if txnOpts.Timeout == 0 {
		return append(options, apitxn.WithTimeout(cc.client.Config().TimeoutOrDefault(timeOutType)))
	}
	return options
}

// Close releases channel client resources (disconnects event hub etc.)
func (cc *ChannelClient) Close() error {
	if cc.eventHub.IsConnected() == true {
		return cc.eventHub.Disconnect()
	}

	return nil
}

// RegisterChaincodeEvent registers chain code event
// @param {chan bool} channel which receives event details when the event is complete
// @returns {object} object handle that should be used to unregister
func (cc *ChannelClient) RegisterChaincodeEvent(notify chan<- *apitxn.CCEvent, chainCodeID string, eventID string) apitxn.Registration {

	// Register callback for CE
	rce := cc.eventHub.RegisterChaincodeEvent(chainCodeID, eventID, func(ce *fab.ChaincodeEvent) {
		notify <- &apitxn.CCEvent{ChaincodeID: ce.ChaincodeID, EventName: ce.EventName, TxID: ce.TxID, Payload: ce.Payload}
	})

	return rce
}

// UnregisterChaincodeEvent removes chain code event registration
func (cc *ChannelClient) UnregisterChaincodeEvent(registration apitxn.Registration) error {

	switch regType := registration.(type) {

	case *fab.ChainCodeCBE:
		cc.eventHub.UnregisterChaincodeEvent(regType)
	default:
		return errors.Errorf("Unsupported registration type: %v", reflect.TypeOf(registration))
	}

	return nil

}
