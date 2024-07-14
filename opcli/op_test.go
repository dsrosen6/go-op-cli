package opcli

import (
	"testing"
)

func TestVaultList(t *testing.T) {
	op := Client{}
	vaults, err := op.GetVaultList()
	if err != nil {
		t.Error(err)
	}

	if len(vaults) < 1 {
		t.Log("No vaults available")
	} else {
		for _, vault := range vaults {
			t.Logf("Vault: %s", vault.Name)
		}
	}
}

func TestGetItemCategory(t *testing.T) {
	op := Client{}
	category, err := op.getItemCategory("Netflix")
	if err != nil {
		t.Error(err)
	}

	t.Logf("Category: %s", category)
}
