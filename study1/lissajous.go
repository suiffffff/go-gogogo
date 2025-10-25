package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//？？？这是人类啊？？？
//你是不是有点太看得起我的数学了？

// 这个不用查都大概知道是引入颜色
var palette = []color.Color{color.White, color.Black}

// 这个是c的define，不过c的const也可以固定指针
const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer) {
	const (
		//周期
		cycles = 5
		//分辨率
		res = 0.001
		//大小，也是像素
		size = 100
		//帧数
		nframes = 64
		//延迟时间，1表示10ms
		delay = 8
	)
	//你不应该期待一个人能看得懂这么抽象的函数
	//随机生成一个帧率
	freq := rand.Float64() * 3.0
	//创建一个gif，循环表示帧率
	anim := gif.GIF{LoopCount: nframes}
	//相位
	phase := 0.0
	//自己了解喵，代码学数学吗？
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	//输出到标准输出
	gif.EncodeAll(out, &anim)
}
func main() {
	lissajous(os.Stdout)
}
