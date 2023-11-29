package rawContracts

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/utils/abiCoder"
)

// TODO: change to prod url
const oracleURL = "https://xc-testnet.pyth.network/api/latest_vaas?ids[]="

// IRawERC7412Contract is a ERC7412 contract interface
type IRawERC7412Contract interface {
	// GetCallFulfillOracleQuery is used to get calldata for fulfillOracleQuery method
	GetCallFulfillOracleQuery(feedID string) ([]byte, error)
	// Address is used to get contract address
	Address() common.Address
}

// ERC7412 is data struct for erc7412 contract implementation
type ERC7412 struct {
	abi      *abi.ABI
	address  common.Address
	provider *ethclient.Client
	contract *bind.BoundContract
}

// NewERC7412 is used to get new ERC7412 instance
func NewERC7412(address common.Address, provider *ethclient.Client) (IRawERC7412Contract, error) {
	c := &ERC7412{}

	abiInstance, err := getABI("./contracts/84531-ERC7412.json")
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

func (p *ERC7412) Address() common.Address {
	return p.address
}

func (p *ERC7412) GetCallFulfillOracleQuery(feedID string) ([]byte, error) {
	feedIDRaw := common.FromHex(feedID)
	feedIDRaw32 := *abi.ConvertType(feedIDRaw, new([32]byte)).(*[32]byte)

	resp, err := http.Get(fmt.Sprintf("%s%s", oracleURL, feedID))
	if err != nil {
		logErr("GetCallFulfillOracleQuery", fmt.Sprintln("err fetch oracle:", err.Error()))
		return nil, errors.GetFetchErr(err, "pyth oracle")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logErr("GetCallFulfillOracleQuery", fmt.Sprintln("err oracle response read:", err.Error()))
		return nil, errors.GetFetchErr(err, "pyth oracle")
	}

	bodyStrings := []string{}

	if err := json.Unmarshal(body, &bodyStrings); err != nil {
		logErr("GetCallFulfillOracleQuery", fmt.Sprintln("err oracle response unmarshal:", err.Error()))
		return nil, errors.GetFetchErr(err, "pyth oracle")
	}

	updateData := make([]byte, base64.StdEncoding.DecodedLen(len(bodyStrings[0])))
	_, err = base64.StdEncoding.Decode(updateData, []byte(bodyStrings[0]))
	if err != nil {
		logErr("GetCallFulfillOracleQuery", fmt.Sprintln("err oracle response decode:", err.Error()))
		return nil, errors.GetFetchErr(err, "pyth oracle")
	}

	coder, err := abiCoder.NewCoder([]string{"uint8", "uint64", "bytes32[]", "bytes[]"})
	if err != nil {
		logErr("GetCallFulfillOracleQuery", fmt.Sprintln("err pack:", err.Error()))
		return nil, errors.GetReadContractErr(err, "ERC7412", "encode call data")
	}

	callDataOracleArgHex, err := coder.Bytes(uint8(1), uint64(30), [][32]byte{feedIDRaw32}, [][]byte{updateData})
	if err != nil {
		logErr("GetCallFulfillOracleQuery", fmt.Sprintln("err pack:", err.Error()))
		return nil, errors.GetReadContractErr(err, "ERC7412", "encode call data")
	}

	callDataOracle, err := p.abi.Pack("fulfillOracleQuery", callDataOracleArgHex)
	if err != nil {
		logErr("GetCallFulfillOracleQuery", fmt.Sprintln("err pack:", err.Error()))
		return nil, errors.GetReadContractErr(err, "ERC7412", "fulfillOracleQuery")
	}

	return callDataOracle, nil
}
