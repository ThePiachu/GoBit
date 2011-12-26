// Copyright 2011 The Go Authors. All rights reserved.
// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ecdsa implements the Elliptic Curve Digital Signature Algorithm, as
// defined in FIPS 186-3.
package bitecdsa

// References:
//   [NSA]: Suite B implementor's guide to FIPS 186-3,
//     http://www.nsa.gov/ia/_files/ecdsa.pdf
//   [SECG]: SECG, SEC1
//     http://www.secg.org/download/aid-780/sec1-v2.pdf

import (
	"big"
	"bitelliptic"
	"io"
	"os"
)

// PublicKey represents an ECDSA public key.
type PublicKey struct {
	*bitelliptic.BitCurve
	X, Y *big.Int
}

// PrivateKey represents a ECDSA private key.
type PrivateKey struct {
	PublicKey
	D *big.Int
}

var one = new(big.Int).SetInt64(1)

// randFieldElement returns a random element of the field underlying the given
// curve using the procedure given in [NSA] A.2.1.
func randFieldElement(c *bitelliptic.BitCurve, rand io.Reader) (k *big.Int, err os.Error) {
	b := make([]byte, c.BitSize/8+8)
	_, err = io.ReadFull(rand, b)
	if err != nil {
		return
	}

	k = new(big.Int).SetBytes(b)
	n := new(big.Int).Sub(c.N, one)
	k.Mod(k, n)
	k.Add(k, one)
	return
}

// GenerateKey generates a public&private key pair.
func GenerateKey(c *bitelliptic.BitCurve, rand io.Reader) (priv *PrivateKey, err os.Error) {
	k, err := randFieldElement(c, rand)
	if err != nil {
		return
	}

	priv = new(PrivateKey)
	priv.PublicKey.BitCurve = c
	priv.D = k
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
	return
}

// hashToInt converts a hash value to an integer. There is some disagreement
// about how this is done. [NSA] suggests that this is done in the obvious
// manner, but [SECG] truncates the hash to the bit-length of the curve order
// first. We follow [SECG] because that's what OpenSSL does.
func hashToInt(hash []byte, c *bitelliptic.BitCurve) *big.Int {
	orderBits := c.N.BitLen()
	orderBytes := (orderBits + 7) / 8
	if len(hash) > orderBytes {
		hash = hash[:orderBytes]
	}

	ret := new(big.Int).SetBytes(hash)
	excess := orderBytes*8 - orderBits
	if excess > 0 {
		ret.Rsh(ret, uint(excess))
	}
	return ret
}

// Sign signs an arbitrary length hash (which should be the result of hashing a
// larger message) using the private key, priv. It returns the signature as a
// pair of integers. The security of the private key depends on the entropy of
// rand.
func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err os.Error) {
	// See [NSA] 3.4.1
	c := priv.PublicKey.BitCurve

	var k, kInv *big.Int
	for {
		for {
			k, err = randFieldElement(c, rand)
			if err != nil {
				r = nil
				return
			}

			kInv = new(big.Int).ModInverse(k, c.N)
			r, _ = priv.BitCurve.ScalarBaseMult(k.Bytes())
			r.Mod(r, priv.BitCurve.N)
			if r.Sign() != 0 {
				break
			}
		}

		e := hashToInt(hash, c)
		s = new(big.Int).Mul(priv.D, r)
		s.Add(s, e)
		s.Mul(s, kInv)
		s.Mod(s, priv.PublicKey.BitCurve.N)
		if s.Sign() != 0 {
			break
		}
	}

	return
}

// Verify verifies the signature in r, s of hash using the public key, pub. It
// returns true iff the signature is valid.
func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool {
	// See [NSA] 3.4.2
	c := pub.BitCurve

	if r.Sign() == 0 || s.Sign() == 0 {
		return false
	}
	if r.Cmp(c.N) >= 0 || s.Cmp(c.N) >= 0 {
		return false
	}
	e := hashToInt(hash, c)
	w := new(big.Int).ModInverse(s, c.N)

	u1 := e.Mul(e, w)
	u2 := w.Mul(r, w)

	x1, y1 := c.ScalarBaseMult(u1.Bytes())
	x2, y2 := c.ScalarMult(pub.X, pub.Y, u2.Bytes())
	if x1.Cmp(x2) == 0 {
		return false
	}
	x, _ := c.Add(x1, y1, x2, y2)
	x.Mod(x, c.N)
	return x.Cmp(r) == 0
}
