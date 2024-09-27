package main

import (
	"strings"
)

func main() {

	clave := ""
	mensaje := "I mensaje"

	encryptMessage(clave, mensaje)

}

func encryptMessage(clave, mensaje string) string {
	if clave == "" || clave == " " {
		clave = "DCJ"
	}
	if mensaje == "" || mensaje == " " {
		return ""
	}

	sliceMessage := strings.Split(mensaje, "")
	var newMessage string
	for i, v := range sliceMessage {

		switch {

		case v == "a" || v == "A":
			key := append([]string{clave}, sliceMessage[i])
			newMessage = newMessage + strings.Join(key, "")
		case v == "e" || v == "E":
			key := append([]string{clave}, sliceMessage[i])
			newMessage = newMessage + strings.Join(key, "")

		case v == "i" || v == "I":
			key := append([]string{clave}, sliceMessage[i])
			newMessage = newMessage + strings.Join(key, "")
		case v == "o" || v == "O":
			key := append([]string{clave}, sliceMessage[i])
			newMessage = newMessage + strings.Join(key, "")
		case v == "u" || v == "U":
			key := append([]string{clave}, sliceMessage[i])
			newMessage = newMessage + strings.Join(key, "")
		default:
			newMessage = newMessage + v
		}

	}

	return newMessage

}
