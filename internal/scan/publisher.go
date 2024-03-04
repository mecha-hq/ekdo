package scan

type Publisher interface {
	Publish(path string) error
}
