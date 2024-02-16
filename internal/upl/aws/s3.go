package aws

func NewS3Publisher() *S3Publisher {
	return &S3Publisher{}
}

type S3Publisher struct {
}

func (p *S3Publisher) Publish(path string) error {
	return nil
}
