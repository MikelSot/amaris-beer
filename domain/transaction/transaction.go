package transaction

type Transaction interface {
	Commit() error
	Rollback() error
}
