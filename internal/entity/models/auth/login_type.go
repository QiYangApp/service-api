package auth

// Type represents an login type.
type Type int

// Note: new type must append to the end of list to maintain compatibility.
const (
	NoType Type = iota
	Plain       // 1
	LDAP        // 2
	SMTP        // 3
	PAM         // 4
	DLDAP       // 5
	OAuth2      // 6
	SSPI        // 7
)

// String returns the string name of the LoginType
func (typ Type) String() string {
	return Names[typ]
}

// Int returns the int value of the LoginType
func (typ Type) Int() int {
	return int(typ)
}

// Names contains the name of LoginType values.
var Names = map[Type]string{
	LDAP:   "LDAP (via BindDN)",
	DLDAP:  "LDAP (simple auth)", // Via direct bind
	SMTP:   "SMTP",
	PAM:    "PAM",
	OAuth2: "OAuth2",
	SSPI:   "SPNEGO with SSPI",
}
