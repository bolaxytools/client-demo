package keystore

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/bolaxy/common"
	"github.com/bolaxy/common/hexutil"
	"github.com/bolaxy/crypto"

	"demo/types"
)


var (
	signer = types.NewEIP155Signer(big.NewInt(1))
)

func GenerateKey() ([]byte, string, error) {
	pk, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, "", err
	}
	addr := crypto.PubkeyToAddress(pk.PublicKey)

	return crypto.FromECDSA(pk), addr.String(), nil
}

func RecoverKey(raw []byte) (*ecdsa.PrivateKey, error) {
	return crypto.ToECDSA(raw)
}

func Sign(to, value, hexkey string, nonce uint64) (string, error) {
	k, err := hexutil.Decode(hexkey)
	if err != nil {
		return "", errors.New(fmt.Sprintf("decode hexkey: %s", err.Error()))
	}
	pk, err := RecoverKey(k)
	if err != nil {
		return "", errors.New(fmt.Sprintf("recover key: %s", err.Error()))
	}
	_to := common.HexToAddress(to)
	amount := new(big.Int)
	amount.SetString(value, 10)
	gasPrice := big.NewInt(1)
	// println(to, value, hexkey, nonce)

	tx := types.NewTransaction(nonce, _to, amount, 30000, gasPrice, nil)
	tx, err = types.SignTx(tx, signer, pk)
	if err != nil {
		return "", err
	}

	buffer := bytes.NewBuffer(nil)
	if err = tx.EncodeRLP(buffer); err != nil {
		return "", err
	}

	return hexutil.Encode(buffer.Bytes()), nil
}