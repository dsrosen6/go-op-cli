package opcli

import (
	"fmt"
	"time"
)

// Short summary of a 1Password vault, output of "vault list" op command
type VaultListItem struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ContentVersion int64  `json:"content_version"`
}

// Full details of a 1Password vault, output of "vault get" op command
type Vault struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	ContentVersion   int64     `json:"content_version"`
	AttributeVersion int64     `json:"attribute_version"`
	Items            int64     `json:"items"`
	Type             string    `json:"type"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// Retrieves a list of available vaults with short summaries including ID, name, and content version.
// Mostly used in order to get ID and name to pass to GetVault function.
func (c *Client) VaultList() ([]*VaultListItem, error) {
	var out []*VaultListItem
	err := c.RunOpUnmarshal("vault", []string{"list"}, &out)
	if err != nil {
		return nil, fmt.Errorf("error getting vaults list: %s", err)
	}

	return out, nil
}

// Retrieves full details of a specific vault. "identifier" argument can take either
// vault ID or vault name.
func (c *Client) VaultGet(identifier string) (Vault, error) {
	var out Vault
	err := c.RunOpUnmarshal("vault", []string{"get", identifier}, &out)
	if err != nil {
		return Vault{}, fmt.Errorf("error getting details for vault %s: %s", identifier, err)
	}

	return out, nil
}
