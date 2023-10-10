# math
A simple repo with math functions. It was published as a module following the [Publishing a module](https://go.dev/doc/modules/publishing) tutorial from the go official website. One example of its use is written below:

```go
package main

import (
	"fmt"
	"log"

	"github.com/carlos-medina/math/utils"
)

func main() {

	r, err := utils.Factorial(-1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)
}

```