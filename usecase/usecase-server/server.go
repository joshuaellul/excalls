package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"hash"
	"io"
	"log"
	"math/big"
	"net/http"
	"strings"
)

var private_key *ecdsa.PrivateKey

func sign_message(s string) string {
	var h hash.Hash
	h = md5.New()
	io.WriteString(h, s)
	signhash := h.Sum(nil)

	sigr, sigs, _ := ecdsa.Sign(rand.Reader, private_key, signhash)
	ret := s + "\n"
	ret += sigr.Text(16) + "\n"
	ret += sigs.Text(16) + "\n"
	ret += private_key.PublicKey.X.Text(16) + "\n"
	ret += private_key.PublicKey.Y.Text(16) + "\n"

	return ret
}

func excall0str(w http.ResponseWriter, r *http.Request) {
	ret := sign_message(strings.Repeat("0", 32))
	fmt.Fprintf(w, ret)
}

func excall1str(w http.ResponseWriter, r *http.Request) {
	//ret := sign_message(strings.Repeat("1", 32))
	ret := sign_message("11111111111111111111111111111111")
	fmt.Fprintf(w, ret)
}

func excallrand(w http.ResponseWriter, r *http.Request) {
	ret, _ := rand.Int(rand.Reader, big.NewInt(2))
	if ret.Int64() == 1 {
		excall1str(w, r)
	} else {
		excall0str(w, r)
	}
}

func init_key() {
	private_key, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

//http://localhost:8080/excall0str
func main() {
	init_key()

	http.HandleFunc("/excall0str", excall0str)
	http.HandleFunc("/excall1str", excall1str)
	http.HandleFunc("/excallrand", excallrand)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
