package cryptoAlgos

import "math/big"

type Alice struct {
	a *big.Int
	p *big.Int
	g *big.Int
	A *big.Int
	B *big.Int
	K *big.Int
}

func NewAlice(a int, p int, g int) *Alice {
	return &Alice{big.NewInt(int64(a)),
		big.NewInt(int64(p)),
		big.NewInt(int64(g)),
		big.NewInt(-1),
		big.NewInt(-1),
		big.NewInt(-1)}
}

type Bob struct {
	b *big.Int
	p *big.Int
	g *big.Int
	A *big.Int
	B *big.Int
	K *big.Int
}

func NewBob(a int, p int, g int) *Bob {
	return &Bob{big.NewInt(int64(a)),
		big.NewInt(int64(p)),
		big.NewInt(int64(g)),
		big.NewInt(-1),
		big.NewInt(-1),
		big.NewInt(-1)}
}

func GetKeys(alice *Alice, bob *Bob) (*big.Int, *big.Int) {
	alice.A.Exp(alice.g, alice.a, alice.p)
	bob.B.Exp(bob.g, bob.b, bob.p)
	alice.B = bob.B
	bob.A = alice.A
	alice.K.Exp(alice.B, alice.a, alice.p)
	bob.K.Exp(bob.A, bob.b, bob.p)

	return alice.K, bob.K
}
