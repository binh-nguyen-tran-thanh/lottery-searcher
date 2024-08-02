package port

type Repository interface {
	Region() RegionRepository
	Lottery() LotteryRepository
}
