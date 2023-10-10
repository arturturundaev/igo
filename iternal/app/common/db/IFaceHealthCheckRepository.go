package db

type IFaceHealthCheckRepository interface {
	Check() bool
}
