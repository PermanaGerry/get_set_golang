package repository

import (
	"test_pointer/inter"
	"test_pointer/model"
)

type InterfaceNota struct {
	inter.InterfaceRepoNota
}

type RepositoryUser []*model.DataNota

// func (repo RepositoryUser) GetAll() *RepositoryUser {
// 	return repo
// }

func (repo RepositoryUser) GetAllNota() []string {
	var list []string

	for _, v := range repo {
		list = append(list, v.GetNota())
	}

	return list
}

func (repo RepositoryUser) GetAllTotal() []int {
	var list []int

	for _, v := range repo {
		list = append(list, v.GetTotal())
	}

	return list
}

func (repo RepositoryUser) FindByNota(nota string) []model.DataNota {
	var filtered []model.DataNota

	for _, v := range repo {
		if v.GetNota() == nota {
			filtered = append(filtered, model.DataNota{
				Nota:  v.GetNota(),
				Total: v.GetTotal(),
			})
		}
	}

	return filtered
}

func (repo RepositoryUser) FindByTotal(total int) []model.DataNota {
	var filtered []model.DataNota

	for _, v := range repo {
		if v.GetTotal() == total {
			filtered = append(filtered, model.DataNota{
				Nota:  v.GetNota(),
				Total: v.GetTotal(),
			})
		}
	}

	return filtered
}
