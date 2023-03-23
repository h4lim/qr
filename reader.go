package reader

import (
	"strconv"
)

const (
	QRIS_VERSION_TAG00              = "00"
	QRIS_TYPE_TAG00                 = "01"
	QRIS_52_TAG00                   = "52"
	QRIS_53_TAG00                   = "53"
	QRIS_58_TAG00                   = "58"
	QRIS_61_TAG00                   = "61"
	QRIS_62_TAG00                   = "62"
	QRIS_AMOUNT_TAG00               = "54"
	QRIS_MERCHANT_OWNER_TAG00       = "59"
	QRIS_MERCHANT_ADDRESS_TAG00     = "60"
	QRIS_CHECKSUM_TAG00             = "63"
	QRIS_OWNER_TAG26                = "00"
	QRIS_MERCHANT_ID_TAG26          = "01"
	QRIS_MERCHANT_ACQUIRER_ID_TAG26 = "02"
	QRIS_MERCHANT_SCALE_TAG26       = "03"
	QRIS_WEB_TAG51                  = "00"
	QRIS_ID_TAG51                   = "02"
	QRIS_SCALE_TAG51                = "03"
)

type QrisTag struct {
	Tag00 Tag00 `json:"tag_00"`
	Tag26 Tag26 `json:"tag_26"`
	Tag51 Tag51 `json:"tag_51"`
}
type Tag00 struct {
	Version         string            `json:"version"`
	Type            string            `json:"type"`
	Tag52           string            `json:"tag_52"`
	Tag53           string            `json:"tag_53"`
	Tag58           string            `json:"tag_58"`
	Tag61           string            `json:"tag_61"`
	Tag62           string            `json:"tag_62"`
	Amount          int64             `json:"amount"`
	MerchantOwner   string            `json:"merchant_owner"`
	MerchantAddress string            `json:"merchant_address"`
	Checksum        string            `json:"checksum"`
	UnknownTag      map[string]string `json:"unknown_tag"`
}

type Tag26 struct {
	QrOwner            string            `json:"qr_owner"`
	MerchantID         string            `json:"merchant_id"`
	MerchantAcquirerID string            `json:"merchant_acquirer_id"`
	MerchantScale      string            `json:"merchant_scale"`
	UnknownTag         map[string]string `json:"unknown_tag"`
}

type Tag51 struct {
	QrisWeb    string            `json:"qris_web"`
	QrisID     string            `json:"qris_id"`
	Scale      string            `json:"scale"`
	UnknownTag map[string]string `json:"unknown_tag"`
}

type IQrReader interface {
	Read() (*QrisTag, *error)
}

type qrisReaderContext struct {
	rawData string
}

func NewQrisReader(rawData string) IQrReader {
	return qrisReaderContext{rawData: rawData}
}

func (q qrisReaderContext) Read() (*QrisTag, *error) {

	// get data from tag 00
	map00 := mapQris(q.rawData)

	// get tag 26 from map
	data26 := map00["26"]

	// get data from tag 26
	map26 := mapQris(data26[4:])

	// get tag 51 from map
	data51 := map00["51"]

	// get data from tag 51
	map51 := mapQris(data51[4:])

	// get all data qris tag
	qrisTag := mapData(map00, map26, map51)

	return &qrisTag, nil
}

func mapData(map00 map[string]string, map26 map[string]string, map51 map[string]string) QrisTag {

	var qrisTag QrisTag
	for k, v := range map00 {
		switch k {
		case QRIS_VERSION_TAG00:
			qrisTag.Tag00.Version = getData(v)
		case QRIS_TYPE_TAG00:
			qrisTag.Tag00.Type = getData(v)
		case QRIS_52_TAG00:
			qrisTag.Tag00.Tag52 = getData(v)
		case QRIS_53_TAG00:
			qrisTag.Tag00.Tag53 = getData(v)
		case QRIS_58_TAG00:
			qrisTag.Tag00.Tag58 = getData(v)
		case QRIS_61_TAG00:
			qrisTag.Tag00.Tag61 = getData(v)
		case QRIS_62_TAG00:
			qrisTag.Tag00.Tag62 = getData(v)
		case QRIS_AMOUNT_TAG00:
			amount, _ := strconv.ParseInt(getData(v), 10, 64)
			qrisTag.Tag00.Amount = amount
		case QRIS_MERCHANT_OWNER_TAG00:
			qrisTag.Tag00.MerchantOwner = getData(v)
		case QRIS_MERCHANT_ADDRESS_TAG00:
			qrisTag.Tag00.MerchantAddress = getData(v)
		case QRIS_CHECKSUM_TAG00:
			qrisTag.Tag00.Checksum = getData(v)
		default:
			mapUnknownTag := make(map[string]string)
			mapUnknownTag[k] = getData(v)
			qrisTag.Tag00.UnknownTag = mapUnknownTag
		}
	}

	for k, v := range map26 {
		switch k {
		case QRIS_OWNER_TAG26:
			qrisTag.Tag26.QrOwner = getData(v)
		case QRIS_MERCHANT_ID_TAG26:
			qrisTag.Tag26.MerchantID = getData(v)
		case QRIS_MERCHANT_ACQUIRER_ID_TAG26:
			qrisTag.Tag26.MerchantAcquirerID = getData(v)
		case QRIS_MERCHANT_SCALE_TAG26:
			qrisTag.Tag26.MerchantScale = getData(v)
		default:
			mapUnknownTag := make(map[string]string)
			mapUnknownTag[k] = getData(v)
			qrisTag.Tag00.UnknownTag = mapUnknownTag
		}
	}

	for k, v := range map51 {
		switch k {
		case QRIS_WEB_TAG51:
			qrisTag.Tag51.QrisWeb = getData(v)
		case QRIS_ID_TAG51:
			qrisTag.Tag51.QrisID = getData(v)
		case QRIS_SCALE_TAG51:
			qrisTag.Tag51.Scale = getData(v)
		default:
			mapUnknownTag := make(map[string]string)
			mapUnknownTag[k] = getData(v)
			qrisTag.Tag00.UnknownTag = mapUnknownTag
		}
	}

	return qrisTag
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
