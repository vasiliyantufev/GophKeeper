package validator

/*
 * Text rules:
 * at least 10 letters
 */

func VerifyText(s []byte) bool {
	var (
		hasMinLen = false
	)
	if len(string(s)) >= 10 {
		hasMinLen = true
	}
	return hasMinLen
}
