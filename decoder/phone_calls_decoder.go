package decoder

import (
	"bytes"
	"encoding/csv"
	"github.com/jszwec/csvutil"
	"time"
)

func GetDecoder(csvInput []byte) (*csvutil.Decoder, error) {
	dec, err := csvutil.NewDecoder(csv.NewReader(bytes.NewReader(csvInput)))
	if err != nil {
		return nil, err
	}
	return dec, nil
}

func SetDecoderTimeFormat(dec *csvutil.Decoder, dateFormat string) *csvutil.Decoder {
	unmarshalTime := func(data []byte, t *time.Time) error {
		tt, err := time.Parse(dateFormat, string(data))
		if err != nil {
			return err
		}
		*t = tt
		return nil
	}
	dec.Register(unmarshalTime)
	return dec
}
