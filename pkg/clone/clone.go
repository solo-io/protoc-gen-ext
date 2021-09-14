package clone

import (
	"google.golang.org/protobuf/proto"
)

type Cloner interface {
	Clone(object proto.Message) proto.Message
}
