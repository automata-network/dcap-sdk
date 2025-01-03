package sp1

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/automata-network/dcap-sdk/packages/godcap/bincode"
	"github.com/chzyer/logex"
)

// GetNonceRequest represents a request to get the nonce for an account.
type GetNonceRequest struct {
	Address []byte `json:"address"`
}

// GetNonceResponse represents the response containing the nonce.
type GetNonceResponse struct {
	Nonce uint64 `json:"nonce,string"`
}

// RpcGetNonce retrieves the nonce for the client's public address.
func (c *Client) RpcGetNonce(ctx context.Context) (uint64, error) {
	addr := c.Public()
	var res GetNonceResponse
	if err := c.api("GetNonce", &GetNonceRequest{Address: addr[:]}, &res); err != nil {
		return 0, logex.Trace(err)
	}
	return res.Nonce, nil
}

// CreateProofRequest represents a request to create a proof.
type CreateProofRequest struct {
	/// The signature of the message.
	Signature []byte `json:"signature"`
	/// The nonce for the account.
	Nonce uint64 `json:"nonce,string"`
	/// The mode for proof generation.
	Mode uint32 `json:"mode"`
	/// The deadline for the proof request, signifying the latest time a fulfillment would be valid.
	Deadline uint64 `json:"deadline,string"`
	/// The SP1 circuit version to use for the proof.
	CircuitVersion string `json:"circuit_version"`
}

// CreateProofResponse represents the response containing proof details.
type CreateProofResponse struct {
	/// The proof identifier.
	ProofId string `json:"proof_id"`
	/// The URL to upload the ELF file.
	ProgramUrl string `json:"program_url"`
	/// The URL to upload the standard input (stdin).
	StdinUrl string `json:"stdin_url"`
}

// RpcCreateProof creates a proof with the given parameters.
func (c *Client) RpcCreateProof(ctx context.Context, nonce uint64, deadline uint64, mode ProofMode) (*CreateProofResponse, error) {
	sig, err := c.auth.SignMessage(&CreateProofMsg{
		Nonce:    nonce,
		Deadline: deadline,
		Mode:     uint32(mode),
		Version:  c.cfg.Version,
	})

	if err != nil {
		return nil, logex.Trace(err)
	}
	var res CreateProofResponse
	if err := c.api("CreateProof", &CreateProofRequest{
		Signature:      sig,
		Nonce:          nonce,
		Deadline:       uint64(deadline),
		Mode:           uint32(mode),
		CircuitVersion: c.cfg.Version,
	}, &res); err != nil {
		return nil, logex.Trace(err)
	}
	return &res, nil
}

// SubmitProofRequest represents a request to submit a proof.
type SubmitProofRequest struct {
	/// The signature of the message.
	Signature []byte `json:"signature"`
	/// The nonce for the account.
	Nonce uint64 `json:"nonce,string"`
	/// The proof identifier.
	ProofId string `json:"proof_id"`
}

// SubmitProofResponse represents the response for submitting a proof, empty on success.
type SubmitProofResponse struct{}

// RpcSubmitProof submits the proof with the given nonce and proof ID.
func (c *Client) RpcSubmitProof(ctx context.Context, nonce uint64, proofId string) error {
	var submitRes SubmitProofResponse
	sig, err := c.auth.SignMessage(&SubmitProofMsg{Nonce: nonce, ProofId: proofId})
	if err != nil {
		return logex.Trace(err)
	}
	if err := c.api("SubmitProof", &SubmitProofRequest{
		Signature: sig,
		Nonce:     nonce,
		ProofId:   proofId,
	}, &submitRes); err != nil {
		return logex.Trace(err)
	}
	return nil
}

// Prove creates and submits a proof, then polls for the proof status.
func (c *Client) Prove(ctx context.Context, elf []byte, stdin *SP1Stdin) (*SP1ProofWithPublicValues, error) {
	proofId, err := c.CreateProof(ctx, elf, stdin, ProofModeGroth16)
	if err != nil {
		return nil, logex.Trace(err)
	}
	proof, err := c.PollProof(ctx, proofId, time.Duration(c.cfg.PollIntervalSecs)*time.Second)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return proof, nil
}

