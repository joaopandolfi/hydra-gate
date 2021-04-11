package security

import (
	"testing"

	"hydra_gate/remotes/aes"
)

func TestHash(t *testing.T) {
	key := "32 character key ..............."
	planeVal := "bananinha"
	crypted1, _ := aes.UnsecureEncrypt(key, planeVal)
	crypted2, _ := aes.UnsecureEncrypt(key, planeVal)
	if crypted1 != crypted2 {
		t.Errorf("expected equal, but got: %s and %s", crypted1, crypted2)
	}
}
