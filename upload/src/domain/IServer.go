package domain

type IServer interface {
	Run() error
	Quit()
}
