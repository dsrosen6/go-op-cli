package opcli

import "fmt"

// Short summary of a 1Password vault, output of "vault list" op command used in GetVaultList
type VaultListItem struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ContentVersion int64  `json:"content_version"`
}

// Options for filtering the list of available vaults in GetVaultList
type VaultListOptions struct {
	Group      string
	Permission string
	User       string
}

// Retrieves a list of available vaults with short summaries including ID, name, and content version.
//
// Mostly used in order to get ID and name to pass to GetVault function.
func (op *Client) GetVaultList(options ...VaultListOptions) ([]*VaultListItem, error) {
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
	if err := op.commandWithUnmarshal("vault", args, &out); err != nil {
		return nil, fmt.Errorf("error getting vault list: %w", err)
	}

	return out, nil
}
