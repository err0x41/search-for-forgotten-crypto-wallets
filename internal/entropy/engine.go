package entropy

import (
	"crypto/rand"
	"math/big"
)

var (
	// Порядок N для secp256k1
	N_secp, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
	// Порядок L для ed25519
	L_ed, _   = new(big.Int).SetString("1000000000000000000000000000000014def9dea2f79cd65812631a5cf5d3ed", 16)
)

func GetSmartEntropy() []byte {
	entropy := make([]byte, 16)
	rand.Read(entropy)

	entInt := new(big.Int).SetBytes(entropy)
	
	// Эвристика: Прыгаем в сектора, близкие к порядкам кривых или их производным
	switch entropy[0] % 3 {
	case 1:
		entInt.Xor(entInt, N_secp)
	case 2:
		entInt.Xor(entInt, L_ed)
	default:
		// Оставляем чистый рандом для покрытия остального пространства
	}

	res := entInt.Bytes()
	if len(res) < 16 {
		padded := make([]byte, 16)
		copy(padded[16-len(res):], res)
		return padded
	}
	return res[:16]
}
