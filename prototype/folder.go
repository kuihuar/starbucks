package prototype

import "fmt"

type Folder struct {
	Childerens []Inode
	Name string
}

func (f *Folder)Print(indentation string)  {
	fmt.Println(indentation + f.Name)
	for _, i := range f.Childerens {
		i.Print(indentation + indentation)
	}
}
func (f *Folder)Clone() Inode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}

	var tempChilders []Inode

	for _, i := range f.Childerens {
		copy := i.Clone()
		tempChilders = append(tempChilders,copy)
	}
	cloneFolder.Childerens = tempChilders

	return cloneFolder
}