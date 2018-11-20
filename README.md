# adcl

adcl is a golang ADC library, covering both low-level protocol aspects as well
as higher-level client and hub functionality.

Resources:

 * [ADC Project Homepage](https://adc.dcbase.org/)
 * [ADC Protocol Specification](https://adc.dcbase.org/Protocol)
 * [ADC Extension Specification](https://adc.dcbase.org/Extensions)

## Development Status

adcl is under active (?) development. As soon as sub-packages reach beta
status (semi-stable) they will be listed here.

## Example: Reading messages from connection

```golang
package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/seoester/adcl/protocol/message"
	"github.com/seoester/adcl/protocol/parser"
)

// streamContent is the content of a simulated connection with three messages.
// The last message is invalid (code is missing) and will result in an error.
const streamContent = `ISTA 144 Invalid\sparameter\svalue! TOqwertzu
BSTA SDFSDF 000 OK
CSTA
`

func main() {
	conn := strings.NewReader(streamContent)
	reader := bufio.NewReader(conn)

	p := parser.New(reader)

	for {
		mes, err := p.ReadMessage()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error occurred:", err)
			continue
		}

		fmt.Printf("Received Message: %+v\n Content: %+v\n", mes, mes.Content)

		handle(&mes)
	}

	fmt.Println("End of stream.")
}

func handle(mes *message.Message) {
	switch mes.Command {
	case message.CommandSTA:
		cnt := mes.Content.(*message.STAContent)
		fmt.Printf("Processing status message with code %d\n", cnt.Code.Code())
	}
}
```
