package main

import (
    "fmt"
    "func_type/model"
)

func main() {

    fmt.Println(model.AnyAdd(func(i int, i2 int) int {
        return i - i2
    }))

    str := model.FuncStruct{MyAdder: Add}
    fmt.Println(str.MyAdder(1, 2))

}

func Add(i int, int2 int) int {
    return 1
}
