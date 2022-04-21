package service

type pingServiceImpl struct {
}

func NewPingService() pingServiceImpl {
	return pingServiceImpl{}
}

func (ps *pingServiceImpl) GetPingData() string {
	return "pong"
}
