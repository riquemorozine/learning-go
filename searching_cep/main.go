package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the CEP: ")
	cep, err := reader.ReadString('\n')

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}

	cep = strings.TrimSpace(cep)
	url := fmt.Sprintf("https://viacep.com.br/ws/%v/json", cep)

	req, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}

	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	}

	var data ViaCEP
	err = json.Unmarshal(res, &data)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: parsing %s: %v\n", url, err)
	}

	fmt.Println(data.Bairro)
}
