package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nathanzeras/go-ws-tks/app/models"
)

func Search(w http.ResponseWriter, r *http.Request) {

	//fmt.Println(mux.Vars(r))

	//fmt.Println(mux.Vars(r)["cpf"])
	//fmt.Println(mux.Vars(r)["procedimento"])
	p, err := ioutil.ReadAll(r.Body)
	
	if err != nil {
		//log.Fatalf("Error reading request: %v", err)
		fmt.Printf("Error reading request: %v", err)
	}
	var envio models.Envio

	if err := json.Unmarshal(p, &envio); err != nil {

		fmt.Printf("Error during unmarshall envio: %v", err)
		errRetorno := err.Error()
		fmt.Fprint(w, string(errRetorno))

	} else {
		//Removing all occorruncies from caracter on string
		output := models.SearchAutorizacao(&envio)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(output))
	}

	//fmt.Println(r.URL.Query())

}


func SearchOld(w http.ResponseWriter, r *http.Request) {

	//fmt.Println(mux.Vars(r))

	//fmt.Println(mux.Vars(r)["cpf"])
	//fmt.Println(mux.Vars(r)["procedimento"])

	output := models.SearchAutorizacaoOld(mux.Vars(r))
	//fmt.Println(r.URL.Query())

	fmt.Fprint(w, string(output))
}

