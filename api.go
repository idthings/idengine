package idengine

// Engine engines
type Engine struct{}

// ValidateSecret authenticates
func (e Engine) ValidateSecret(id string, secret string) bool {

	return false
}

// ValidateDigestSHA256 authenticates
func (e Engine) ValidateDigestSHA256(id string, digest string, content string) bool {

	return false
}

// NewIdentity creates an identity
func (e Engine) NewIdentity() (id string, secret string) {

	return "id", "secret"
}

// UpdateSecret updates
func (e Engine) UpdateSecret(id string, secret string) {

}
