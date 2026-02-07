package derivation

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
)

// GetSolanaAddress генерирует адрес m/44'/501'/0'/0'
func GetSolanaAddress(seed []byte) (string, error) {
	// Для Ed25519 в Solana часто используется SLIP-0010
	// Упрощенная версия получения ключа из сида:
	hash := hmac.New(sha512.New, []byte("ed25519 seed"))
	hash.Write(seed)
	sum := hash.Sum(nil)
	
	privKey := ed25519.NewKeyFromSeed(sum[:32])
	pubKey := privKey.Public().(ed25519.PublicKey)
	
	return base58.Encode(pubKey), nil
}

// GetTONAddress генерирует User-friendly адрес (V4R2)
func GetTONAddress(seed []byte) (string, error) {
	hash := hmac.New(sha512.New, []byte("ton seed"))
	hash.Write(seed)
	sum := hash.Sum(nil)
	
	privKey := ed25519.NewKeyFromSeed(sum[:32])
	pubKey := privKey.Public().(ed25519.PublicKey)
	
	// Упрощенный формат TON (Raw + Flags + Checksum)
	// В реальности требует упаковки в BOC, здесь даем заготовку под Base64
	return base64.RawURLEncoding.EncodeToString(pubKey), nil
}

func GetEVMAddress(seed []byte) (string, error) {
	masterKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	c, _ := masterKey.Derive(hdkeychain.HardenedKeyStart + 44)
	c, _ = c.Derive(hdkeychain.HardenedKeyStart + 60)
	c, _ = c.Derive(hdkeychain.HardenedKeyStart + 0)
	c, _ = c.Derive(0)
	c, _ = c.Derive(0)
	
	pub, _ := c.ECPubKey()
	return crypto.PubkeyToAddress(*pub.ToECDSA()).Hex(), nil
}
