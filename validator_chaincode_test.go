package chaincode

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

func getTxBytes() []byte {
	bytes, err := ioutil.ReadFile("1g_repost_fix_kernel.json")
	if err != nil {
		log.Panic("cannot open json file with test transaction")
	}

	return bytes
}

func TestInit(t *testing.T) {
	cc := new(ValidatorChaincode)
	stub := shim.NewMockStub("chaincode", cc)
	res := stub.MockInit("1", [][]byte{[]byte("init")})
	log.Print(res)

	assert.EqualValues(t, shim.OK, res.Status)
}

func TestValidate(t *testing.T) {
	cc := new(ValidatorChaincode)
	stub := shim.NewMockStub("chaincode", cc)
	res := stub.MockInvoke("1", [][]byte{[]byte("validate"), getTxBytes()})
	log.Print(res)

	assert.EqualValues(t, shim.OK, res.Status)
}
