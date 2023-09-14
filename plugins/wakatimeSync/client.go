package wakatimeSync

import "wakatimeSync/net"

func New(token string) net.Api {
	return net.Api{Token: token}
}
