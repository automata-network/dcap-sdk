package sp1

import (
	_ "embed"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/automata-network/dcap-sdk/packages/godcap/bincode"
	"github.com/chzyer/logex"
)

//go:embed test_proof.hex
var testProof string

func TestProof(t *testing.T) {
	proofBytes, err := hex.DecodeString(testProof)
	if err != nil {
		t.Fatal(err)
	}
	proof, err := bincode.Unmarshal[*SP1ProofWithPublicValues](proofBytes)
	if err != nil {
		logex.Error(err)
		t.Fatal(err)
	}
	fmt.Println(proof)
}
