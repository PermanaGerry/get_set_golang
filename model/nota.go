package model

type DataNota struct {
	Nota  string
	Total int
}

func (dataNota *DataNota) SetNota(noNota string) {
	dataNota.Nota = noNota
}

func (dataNota *DataNota) GetNota() string {
	return dataNota.Nota
}

func (dataNota *DataNota) SetTotal(total int) {
	dataNota.Total = total
}

func (dataNota *DataNota) GetTotal() int {
	return dataNota.Total
}
