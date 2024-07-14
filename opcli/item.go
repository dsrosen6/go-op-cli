package opcli

import (
	"encoding/json"
	"fmt"
)

func (op *Client) getItemCategory(name string) (string, error) {
	args := []string{"get", name}
	var result map[string]json.RawMessage
	err := op.commandWithUnmarshal("item", args, &result)
	if err != nil {
		return "", fmt.Errorf("error getting item from 1Password: %w", err)
	}

	// Output the "category" field from the json output
	category := string(result["category"])
	return category, nil
}
