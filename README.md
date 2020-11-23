# Fork of [go-itg](https://deadbeef.codes/steven/go-itg)
Thanks Steven
## go-itg
IT Glue API Structs + Methods for the Go

This is nowhere near complete, and at least for now, I will only be adding support the functionality required for use within my organization.  Feel free to send a pull request if you'd like to contribute anything that's missing. 

## Installation
```
go get github.com/Private-Universe/itglue
```

## Usage
```
package main

import (
	"fmt"
	"log"

	"github.com/Private-Universe/itglue"
)

func main() {
	fmt.Println("Example IT Glue Application")
	itg := itglue.NewITGAPI("ITG.XXXXXXXXXXXXXXXXXXXXXXXX.XXXXXXXXXXXXXXXXXXXXXXXXXXXX-XXXXXXXXXXXX")

	nd, err := itg.GetOrganizationByName("Next Digital Inc.", 1) //Returns page one of an organization list
	if err != nil {
		log.Fatalf("could not get nd: %s", err)
	}
	log.Printf("%s - %s\n", nd.Data[0].ID, nd.Data[0].Attributes.Name)

}

```
