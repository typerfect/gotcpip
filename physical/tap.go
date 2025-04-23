package physical

const (
	tapName = "tap0"

)
var (
	GlobalTapIO * tapIO
)

type tapIO struct {

}




func(tap * tapIO)	Write([]byte)(int, error) {

	return 0, nil
}

func(tap * tapIO) Read()([]byte, error){

	return nil, nil
}
func(tap * tapIO) getInput()(chan []byte, error){
	return nil, nil

}