// CreateProof creates a proof and uploads the necessary files.
func (c *Client) CreateProof(ctx context.Context, elf []byte, stdin *SP1Stdin, mode ProofMode) (string, error) {
	nonce, err := c.RpcGetNonce(ctx)
	if err != nil {
		return "", logex.Trace(err)
	}
	logex.Infof("account=%v, nonce: %v", c.Public(), nonce)
	now := time.Now().Add(10 * time.Second).Unix()
	createProofRes, err := c.RpcCreateProof(ctx, nonce, uint64(now), mode)
	if err != nil {
		return "", logex.Trace(err)
	}

	elfBytes := bincode.Bytes(elf).Bincode()
	logex.Infof("upload program: %v", len(elfBytes))
	if _, err := c.s3(http.MethodPut, createProofRes.ProgramUrl, bytes.NewReader(elfBytes)); err != nil {
		return "", logex.Trace(err)
	}
	logex.Infof("upload stdin: %v", len(stdin.Bincode()))
	if _, err := c.s3(http.MethodPut, createProofRes.StdinUrl, bytes.NewReader(stdin.Bincode())); err != nil {
		return "", logex.Trace(err)
	}
	nonce, err = c.RpcGetNonce(ctx)
	if err != nil {
		return "", logex.Trace(err)
	}
	logex.Infof("account=%v, nonce: %v", c.Public(), nonce)
	if err := c.RpcSubmitProof(ctx, nonce, createProofRes.ProofId); err != nil {
		return "", logex.Trace(err)
	}
	return createProofRes.ProofId, nil
}

// GetProofStatusRequest represents a request to get the status of a proof.
type GetProofStatusRequest struct {
	ProofId string `json:"proof_id"`
}

// GetProofStatusResponse represents the response containing the proof status.
type GetProofStatusResponse struct {
	/// The status of the proof request.
	Status string `json:"status"`
	/// Optional proof URL, where you can download the result of the proof request. Only included if
	/// the proof has been fulfilled.
	ProofUrl string `json:"proof_url"`
	/// If the proof was unclaimed, the reason why.
	UnclaimReason string `json:"unclaim_reason"`
	/// If the proof was unclaimed, the description detailing why.
	UnclaimDescription string `json:"unclaim_description"`
}

// RpcGetProofStatus retrieves the status of the proof with the given proof ID.
func (c *Client) RpcGetProofStatus(ctx context.Context, proofId string) (*GetProofStatusResponse, error) {
	var res GetProofStatusResponse
	if err := c.api("GetProofStatus", &GetProofStatusRequest{ProofId: proofId}, &res); err != nil {
		return nil, logex.Trace(err)
	}
	return &res, nil
}

// SP1ProofWithPublicValues represents a proof along with its public values.
type SP1ProofWithPublicValues struct {
	Proof        SP1Proof
	Stdin        SP1Stdin
	PublicValues SP1PublicValues
	Sp1Version   bincode.String
}

// Bytes serializes the proof with public values into bytes.
func (p *SP1ProofWithPublicValues) Bytes() ([]byte, error) {
	switch p.Proof.Type.Raw() {
	case 3: // Groth16
		proof := p.Proof.Groth16
		bytes := make([]byte, 0, 4+len(proof.EncodedProof))
		bytes = append(bytes, proof.Groth16VkeyHash[:4]...)
		decodedProof, err := hex.DecodeString(string(proof.EncodedProof))
		if err != nil {
			return nil, logex.Trace(err)
		}
		bytes = append(bytes, decodedProof...)
		return bytes, nil
	default:
		return nil, logex.NewErrorf("unsupported proof mode: %v", p.Proof.Type)
	}
}

