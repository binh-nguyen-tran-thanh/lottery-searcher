package port

type Service interface {
	Ping() IPing
}
