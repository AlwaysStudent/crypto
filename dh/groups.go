// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package dh

import (
	"math/big"
	"sync"
)

// DH groups defined in https://www.ietf.org/rfc/rfc3526.txt
const (
	// The 2048 bit prime form 3.
	rfc3526_2048G = "02"
	rfc3526_2048P = "FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3" +
		"404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BF" +
		"B5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C6" +
		"2F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28" +
		"FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AACAA68FFFFFFFFFFFFFFFF"

	// The 3072 bit prime form 4.
	rfc3526_3072G = "02"
	rfc3526_3072P = "FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3" +
		"404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BF" +
		"B5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C6" +
		"2F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28" +
		"FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AAAC42DAD33170D04507A33A85521ABDF1" +
		"CBA64ECFB850458DBEF0A8AEA71575D060C7DB3970F85A6E1E4C7ABF5AE8CDB0933D71E8C94E04A25619DCEE3D2261AD2EE6BF12FFA0" +
		"6D98A0864D87602733EC86A64521F2B18177B200CBBE117577A615D6C770988C0BAD946E208E24FA074E5AB3143DB5BFCE0FD108E4B8" +
		"2D120A93AD2CAFFFFFFFFFFFFFFFF"

	// The 4096 bit prime form 5.
	rfc3526_4096G = "02"
	rfc3526_4096P = "FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3" +
		"404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BF" +
		"B5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C6" +
		"2F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28" +
		"FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AAAC42DAD33170D04507A33A85521ABDF1" +
		"CBA64ECFB850458DBEF0A8AEA71575D060C7DB3970F85A6E1E4C7ABF5AE8CDB0933D71E8C94E04A25619DCEE3D2261AD2EE6BF12FFA0" +
		"6D98A0864D87602733EC86A64521F2B18177B200CBBE117577A615D6C770988C0BAD946E208E24FA074E5AB3143DB5BFCE0FD108E4B8" +
		"2D120A92108011A723C12A787E6D788719A10BDBA5B2699C327186AF4E23C1A946834B6150BDA2583E9CA2AD44CE8DBBBC2DB04DE8EF" +
		"92E8EFC141FBECAA6287C59474E6BC05D99B2964FA090C3A2233BA186515BE7ED1F612970CEE2D7AFB81BDD762170481CD0069127D5B" +
		"05AA993B4EA988D8FDDC186FFB7DC90A6C08F4DF435C934063199FFFFFFFFFFFFFFFF"
)

var once sync.Once
var rfc3526_2048 *Group
var rfc3526_3072 *Group
var rfc3526_4096 *Group

func initRFC3526_2048() {
	rfc3526_2048 = &Group{}
	rfc3526_2048.G, _ = new(big.Int).SetString(rfc3526_2048G, 16)
	rfc3526_2048.P, _ = new(big.Int).SetString(rfc3526_2048P, 16)

}

func initRFC3526_3072() {
	rfc3526_3072 = &Group{}
	rfc3526_3072.G, _ = new(big.Int).SetString(rfc3526_3072G, 16)
	rfc3526_3072.P, _ = new(big.Int).SetString(rfc3526_3072P, 16)

}

func initRFC3526_4096() {
	rfc3526_4096 = &Group{}
	rfc3526_4096.G, _ = new(big.Int).SetString(rfc3526_4096G, 16)
	rfc3526_4096.P, _ = new(big.Int).SetString(rfc3526_4096P, 16)
}

// Initialize the primes and generators
func initAll() {
	initRFC3526_2048()
	initRFC3526_3072()
	initRFC3526_4096()
}

// RFC3526_2048 creates a new dh.Group consisting of the prime
// and the generator. The prime (and generator) are
// described in RFC 3526 (3.). The prime is a 2048 bit value.
func RFC3526_2048() *Group {
	once.Do(initAll)
	g := &Group{
		P: new(big.Int).Set(rfc3526_2048.P),
		G: new(big.Int).Set(rfc3526_2048.G),
	}
	return g
}

// RFC3526_3072 creates a new dh.Group consisting of the prime
// and the generator. The prime (and generator) are
// described in RFC 3526 (4.). The prime is a 3072 bit value.
func RFC3526_3072() *Group {
	once.Do(initAll)
	g := &Group{
		P: new(big.Int).Set(rfc3526_3072.P),
		G: new(big.Int).Set(rfc3526_3072.G),
	}
	return g
}

// RFC3526_4096 creates a new dh.Group consisting of the prime
// and the generator. The prime (and generator) are
// described in RFC 3526 (5.). The prime is a 4096 bit value.
func RFC3526_4096() *Group {
	once.Do(initAll)
	g := &Group{
		P: new(big.Int).Set(rfc3526_4096.P),
		G: new(big.Int).Set(rfc3526_4096.G),
	}
	return g
}
