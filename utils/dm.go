package utils

import (
	"fmt"
	"github.com/redzl/go-dmsoft"
	"log"
	"os"
	"time"
	"unsafe"
)

const (
	// 填写自己的注册码
	DmRegCode   = ""
	DmExtraCode = ""
)

var dm *dmsoft.Dmsoft

type _bytes struct {
	Data int
	Len  int
}

func init() {
	dm = CreateDmObj()
	log.Printf("插件版本:%s", dm.Ver())
	dm.Reg(DmRegCode, DmExtraCode)
}

func FindZd() {

	var data, size int
	dm.GetScreenDataBmp(0, 0, 3840, 2160, &data, &size)
	bs := *(*[]byte)(unsafe.Pointer(&_bytes{
		data,
		size,
	}))
	os.WriteFile("test.bmp", bs, os.ModePerm)
	start := time.Now()
	// x，y接收返回的坐标
	var x, y int
	ret := dm.FindPic(0, 0, 3840, 2160, "./static/ZD.png", "393400", 0.95, 0, &x, &y)
	if ret != -1 {
		dm.MoveTo(x, y)
	}
	fmt.Println(x, y)
	end := time.Now()
	log.Printf("雷决检测测试：%v2d2\n", end.Sub(start))
}

func CreateDmObj() *dmsoft.Dmsoft {
	// 获取当前工作目录
	dir, _ := os.Getwd()
	// 设置dm.dll路径,并进行注册
	ret := dmsoft.SetDllPathW(dir+"\\dll\\dm.dll", 0)
	if ret {
		log.Println("插件注册成功！")
	} else {
		log.Println("插件注册失败！")
	}
	return dmsoft.NewDmsoft()
}
