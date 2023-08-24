package errors

import "fmt"

var (
	// BlankRPCURLErr is used when blank rpc url is received
	BlankRPCURLErr = fmt.Errorf("blank rpc url")
	// DialPRCErr is used when error occurred while connecting to rpc provider
	DialPRCErr = fmt.Errorf("dial rpc error")
	// BlankContractAddrErr is used when blank contract address is received
	BlankContractAddrErr = fmt.Errorf("contract address is blank")
	// InvalidContractAddrErr is used when invalid contact address is received
	InvalidContractAddrErr = fmt.Errorf("contract address invalid")
	//InitContractErr is used when error occurred while contract initialization
	InitContractErr = fmt.Errorf("contract initialization error")
)

func GetDialRPCErr(err error) error {
	return fmt.Errorf("%w: %w", DialPRCErr, err)
}

func GetInitContractErr(err error) error {
	return fmt.Errorf("%w: %w", InitContractErr, err)
}
