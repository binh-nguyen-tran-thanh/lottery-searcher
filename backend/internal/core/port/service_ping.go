package port

type IPing interface {
	Ping() (map[string]any, error)
	Healthy() (map[string]any, error)
}
