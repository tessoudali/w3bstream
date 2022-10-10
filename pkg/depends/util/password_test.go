package util_test

import (
	"testing"

	"github.com/iotexproject/w3bstream/pkg/depends/util"
	. "github.com/onsi/gomega"
)

var (
	gPassword  = "12345"
	gAccountID = "100"
)

func TestHashOfAccountPassword(t *testing.T) {
	cryptPasswd := util.HashOfAccountPassword(gAccountID, gPassword)
	t.Log(cryptPasswd)
	password, err := util.ExtractRawPasswordByAccountAndPassword(gAccountID, cryptPasswd)
	NewWithT(t).Expect(err).To(BeNil())
	NewWithT(t).Expect(password).To(Equal(gPassword))
}
