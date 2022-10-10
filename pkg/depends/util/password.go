package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func GenRandomPassword(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll {
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

func HashOfAccountPassword(accountID string, password string) string {
	return string(toMD5(toMD5([]byte(fmt.Sprintf("%s-%s", accountID, password)))))
}

func toMD5(src []byte) []byte {
	m := md5.New()
	_, _ = m.Write(src)
	cipherStr := m.Sum(nil)
	return []byte(hex.EncodeToString(cipherStr))
}

func ExtractRawPasswordByAccountAndPassword(accountID, passwordMD5 string) (string, error) {
	return "", nil // TODO
}
