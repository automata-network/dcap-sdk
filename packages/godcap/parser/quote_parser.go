package parser

import (
	"context"
	"crypto/x509"
	"encoding/asn1"
	"encoding/binary"
	"encoding/hex"
	"encoding/pem"

	"github.com/automata-network/dcap-sdk/packages/godcap/pccs"
	"github.com/chzyer/logex"
)

var led = binary.LittleEndian

var ErrInvalidPemType = logex.Define("Invalid PEM type: %v")
var OidFmpsc = asn1.ObjectIdentifier{1, 2, 840, 113741, 1, 13, 1, 4}

const SGX_TEE_TYPE = uint32(0x00000000)
const TDX_TEE_TYPE = uint32(0x00000081)
const V4_QUOTE = uint16(0x04)
const V3_QUOTE = uint16(0x03)

type QuoteParser struct {
	spec  QuoteSpec
	quote []byte
}

func NewQuoteParser(quote []byte) *QuoteParser {
	spec := DetectQuoteSpec(quote)
	return &QuoteParser{spec: spec, quote: quote}
}

func (q *QuoteParser) CertData() []byte {
	return q.quote[q.CertDataOffset():]
}

func (q *QuoteParser) Quote() []byte {
	return q.quote
}

func (q *QuoteParser) CertDataOffset() int {
	offset := q.spec.AuthDataSizeOffset()
	authDataSize := led.Uint16(q.quote[offset:])
	return offset + 2 + int(authDataSize) + 2 + 4
}

func (q *QuoteParser) PckIssuer(cert *x509.Certificate) string {
	return cert.Issuer.CommonName
}

func (q *QuoteParser) SgxExt(pck *x509.Certificate) ([]SgxExt, error) {
	for _, ext := range pck.Extensions {
		if ext.Id.Equal(asn1.ObjectIdentifier{1, 2, 840, 113741, 1, 13, 1}) {
			var exts []SgxExt
			if _, err := asn1.Unmarshal(ext.Value, &exts); err != nil {
				return nil, logex.Trace(err)
			}
			return exts, nil
		}
	}
	return nil, nil
}

func (q *QuoteParser) TcbInfo(ctx context.Context, ps *pccs.Client, fmspc string) (*pccs.TcbInfo, error) {
	tcbType := q.spec.TcbType()
	tcbVersion := q.spec.TcbVersion()

	tcbInfo, err := ps.GetTcbInfo(ctx, tcbType, fmspc, tcbVersion)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return tcbInfo, nil
}

func (q *QuoteParser) EnclaveID(ctx context.Context, ps *pccs.Client) (*pccs.EnclaveIdentityInfo, error) {
	info, err := ps.GetEnclaveID(ctx, q.spec.EnclaveIDType(), q.spec.Version())
	if err != nil {
		return nil, logex.Trace(err)
	}
	return info, nil
}

func (q *QuoteParser) Fmpsc(exts []SgxExt) string {
	for _, ext := range exts {
		if ext.OID.Equal(OidFmpsc) {
			return hex.EncodeToString(ext.Value.Bytes)
		}
	}
	return ""
}

func (q *QuoteParser) PckType(pck *x509.Certificate) (uint8, error) {
	var pckType uint8
	switch name := q.PckIssuer(pck); name {
	case "Intel SGX PCK Platform CA":
		pckType = pccs.CA_PLATFORM
	case "Intel SGX PCK Processor CA":
		pckType = pccs.CA_PROCESSOR
	default:
		return 0, logex.NewErrorf("unknown pck issuer: %v", name)
	}
	return pckType, nil
}

func (q *QuoteParser) Certificates() ([]*x509.Certificate, error) {
	certData := q.CertData()
	var certs []*x509.Certificate

parseCert:
	pemBlock, certData := pem.Decode(certData)
	if pemBlock != nil {
		if pemBlock.Type != "CERTIFICATE" {
			return nil, ErrInvalidPemType.Format(pemBlock.Type)
		}
		cert, err := x509.ParseCertificate(pemBlock.Bytes)
		if err != nil {
			return nil, logex.Trace(err)
		}
		certs = append(certs, cert)
		goto parseCert
	}
	return certs, nil
}

type V3QuoteSpec struct{}

func (q *V3QuoteSpec) AuthDataSizeOffset() int {
	// 48 + 384 + 4 + 64 + 64 + 384 + 64
	return 1012
}

func (q *V3QuoteSpec) TcbType() uint8 {
	return 0
}

func (q *V4QuoteSpec) TcbType() uint8 {
	switch q.TeeType {
	case TDX_TEE_TYPE:
		return 1
	case SGX_TEE_TYPE:
		return 0
	default:
		panic("unknown teeType")
	}
}

func (q *V3QuoteSpec) TcbVersion() uint32 {
	return 2
}

func (q *V4QuoteSpec) TcbVersion() uint32 {
	return 3
}

func (q *V3QuoteSpec) Version() uint32 {
	return 3
}

func (q *V4QuoteSpec) Version() uint32 {
	return 4
}

func (q *V3QuoteSpec) EnclaveIDType() uint8 {
	return pccs.ENCLAVE_ID_QE
}

func (q *V4QuoteSpec) EnclaveIDType() uint8 {
	switch q.TeeType {
	case TDX_TEE_TYPE:
		return pccs.ENCLAVE_ID_TDQE
	case SGX_TEE_TYPE:
		return pccs.ENCLAVE_ID_QE
	default:
		panic("unknown teeType")
	}
}

type V4QuoteSpec struct {
	TeeType uint32
}

func (q *V4QuoteSpec) AuthDataSizeOffset() int {
	switch q.TeeType {
	case SGX_TEE_TYPE:
		// 48 + 384 + 4 + 64 + 64 + 2 + 4 + 384 + 64
		return 1018
	case TDX_TEE_TYPE:
		// 48 + 584 + 4 + 64 + 64 + 2 + 4 + 384 + 64
		return 1218
	default:
		panic("invalid TEE Type")
	}
}

type QuoteSpec interface {
	AuthDataSizeOffset() int
	TcbType() uint8
	TcbVersion() uint32
	EnclaveIDType() uint8
	Version() uint32
}

func DetectQuoteSpec(quote []byte) QuoteSpec {
	off := 0
	ed := binary.LittleEndian
	version := ed.Uint16(quote)
	off += 4
	teeType := ed.Uint32(quote[off:])
	if version == 3 {
		return &V3QuoteSpec{}
	} else if version == 4 {
		return &V4QuoteSpec{
			TeeType: teeType,
		}
	} else {
		panic("unexpected quote version")
	}
}

type SgxExt struct {
	OID   asn1.ObjectIdentifier
	Value asn1.RawValue
}
