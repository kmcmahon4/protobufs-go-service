package pkg

type ProtobufGoService interface {
	// declare functions here

}

type protobufGoService struct {
}

var _ ProtobufGoService = (*protobufGoService)(nil)

func NewProtobufGoService() ProtobufGoService {
	return &protobufGoService{}
}

// function declarations
// do things
