package rpc

import (
	"encoding/json"
	"math/big"
)

type Int struct {
	*big.Int
}

func (num *Int) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		var value big.Int
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}

		num.Int = &value
		return nil
	}

	var value int64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	num.Int = big.NewInt(value)
	return nil
}
