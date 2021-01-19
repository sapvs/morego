package ds

import "io"

//DataStructure collection of data nodes/ values
type DataStructure interface {
	//Visualize prints to the writer in good enough format
	Visualize(ds DataStructure, w *io.Writer)
	//Add adds value to the DataStructure
	Add()
	//Remove removes supplied value or from the index from the DataStructure
	Remove(value interface{}, index int)
	//Size get the size of the DataStructure
	Size() int
}
