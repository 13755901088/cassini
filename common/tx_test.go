package common

import (
	"encoding/hex"
	"strings"
	"testing"

	motx "github.com/QOSGroup/cassini/mock/tx"
	"github.com/QOSGroup/cassini/types"
	"github.com/QOSGroup/qbase/example/basecoin/app"
	"github.com/QOSGroup/qbase/qcp"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmtypes "github.com/tendermint/tendermint/types"
)

func TestTransform(t *testing.T) {
	tx := motx.NewTxQcpMock("abc", "xyz", 1, 99)
	e, err := Transform(tx)

	assert.NoError(t, err)

	assert.Equal(t, tx.BlockHeight, e.Height, "!!! Transform Tx error !!!")
	assert.Equal(t, tx.TxIndex, int64(e.Index), "!!! Transform Tx error !!!")

	ca := types.CassiniEventDataTx{}
	err = ca.ConstructFromTags(e.Result.Tags)

	assert.NoError(t, err)

	assert.Equal(t, ca.From, tx.From, "!!! Transform Tx error !!!")
	assert.Equal(t, ca.To, tx.To, "!!! Transform Tx error !!!")
	assert.Equal(t, ca.Sequence, tx.Sequence, "!!! Transform Tx error !!!")
}

func TestSignTxQcp(t *testing.T) {
	tx := motx.NewTxQcpMock("abc", "xyz", 1, 99)
	pri := "a328891040aa18a9fea8baf6ad2b1502391969324258ec8562163adf4e138eb83e0dbd63b60f3521e8dfd13d533b901aaadaedf345b26d400a0fd5fd65c24f7bf66cbfef81"
	cdc := app.MakeCodec()
	err := SignTxQcp(tx, pri, cdc)

	assert.NoError(t, err)

	pub := "1624de64200f3521e8dfd13d533b901aaadaedf345b26d400a0fd5fd65c24f7bf66cbfef81"

	pubHex, err := hex.DecodeString(pub)
	assert.NoError(t, err)

	var pubkey ed25519.PubKeyEd25519
	cdc.MustUnmarshalBinaryBare(pubHex, &pubkey)

	assert.Equal(t, true, tx.Sig.Pubkey.Equals(pubkey), "!!! Sign Tx error !!!")
}

func TestQcpKey(t *testing.T) {
	assert.Equal(t, "qcp.to", qcp.QcpTo, "!!! Qcp Key changed !!!")
	assert.Equal(t, "qcp.from", qcp.QcpFrom, "!!! Qcp Key changed !!!")
	assert.Equal(t, "qcp.sequence", qcp.QcpSequence, "!!! Qcp Key changed !!!")
	assert.Equal(t, "qcp.hash", qcp.QcpHash, "!!! Qcp Key changed !!!")
}

func TestGetTxQcpHashCheck(t *testing.T) {
	tx := motx.NewTxQcpMock("abc", "xyz", 1, 99)
	event, err := Transform(tx)
	assert.NoError(t, err)

	hashStr := getHashStr(event, qcp.QcpHash)

	txo := motx.NewTxQcpMock("abc", "xyz", 1, 99)
	hashTxoStr := Bytes2HexStr(HashTx(txo))
	assert.Equal(t, hashStr, hashTxoStr)
}

func TestGetTxQcpHashCheckHeight(t *testing.T) {
	tx := motx.NewTxQcpMock("abc", "xyz", 1, 99)
	event, err := Transform(tx)
	assert.NoError(t, err)

	hashStr := getHashStr(event, qcp.QcpHash)

	txo := motx.NewTxQcpMock("abc", "xyz", 2, 99)
	hashTxoStr := Bytes2HexStr(HashTx(txo))
	assert.NotEqual(t, hashStr, hashTxoStr)
}

func TestGetTxQcpHashCheckFrom(t *testing.T) {
	tx := motx.NewTxQcpMock("abc", "xyz", 1, 99)
	event, err := Transform(tx)
	assert.NoError(t, err)

	hashStr := getHashStr(event, qcp.QcpHash)

	txo := motx.NewTxQcpMock("abcd", "xyz", 1, 99)
	hashTxoStr := Bytes2HexStr(HashTx(txo))
	assert.NotEqual(t, hashStr, hashTxoStr)
}

func TestGetTxQcpHashCheckTo(t *testing.T) {
	tx := motx.NewTxQcpMock("abc", "xyz", 1, 99)
	event, err := Transform(tx)
	assert.NoError(t, err)

	hashStr := getHashStr(event, qcp.QcpHash)

	txo := motx.NewTxQcpMock("abc", "axyz", 1, 99)
	hashTxoStr := Bytes2HexStr(HashTx(txo))
	assert.NotEqual(t, hashStr, hashTxoStr)
}

func TestGetTxQcpHashCheckSequence(t *testing.T) {
	tx := motx.NewTxQcpMock("abc", "xyz", 1, 99)
	event, err := Transform(tx)
	assert.NoError(t, err)

	hashStr := getHashStr(event, qcp.QcpHash)

	txo := motx.NewTxQcpMock("abc", "xyz", 1, 11)
	hashTxoStr := Bytes2HexStr(HashTx(txo))
	assert.NotEqual(t, hashStr, hashTxoStr)
}

func getHashStr(e *tmtypes.EventDataTx, key string) string {
	for _, kv := range e.Result.Tags {
		if strings.EqualFold(key, string(kv.Key)) {
			return Bytes2HexStr(kv.Value)
		}
	}
	return ""
}
