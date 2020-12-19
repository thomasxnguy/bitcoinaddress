package p2sh

import (
	"encoding/hex"
	"errors"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
	apierrors "github.com/thomasxnguy/bitcoinaddress/errors"
	"net/http"
)

// Service to generate and manage bitcoin p2sh address
type Service struct {
	Net *chaincfg.Params
}

// NewService create a new p2sh service
func NewService() *Service {
	var net *chaincfg.Params
	testnet := viper.GetBool("testnet")
	if testnet {
		net = &chaincfg.TestNet3Params
	} else {
		net = &chaincfg.MainNetParams
	}
	return &Service{
		Net: net,
	}
}

// Endpoint to generate a p2sh address from n, m parameters and public keys
func (rs *Service) generateP2SHAddress(w http.ResponseWriter, r *http.Request) {
	body := &P2shRequest{}
	if err := render.Bind(r, body); err != nil {
		switch err.(type) {
		case validation.Errors:
			render.Render(w, r, apierrors.ErrValidation(errors.New("p2sh validation errors"), err.(validation.Errors)))
			return
		}
		render.Render(w, r, apierrors.ErrInvalidRequest(err))
		return
	}

	// Check that m >= n in a n-of-m scheme
	if body.Req > len(body.PublicKeys) {
		render.Render(
			w,
			r,
			apierrors.ErrInvalidRequest(
				errors.New("Req needs to be less or equal than number of public key provided")))
		return
	}

	// 1. Create an AddressPubKey from the public key
	var addressPubKeys []*btcutil.AddressPubKey

	for _, s := range body.PublicKeys {
		pubKeyBytes, err := hex.DecodeString(s)
		if err != nil {
			render.Render(w, r, apierrors.ErrInternalError(err))
			return
		}
		addressPubKey, err := btcutil.NewAddressPubKey(pubKeyBytes, rs.Net)
		if err != nil {
			render.Render(w, r, apierrors.ErrInternalError(err))
			return
		}
		addressPubKeys = append(addressPubKeys, addressPubKey)
	}

	// 2. Create a redeem Script
	redeemScript, err := txscript.MultiSigScript(addressPubKeys, body.Req)
	if err != nil {
		render.Render(w, r, apierrors.ErrInternalError(err))
		return
	}
	// 3. Create p2sh adddress from redeem Script
	p2shAddress, err := btcutil.NewAddressScriptHash(redeemScript, rs.Net)
	if err != nil {
		render.Render(w, r, apierrors.ErrInternalError(err))
		return
	}

	render.Respond(w, r, newP2shResponse(p2shAddress.EncodeAddress()))
}
