package types

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	hsp "github.com/CovenantSQL/HashStablePack/marshalhash"
)

// MarshalHash marshals for hash
func (z *Account) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 4
	o = append(o, 0x84, 0x84)
	o = hsp.AppendArrayHeader(o, uint32(SupportTokenNumber))
	for za0001 := range z.TokenBalance {
		o = hsp.AppendUint64(o, z.TokenBalance[za0001])
	}
	o = append(o, 0x84)
	o = hsp.AppendFloat64(o, z.Rating)
	o = append(o, 0x84)
	if oTemp, err := z.NextNonce.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x84)
	if oTemp, err := z.Address.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Account) Msgsize() (s int) {
	s = 1 + 13 + hsp.ArrayHeaderSize + (int(SupportTokenNumber) * (hsp.Uint64Size)) + 7 + hsp.Float64Size + 10 + z.NextNonce.Msgsize() + 8 + z.Address.Msgsize()
	return
}

// MarshalHash marshals for hash
func (z *MinerInfo) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 9
	o = append(o, 0x89, 0x89)
	o = hsp.AppendInt32(o, int32(z.Status))
	o = append(o, 0x89)
	o = hsp.AppendArrayHeader(o, uint32(len(z.UserArrears)))
	for za0001 := range z.UserArrears {
		if z.UserArrears[za0001] == nil {
			o = hsp.AppendNil(o)
		} else {
			// map header, size 2
			o = append(o, 0x82, 0x82)
			if oTemp, err := z.UserArrears[za0001].User.MarshalHash(); err != nil {
				return nil, err
			} else {
				o = hsp.AppendBytes(o, oTemp)
			}
			o = append(o, 0x82)
			o = hsp.AppendUint64(o, z.UserArrears[za0001].Arrears)
		}
	}
	o = append(o, 0x89)
	if oTemp, err := z.Address.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x89)
	if oTemp, err := z.NodeID.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x89)
	o = hsp.AppendString(o, z.Name)
	o = append(o, 0x89)
	o = hsp.AppendString(o, z.EncryptionKey)
	o = append(o, 0x89)
	o = hsp.AppendUint64(o, z.PendingIncome)
	o = append(o, 0x89)
	o = hsp.AppendUint64(o, z.ReceivedIncome)
	o = append(o, 0x89)
	o = hsp.AppendUint64(o, z.Deposit)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MinerInfo) Msgsize() (s int) {
	s = 1 + 7 + hsp.Int32Size + 12 + hsp.ArrayHeaderSize
	for za0001 := range z.UserArrears {
		if z.UserArrears[za0001] == nil {
			s += hsp.NilSize
		} else {
			s += 1 + 5 + z.UserArrears[za0001].User.Msgsize() + 8 + hsp.Uint64Size
		}
	}
	s += 8 + z.Address.Msgsize() + 7 + z.NodeID.Msgsize() + 5 + hsp.StringPrefixSize + len(z.Name) + 14 + hsp.StringPrefixSize + len(z.EncryptionKey) + 14 + hsp.Uint64Size + 15 + hsp.Uint64Size + 8 + hsp.Uint64Size
	return
}

// MarshalHash marshals for hash
func (z *ProviderProfile) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 9
	o = append(o, 0x89, 0x89)
	if oTemp, err := z.TokenType.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x89)
	o = hsp.AppendArrayHeader(o, uint32(len(z.TargetUser)))
	for za0001 := range z.TargetUser {
		if oTemp, err := z.TargetUser[za0001].MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x89)
	o = hsp.AppendFloat64(o, z.LoadAvgPerCPU)
	o = append(o, 0x89)
	if oTemp, err := z.Provider.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x89)
	if oTemp, err := z.NodeID.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x89)
	o = hsp.AppendUint64(o, z.Deposit)
	o = append(o, 0x89)
	o = hsp.AppendUint64(o, z.GasPrice)
	o = append(o, 0x89)
	o = hsp.AppendUint64(o, z.Space)
	o = append(o, 0x89)
	o = hsp.AppendUint64(o, z.Memory)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ProviderProfile) Msgsize() (s int) {
	s = 1 + 10 + z.TokenType.Msgsize() + 11 + hsp.ArrayHeaderSize
	for za0001 := range z.TargetUser {
		s += z.TargetUser[za0001].Msgsize()
	}
	s += 14 + hsp.Float64Size + 9 + z.Provider.Msgsize() + 7 + z.NodeID.Msgsize() + 8 + hsp.Uint64Size + 9 + hsp.Uint64Size + 6 + hsp.Uint64Size + 7 + hsp.Uint64Size
	return
}

