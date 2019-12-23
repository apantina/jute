// Autogenerated jute compiler
// @generated from '/home/bbennett/src/jute/testdata/zookeeper.jute'

package txn // github.com/go-zookeeper/zk/internal/txn

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type Txn struct {
	Type int32  // type
	Data []byte // data
}

func (r *Txn) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.Type, err = dec.ReadInt()
	if err != nil {
		return err
	}
	r.Data, err = dec.ReadBuffer()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *Txn) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteInt(r.Type); err != nil {
		return err
	}
	if err := enc.WriteBuffer(r.Data); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *Txn) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Txn(%+v)", *r)
}