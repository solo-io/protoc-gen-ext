package clone

import (
	"github.com/golang/protobuf/proto"

)

type Cloner interface {
	Clone() proto.Message
}