// New creates a new instance of SP1ProofWithPublicValues.
func (p *SP1ProofWithPublicValues) New() bincode.FromBin {
	return new(SP1ProofWithPublicValues)
}

// String returns a string representation of SP1ProofWithPublicValues.
func (p *SP1ProofWithPublicValues) String() string {
	return fmt.Sprintf("SP1ProofWithPublicValues{proof: %v, stdin: %v, public_values: %v, sp1_version: %v}", p.Proof.String(), p.Stdin.String(), p.PublicValues.String(), p.Sp1Version.String())
}

// FromBin deserializes the proof with public values from bytes.
func (p *SP1ProofWithPublicValues) FromBin(data []byte) ([]byte, error) {
	var err error
	data, err = p.Proof.FromBin(data)
	if err != nil {
		return nil, logex.Trace(err)
	}
	data, err = p.Stdin.FromBin(data)
	if err != nil {
		return nil, logex.Trace(err)
	}
	data, err = p.PublicValues.FromBin(data)
	if err != nil {
		return nil, logex.Trace(err)
	}
	data, err = p.Sp1Version.FromBin(data)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return data, nil
}

// SP1Proof represents a proof with its type and specific proof data.
type SP1Proof struct {
	Type    bincode.U32
	Groth16 *Groth16Bn254Proof
}

// New creates a new instance of SP1Proof.
func (p *SP1Proof) New() bincode.FromBin {
	return new(SP1Proof)
}

// String returns a string representation of SP1Proof.
func (p *SP1Proof) String() string {
	if uint32(p.Type) == 3 {
		return fmt.Sprintf("SP1Proof:Groth16(%v)", p.Groth16.String())
	} else {
		return "unknown SP1Proof"
	}
}

// FromBin deserializes the proof from bytes.
func (p *SP1Proof) FromBin(data []byte) ([]byte, error) {
	var err error
	data, err = p.Type.FromBin(data)
	if err != nil {
		return nil, logex.Trace(err)
	}
	switch p.Type.Raw() {
	case 3:
		p.Groth16 = p.Groth16.New().(*Groth16Bn254Proof)
		data, err = p.Groth16.FromBin(data)
		if err != nil {
			return nil, logex.Trace(err)
		}
	default:
		return nil, bincode.ErrUnexpectEnum.Format(p, p.Type)
	}
	return data, nil
}

// Groth16Bn254Proof represents a Groth16 proof with specific data.
type Groth16Bn254Proof struct {
	PublicInputs    [2]bincode.String
	EncodedProof    bincode.String
	RawProof        bincode.String
	Groth16VkeyHash bincode.Bytes32
}

// New creates a new instance of Groth16Bn254Proof.
func (p *Groth16Bn254Proof) New() bincode.FromBin {
	return new(Groth16Bn254Proof)
}

// String returns a string representation of Groth16Bn254Proof.
func (p *Groth16Bn254Proof) String() string {
	return fmt.Sprintf("Groth16Bn254Proof{public_inputs: %v, encoded_proof: %v, raw_proof: %v, groth16_vkey_hash: %v}", p.PublicInputs, p.EncodedProof, p.RawProof, p.Groth16VkeyHash)
}

// FromBin deserializes the Groth16 proof from bytes.
func (p *Groth16Bn254Proof) FromBin(data []byte) ([]byte, error) {
	return bincode.UnmarshalFields(data, []bincode.FromBin{&p.PublicInputs[0], &p.PublicInputs[1], &p.EncodedProof, &p.RawProof, &p.Groth16VkeyHash})
}

// Buffer represents a buffer with binary data.
type Buffer struct {
	Data bincode.Bytes
}

// String returns a string representation of Buffer.
func (b *Buffer) String() string {
	return fmt.Sprintf("Buffer{data: %v}", b.Data.String())
}

// FromBin deserializes the buffer from bytes.
func (b *Buffer) FromBin(data []byte) ([]byte, error) {
	return bincode.UnmarshalFields(data, []bincode.FromBin{&b.Data})
}

