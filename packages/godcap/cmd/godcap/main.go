package main

import (
	"fmt"
	"os"

	"github.com/automata-network/dcap-sdk/packages/godcap"
	"github.com/chzyer/flagly"
	"github.com/chzyer/logex"
)

type GoDcap struct {
	Config   *GoDcapConfig   `flagly:"handler"`
	Examples *GoDcapExamples `flagly:"handler"`
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

func main() {
	if err := flagly.RunByArgs(&GoDcap{}, os.Args); err != nil {
		logex.Fatal(err)
	}
}
