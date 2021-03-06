// Code generated by github.com/spacemeshos/go-scale/gen. DO NOT EDIT.

package types

import (
	"github.com/spacemeshos/go-scale"
)

func (t *SelfSpawn) EncodeScale(enc *scale.Encoder) (total int, err error) {
	// field Type (0)
	if n, err := scale.EncodeCompact8(enc, t.Type); err != nil {
		return total, err
	} else {
		total += n
	}

	// field Body (1)
	if n, err := t.Body.EncodeScale(enc); err != nil {
		return total, err
	} else {
		total += n
	}

	return total, nil
}

func (t *SelfSpawn) DecodeScale(dec *scale.Decoder) (total int, err error) {
	// field Type (0)
	if field, n, err := scale.DecodeCompact8(dec); err != nil {
		return total, err
	} else {
		total += n
		t.Type = field
	}

	// field Body (1)
	if n, err := t.Body.DecodeScale(dec); err != nil {
		return total, err
	} else {
		total += n
	}

	return total, nil
}

func (t *SelfSpawnBody) EncodeScale(enc *scale.Encoder) (total int, err error) {
	// field Address (0)
	if n, err := t.Address.EncodeScale(enc); err != nil {
		return total, err
	} else {
		total += n
	}

	// field Selector (1)
	if n, err := scale.EncodeCompact8(enc, t.Selector); err != nil {
		return total, err
	} else {
		total += n
	}

	// field Payload (2)
	if n, err := t.Payload.EncodeScale(enc); err != nil {
		return total, err
	} else {
		total += n
	}

	return total, nil
}

func (t *SelfSpawnBody) DecodeScale(dec *scale.Decoder) (total int, err error) {
	// field Address (0)
	if n, err := t.Address.DecodeScale(dec); err != nil {
		return total, err
	} else {
		total += n
	}

	// field Selector (1)
	if field, n, err := scale.DecodeCompact8(dec); err != nil {
		return total, err
	} else {
		total += n
		t.Selector = field
	}

	// field Payload (2)
	if n, err := t.Payload.DecodeScale(dec); err != nil {
		return total, err
	} else {
		total += n
	}

	return total, nil
}

func (t *SelfSpawnPayload) EncodeScale(enc *scale.Encoder) (total int, err error) {
	// field Template (0)
	if n, err := t.Template.EncodeScale(enc); err != nil {
		return total, err
	} else {
		total += n
	}

	// field Arguments (1)
	if n, err := t.Arguments.EncodeScale(enc); err != nil {
		return total, err
	} else {
		total += n
	}

	// field GasPrice (2)
	if n, err := scale.EncodeCompact32(enc, t.GasPrice); err != nil {
		return total, err
	} else {
		total += n
	}

	// field Signature (3)
	if n, err := t.Signature.EncodeScale(enc); err != nil {
		return total, err
	} else {
		total += n
	}

	return total, nil
}

func (t *SelfSpawnPayload) DecodeScale(dec *scale.Decoder) (total int, err error) {
	// field Template (0)
	if n, err := t.Template.DecodeScale(dec); err != nil {
		return total, err
	} else {
		total += n
	}

	// field Arguments (1)
	if n, err := t.Arguments.DecodeScale(dec); err != nil {
		return total, err
	} else {
		total += n
	}

	// field GasPrice (2)
	if field, n, err := scale.DecodeCompact32(dec); err != nil {
		return total, err
	} else {
		total += n
		t.GasPrice = field
	}

	// field Signature (3)
	if n, err := t.Signature.DecodeScale(dec); err != nil {
		return total, err
	} else {
		total += n
	}

	return total, nil
}

func (t *SelfSpawnArguments) EncodeScale(enc *scale.Encoder) (total int, err error) {
	// field Key (0)
	if n, err := t.Key.EncodeScale(enc); err != nil {
		return total, err
	} else {
		total += n
	}

	return total, nil
}

func (t *SelfSpawnArguments) DecodeScale(dec *scale.Decoder) (total int, err error) {
	// field Key (0)
	if n, err := t.Key.DecodeScale(dec); err != nil {
		return total, err
	} else {
		total += n
	}

	return total, nil
}
