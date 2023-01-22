package contracts

type Output interface {
	Print(text string) error
	Flush() error
}
