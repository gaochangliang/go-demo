package store

import "time"

func init() {
	Register("mysql", func(opt Opt) (IStore, error) {
		return newMySqldbStroe(opt)
	})
}

type mongodbStore struct {
	dup            time.Duration
	maxIdleConnNum int
}

func newMySqldbStroe(opt Opt) (*sqldbStore, error) {
	ms := &sqldbStore{
		maxIdleConnNum: 3,
	}
	return ms, nil
}

func (s *mongodbStore) SavePersonInfo() error {
	return nil
}

func (s *mongodbStore) GetPersonInfoById(id string) {
	return
}
