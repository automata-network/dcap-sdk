package sp1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type ProofMode uint8

const (
	ProofModeUnspecified ProofMode = iota
	ProofModeCore
	ProofModeCompressed
	ProofModePlonk
	ProofModeGroth16
)

type ProofStatus uint32

const (
	/// Unspecified or invalid status.
	ProofUnspecifiedStatus ProofStatus = iota
	/// The proof request has been created but is awaiting the requester to submit it.
	ProofPreparing
	/// The proof request has been submitted and is awaiting a prover to claim it.
	ProofRequested
	/// The proof request has been claimed and is awaiting a prover to fulfill it.
	ProofClaimed
	/// The proof request was previously claimed but has now been unclaimed.
	ProofUnclaimed
	/// The proof request has been fulfilled and is available for download.
	ProofFulfilled
)

type Config struct {
	Rpc              string `json:"rpc"`
	PrivateKey       string `json:"private_key"`
	TimeoutSecs      int    `json:"timeout_secs"`
	PollIntervalSecs int    `json:"poll_interval_secs"`
	Version          string `json:"version"`
}

func (c *Config) Init() error {
	if c.Rpc == "" {
		c.Rpc = os.Getenv("PROVER_NETWORK_RPC")
	}
	if c.Rpc == "" {
		c.Rpc = "https://rpc.succinct.xyz/"
	}
	if c.PrivateKey == "" {
		c.PrivateKey = os.Getenv("SP1_PRIVATE_KEY")
	}
	if c.TimeoutSecs == 0 {
		c.TimeoutSecs = 30
	}
	if c.PollIntervalSecs == 0 {
		c.PollIntervalSecs = 5
	}
	if c.Version == "" {
		c.Version = "v3.0.0"
	}
	return nil
}

type Client struct {
	cfg  *Config
	auth *EIP712Auth
}

func NewClient(cfg *Config) (*Client, error) {
	if err := cfg.Init(); err != nil {
		return nil, logex.Trace(err)
	}
	key, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return nil, logex.Trace(err)
	}
	client := &Client{
		cfg:  cfg,
		auth: NewEIP712Auth(key),
	}
	return client, nil
}

func (c *Client) Public() common.Address {
	return crypto.PubkeyToAddress(c.auth.key.PublicKey)
}

func (c *Client) s3(method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, logex.Trace(err)
	}
	httpResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, logex.Trace(err)
	}
	defer httpResponse.Body.Close()
	httpBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, logex.Trace(err)
	}
	if httpResponse.StatusCode/100 != 2 {
		return nil, logex.NewErrorf("http remote error: %v", string(httpBody))
	}
	return httpBody, nil
}

func (c *Client) api(method string, req interface{}, res interface{}) error {
	data, err := json.Marshal(req)
	if err != nil {
		return logex.Trace(err)
	}
	resp, err := http.Post(fmt.Sprintf("%vnetwork.NetworkService/%v", c.cfg.Rpc, method), "application/json", bytes.NewReader(data))
	if err != nil {
		return logex.Trace(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return logex.Trace(err)
	}
	if resp.StatusCode/100 != 2 {
		return logex.NewErrorf("remote error on [%v]: %v", method, string(body))
	}
	if err := json.Unmarshal(body, res); err != nil {
		return logex.Trace(err)
	}
	return nil
}
