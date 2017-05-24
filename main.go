//    keypairGen
//    Copyright (C) 2017  boboliu
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.
//
//    You should have received a copy of the GNU General Public License
//    along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"flag"
	"fmt"
	"os"
	"crypto/x509"
	"encoding/pem"
	"strconv"
	"strings"
)

func main() {
	var bits int
	flag.IntVar(&bits, "b", 1024, "密钥长度，默认为1024位")
	GenRsaKey(bits)
}

func GenRsaKey(bits int) {

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	check(err)
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	check(err)
	err = pem.Encode(file, block)
	check(err)
	publicKey := &privateKey.PublicKey
	pubFile, err := os.Create("public.key")
	check(err)
	pubBytes := make([]string, len(publicKey.N.Bytes()))
	for c, i := range publicKey.N.Bytes() {
		pubBytes[c] = strconv.Itoa(int(i))
	}
	pubString := strings.Join(pubBytes, "/")
	fmt.Fprintln(pubFile, "N:", pubString)
	fmt.Fprintln(pubFile, "E:", strconv.Itoa(publicKey.E))
}

func check(err error) {
	if err != nil {
		fmt.Println("密钥生成失败")
		panic(err)
	}
}
