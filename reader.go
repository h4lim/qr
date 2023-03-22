package reader

import (
	"strconv"
)

const (
	QRIS_VERSION_TAG00          = "00"
	QRIS_TYPE_TAG00             = "01"
	QRIS_AMOUNT_TAG00           = "54"
	QRIS_MERCHANT_OWNER00       = "59"
	QRIS_MERCHANT_ADDRESS00     = "60"
	QRIS_CHECKSUM00             = "63"
	QRIS_OWNER26                = "00"
	QRIS_MERCHANT_ID26          = "01"
	QRIS_MERCHANT_ACQUIRER_ID26 = "02"
	QRIS_MERCHANT_SCALE26       = "03"
)

type Tag00 struct {
	Version         string
	Type            string
	Amount          int64
	MerchantOwner   string
	MerchantAddress string
	Checksum        string
	Tag26           Tag26
}

type Tag26 struct {
	QrOwner            string
	MerchantID         string
	MerchantAcquirerID string
	MerchantScale      string
}

type IQrReader interface {
	Read() (*Tag00, *error)
}

type qrisReaderContext struct {
	rawData string
}

func NewQrisReader(rawData string) IQrReader {
	return qrisReaderContext{rawData: rawData}
}

func (q qrisReaderContext) Read() (*Tag00, *error) {

	map00 := mapQris(q.rawData)
	data26 := map00["26"]
	map26 := mapQris(data26[4:])
	tag00 := mapData(map00, map26)
	return &tag00, nil
}

func mapData(map00 map[string]string, map26 map[string]string) Tag00 {

	var tag00 Tag00
	for k, v := range map00 {
		switch k {
		case QRIS_VERSION_TAG00:
			tag00.Version = getData(v)
		case QRIS_TYPE_TAG00:
			tag00.Type = getData(v)
		case QRIS_AMOUNT_TAG00:
			amount, _ := strconv.ParseInt(getData(v), 10, 64)
			tag00.Amount = amount
		case QRIS_MERCHANT_OWNER00:
			tag00.MerchantOwner = getData(v)
		case QRIS_MERCHANT_ADDRESS00:
			tag00.MerchantAddress = getData(v)
		case QRIS_CHECKSUM00:
			tag00.Checksum = getData(v)
		}
	}

	for k, v := range map26 {
		switch k {
		case QRIS_OWNER26:
			tag00.Tag26.QrOwner = getData(v)
		case QRIS_MERCHANT_ID26:
			tag00.Tag26.MerchantID = getData(v)
		case QRIS_MERCHANT_ACQUIRER_ID26:
			tag00.Tag26.MerchantAcquirerID = getData(v)
		case QRIS_MERCHANT_SCALE26:
			tag00.Tag26.MerchantScale = getData(v)
		}
	}

	return tag00
}

func mapQris(rawData string) map[string]string {

	n := 0
	mapData := make(map[string]string)
	for n < len(rawData) {
		key := rawData[n : n+2]
		lengthData := getLengthData(rawData[n : n+4])
		indexValue := n + 4 + lengthData
		mapData[key] = rawData[n:indexValue]
		n = indexValue
	}
	return mapData
}

func getData(data string) string {
	return data[4:]
}

func getLengthData(data string) int {
	lengthData, _ := strconv.Atoi(data[2:])

	return lengthData
}

