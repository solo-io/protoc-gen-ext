package clone

import (
	"google.golang.org/protobuf/proto"
)

type Cloner interface {
	Clone() proto.Message
}
