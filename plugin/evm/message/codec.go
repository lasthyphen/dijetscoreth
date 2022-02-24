// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"github.com/lasthyphen/dijetsnetgo1.2/codec"
	"github.com/lasthyphen/dijetsnetgo1.2/codec/linearcodec"
	"github.com/lasthyphen/dijetsnetgo1.2/codec/reflectcodec"
	"github.com/lasthyphen/dijetsnetgo1.2/utils/units"
	"github.com/lasthyphen/dijetsnetgo1.2/utils/wrappers"
)

const (
	codecVersion   uint16 = 0
	maxMessageSize        = 512 * units.KiB
	maxSliceLen           = maxMessageSize
)

// Codec does serialization and deserialization
var c codec.Manager

func init() {
	c = codec.NewManager(maxMessageSize)
	lc := linearcodec.New(reflectcodec.DefaultTagName, maxSliceLen)

	errs := wrappers.Errs{}
	errs.Add(
		lc.RegisterType(&AtomicTx{}),
		lc.RegisterType(&EthTxs{}),
		c.RegisterCodec(codecVersion, lc),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
}
