package main

import (
	"fmt"
	"timestamps/hnet"
)

//=============================== MAIN ==============================

func main() {

	var url string
	fmt.Println("Please enter the link from where you want to get the data: ")
	fmt.Scanln(&url)

	hnet.Url_selection(url)

	//Example:
	//The program serach for tag "START" in github .md to do this work.
	// insert any of the following links
	// url1 := "https://www.youtube.com/watch?v=Pfk-KxO3in8"
	// url := "https://github.com/rwxrob/boost/blob/2022/03/README.md"
	// and offset 00:00:09
	// The program also test the valid link, valid data in the link and valid offset format.
	// Sure there are possible errors but is the first version.
}

//===================================================================
//===================================================================
