package store

import "time"

func init() {
	Register("mogodb", func(opt Opt) (IStore, error) {
		return newSQLdbStroe(opt)
	})
}

type sqldbStore struct {
	dup            time.Duration
	maxIdleConnNum int
}

func newSQLdbStroe(opt Opt) (*sqldbStore, error) {
	ms := &sqldbStore{
		maxIdleConnNum: 3,
	}
	return ms, nil
}

func (s *sqldbStore) SavePersonInfo() error {
	return nil
}

func (s *sqldbStore) GetPersonInfoById(id string) {
	return
}
