package opcli

// 1Password CLI has multiple optional flags which must be passed as strings. In order
// to handle this, BoolPtrString translates a bool pointer to the strings "true",
// "false", or "" if nil.
func BoolPtrString(b *bool) string {
	if b != nil {
		if *b {
			return "true"
		}
		return "false"
	}

	return ""
}
