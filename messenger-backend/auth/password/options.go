package password

type EncryptOption interface {
	apply(*Argon2Gen)
}

type saltOption []byte

func (s saltOption) apply(ag *Argon2Gen) {
	ag.defaultSalt = s
}

func WithSalt(salt string) EncryptOption {
	return saltOption(salt)
}

type defaultSaltOption []byte

func (d defaultSaltOption) apply(ag *Argon2Gen) {
	ag.defaultSalt = d
}

func WithDefaultSalt() EncryptOption {
	return defaultSaltOption("default-salt")
}

func NewArgon2Gen(opts ...EncryptOption) *Argon2Gen {
	ag := &Argon2Gen{}
	for _, opt := range opts {
		opt.apply(ag)
	}
	return ag
}
