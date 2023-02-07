package repository

import (
	"fmt"
	"math/rand"
)

type Repo interface {
	Get(url string) (value string, found bool)
	Delete(url string) (err error)
	Insert(url string) (short string, err error)
}

type Repository struct {
	linkMap map[string]string
}

func New() Repository {
	return Repository{
		linkMap: map[string]string{},
	}
}

func (r Repository) Get(url string) (value string, found bool) {
	value, found = r.linkMap[url]
	if !found {
		return
	}
	return
}

func (r Repository) Delete(url string) (err error) {
	return
}

func (r Repository) Insert(url string) (short string, err error) {
	genString := fmt.Sprint(rand.Int63n(1000))
	r.linkMap[genString] = url
	return genString, nil
}
