package epikauth

import (
	"encoding/hex"
	"encoding/json"

	_ "github.com/EpiK-Protocol/epik-auth-sdk/bls" // enable bls signatures
	"github.com/EpiK-Protocol/go-epik/chain/types"
	"github.com/EpiK-Protocol/go-epik/lib/sigs"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)

func Sign(plain []byte, epikWalletPrivateKey string) (signature []byte, addr address.Address, err error) {
	data, err := hex.DecodeString(epikWalletPrivateKey)
	if err != nil {
		return
	}
	keyInfo := &types.KeyInfo{}
	err = json.Unmarshal(data, keyInfo)
	if err != nil {
		return
	}
	t := crypto.SigTypeBLS
	switch keyInfo.Type {
	case "bls":
		t = crypto.SigTypeBLS
	case "secp256k1":
		t = crypto.SigTypeSecp256k1
	}
	pubKey, err := sigs.ToPublic(t, keyInfo.PrivateKey)
	if err != nil {
		return
	}
	addr, err = address.NewBLSAddress(pubKey)
	if err != nil {
		return
	}
	s, err := sigs.Sign(t, keyInfo.PrivateKey, plain)
	if err != nil {
		return
	}
	signature = s.Data
	return
}

func Verify(plain, signature []byte, addr address.Address) (err error) {
	sig := &crypto.Signature{
		Type: crypto.SigTypeBLS,
		Data: signature,
	}
	return sigs.Verify(sig, addr, plain)
}