// SP1PublicValues represents public values associated with a proof.
type SP1PublicValues struct {
	Buffer Buffer
}

// New creates a new instance of SP1PublicValues.
func (v *SP1PublicValues) New() bincode.FromBin {
	return new(SP1PublicValues)
}

// String returns a string representation of SP1PublicValues.
func (v *SP1PublicValues) String() string {
	return fmt.Sprintf("SP1PublicValues{buffer: %v}", v.Buffer.String())
}

// FromBin deserializes the public values from bytes.
func (v *SP1PublicValues) FromBin(data []byte) ([]byte, error) {
	return v.Buffer.FromBin(data)
}

// PollProof polls the status of the proof until it is fulfilled or an error occurs.
func (c *Client) PollProof(ctx context.Context, proofId string, interval time.Duration) (*SP1ProofWithPublicValues, error) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	errRetryTime := 3
	isClaimed := false
	for {
		select {
		case <-ctx.Done():
			return nil, logex.Trace(ctx.Err())
		case <-ticker.C:
			status, err := c.RpcGetProofStatus(ctx, proofId)
			if err != nil {
				if errRetryTime == 0 {
					return nil, logex.Trace(err)
				}
				logex.Error(err)
				errRetryTime--
				continue
			}
			switch status.Status {
			case "PROOF_FULFILLED":
				if status.ProofUrl == "" {
					return nil, logex.NewErrorf("missing receipt: %v", status)
				}
				proofBytes, err := c.s3(http.MethodGet, status.ProofUrl, nil)
				if err != nil {
					return nil, logex.Trace(err)
				}
				res, err := bincode.Unmarshal[*SP1ProofWithPublicValues](proofBytes)
				if err != nil {
					return nil, logex.Trace(err)
				}
				return res, nil
			case "PROOF_CLAIMED":
				if !isClaimed {
					logex.Info("Proof request claimed, proving...")
					isClaimed = true
				}
			case "PROOF_UNCLAIMED":
				return nil, logex.NewErrorf(
					"Proof generation failed: [%v] %v",
					status.UnclaimReason,
					status.UnclaimDescription,
				)
			default:
				logex.Infof("Session %v is running: %v", proofId, status.Status)
			}
		}
	}
}

// SP1Stdin represents the standard input for SP1.
type SP1Stdin struct {
	Buffer bincode.Collection[*bincode.Bytes]
	Ptr    bincode.U64
	Proofs bincode.Collection[*bincode.U32]
}

// NewSP1StdinFromInput creates a new SP1Stdin from the given input bytes.
func NewSP1StdinFromInput(input []byte) *SP1Stdin {
	return &SP1Stdin{
		Buffer: bincode.Collection[*bincode.Bytes]([]*bincode.Bytes{(*bincode.Bytes)(&input)}),
	}
}

// New creates a new instance of SP1Stdin.
func (s *SP1Stdin) New() bincode.FromBin {
	return new(SP1Stdin)
}

// String returns a string representation of SP1Stdin.
func (s *SP1Stdin) String() string {
	return fmt.Sprintf("SP1Stdin(%v)", len(s.Buffer))
}

// FromBin deserializes the standard input from bytes.
func (s *SP1Stdin) FromBin(data []byte) ([]byte, error) {
	return bincode.UnmarshalFields(data, []bincode.FromBin{
		&s.Buffer, &s.Ptr, &s.Proofs,
	})
}

// Bincode serializes the standard input into bytes.
func (s *SP1Stdin) Bincode() []byte {
	var buf []byte
	var led = binary.LittleEndian
	// collection
	buf = led.AppendUint64(buf, uint64(len(s.Buffer)))
	for _, buffer := range s.Buffer {
		buf = append(buf, buffer.Bincode()...)
	}
	// ptr
	buf = led.AppendUint64(buf, 0)
	// proofs
	buf = led.AppendUint64(buf, 0)
	return buf
}
