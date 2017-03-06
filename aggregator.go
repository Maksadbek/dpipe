package dpipe

// Aggregator sorts/groups
// collected data before writing to output
type Aggregator interface {
	Add(Hotel) error
	Do() (Hotel, error)
}
