package types

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	hsp "github.com/CovenantSQL/HashStablePack/marshalhash"
)

// MarshalHash marshals for hash
func (z *Account) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 6
	o = append(o, 0x86, 0x86)
	o = hsp.AppendArrayHeader(o, uint32(len(z.Profiles)))
	for za0001 := range z.Profiles {
		if z.Profiles[za0001] == nil {
			o = hsp.AppendNil(o)
		} else {
			if oTemp, err := z.Profiles[za0001].MarshalHash(); err != nil {
				return nil, err
			} else {
				o = hsp.AppendBytes(o, oTemp)
			}
		}
	}
	o = append(o, 0x86)
	o = hsp.AppendArrayHeader(o, uint32(len(z.TxBillings)))
	for za0002 := range z.TxBillings {
		if z.TxBillings[za0002] == nil {
			o = hsp.AppendNil(o)
		} else {
			if oTemp, err := z.TxBillings[za0002].MarshalHash(); err != nil {
				return nil, err
			} else {
				o = hsp.AppendBytes(o, oTemp)
			}
		}
	}
	o = append(o, 0x86)
	o = hsp.AppendFloat64(o, z.Rating)
	o = append(o, 0x86)
	if oTemp, err := z.Address.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x86)
	o = hsp.AppendUint64(o, z.StableCoinBalance)
	o = append(o, 0x86)
	o = hsp.AppendUint64(o, z.CovenantCoinBalance)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Account) Msgsize() (s int) {
	s = 1 + 9 + hsp.ArrayHeaderSize
	for za0001 := range z.Profiles {
		if z.Profiles[za0001] == nil {
			s += hsp.NilSize
		} else {
			s += z.Profiles[za0001].Msgsize()
		}
	}
	s += 11 + hsp.ArrayHeaderSize
	for za0002 := range z.TxBillings {
		if z.TxBillings[za0002] == nil {
			s += hsp.NilSize
		} else {
			s += z.TxBillings[za0002].Msgsize()
		}
	}
	s += 7 + hsp.Float64Size + 8 + z.Address.Msgsize() + 18 + hsp.Uint64Size + 19 + hsp.Uint64Size
	return
}

// MarshalHash marshals for hash
func (z *SQLChainProfile) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 3
	o = append(o, 0x83, 0x83)
	o = hsp.AppendByte(o, byte(z.Role))
	o = append(o, 0x83)
	if oTemp, err := z.ID.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x83)
	o = hsp.AppendUint64(o, z.Deposit)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SQLChainProfile) Msgsize() (s int) {
	s = 1 + 5 + hsp.ByteSize + 3 + z.ID.Msgsize() + 8 + hsp.Uint64Size
	return
}

// MarshalHash marshals for hash
func (z SQLChainRole) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	o = hsp.AppendByte(o, byte(z))
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z SQLChainRole) Msgsize() (s int) {
	s = hsp.ByteSize
	return
}
