package main

import (
	"api/src/config"
	"api/src/router"

	//"crypto/rand" *descomentar caso vá gerar o token JWT
	//"encoding/base64" *descomentar caso vá gerar o token JWT

	"fmt"
	"log"
	"net/http"
)

// *descomentar caso vá gerar o token JWT
// func init() {
// 	chave := make([]byte, 64)

// 	if _, erro := rand.Read(chave); erro != nil {
// 		log.Fatal()
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	fmt.Println(stringBase64)
// }

func main() {

	config.Carregar()

	fmt.Printf("API Listening on port %d\n", config.Porta)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
