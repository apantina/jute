// Autogenerated jute compiler
// @generated from '/home/bbennett/src/jute/testdata/zookeeper.jute'

package proto // github.com/go-zookeeper/zk/internal/proto

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type GetEphemeralsRequest struct {
	PrefixPath *string // prefixPath
}

func (r *GetEphemeralsRequest) GetPrefixPath() string {
	if r != nil && r.PrefixPath != nil {
		return *r.PrefixPath
	}
	return ""
}

func (r *GetEphemeralsRequest) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.PrefixPath, err = dec.ReadString()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *GetEphemeralsRequest) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteString(r.PrefixPath); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *GetEphemeralsRequest) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetEphemeralsRequest(%+v)", *r)
}
