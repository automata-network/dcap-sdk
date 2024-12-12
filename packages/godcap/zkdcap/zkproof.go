package zkdcap

import (
	"context"

	"github.com/automata-network/dcap-sdk/packages/godcap/bonsai"
	"github.com/automata-network/dcap-sdk/packages/godcap/pccs"
	"github.com/automata-network/dcap-sdk/packages/godcap/sp1"
	"github.com/chzyer/logex"
)

type ZkType uint8

var (
	ZkTypeRiscZero = ZkType(1)
	ZkTypeSuccinct = ZkType(2)
)

type ZkProof struct {
	Type   ZkType
	Output []byte
	Proof  []byte
}

type ZkProofConfig struct {
	Bonsai *bonsai.Config `json:"bonsai"`
	Sp1    *sp1.Config    `json:"sp1"`
}

type ZkProofClient struct {
	Bonsai *bonsai.Client
	Sp1    *sp1.Client
	ps     *pccs.Server
}

func NewZkProofClient(cfg *ZkProofConfig, ps *pccs.Server) (*ZkProofClient, error) {
	if cfg == nil {
		cfg = new(ZkProofConfig)
	}
	if cfg.Bonsai == nil {
		cfg.Bonsai = new(bonsai.Config)
	}
	if cfg.Sp1 == nil {
		cfg.Sp1 = new(sp1.Config)
	}
	if err := cfg.Bonsai.Init(); err != nil {
		return nil, logex.Trace(err)
	}
	if err := cfg.Sp1.Init(); err != nil {
		return nil, logex.Trace(err)
	}
	bonsaiClient, err := bonsai.NewClient(cfg.Bonsai)
	if err != nil {
		return nil, logex.Trace(err)
	}
	sp1Client, err := sp1.NewClient(cfg.Sp1)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return &ZkProofClient{
		Bonsai: bonsaiClient,
		Sp1:    sp1Client,
		ps:     ps,
	}, nil
}

func (c *ZkProofClient) ProveQuote(ctx context.Context, ty ZkType, quote []byte, collateral *Collateral) (*ZkProof, error) {
	proof := &ZkProof{Type: ty}
	switch ty {
	case ZkTypeRiscZero:
		input := BonsaiGenerateInput(quote, collateral)
		proveInfo, err := c.Bonsai.Prove(ctx, BONSAI_IMAGE_ID, input, bonsai.ReceiptGroth16)
		if err != nil {
			return nil, logex.Trace(err)
		}
		proof.Output = []byte(proveInfo.Receipt.Journal.Bytes)
		proof.Proof = bonsai.Groth16Encode([]byte(proveInfo.Receipt.Inner.Groth16.Seal))
	case ZkTypeSuccinct:
		stdin := sp1.NewSP1StdinFromInput(Sp1GenerateInput(quote, collateral))
		res, err := c.Sp1.Prove(ctx, SUCCINCT_ZKVM_ELF, stdin)
		if err != nil {
			return nil, logex.Trace(err)
		}
		proofBytes, err := res.Bytes()
		if err != nil {
			return nil, logex.Trace(err)
		}
		proof.Output = []byte(res.PublicValues.Buffer.Data)
		proof.Proof = proofBytes
	}
	return proof, nil
}
