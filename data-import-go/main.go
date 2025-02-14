package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <xmlFile> <xsltFile>")
		return
	}

	xmlFile := os.Args[1]
	xsltFile := os.Args[2]

	cmd := exec.Command("xsltproc", xsltFile, xmlFile)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing xsltproc:", err)
		return
	}

	transformedXML := out.Bytes()

	var data interface{}
	if err := xml.Unmarshal(transformedXML, &data); err != nil {
		fmt.Println("Error parsing transformed XML:", err)
		return
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
