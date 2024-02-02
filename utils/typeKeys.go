package utils

import (
	"context"
	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/bitmap"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// KeyPresser 模拟键盘输入
type KeyPresser struct {
	mu sync.Mutex
}

// NewKeyPresser 创建一个 KeyPresser
func NewKeyPresser() *KeyPresser {
	return &KeyPresser{}
}

// KeyTap 模拟键盘输入
func (k *KeyPresser) KeyTap(key string, sleep ...int) {
	k.mu.Lock()
	defer k.mu.Unlock()
	_ = robotgo.KeyDown(key)
	if sleep != nil && len(sleep) > 0 && sleep[0] > 0 {
		log.Println("按键独立 sleep:", sleep[0])
		robotgo.MilliSleep(sleep[0])
	} else {
		robotgo.MilliSleep(BTime)
	}
	_ = robotgo.KeyUp(key)
}

// AtomicBool 原子布尔值
type AtomicBool struct {
	val        int32
	cancelFunc context.CancelFunc
}

func (b *AtomicBool) Set(value bool) {
	var i int32 = 0
	if value {
		i = 1
	}
	atomic.StoreInt32(&(b.val), i)
}

func (b *AtomicBool) Get() bool {
	return atomic.LoadInt32(&(b.val)) != 0
}

// AfterFalse 以定时器设定 AtomicBool 的值，并在指定时长后自动清除
func (b *AtomicBool) AfterFalse(duration time.Duration) {
	b.Set(true)

	// 如果已经有一个定时器在运行，取消它
	if b.cancelFunc != nil {
		b.cancelFunc()
	}

	ctx, cancel := context.WithCancel(context.Background())
	b.cancelFunc = cancel

	go func() {
		select {
		case <-time.After(duration):
			b.Set(false)
		case <-ctx.Done():
			// 如果收到取消信号，不做任何事情
			return
		}
	}()
}

// typeKeys 模拟键盘输入
func typeKeys(status *bool) {
	if cname, err := GetActiveWindowClassName(); err == nil {
		if cname == "LaunchUnrealUWindowsClient" {

		} else {
			log.Println("NO")
			return
		}
	} else {
		log.Println(err.Error())
		return
	}
	// 一些运行参数信息
	var (
		isLj          AtomicBool
		isYs          AtomicBool
		isXY          AtomicBool
		isCuiDuBiShou AtomicBool
		isYB          AtomicBool
		isZD          AtomicBool
		isZL          AtomicBool
		isGl          AtomicBool
		isBc          AtomicBool
		isSS          AtomicBool
		num           int
	)
	keyPresser := NewKeyPresser()
	for {
		var start, end time.Time
		var fx, fy int
		if !*status {
			return
		}
		num++
		log.Println("num:", num)
		x, y, w, h := robotgo.GetBounds(robotgo.GetPid())
		if w == 0 || h == 0 {
			log.Println("应用基础数据获取异常......")
			continue
		}
		bit := robotgo.CaptureScreen(int(float64(x)*Scale+float64(w)*Scale/3), int(float64(y)*Scale+float64(h)*Scale)/3*2, int(float64(w)*Scale)/3, int(float64(h)*Scale)/3)
		bitTop := robotgo.CaptureScreen(int(float64(x)*Scale+float64(w)*Scale/3), 0, int(float64(w)*Scale)/3, int(float64(h)*Scale)/3)
		if bitTop == nil || bit == nil {
			log.Println("截图失败，跳过本次循环......")
			continue
		}
		allStart := time.Now()
		// 检测是否可以隐身
		start = time.Now()
		ys := robotgo.ImgToCBitmap(YsTime)
		fx, fy = bitmap.Find(ys, bit, Tolerance)
		if fx != -1 && fy != -1 {
			isYs.Set(true)
			log.Println("正在隐身中")
		} else {
			isYs.Set(false)
			log.Println("结束隐身")
		}
		end = time.Now()
		log.Printf("隐身检测：%v\n", end.Sub(start))
		// 检测是否可以雷决
		start = time.Now()
		l := robotgo.ImgToCBitmap(LJ)
		fx, fy = bitmap.Find(l, bit, Tolerance)
		if fx != -1 && fy != -1 {
			isLj.Set(true)
			log.Println("雷决启动ForBit")
		} else {
			isLj.Set(false)
		}
		end = time.Now()
		log.Printf("雷决检测：%v\n", end.Sub(start))
		// 检测是否可以使用影匕
		start = time.Now()
		yb := robotgo.ImgToCBitmap(YB)
		fx, fy = bitmap.Find(yb, bit, Tolerance)
		if fx != -1 && fy != -1 {
			isYB.Set(true)
			log.Println("可以影匕ForBit")
		} else {
			isYB.Set(false)
		}
		end = time.Now()
		log.Printf("影匕检测：%v\n", end.Sub(start))
		// 检测是否可以掷毒
		start = time.Now()
		zd := robotgo.ImgToCBitmap(ZD)
		fx, fy = bitmap.Find(zd, bit, Tolerance)
		if fx != -1 && fy != -1 {
			isZD.Set(true)
			log.Println("可以掷毒ForBit")
		} else {
			isZD.Set(false)
		}
		end = time.Now()
		log.Printf("掷毒检测：%v\n", end.Sub(start))
		// 检测是否可以掷毒雷
		start = time.Now()
		zl := robotgo.ImgToCBitmap(ZL)
		fx, fy = bitmap.Find(zl, bit, Tolerance)
		if fx != -1 && fy != -1 {
			isZL.Set(true)
			log.Println("可以掷毒雷ForBit")
		} else {
			isZL.Set(false)
		}
		end = time.Now()
		log.Printf("掷毒雷检测：%v\n", end.Sub(start))
		// 检测是否可以淬毒匕首
		start = time.Now()
		cdbs := robotgo.ImgToCBitmap(CDBS)
		fx, fy = bitmap.Find(cdbs, bit, Tolerance)
		if fx != -1 && fy != -1 {
			isCuiDuBiShou.Set(true)
			log.Println("可以淬毒匕首ForBit")
		} else {
			isCuiDuBiShou.Set(false)
		}
		end = time.Now()
		log.Printf("淬毒匕首检测：%v\n", end.Sub(start))
		// 检测是否可以吸影
		start = time.Now()
		xy := robotgo.ImgToCBitmap(XY)
		fx, fy = bitmap.Find(xy, bit, Tolerance)
		if fx != -1 && fy != -1 {
			isXY.Set(true)
			log.Println("检测到吸影可使用")
		} else {
			isXY.Set(false)
		}
		end = time.Now()
		log.Printf("吸影检测：%v\n", end.Sub(start))
		// 检测是否可以背刺
		start = time.Now()
		bc := robotgo.ImgToCBitmap(B)
		fx, fy = bitmap.Find(bc, bit, Tolerance)
		if fx != -1 && fy != -1 {
			isBc.Set(true)
			log.Println("检测到背刺")
		} else {
			isBc.Set(false)
		}
		end = time.Now()
		log.Printf("背刺检测：%v\n", end.Sub(start))
		allEnd := time.Now()
		log.Printf("检测耗时：%v\n", allEnd.Sub(allStart))
		if isYs.Get() && !isGl.Get() {
			robotgo.MilliSleep(BTime)
			keyPresser.KeyTap(robotgo.Key2)
			log.Println("挂雷成功")
			isGl.AfterFalse(time.Second * 5)
		}
		if isYs.Get() {
			if isZL.Get() {
				keyPresser.KeyTap(robotgo.KeyZ, 60)
				log.Println("掷毒雷启动ForBit")
				printTime(allStart)
				continue
			} else {
				if isYB.Get() {
					keyPresser.KeyTap(robotgo.KeyX, 80)
					log.Println("影匕启动ForBit")
					printTime(allStart)
					continue
				}
			}
		}
		if isZD.Get() {
			keyPresser.KeyTap(robotgo.Key4)
			isZD.Set(false)
			log.Println("掷毒启动ForBit")
		}
		log.Println("隐身状态", isYs.Get(), "影匕状态：", isYB.Get(), "吸影状态：", isXY.Get(), "背刺状态：", isBc.Get(), "SS状态：", isSS.Get(), "掷毒状态：", isZD.Get(), "掷毒雷状态：", isZL.Get(), "淬毒匕首状态：", isCuiDuBiShou.Get(), "雷决状态：", isLj.Get())
		if !isYs.Get() && isXY.Get() {
			isXY.AfterFalse(time.Millisecond * 200)
			keyPresser.KeyTap(robotgo.Key1, 80)
			log.Println("吸影成功")
			robotgo.MilliSleep(BTime)
			printTime(allStart)
			continue
		}
		if !isYs.Get() && !isXY.Get() && !isSS.Get() {
			log.Printf("SS执行前：%v\n", isSS.Get())
			isSS.AfterFalse(time.Millisecond * 500)
			keyPresser.KeyTap(robotgo.KeyS, 30)
			robotgo.MilliSleep(90)
			keyPresser.KeyTap(robotgo.KeyS, 30)
			log.Println("SS执行成功")
			printTime(allStart)
			continue
		}
		if isSS.Get() && !isYs.Get() && !isXY.Get() {
			keyPresser.KeyTap(robotgo.Key1, 80)
			log.Println("SS后潜行执行成功")
			printTime(allStart)
			continue
		}
		if isLj.Get() {
			keyPresser.KeyTap(robotgo.Key4, 50)
			isLj.Set(false)
			log.Println("雷决启动ForBit")
			robotgo.MilliSleep(BTime)
			printTime(allStart)
			continue
		}
		if isCuiDuBiShou.Get() {
			keyPresser.KeyTap(robotgo.KeyX)
			isCuiDuBiShou.Set(false)
			log.Println("淬毒匕首启动")
			printTime(allStart)
			continue
		}
		if isYs.Get() && isBc.Get() && !isXY.Get() {
			keyPresser.KeyTap(robotgo.KeyR, 80)
			log.Println("背刺启动ForBit")
			robotgo.MilliSleep(BTime)
			printTime(allStart)
			continue
		}
		keyPresser.KeyTap(T)
		keyPresser.KeyTap(F)
		//销毁数据
		robotgo.FreeBitmap(bit)
		robotgo.FreeBitmap(bitTop)
		printTime(allStart)
	}
}

// printTime 打印耗时
func printTime(allStart time.Time) {
	allEnd := time.Now()
	log.Printf("单次循环耗时：%v\n", allEnd.Sub(allStart))
	log.Println("----------------------")
}
