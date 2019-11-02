package chaincode

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	cc := new(ValidatorChaincode)
	stub := shim.NewMockStub("chaincode", cc)
	res := stub.MockInit("1", [][]byte{[]byte("init")})

	assert.Equal(t, shim.OK, res.Status)
}
