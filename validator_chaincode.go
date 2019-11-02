package chaincode

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	mw "github.com/olegabu/go-mimblewimble"
)

var logger = shim.NewLogger("ValidatorChaincode")

// Chaincode is the definition of the chaincode structure.
type ValidatorChaincode struct {
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *ValidatorChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	function, params := stub.GetFunctionAndParameters()
	logger.Info(function, params)
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *ValidatorChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, params := stub.GetFunctionAndParameters()
	logger.Info(function, params)

	txBytes := []byte(params[0])

	if function == "validate" {
		err := mw.ValidateSignature(txBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		//err = mw.ValidateCommitmentsSum(txBytes)
		//if err != nil {
		//	return shim.Error(err.Error())
		//}

		return shim.Success(nil)
	}

	return shim.Error("unknown function")
}
