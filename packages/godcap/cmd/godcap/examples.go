package main

import (
	"context"

	"github.com/automata-network/dcap-sdk/packages/godcap"
	"github.com/automata-network/dcap-sdk/packages/godcap/mock"
	"github.com/automata-network/dcap-sdk/packages/godcap/zkdcap"
	"github.com/chzyer/logex"
)

type GoDcapExamples struct {
	CheckDcapQuoteOnChain      *GoDcapExamplesCheckDcapQuoteOnChain      `flagly:"handler" name:"check-dcap-quote-on-chain"`
	CheckDcapQuoteWithRisc0    *GoDcapExamplesCheckDcapQuoteWithRisc0    `flagly:"handler" name:"check-dcap-quote-with-risc0"`
	CheckDcapQuoteWithSuccinct *GoDcapExamplesCheckDcapQuoteWithSuccinct `flagly:"handler" name:"check-dcap-quote-with-succinct"`
}

type GoDcapExamplesCheckDcapQuoteOnChain struct {
}

func (h *GoDcapExamplesCheckDcapQuoteOnChain) FlaglyHandle() error {
	ctx := context.Background()
	portal, err := godcap.NewDcapPortal(ctx)
	if err != nil {
		return logex.Trace(err)
	}
	pass, err := portal.CheckQuote(ctx, mock.Quotes[0])
	if err != nil {
		return logex.Trace(err)
	}
	logex.Infof("verify quote pass: %v", pass)
	return nil
}

type GoDcapExamplesCheckDcapQuoteWithRisc0 struct {
}

func (h *GoDcapExamplesCheckDcapQuoteWithRisc0) FlaglyHandle() error {
	ctx := context.Background()
	portal, err := godcap.NewDcapPortal(ctx)
	if err != nil {
		return logex.Trace(err)
	}
	zkproof, err := portal.GenerateZkProof(ctx, zkdcap.ZkTypeRiscZero, mock.Quotes[0])
	if err != nil {
		return logex.Trace(err)
	}
	pass, err := portal.CheckZkProof(ctx, zkproof)
	if err != nil {
		return logex.Trace(err)
	}
	logex.Infof("verify quote pass: %v", pass)
	return nil
}

type GoDcapExamplesCheckDcapQuoteWithSuccinct struct {
}

func (h *GoDcapExamplesCheckDcapQuoteWithSuccinct) FlaglyHandle() error {
	ctx := context.Background()
	portal, err := godcap.NewDcapPortal(ctx)
	if err != nil {
		return logex.Trace(err)
	}
	zkproof, err := portal.GenerateZkProof(ctx, zkdcap.ZkTypeSuccinct, mock.Quotes[0])
	if err != nil {
		return logex.Trace(err)
	}
	pass, err := portal.CheckZkProof(ctx, zkproof)
	if err != nil {
		return logex.Trace(err)
	}
	logex.Infof("verify quote pass: %v", pass)
	return nil
}
