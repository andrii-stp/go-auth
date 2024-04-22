package auth

type IdpProvider struct {
}

type Opts struct {
}

func NewIdpProvider(opts Opts) *IdpProvider {
	return &IdpProvider{}
}
