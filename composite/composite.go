package composite

import "fmt"

type component interface {
	search(string)
}

type folder struct {
	components []component
	name string
}
func (f *folder)search(keyword string) {
	fmt.Printf("Searchig recurisive for keyword %s in folder %s\n", keyword, f.name)
	for _,composite := range f.components {
		composite.search(keyword)
	}
}

func (f *folder)add (c component) {
	f.components= append(f.components,c)
}
type file struct {
	name string
}
func(f *file) search(keyword string) {
	fmt.Printf("Searchig for keyword %s in file %s\n", keyword, f.name)
}
func (f *file)getName()string {
	return f.name
}