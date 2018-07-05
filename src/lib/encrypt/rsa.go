package encrypt

import (
	"math/big"
	"crypto/rsa"
	"crypto/rand"
	"encoding/hex"
)

func get(pubkey string) {
	int := new(big.Int)
	int.SetString(pubkey, 16)
	
	pub := rsa.PublicKey{
		N: int,
		E: 65537,
	}
	encryString := preLoginData.servertime + "\t" + preLoginData.nonce + "\n" + password
	
	encryResult, _ := rsa.EncryptPKCS1v15(rand.Reader, &pub, []byte(encryString))
	sp := hex.EncodeToString(encryResult)
}
