// Autogenerated jute compiler
// @generated from '/home/bbennett/src/jute/testdata/zookeeper.jute'

package proto // github.com/go-zookeeper/zk/internal/proto

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
	"github.com/go-zookeeper/zk/internal/data"
)

type Create2Response struct {
	Path *string    // path
	Stat *data.Stat // stat
}

func (r *Create2Response) GetPath() string {
	if r != nil && r.Path != nil {
		return *r.Path
	}
	return ""
}

func (r *Create2Response) GetStat() *data.Stat {
	if r != nil && r.Stat != nil {
		return r.Stat
	}
	return nil
}

func (r *Create2Response) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.Path, err = dec.ReadString()
	if err != nil {
		return err
	}
	if err = dec.ReadRecord(r.Stat); err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *Create2Response) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteString(r.Path); err != nil {
		return err
	}
	if err := enc.WriteRecord(r.Stat); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *Create2Response) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Create2Response(%+v)", *r)
}
