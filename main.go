package main

import "github.com/squaaat/squaaat-api/cmd"


// swagger:route POST /foobar foobar-tag idOfFoobarEndpoint
// Foobar does some amazing stuff.
// responses:
//   200: foobarResponse

// This text will appear as description of your response body.
// swagger:response foobarResponse

func main() {
	cmd.Start()
}
