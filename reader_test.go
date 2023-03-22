package reader_test

import (
	reader "github.com/h4lim/qr"
	"testing"
)

func TestQrisReader(t *testing.T) {
	rawData := "00020101021226640013COM.MYWEB.WWW01181234567890123456780214123456789012340303UKE5405100005912QRIS WANTUNO6013Jakarta Pusat6304XXXX"
	qrisReader := reader.NewQrisReader(rawData)
	tag00, err := qrisReader.Read()
	if err != nil {
		t.Error("ERROR",*err)
		return
	}

	t.Log("SUCCESS ",tag00)
}
