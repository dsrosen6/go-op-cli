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

// Options for filtering the list of available vaults
type VaultListOption struct {
	Group      string
	Permission string
	User       string
}

// Options for creation of a new vault.
// If AllowAdminsToManage is not provided, the default policy for the account applies.
type VaultCreateOption struct {
	// TODO: Find clean solution for optional bools
	// AllowAdminsToManage *bool
	Description string
	Icon        string
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
func (c *Client) VaultList(options ...VaultListOption) ([]*VaultListItem, error) {
	args := []string{"list"}
	for _, option := range options {
		if option.Group != "" {
			args = append(args, "--group", option.Group)
		}
		if option.Permission != "" {
			args = append(args, "--permissions", option.Permission)
		}
		if option.User != "" {
			args = append(args, "--user", option.User)
		}
	}

	var out []*VaultListItem
	err := c.RunOpUnmarshal("vault", args, &out)
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

func (c *Client) VaultCreate(vaultName string, options ...VaultCreateOption) (Vault, error) {
	args := []string{"create", vaultName}
	for _, option := range options {
		// If AllowAdminsToManage is set to true or false, add "--allow-admins-to-manage" and "true" or "false" to args

		// TODO: Find clean solution for optional bools
		// if option.AllowAdminsToManage != nil {
		// 	boolStr := strconv.FormatBool(*option.AllowAdminsToManage)
		// 	args = append(args, "--allow-admins-to-manage", boolStr)
		// }
		if option.Description != "" {
			args = append(args, "--description", option.Description)
		}
		if option.Icon != "" {
			args = append(args, "--description", option.Icon)
		}
	}

	var out Vault
	err := c.RunOpUnmarshal("vault", args, &out)
	if err != nil {
		return Vault{}, fmt.Errorf("error creating vault %s: %s", vaultName, err)
	}

	return out, nil
}
