package tempconv

import "fmt"

// 需要注意的是，一个文件夹内只能包含一个包（即包名字要一致，如果一个文件夹有两个包名，就会出现问题）
type Celsius float64
type Fahrenheit float64

const (
	AbosoluteZeroC Celsius = -273.15
	Freezing       Celsius = 0
	Boilling       Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g °C", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
