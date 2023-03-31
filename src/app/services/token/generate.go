package token

// var tokenGeneratorInstance *tokenGenerator

// type GeneratorInterface interface{
// 	Generator
// 	init() *GeneratorInterface
// 	generator() string
// 	parser() interface{}
// 	updateAt() bool
// }

// type Generator struct {
// 	ExpiresAt int
// 	Issuer    string
// 	Subject   string
// 	Id        int
// 	Audience  []string
// 	jwt.RegisteredClaims
// }

// type TokenGenerator struct {
// 	Generator
// }

// func (t *TokenGenerator) init() *GeneratorInterface {
// 	// Create the claims
// 	jwt.RegisteredClaims{
// 		// A usual scenario is to set the expiration time relative to the current time
// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.ExpiresAt * time.Minute)),
// 		IssuedAt:  jwt.NewNumericDate(time.Now()),
// 		NotBefore: jwt.NewNumericDate(time.Now()),
// 		Issuer:    t.Issuer,
// 		Subject:   t.Subject,
// 		ID:        t.ID,
// 		Audience:  t.Audience,
// 	},

// 	return t
// }

// func (t *TokenGenerator) generate() string {
// 	return ""
// }
// func (t *TokenGenerator) parser() interface{} {
// 	return new(interface{})
// }
// func (t *TokenGenerator) updateAt() bool {
// 	return true
// }
