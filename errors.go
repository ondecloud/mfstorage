package mfstorage

type errStoreNotFound struct {
	storeName string
}

func (e errStoreNotFound) Error() string {
	return "store not found: " + e.storeName
}

func ErrStoreNotFound(storeName string) error {
	return &errStoreNotFound{
		storeName: storeName,
	}
}
