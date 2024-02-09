package integration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cezarcruz/gym-address/domain"
)

func GetAddress(zipcode string) domain.Address {
	res, err := http.Get("https://viacep.com.br/ws/" + zipcode + "/json/")

	if err != nil {
		fmt.Println("dont know how to handle yet")
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var a domain.AddressViaCepResponse
	json.Unmarshal(body, &a)

	return domain.Address{
		Zipcode:      a.Cep,
		Street:       a.Logradouro,
		Complement:   a.Complemento,
		Neighborhood: a.Bairro,
		City:         a.Localidade,
		State:        a.Uf,
	}

}
