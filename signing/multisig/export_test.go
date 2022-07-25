package multisig

import crypto "github.com/ElrondNetwork/elrond-go-crypto"

// ConvertBytesToPubKeys -
func ConvertBytesToPubKeys(pubKeys [][]byte, kg crypto.KeyGenerator) ([]crypto.PublicKey, error) {
	return convertBytesToPubKeys(pubKeys, kg)
}

// ConvertBytesToPubKey -
func ConvertBytesToPubKey(pubKeyBytes []byte, kg crypto.KeyGenerator) (crypto.PublicKey, error) {
	return convertBytesToPubKey(pubKeyBytes, kg)
}

// ConvertBytesToPrivateKey -
func ConvertBytesToPrivateKey(privateKey []byte, kg crypto.KeyGenerator) (crypto.PrivateKey, error) {
	return convertBytesToPrivateKey(privateKey, kg)
}
