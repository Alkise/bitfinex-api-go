package summary

import (
	"fmt"

	"github.com/bitfinexcom/bitfinex-api-go/pkg/convert"
)

type Summary struct {
	MakerFee         float64
	DerivRebate      float64
	TakerFeeToCrypto float64
	TakerFeeToStable float64
	TakerFeeToFiat   float64
	DerivTakerFee    float64
}

func FromRaw(raw []interface{}) (s *Summary, err error) {
	if len(raw) < 9 {
		return s, fmt.Errorf("data slice too short for summary: %#v", raw)
	}

	if raw[4] == nil {
		return
	}

	fraw := raw[4].([]interface{})

	if len(fraw) < 2 {
		return s, fmt.Errorf("data slice too for fee info: %#v", fraw)
	}

	mraw := fraw[0].([]interface{})

	if len(mraw) < 6 {
		return s, fmt.Errorf("data slice too for maker fee: %#v", mraw)
	}

	traw := fraw[1].([]interface{})

	if len(traw) < 6 {
		return s, fmt.Errorf("data slice too for taker fee: %#v", traw)
	}

	s = &Summary{
		MakerFee:         convert.F64ValOrZero(mraw[0]),
		DerivRebate:      convert.F64ValOrZero(mraw[5]),
		TakerFeeToCrypto: convert.F64ValOrZero(traw[0]),
		TakerFeeToStable: convert.F64ValOrZero(traw[1]),
		TakerFeeToFiat:   convert.F64ValOrZero(traw[2]),
		DerivTakerFee:    convert.F64ValOrZero(traw[5]),
	}

	return
}
