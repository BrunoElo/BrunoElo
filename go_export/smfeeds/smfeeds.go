// A program that exports social media data to external files
// of the format "txt", "json", "xml" and "yaml"

package main

import (
	"encoding/xml"
	"io"
	"os"
	"bytes"
	"encoding/json"
	"go_export/exporter"
	"fmt"
	"go_export/smediafeeds"
	"gopkg.in/yaml.v2"
)

func main() {
	fb := new(smediafeeds.Facebook)
	twtr := new(smediafeeds.Twitter)
	lnkdn := new(smediafeeds.Linkedin)
	
	// Exporting to "txt" file
	err := exporter.Export(fb, "fbdata.txt")
	err = exporter.Export(twtr, "twtrdata.txt")
	err = exporter.Export(lnkdn, "lnkdndata.txt")
	
	// Array of slices was made to hold the social Media slices
	smfeeds := [3][]string{fb.Feed(),twtr.Feed(),lnkdn.Feed()}
	l := len(smfeeds)

	// file made in global scope to act as global declaration of "file"
	file, err := os.Create("declaration")
		if err != nil {
			panic(err)
		}
		defer file.Close()

	//To make json exports to json files(in the specified format)
	//slices had to be converted to maps using slicetomap function
	
	buf := new(bytes.Buffer)
	jencoder := json.NewEncoder(buf)	

	for i := 0; i < l; i++ {
		if i == 0 {
			// Exporting fbmap to "json" file
			fbmap := slicetomap(smfeeds[i])
			jencoder.Encode(fbmap)
			file, err = os.Create("fbdata.json")
		} else if i == 1 {
			// Exporting twtrmap to "json" file
			twtrmap := slicetomap(smfeeds[i])
			jencoder.Encode(twtrmap)
			file, err = os.Create("twtrdata.json")
		} else {
			// Exporting lnkdnmap to "json" file
			lnkdnmap := slicetomap(smfeeds[i])
			jencoder.Encode(lnkdnmap)
			file, err = os.Create("lnkdndata.json")
		}
		// Checks for errors
		if err != nil {
			panic(err)
		}		
		defer file.Close()  // Ensures that file is closed		
		io.Copy(file, buf)  // Copies data to file
	}
	
	//XML	
	buf = new(bytes.Buffer)
	xencoder := xml.NewEncoder(buf)
	
	for i := 0; i < l; i++ {
		if i == 0 {
			// Exporting fb slice to "xml" file
			xencoder.Encode(smfeeds[i])
			file, err = os.Create("fbdata.xml")
		} else if i == 1 {
			// Exporting twtr slice to "xml" file
			xencoder.Encode(smfeeds[i])
			file, err = os.Create("twtrdata.xml")
		} else {
			// Exporting lnkdn slice to "xml" file
			xencoder.Encode(smfeeds[i])
			file, err = os.Create("lnkdndata.xml")
		}
		if err != nil {
			panic(err)
		}
		defer file.Close()
		io.Copy(file, buf)
	}

	//YAML		
	buf = new(bytes.Buffer)
	yencoder := yaml.NewEncoder(buf)
	
	for i := 0; i < l; i++ {
		if i == 0 {
			// Exporting fb to "yaml" file
			yencoder.Encode(smfeeds[i])
			file, err = os.Create("fbdata.yaml")
		} else if i == 1 {
			// Exporting twtr to "yaml" file
			yencoder.Encode(smfeeds[i])
			file, err = os.Create("twtrdata.yaml")
		} else {
			// Exporting lnkdn to "yaml" file
			yencoder.Encode(smfeeds[i])
			file, err = os.Create("lnkdndata.yaml")
		}
		if err != nil {
			panic(err)
		}
		defer file.Close()
		io.Copy(file, buf)
	}
	PrintFeeds(fb, twtr, lnkdn)
}

// PrintFeeds prints the social media feeds on the terminal
func PrintFeeds(platforms ...smediafeeds.SocialMedia) {
	fmt.Println("*******************************")
	for _, sm := range platforms {
		for _, fd := range sm.Feed() {
			fmt.Println(fd)
		}
		fmt.Println("*******************************")
	}
}

// function slicetomap converts the slice values returned from the
// .Feed() method into a map of string values
func slicetomap(e []string) map[string] string {		
	elements := e
	slimap := make(map[string]string)
	str := [3]string{"1","2","3"}
	for i := 0; i< len(elements); i++ {
		slimap[str[i]] = elements[i]
	}
	fmt.Println(slimap)		
	return slimap
}
	
	