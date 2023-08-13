# saia-go
Go client for https://saia.3dlook.me/docs/

## Example

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shing-dev/saia-go"
)

func main() {
	apiKey := os.Getenv("SAIA_API_KEY")
	saiaClient := saia.NewClient(apiKey)

	resp, err := saiaClient.Person.CreatePerson(context.Background(), &saia.CreatePersonParams{
		Gender: saia.GenderMale,
		Height: 170,
		Weight: 62.5,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
```