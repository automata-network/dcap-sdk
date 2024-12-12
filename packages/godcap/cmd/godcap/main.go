package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/automata-network/dcap-sdk/packages/godcap"
	"github.com/automata-network/dcap-sdk/packages/godcap/bonsai"
	"github.com/automata-network/dcap-sdk/packages/godcap/zkdcap"
	"github.com/chzyer/flagly"
	"github.com/chzyer/logex"
)

type GoDcap struct {
	Config *GoDcapConfig `flagly:"handler"`
	Bonsai *GoDcapBonsai `flagly:"handler"`
}

type GoDcapConfig struct {
	Contract *GoDcapConfigContract `flagly:"handler"`
}

type GoDcapConfigContract struct {
	ChainId int64 `type:"[0]"`
}

func (g *GoDcapConfigContract) FlaglyHandle() error {
	chain := godcap.ChainConfigFromChainId(g.ChainId)
	if g.ChainId == 0 {
		return flagly.ErrShowUsage
	}
	if chain == nil {
		return logex.NewErrorf("chain_id=%v not found", g.ChainId)
	}
	fmt.Println(chain.AutomataDcapAttestationFee)
	return nil
}

type GoDcapBonsai struct {
	ApiKey string
	Rpc    string
}

func (b *GoDcapBonsai) FlaglyHandle() error {
	ctx := context.Background()

	portal, err := godcap.NewDcapPortal(ctx, b.Rpc)
	if err != nil {
		return logex.Trace(err)
	}
	proof, err := portal.GenerateZkProof(ctx, zkdcap.ZkTypeRiscZero, godcap.MockQuotes[1])
	if err != nil {
		return logex.Trace(err)
	}
	ps := portal.Pccs()

	collateral, err := zkdcap.NewCollateralFromQuoteParser(ctx, parser, ps)
	if err != nil {
		return logex.Trace(err)
	}

	input := zkdcap.BonsaiGenerateInput(parser.Quote(), collateral)
	logex.Info("generated input")

	proveInfo, err := client.Prove(ctx, zkdcap.BONSAI_IMAGE_ID, input, 5*time.Second, bonsai.ReceiptGroth16)
	if err != nil {
		return logex.Trace(err)
	}

	output := proveInfo.Receipt.Journal.Bytes
	seal := bonsai.Groth16Encode(proveInfo.Receipt.Inner.Groth16.Seal)

	logex.Infof("Output: 0x%x", output)
	logex.Infof("Proof: 0x%x", seal)

	return nil
}

func main() {
	if err := flagly.RunByArgs(&GoDcap{}, os.Args); err != nil {
		logex.Fatal(err)
	}
}
