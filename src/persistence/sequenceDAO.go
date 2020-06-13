package persistence

import (
	"errors"
	"gopetstore/src/domain"
	"gopetstore/src/util"
)

const (
	getSequenceSQL    = `SELECT name, nextid FROM SEQUENCE WHERE NAME = ?`
	updateSequenceSQL = `UPDATE SEQUENCE SET NEXTID = ? WHERE NAME = ?`
)

func GetSequence(name string) (*domain.Sequence, error) {
	d, err := util.GetConnection()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = d.Close()
	}()
	r, err := d.Query(getSequenceSQL, name)
	if err != nil {
		return nil, err
	}
	if r.Next() {
		var name string
		var nextId int
		err := r.Scan(&name, &nextId)
		if err != nil {
			return nil, err
		}
		return &domain.Sequence{
			Name:   name,
			NextId: nextId,
		}, nil
	}
	return nil, errors.New("can not find sequence by this name ")
}

func UpdateSequence(s *domain.Sequence) error {
	d, err := util.GetConnection()
	if err != nil {
		return err
	}
	defer func() {
		_ = d.Close()
	}()
	r, err := d.Exec(updateSequenceSQL, s.NextId, s.Name)
	if err != nil {
		return err
	}
	row, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if row > 0 {
		return nil
	}
	return errors.New("can not update sequence")
}
