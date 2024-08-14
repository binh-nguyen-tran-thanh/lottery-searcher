package port

type Service interface {
	Ping() IPing
	Region() RegionServicer
	Result() ResultService
}
