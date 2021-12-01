package main

import (
	"encoding/json"
	"fmt"
	"get_set_golang/repository"
)

var ident repository.RepositoryUser

func main() {

	data := `[{"Nota":"testkasus1","Total":23},{"Nota":"testkasus2","Total":34},{"Nota":"testkasus3","Total":67}]`

	json.Unmarshal([]byte(data), &ident)

	ident[1].SetNota("change1")

	fmt.Println(ident[1].GetNota())
	fmt.Println(ident.GetAllNota())
	fmt.Println(ident.FindByNota("testkasus2"))
}
