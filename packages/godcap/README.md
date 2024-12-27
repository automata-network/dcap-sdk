# GoDcap

Go SDK for interactive with [Automata DCAP attestation](http://github.com/automata-network/automata-dcap-attestation)


# Examples

## Verify on chain
```go
func VerifyOnChain(ctx context.Context, quote []byte, privateKeyStr string) error {
    privateKey, err := crypto.HexToECDSA(privateKeyStr)
    if err != nil {
        return err
    }

    portal, err := godcap.NewDcapPortal(ctx, RPC_URL)
    if err != nil {
        return err
    }

    // setup a callback function when the verification success
    //  function setNumber(uint256 newNumber) public fromDcapPortal
    callback := NewCallbackFromAbiJSON(VerifiedCounter.VerifiedCounterABI)
        .WithParams("setNumber", big.NewInt(10))
        .WithTo(verifiedCounterAddr)

    opts, err := portal.BuildTransactOpts(ctx, privateKey)
    if err != nil {
        return err
    }

    tx, err := portal.VerifyOnChain(opts, quote, callback)
    if err != nil {
        return err
    }

    // waiting 
    receipt := <-portal.WaitTx(ctx, tx)
    fmt.Printf("%#v\n", receipt)
}
```

## Verify with ZkProof (Risc0 and Succinct)
```go
func VerifyOnRisc0(ctx context.Context, quote []byte, privateKeyStr string) error {
    privateKey, err := crypto.HexToECDSA(privateKeyStr)
    if err != nil {
        return err
    }

    portal, err := godcap.NewDcapPortal(ctx, RPC_URL)
    if err != nil {
        return err
    }
    if err := portal.EnableZkProof(&zkdcap.ZkProofConfig{
        Bonsai: &bonsai.Config {
            ApiKey: $BONSAI_API_KEY,
        },
        Sp1: &sp1.Config {
            PrivateKey: $SP1_PRIVATE_KEY,
        },
    }); err != nil {
        return err
    }

    zkType := zkdcap.ZkTypeRiscZero // or zkdcap.ZkTypeSuccinct
    zkproof, err := portal.GenerateZkProof(ctx, zkdcap.ZkTypeRiscZero, quote)
    if err != nil {
        return err
    }

    // setup a callback function when the verification success
    //  function setNumber(uint256 newNumber) public fromDcapPortal
    callback := NewCallbackFromAbiJSON(VerifiedCounter.VerifiedCounterABI)
        .WithParams("setNumber", big.NewInt(10))
        .WithTo(verifiedCounterAddr)

    opts, err := portal.BuildTransactOpts(ctx, privateKey)
    if err != nil {
        return err
    }

    tx, err := portal.VerifyAndAttestWithZKProof(opts, zkproof, callback)
    if err != nil {
        return err
    }

    // waiting 
    receipt := <-portal.WaitTx(ctx, tx)
    fmt.Printf("%#v\n", receipt)
}
```
