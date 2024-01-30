package utils

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/bitmap"
	"time"

	_ "image/png"

	"github.com/deluan/lookup"
)

// Helper function to load an image from the filesystem

func Find() {
	// Load full image
	img := loadImage("./test.png")
	// Create a lookup for that image
	l := lookup.NewLookup(img)

	// Load a template to search inside the image
	template := loadImage("./static/Y.png")
	start := time.Now()
	// Find all occurrences of the template in the image
	pp, _ := l.FindAll(template, 0.9)
	end := time.Now()
	fmt.Printf("耗时：%v\n", end.Sub(start))
	// Print the results
	if len(pp) > 0 {
		fmt.Printf("Found %d matches:\n", len(pp))
		for _, p := range pp {
			fmt.Printf("- (%d, %d) with %f accuracy\n", p.X, p.Y, p.G)
		}
	} else {
		println("No matches found")
	}
	start = time.Now()
	bitImg := loadImage("./test.png")
	fx, fy := bitmap.Find(robotgo.ImgToCBitmap(loadImage("./static/B.png")), robotgo.ImgToCBitmap(bitImg))
	//_ = robotgo.Save(robotgo.ToImage(bit), "test.png")
	fmt.Println(fx, fy)
	end = time.Now()
	fmt.Printf("耗时：%v\n", end.Sub(start))
}
