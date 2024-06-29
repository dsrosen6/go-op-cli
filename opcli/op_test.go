package opcli

import "testing"

func TestGetVaults(t *testing.T) {
	client := Client{}
	vaults, err := client.VaultList()
	if err != nil {
		t.Fatal(err)
	}

	if len(vaults) < 1 {
		t.Log("No vaults available to list")
	} else {
		for _, vault := range vaults {
			t.Logf("Vault: %s", vault.Name)
		}
	}
}

func TestGetVault(t *testing.T) {
	client := Client{}
	vault, err := client.VaultGet("kfoesdlf")
	if err != nil {
		t.Fatal(err)
	}

	// Output vault if not empty struct
	if vault != (Vault{}) {
		t.Logf("Vault: %s", vault.Name)
	}
}
