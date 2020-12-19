package common

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec"
	validation "github.com/go-ozzo/ozzo-validation"
)

var (
	ValidatePublicKey = validation.NewStringRule(validatePublicKey, "must be a correct public key")
)

// validatePublicKey verify its a valid secp256k1 pubkey.
func validatePublicKey(s string) bool {
	pubKeyBytes, err := hex.DecodeString(s)
	if err != nil {
		return false
	}
	_, err = btcec.ParsePubKey(pubKeyBytes, btcec.S256())
	if err != nil {
		return false
	}
	return true
}
