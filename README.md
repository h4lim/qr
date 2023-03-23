A simple QRIS Reader written from golang

### Getting the library
With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/h4lim/qr"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `qr` package:

```sh
$ go get -u github.com/h4lim/qr
```

### Use QR Reader from raw data

First you need to import qr , one simplest example likes the follow `example.go`:

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/h4lim/qr"
)

func main() {
	rawData := "00020101021126650013ID.CO.BCA.WWW011893600014000094045302150008850009404530303UKE51440014ID.CO.QRIS.WWW0215ID20200340731930303UKE5204507253033605802ID5910PERKAKASKU6007BANDUNG61054027162070703A0163044D4A"
	qrisReader := qr.NewQrisReader(rawData)
	qrisTag, err := qrisReader.Read()
	if err != nil {
	    fmt.Println(*err)
	    return
	}
	
    b, err2 := json.Marshal(qrisTag)
    if err2 != nil {
        fmt.Println(err2)
        return
    }
    fmt.Println(string(b))
}
```

And use the Go command to run the demo:

```
# run example.go
$ go run example.go
```

The output will be like below:

```
{"tag_00":{"version":"01","type":"11","tag_52":"5072","tag_53":"360","tag_58":"ID","tag_61":"40271","tag_62":"0703A01","amount":0,"merchant_owner":"PERKAKASKU","merchant_address":"BANDUNG","checksum":"4D4A"},"tag_26":{"qr_owner":"ID.CO.BCA.WWW","merchant_id":"936000140000940453","merchant_acquirer_id":"000885000940453","merchant_scale":"UKE"},"tag_51":{"qris_web":"ID.CO.QRIS.WWW","qris_id":"ID2020034073193","scale":"UKE"}}
```
