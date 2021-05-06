package main

import (
	"fmt"

	messagepacker "github.com/wagner-aos/go-binary-message-packer/message"
)

func main() {

	str := "123"
	byteArray := []byte("123")

	fmt.Println("CONTENT:")
	fmt.Printf("\nstring: %v", str)
	fmt.Printf("\nbyte array: %v\n", byteArray)

	fmt.Printf("\nstring em hexa: %x", str)
	fmt.Printf("\nbyte array em hexa: %x\n", byteArray)

	fmt.Printf("\nlength string em hexa: %d", len(str))
	fmt.Printf("\nlength byte array em hexa: %d\n", len(byteArray))

	//=================================================================

	message := []byte("020630323030feff4641a8e1e20a000000000000000431363532363332363030303234353736313330303030303030303030303031303030303030303030303031303030303030303030303031393534303030353036303831353433363130303030303036313935343030303030303030393038313534333035303632373132303530363035303635353431303730303030303430393938373635343939393130393030303030323230323337353236333236303030323435373631333d32373132323031303030303031353330303030303334353036323930303030315465726d3030303949442d436f6465313120202020202047617353746174696f6e303120202020202020202020204f2746616c6c6f6e202020202020204d4f3030395236343034303130303834303834303938363132315f2a020840820258008407a0000000041010950500000080009a032105069c01009f02060000000010009f101c0114250000044000dac10000000000000000000000000000000000109f1a0208409f260865fcb7b725611fcd9f2701809f3303e0e8e89f34034103029f360200239f37043afe4c619f530152303139303030303030303030303330303834303132333031324d53313030303030373030333035303030302020202020202020202059593132334142434445464142434445464748494a4b4c4d4e4f505152535455565758595a")
	fmt.Printf("MESSAGE: %s", message)
	//PACKER
	packer := &messagepacker.MessagePack{}
	err := packer.NewMessagePack("Mastercard Interface Processor", 4, "hex", "hex", "hex", "hex")
	//err := packer.NewMessagePack("Mastercard Interface Processor", 4, "hex", "decimal", "hex", "ascii")
	if err != nil {
		fmt.Printf("\nPACKER ERROR: %s\n", err)
	}
	fmt.Printf("\nPACKER: %#v\n", packer)
	//HEADER
	header, err := packer.ConvertMessageHeader(message)
	if err != nil {
		fmt.Printf("\nHEADER ERROR: %s\n", err)
	}
	fmt.Printf("\nHEADER CONVERTED: %s\n", header)
	//PAYLOAD
	payload, err := packer.ConvertMessagePayload(message[4:])
	if err != nil {
		fmt.Printf("\nPAYLOAD ERROR: %s\n", err)
	}
	fmt.Printf("\nPAYLOAD CONVERTED: %s\n", payload)

}
