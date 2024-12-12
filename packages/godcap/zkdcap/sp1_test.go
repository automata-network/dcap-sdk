package zkdcap

import (
	"context"
	"testing"
	"time"

	"github.com/automata-network/dcap-sdk/packages/godcap/dcap_parser"
	"github.com/automata-network/dcap-sdk/packages/godcap/mock"
	"github.com/automata-network/dcap-sdk/packages/godcap/sp1"
	"github.com/chzyer/logex"
)

func TestClient(t *testing.T) {
	ctx := context.Background()
	client, err := sp1.NewClient(&sp1.Config{})
	if err != nil {
		t.Fatal(err)
	}

	quote := mock.Quotes[1]

	parser := dcap_parser.NewQuoteParser(quote)
	collateral, err := NewCollateralFromQuoteParser(ctx, parser, ps)
	if err != nil {
		t.Fatal(err)
	}

	stdin := sp1.NewSP1StdinFromInput(Sp1GenerateInput(quote, collateral))

	proofId, err := client.CreateProof(ctx, SUCCINCT_ZKVM_ELF, stdin, sp1.ProofModeGroth16)
	if err != nil {
		logex.Error(err)
		t.Fatal(err)
	}
	res, err := client.PollProof(ctx, proofId, 5*time.Second)
	if err != nil {
		logex.Error(err)
		t.Fatal(err)
	}
	_ = res
}
