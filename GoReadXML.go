package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

/*The xml package includes Unmarshal() function that supports decoding data
from a byte slice into values.
The xml.Unmarshal() function
is used to decode the values from the XML formatted file into a Notes struct.*/

type Notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

func main() {
	data, _ := ioutil.ReadFile("notes.xml")

	note := &Notes{}

	_ = xml.Unmarshal([]byte(data), &note)

	fmt.Println(note.To)
	fmt.Println(note.From)
	fmt.Println(note.Heading)
	fmt.Println(note.Body)
}
