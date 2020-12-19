package common

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/spf13/viper"
	"github.com/tyler-smith/go-bip39"
)

// KeyManager is in charge of generating the addresses in the application.
type KeyManager struct {
	// The master key are generate using a mnemonic sentences following BIP39.
	// We use two master key, one for BIP49 (segwit) and the other for BIP84 (native segwit).
	BIP49MasterKey *hdkeychain.ExtendedKey
	BIP84MasterKey *hdkeychain.ExtendedKey

	Net *chaincfg.Params
}

// NewKeyManager returns an KeyManager.
func NewKeyManager() *KeyManager {
	mnemonic := viper.GetString("mnemonic")
	password := viper.GetString("password")

	var net *chaincfg.Params
	var coinType uint32
	testnet := viper.GetBool("testnet")
	if testnet {
		net = &chaincfg.TestNet3Params
		coinType = 1
	} else {
		net = &chaincfg.MainNetParams
		coinType = 0
	}

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password.
	seed := bip39.NewSeed(mnemonic, password)

	// Generate a new master node using the seed.
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		panic(err)
	}

	// BIP49 refers to the accepted common standard of deriving segwit "compatibility" addresses.
	// These addresses begin with a 3.
	// Start path m/49' (49 for BIP49).
	bip49purpose, err := masterKey.Child(49 + hdkeychain.HardenedKeyStart)
	if err != nil {
		panic(err)
	}
	// m/49'/coin_type'.
	bip49master, err := bip49purpose.Child(coinType + hdkeychain.HardenedKeyStart)
	if err != nil {
		panic(err)
	}

	//BIP84 refers to the accepted common standard of deriving native segwit addresses.
	//These addresses always begin with bc1 - and are referred to bech32 addresses.
	// m/84'  (84 for BIP84)
	bip84purpose, err := masterKey.Child(84 + hdkeychain.HardenedKeyStart)
	if err != nil {
		panic(err)
	}
	// m/84'/coin_type'
	bip84master, err := bip84purpose.Child(coinType + hdkeychain.HardenedKeyStart)
	if err != nil {
		panic(err)
	}

	// return new manager
	return &KeyManager{
		BIP49MasterKey: bip49master,
		BIP84MasterKey: bip84master,
		Net:            net,
	}
}

func (km *KeyManager) GetSegWitAddressForAccountAt(index uint32) (string, error) {
	// m/49'/coin_type'/index'
	account, err := km.BIP49MasterKey.Child(index + hdkeychain.HardenedKeyStart)
	if err != nil {
		return "", err
	}

	// m/49'/coin_type'/index'/0 (first external change)
	accountChange, err := account.Child(0)
	if err != nil {
		return "", err
	}

	// m/49'/coin_type'/index'/0/0 (first external change index)
	accountChangeIndex, err := accountChange.Child(0)
	if err != nil {
		return "", err
	}

	accountExternalPub, err := accountChangeIndex.Neuter()
	if err != nil {
		return "", err
	}

	// BIP49 segwit pay-to-script-hash style address.
	pubKey, err := accountExternalPub.ECPubKey()
	if err != nil {
		return "", err
	}
	witnessProg := btcutil.Hash160(pubKey.SerializeCompressed())
	scriptSig, err := txscript.NewScriptBuilder().AddOp(txscript.OP_0).AddData(witnessProg).Script()
	if err != nil {
		return "", err
	}

	segwitAddress, err := btcutil.NewAddressScriptHash(scriptSig, km.Net)
	if err != nil {
		panic(err)
	}

	return segwitAddress.EncodeAddress(), nil
}

func (km *KeyManager) GetNativeSegWitAddressForAccountAt(index uint32) (string, error) {
	// m/84'/coin_type'/index'
	account, err := km.BIP84MasterKey.Child(index + hdkeychain.HardenedKeyStart)
	if err != nil {
		return "", err
	}

	// m/84'/coin_type'/index'/0 (first external change)
	accountChange, err := account.Child(0)
	if err != nil {
		return "", err
	}

	// m/84'/coin_type'/index'/0/0 (first external change index)
	accountChangeIndex, err := accountChange.Child(0)
	if err != nil {
		return "", err
	}

	accountExternalPub, err := accountChangeIndex.Neuter()
	if err != nil {
		return "", err
	}

	// generate a normal p2wkh address from the pubkey hash
	pubKey, err := accountExternalPub.ECPubKey()
	if err != nil {
		return "", err
	}
	witnessProg := btcutil.Hash160(pubKey.SerializeCompressed())

	nativeSegwitaddress, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, km.Net)
	if err != nil {
		return "", err
	}

	return nativeSegwitaddress.EncodeAddress(), nil
}
