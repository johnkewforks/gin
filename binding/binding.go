// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"net/http"

	"gopkg.in/joeybloggs/go-validate-yourself.v4"
)

const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
)

type Binding interface {
	Name() string
	Bind(*http.Request, interface{}) error
}

var _validator = validator.NewValidator("binding", validator.BakedInValidators)

var (
	JSON     = jsonBinding{}
	XML      = xmlBinding{}
	GETForm  = getFormBinding{}
	POSTForm = postFormBinding{}
)

func Default(method, contentType string) Binding {
	if method == "GET" {
		return GETForm
	} else {
		switch contentType {
		case MIMEPOSTForm:
			return POSTForm
		case MIMEJSON:
			return JSON
		case MIMEXML, MIMEXML2:
			return XML
		default:
			return GETForm
		}
	}
}
