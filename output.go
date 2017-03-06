package dpipe

type Output interface {
	Write(Hotel) error
}
