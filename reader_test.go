package reader_test

import (
	"encoding/json"
	"fmt"
	reader "github.com/h4lim/qr"
	"testing"
)

func TestQrisReader(t *testing.T) {
	rawData := "00020101021126650013ID.CO.BCA.WWW011893600014000094045302150008850009404530303UKE51440014ID.CO.QRIS.WWW0215ID20200340731930303UKE5204507253033605802ID5910PERKAKASKU6007BANDUNG61054027162070703A0163044D4A"
	qrisReader := reader.NewQrisReader(rawData)
	qrisTag, err := qrisReader.Read()
	if err != nil {
		t.Error("ERROR",*err)
		return
	}

	b, err2 := json.Marshal(qrisTag)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(string(b))
	t.Log("SUCCESS ", string(b))
}
