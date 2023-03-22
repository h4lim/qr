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
    rawData := "00020101021226640013COM.MYWEB.WWW01181234567890123456780214123456789012340303UKE5405100005912QRIS WANTUNO6013Jakarta Pusat6304XXXX"
	qrisReader := qr.NewQrisReader(rawData)
	tag00, err := qrisReader.Read()
	if err != nil {
	    fmt.Println(*err)
	    return
	}
	
    b, err2 := json.Marshal(tag00)
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
{"Version":"01","Type":"12","Amount":10000,"MerchantOwner":"QRIS WANTUNO","MerchantAddress":"Jakarta Pusat","Checksum":"XXXX","Tag26":{"QrOwner":"COM.MYWEB.WWW","MerchantID":"123456789012345678","MerchantAcquirerID":"12345678901234","MerchantScale":"UKE"}}
```
