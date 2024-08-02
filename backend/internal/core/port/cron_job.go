package port

type CronJob interface {
	Start() error
	Stop() error
}
