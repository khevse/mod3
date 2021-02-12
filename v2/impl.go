package mod3

type Impl struct {
}

func New() *Impl {
	return &Impl{}
}

func (i Impl) String() string {
	return "v2.0.1"
}
