package services


type LinterService struct {}

func NewLinterService() *LinterService {
	return &LinterService{}
}

func (ls *LinterService) Lint() {
	panic("")
}