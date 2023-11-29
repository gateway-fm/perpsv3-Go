package rawContracts

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gateway-fm/perpsv3-Go/contracts/forwarder"
	"github.com/gateway-fm/perpsv3-Go/errors"
)

// IRawForwarderContract is a trustedMulticallForwarder interface
type IRawForwarderContract interface {
	// Aggregate3Value is used to call aggregate3Value contract method
	Aggregate3Value(arg []forwarder.TrustedMulticallForwarderCall3Value) ([]ForwarderResult, error)
	// Address is used to get contract address
	Address() common.Address
}

// Forwarder is an implementation of the trustedMulticallForwarder contract
type Forwarder struct {
	abi      *abi.ABI
	address  common.Address
	provider *ethclient.Client
	contract *bind.BoundContract
}

// ForwarderResult is a data struct for the trustedMulticallForwarder method response
type ForwarderResult struct {
	Success    bool
	ReturnData []byte
}

// NewForwarder is used to get new Forwarder instance
func NewForwarder(address common.Address, provider *ethclient.Client) (IRawForwarderContract, error) {
	c := &Forwarder{}

	abiInstance, err := getABI("./contracts/84531-TrustedMulticallForwarder.json")
	if err != nil {
		return nil, err
	}

	c.abi = abiInstance
	c.provider = provider
	c.address = address
	c.contract = bind.NewBoundContract(
		address,
		*abiInstance,
		provider,
		provider,
		provider,
	)

	return c, nil
}

func (p *Forwarder) Address() common.Address {
	return p.address
}

func (p *Forwarder) Aggregate3Value(arg []forwarder.TrustedMulticallForwarderCall3Value) ([]ForwarderResult, error) {
	out := []interface{}{}

	if err := p.rawCall(1, &out, "aggregate3Value", arg); err != nil {
		logErr("Aggregate3Value", fmt.Sprintln("err call contract method:", err.Error()))
		return nil, errors.GetReadContractErr(err, "RawForwarder", "aggregate3Value")
	}

	return *abi.ConvertType(out[0], new([]ForwarderResult)).(*[]ForwarderResult), nil
}

func (p *Forwarder) rawCall(value uint64, results *[]interface{}, method string, params ...interface{}) error {
	input, err := p.abi.Pack(method, params...)
	if err != nil {
		return err
	}

	cl := p.provider.Client()

	var hex hexutil.Bytes

	arg := map[string]interface{}{
		"from":  common.HexToAddress("0x0000000000000000000000000000000000000000"),
		"to":    &p.address,
		"input": hexutil.Bytes(input),
		"value": hexutil.EncodeUint64(value),
	}

	if err := cl.Call(&hex, "eth_call", arg, "latest"); err != nil {
		return err
	}

	res, err := p.abi.Unpack(method, hex)
	if err != nil {
		return err
	}
	*results = res

	return nil
}
