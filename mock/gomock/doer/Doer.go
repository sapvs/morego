package doer

//go:generate mockgen -source=$GOFILE -destination=$PWD/mocks/${GOFILE} -package=mocks

type Doer interface {
	Do(int, string) error
}