// MarshalHash marshals for hash
func (z *SQLChainProfile) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 11
	o = append(o, 0x8b, 0x8b)
	if oTemp, err := z.Meta.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x8b)
	if oTemp, err := z.TokenType.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x8b)
	o = hsp.AppendArrayHeader(o, uint32(len(z.Miners)))
	for za0001 := range z.Miners {
		if z.Miners[za0001] == nil {
			o = hsp.AppendNil(o)
		} else {
			if oTemp, err := z.Miners[za0001].MarshalHash(); err != nil {
				return nil, err
			} else {
				o = hsp.AppendBytes(o, oTemp)
			}
		}
	}
	o = append(o, 0x8b)
	o = hsp.AppendArrayHeader(o, uint32(len(z.Users)))
	for za0002 := range z.Users {
		if z.Users[za0002] == nil {
			o = hsp.AppendNil(o)
		} else {
			if oTemp, err := z.Users[za0002].MarshalHash(); err != nil {
				return nil, err
			} else {
				o = hsp.AppendBytes(o, oTemp)
			}
		}
	}
	o = append(o, 0x8b)
	o = hsp.AppendBytes(o, z.EncodedGenesis)
	o = append(o, 0x8b)
	if oTemp, err := z.Owner.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x8b)
	if oTemp, err := z.Address.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x8b)
	if oTemp, err := z.ID.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x8b)
	o = hsp.AppendUint32(o, z.LastUpdatedHeight)
	o = append(o, 0x8b)
	o = hsp.AppendUint64(o, z.Period)
	o = append(o, 0x8b)
	o = hsp.AppendUint64(o, z.GasPrice)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SQLChainProfile) Msgsize() (s int) {
	s = 1 + 5 + z.Meta.Msgsize() + 10 + z.TokenType.Msgsize() + 7 + hsp.ArrayHeaderSize
	for za0001 := range z.Miners {
		if z.Miners[za0001] == nil {
			s += hsp.NilSize
		} else {
			s += z.Miners[za0001].Msgsize()
		}
	}
	s += 6 + hsp.ArrayHeaderSize
	for za0002 := range z.Users {
		if z.Users[za0002] == nil {
			s += hsp.NilSize
		} else {
			s += z.Users[za0002].Msgsize()
		}
	}
	s += 15 + hsp.BytesPrefixSize + len(z.EncodedGenesis) + 6 + z.Owner.Msgsize() + 8 + z.Address.Msgsize() + 3 + z.ID.Msgsize() + 18 + hsp.Uint32Size + 7 + hsp.Uint64Size + 9 + hsp.Uint64Size
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

// MarshalHash marshals for hash
func (z *SQLChainUser) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 6
	o = append(o, 0x86, 0x86)
	o = hsp.AppendInt32(o, int32(z.Status))
	o = append(o, 0x86)
	o = hsp.AppendInt32(o, int32(z.Permission))
	o = append(o, 0x86)
	if oTemp, err := z.Address.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x86)
	o = hsp.AppendUint64(o, z.AdvancePayment)
	o = append(o, 0x86)
	o = hsp.AppendUint64(o, z.Arrears)
	o = append(o, 0x86)
	o = hsp.AppendUint64(o, z.Deposit)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SQLChainUser) Msgsize() (s int) {
	s = 1 + 7 + hsp.Int32Size + 11 + hsp.Int32Size + 8 + z.Address.Msgsize() + 15 + hsp.Uint64Size + 8 + hsp.Uint64Size + 8 + hsp.Uint64Size
	return
}

// MarshalHash marshals for hash
func (z Status) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	o = hsp.AppendInt32(o, int32(z))
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Status) Msgsize() (s int) {
	s = hsp.Int32Size
	return
}

// MarshalHash marshals for hash
func (z *UserArrears) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 2
	o = append(o, 0x82, 0x82)
	if oTemp, err := z.User.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x82)
	o = hsp.AppendUint64(o, z.Arrears)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *UserArrears) Msgsize() (s int) {
	s = 1 + 5 + z.User.Msgsize() + 8 + hsp.Uint64Size
	return
}

// MarshalHash marshals for hash
func (z UserPermission) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	o = hsp.AppendInt32(o, int32(z))
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z UserPermission) Msgsize() (s int) {
	s = hsp.Int32Size
	return
}
