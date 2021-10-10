package nullobject

import "fmt"

type department interface {
	getNumberOfprofessors()int
	getName()string
}

type computerscience struct {
	numberOfprofessors int
	}
	func(c *computerscience) getNumberOfprofessors()int {
	return c.numberOfprofessors
	}
	func(c *computerscience) getName()string {
	return "computerscience"
	}

type mechanical struct {
	numberOfprofessors int
}
func(c *mechanical) getNumberOfprofessors()int {
	return c.numberOfprofessors
}
func(c *mechanical) getName()string {
	return "mechanical"
}
type college struct {
	departmemts []department
}
type nullDepartment struct {
	numberOfprofessors int
}
func(c *nullDepartment) getNumberOfprofessors()int {
	return 0
}
func(c *nullDepartment) getName()string {
	return "nullDepartment"
}

func (c *college)addDepartment(departmentName string, numofProcessors int)  {
	if departmentName == "computerscience" {
		computerscienceDepartment := &computerscience{numberOfprofessors: numofProcessors}
		c.departmemts = append(c.departmemts, computerscienceDepartment)
	}
	if departmentName == "mechanical" {
		mechanicalDepartment := &mechanical{numberOfprofessors: numofProcessors}
		c.departmemts = append(c.departmemts, mechanicalDepartment)
	}
	return
}
func (c *college)getDepartment(departmentName string) department {
	for _, department := range c.departmemts {
		if department.getName() == departmentName {
			return department
		}
	}
	return &nullDepartment{}
}

func TestNullObj()  {
	college1 :=createCollege1()
	college2 :=createCollege2()
	totalProfessors := 0
	departmentSlice := []string{"computerscience", "mechanical", "civil"}
	for _, departmentname := range departmentSlice {
		d := college1.getDepartment(departmentname)
		totalProfessors += d.getNumberOfprofessors()
	}
	fmt.Printf("Tocal number of professors in college1 is %d\n",totalProfessors)
	totalProfessors = 0
	for _, departmentname := range departmentSlice {
		d := college2.getDepartment(departmentname)
		totalProfessors += d.getNumberOfprofessors()
	}
	fmt.Printf("Tocal number of professors in college2 is %d\n",totalProfessors)
}

func createCollege1() *college {
	college:= &college{}
	college.addDepartment("computerscience", 4)
	college.addDepartment("mechanical", 6)
	return college
}
func createCollege2() *college {
	college:= &college{}
	college.addDepartment("computerscience", 2)
	return college
}