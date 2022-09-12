// [[file:docs.org::Code/Source/main.go][Code/Source/main.go]]
/*
Copyright
 © 2022 NAME HERE <EMAIL ADDRESS>
*/

package main

import (
	"fmt"
	"github.com/lloydlobo/godadjoke/cmd"
)

func main() {
	cmd.Execute()

	// Testing fmt import after cmd is executed.
	fmt.Println("fmt.Println: dadjoke")
}
// Code/Source/main.go ends here
