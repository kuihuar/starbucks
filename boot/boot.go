package boot

import(
	"starbucks/configs"
)

func init(){
	configs.InitApp()
	configs.InitConf()
	configs.InitResource()
}