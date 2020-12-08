package model

type MyAdd func(int, int) int

//Uses the type MyAdd
func AnyAdd(myAdd MyAdd) int {
    return myAdd(2, 3)
}

type FuncStruct struct {
    MyAdder MyAdd
}
