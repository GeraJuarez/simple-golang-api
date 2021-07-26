package interactor

type KVStoreInteractor interface {
	Get(key string) (string, error)
}
