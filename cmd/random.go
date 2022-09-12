// [[file:../docs.org::Code/Source/cmd/random.go][Code/Source/cmd/random.go]]
/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dad joke.",
	Long:  `This command fetches a random dad joke from the icanhazdadjoke api.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
		// fmt.Println("random called") // works
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

/*
  * getRandomJoke FUNCTION
  *
	* Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
  * If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
  * func Unmarshal(data []byte, v any) error // only returns an error
	* https://pkg.go.dev/encoding/json#Unmarshal
  * In order to see this in action, call it in the randomCmd() run command.
*/
func getRandomJoke() {
	// Setup data to parse.
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url) // first arg_1
	// Set up arg_2 - an interface to store parsed values
	joke := Joke{} // joke variable is a Joke struct

	// Pass joke as a pointer as the second argument
	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		// Since err is the only returned value, use this syntax to
		// return it and declare at the same time
		log.Printf("Could not Unmarshal response - %v", err)
	}

	// PRINT result to console & Convert to string, access from joke variable with dot notation
	fmt.Println(string(joke.Joke))

} // ARCHIVE(docs.org) => fmt.Println("Get Random Joke :P") // redundant after return of responseBytes

// https://pkg.go.dev/net/http#Get
// baseAPI is url
func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	// Handle the errors.
	if err != nil {
		// log.Fatal("Could not request a dadjoke - %v", err) // Fatal leads to exit
		log.Printf("Could not request a dadjoke - %v", err)
	}

	// Let the API know we want the header encoded as JSON
	request.Header.Add("Accept", "application/json")
	// Add custum user agent header on how to use the API
	request.Header.Add("User-Agent", "Dadjoke CLI (github.com/example/dadjoke)")

	// Pass it into the default client.do method
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request - %v", err)
	}
	// Get responseBytes with io.ReadAll method on response.Body
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body - %v", err)
	}

	// After returning work on the getRandomJoke()
	return responseBytes

}

// Code/Source/cmd/random.go ends here
