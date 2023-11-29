package abiCoder

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

// Coder is an ABI coder struct
type Coder struct {
	args abi.Arguments
}

// NewCoder is used to get new Coder instance
func NewCoder(types []string) (*Coder, error) {
	c := &Coder{}

	args := []abi.Argument{}

	for _, s := range types {
		t, err := abi.NewType(s, "", nil)
		if err != nil {
			return nil, err
		}

		args = append(args, abi.Argument{Type: t})
	}

	c.args = args

	return c, nil
}

// Bytes is used to get ABI encoded bytes that can be decoded by the smart-contract
func (c *Coder) Bytes(args ...interface{}) ([]byte, error) {
	return c.args.Pack(args...)
}

// KeccakHash is used to get ABI byte32 type hash from given args
func (c *Coder) KeccakHash(args ...interface{}) ([32]byte, error) {
	b, err := c.args.Pack(args...)
	if err != nil {
		return [32]byte{}, err
	}

	hash := crypto.Keccak256(b)
	hash32 := *abi.ConvertType(hash, new([32]byte)).(*[32]byte)

	return hash32, nil
}
