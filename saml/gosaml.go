package main

import (
	"fmt"

	"github.com/mattbaird/gosaml"
)

func main() {
	// Configure the app and account settings
	appSettings := saml.NewAppSettings("http://www.onelogin.net", "issuer")
	accountSettings := saml.NewAccountSettings("cert", "http://www.onelogin.net")

	// Construct an AuthnRequest
	authRequest := saml.NewAuthorizationRequest(*appSettings, *accountSettings)

	// Return a SAML AuthnRequest as a string
	saml, err := authRequest.GetRequest(false)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(saml)
}
