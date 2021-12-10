package cryptoAlgos

import "math/big"

func GetPower(b int, y int, p int) int {
	bigB := big.NewInt(int64(b))
	bigP := big.NewInt(int64(p))

	for i := 0; i < p; i++ {
		var ans big.Int
		ans.Exp(bigB, big.NewInt(int64(i)), bigP).Int64()
		if int(ans.Int64()) == y {
			return i
		}
	}

	return -1
}
