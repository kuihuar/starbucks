package factorygun

import "fmt"

func GetGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "moverick" {
		return newMoverick(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
func PrintDetails(g iGun)  {
	fmt.Printf("Gun: %s\n", g.GetName())
	fmt.Printf("Power: %d\n", g.GetPower())
}
