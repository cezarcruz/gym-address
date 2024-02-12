package integration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	m "github.com/cezarcruz/gym-address/internal/domain/model"
)

type AddressViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

func GetAddressBy(zipcode string) *m.Address {
	res, err := http.Get("https://viacep.com.br/ws/" + zipcode + "/json/")

	if err != nil {
		fmt.Println("dont know how to handle yet")
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var a AddressViaCepResponse
	json.Unmarshal(body, &a)

	return &m.Address{
		Zipcode:      a.Cep,
		Street:       a.Logradouro,
		Complement:   a.Complemento,
		Neighborhood: a.Bairro,
		City:         a.Localidade,
		State:        a.Uf,
	}
}
