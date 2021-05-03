![](images/tcp-streams.jpeg)

# Go Binary Message Packer

This module helps to convert binary messages received on the TCP socket stream to be processed by another server with a different encoding.

| Date | Author | Comments | Version |
| --- | --- | --- | --- |
| 01/04/2020 | Wagner aka Barão | This is the Beginning of success | 0.0.1 |


### Message Components

|  | **Binary, Hex or Decimal** | **Hex or ASCII** | **Description** |
| --- | --- | --- | ---- | 
| **Header** | 0251 | 0593 | The size of the message |
| **Content** | 6632336336373031... | f23c6701... | The payload (message content) |

### How to use in your code:

* Get from repository

```sh
    go get -v -u "github.com/wagner-aos/go-binary-message-packer"
```


* Import in your go file:

```go

    import "github.com/wagner-aos/go-binary-message-packer"
```

* Create the object that defines the conversions in order to do some operations

```go
    /*This object 'MessagePack' defines how the message will be converted:
    * header  -> from 'decimal' to 'hex'
    * payload -> from  'ascii' to 'hex'
    
    E.g.: According to object MessagePack{} below, the message (header+payload) will be received by socket (decimal + ascii), and will converted to (hex + hex) in order to be written to socket.
    */
    packer := &MessagePack{}
    packer.NewMessagePack("Name of the client", "decimal", "hex", "ascii", "hex")

    //You can validate message size according to header and payload
    isValid, err := packer.ValidateMessageSize(header, payload)

    //Convert message header to decimal in order to do some calcs.
    decimalHeader := packer.GetDecimalHeader(header)
			
    //Get the payload size in order to do some calcs.
    decimalPayloadSize := packer.GetDecimalPayloadSize(payload)

    //Convert message header in order to send to tcp socket for example
    headerConverted := packer.ConvertMessageHeader(header)

    //Convert message payload in order to send to tcp socket for example
    payloadConverted := packer.ConvertMessagePayload(payload)

```

### Message examples

1. HEX Header (0251) + HEX Payload(3032303066323363363730313065653139...)  
```text
    0251303230306632336336373031306565313961343030303030303030303030303030303032313634353334353637383132333435363738303032303030303030303030303030383831303430383230333030383030303030313230333030383034303832353038373130303037363035393131313130303131303030303030303030303130303030303030303030303030314f4b383830303132333435363738313233343536373839313233343536416d617a6f6e2053686f7020202020202020202020202053414f205041554c4f20202020203037363032312a434e50303134303030303030303030303030303039363831313131313131313030303330303130303030303030303032353539463236303835344244313541453231304135313137394632373031383039463130313230323130413530303032303230303030303030303030303030303030303030303030464639463337303442353736443131333946333630323030463639353035343230303030383030303941303331393033323239433031303039463032303630303030303030323538303039463033303630303030303030303030303035463241303230393836383230323538303039463141303230303736394633353031323239463334303334313033303235463234303332323035333139463333303345304630433835463238303230303736383430374130303030303030353941762052696f204e6567726f2c20353030202020202032333435362d383930303736343536373853686f7070696e67204967756174656d692020203030353030303031    
```

### The Message above can be converted from HEX to ASCII(Text) using a online tools such as:

* https://codebeautify.org/hex-string-converter

And the result obtained will be:

```text
    05930200f23c67010ee19a4000000000000000021645345678123456780020000000000008810408203008000001203008040825087100076059111100110000000000100000000000001OK880012345678123456789123456Amazon Shop            SAO PAULO     076021*CNP014000000000000009681111111100030010000000002559F260854BD15AE210A51179F2701809F10120210A50002020000000000000000000000FF9F3704B576D1139F360200F6950542000080009A031903229C01009F02060000000258009F03060000000000005F2A020986820258009F1A0200769F3501229F34034103025F24032205319F3303E0F0C85F280200768407A000000059Av Rio Negro, 500     23456-89007645678Shopping Iguatemi   00500001
```

### Change log:

| Date | Feature | Comments | Version |
| --- | --- | --- | --- |
| 01/04/2020 | Message Converter | Code | O.O.1|


### Useful links and docs

[You can test the message conversion here](https://codebeautify.org/hex-string-converter)


### Useful commands:

###  Creating a version of this lib:

* 1- Test it all!
```sh
    make test
```

* 2- Create a git tag with the version number.
```sh
    git tag -a v0.0.2 -m "Describe the feature or hotfix here"
```

* 3- Push the tag to git repo.
```sh
    git push origin v0.0.2
```

* 4- Now the new version can be dowloaded to the other go projects.
```sh
    go get -u -v "github.com/wagner-aos/go-binary-message-packer"    
```

### I hope you enjoy it ;-)


