package main

import "encoding/xml"

type RSS struct {
	XmlName xml.Name `xml:"html"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	XmlName xml.Name `xml:"body"`
}
