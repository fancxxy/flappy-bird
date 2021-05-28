package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type paint struct {
	src       map[string]*sdl.Rect
	dst       map[string]*sdl.Rect
	bird      *bird
	land      *land
	pipes     *pipes
	moveSpeed int32
	day       string

	// 渐变效果
	splashAlphaVal uint8
	scoreAlphaVal  uint8
	panelAlphaVal  uint8
	medalAlphaVal  uint8
	medalReduce    int32
	fadeRate       uint8
	panelOffset    int32
}

func newPaint(position map[string]*sdl.Rect) *paint {
	pt := &paint{
		src: make(map[string]*sdl.Rect),
		dst: make(map[string]*sdl.Rect),
	}

	pt.moveSpeed = 2
	pt.fadeRate = 5

	pt.src["day"] = position["bg_day"]
	pt.src["night"] = position["bg_night"]
	pt.src["title"] = position["title"]
	pt.src["ready"] = position["text_ready"]
	pt.src["tutorial"] = position["tutorial"]
	pt.src["over"] = position["text_game_over"]
	pt.src["panel"] = position["score_panel"]
	pt.src["play"] = position["button_play"]
	pt.src["pause"] = position["button_pause"]
	pt.src["new"] = position["new"]

	pt.dst["day"] = &sdl.Rect{X: 0, Y: 0, W: pt.src["day"].W, H: pt.src["day"].H}
	pt.dst["night"] = pt.dst["day"]

	src := pt.src["title"]
	pt.dst["title"] = &sdl.Rect{X: (windowWidth - src.W) / 2, Y: 50, W: src.W, H: src.H}
	src = pt.src["ready"]
	pt.dst["ready"] = &sdl.Rect{X: (windowWidth - src.W) / 2, Y: 160, W: src.W, H: src.H}
	src = pt.src["tutorial"]
	pt.dst["tutorial"] = &sdl.Rect{X: (windowWidth - src.W) / 2, Y: 230, W: src.W, H: src.H}
	src = pt.src["over"]
	pt.dst["over"] = &sdl.Rect{X: (windowWidth - src.W) / 2, Y: 150, W: src.W, H: src.H}
	src = pt.src["panel"]
	pt.dst["panel"] = &sdl.Rect{X: (windowWidth - src.W) / 2, Y: 220, W: src.W, H: src.H}
	src = pt.src["play"]
	pt.dst["play"] = &sdl.Rect{X: (windowWidth - src.W) / 2, Y: 380, W: src.W, H: src.H}

	for i := 0; i < 10; i++ {
		name := "font_0" + strconv.FormatInt(int64(i+48), 10)
		pt.src["big"+strconv.Itoa(i)] = position[name]

		name = "number_context_0" + strconv.FormatInt(int64(i), 10)
		pt.src["small"+strconv.Itoa(i)] = position[name]
	}

	pt.src["platina"] = position["medals_0"]
	pt.src["gold"] = position["medals_1"]
	pt.src["silver"] = position["medals_2"]
	pt.src["bronze"] = position["medals_3"]

	pt.dst["big"] = &sdl.Rect{
		X: (windowWidth - pt.src["big0"].W) / 2,
		Y: 50,
		W: pt.src["big0"].W,
		H: pt.src["big0"].H,
	}

	pt.dst["small"] = &sdl.Rect{
		X: windowWidth/2 + 78,
		Y: 258,
		W: pt.src["small0"].W,
		H: pt.src["small0"].H,
	}

	pt.dst["medal"] = &sdl.Rect{X: windowWidth/2 - 137, Y: 130, W: pt.src["gold"].W + 102, H: pt.src["gold"].H + 102}
	pt.dst["pause"] = &sdl.Rect{X: 30, Y: 30, W: pt.src["pause"].W, H: pt.src["pause"].H}
	pt.dst["new"] = &sdl.Rect{X: windowWidth/2 + 5, Y: 300, W: pt.src["new"].W, H: pt.src["new"].H}

	pt.bird = newBird(position)
	pt.pipes = newPipes(position)
	pt.land = newLand(position)

	return pt
}

func (pt *paint) init() {
	pt.bird.init()
	pt.pipes.init()
	pt.day = random([]string{"day", "night"})
	pt.splashAlphaVal = 255
	pt.scoreAlphaVal = 0
	pt.panelAlphaVal = 0
	pt.panelOffset = 0
	pt.medalReduce = 0
	pt.medalAlphaVal = 0
}

func (pt *paint) ready(render *sdl.Renderer, image *sdl.Texture) {
	pt.background(render, image)
	pt.land.move()
	pt.land.paint(render, image)
	pt.splash(render, image)
	pt.bird.wing()
	pt.bird.flow()
	pt.bird.paint(render, image)
}

func (pt *paint) run(render *sdl.Renderer, image *sdl.Texture, score int64) {
	pt.runEffect()
	pt.background(render, image)
	pt.land.move()
	pt.land.paint(render, image)
	pt.splash(render, image)
	pt.pipes.move()
	pt.pipes.paint(render, image)
	pt.score(render, image, score)
	pt.bird.wing()
	pt.bird.fall()
	pt.bird.paint(render, image)
}

func (pt *paint) over(render *sdl.Renderer, image *sdl.Texture, score, best int64, record bool) {
	pt.overEffect()
	pt.background(render, image)
	pt.land.paint(render, image)
	pt.pipes.paint(render, image)
	pt.bird.fall()
	pt.bird.paint(render, image)
	pt.panel(render, image, score, best, record)
	pt.medal(render, image, score)
	pt.bird.paint(render, image)
}

func (pt *paint) background(render *sdl.Renderer, image *sdl.Texture) {
	src, dst := pt.src[pt.day], pt.dst[pt.day]
	for i := 0; i < 3; i++ {
		dst.X = int32(i) * dst.W
		render.Copy(image, src, dst)
	}
}

func (pt *paint) panel(render *sdl.Renderer, image *sdl.Texture, score, best int64, record bool) {
	image.SetAlphaMod(pt.panelAlphaVal)
	for _, scene := range []string{"over", "panel", "play"} {
		dst := *pt.dst[scene]
		dst.Y -= int32(pt.panelOffset)
		render.Copy(image, pt.src[scene], &dst)
	}

	var digits []string
	for number := score; number != 0; number = number / 10 {
		digits = append(digits, strconv.FormatInt(number%10, 10))
	}
	if len(digits) == 0 {
		digits = append(digits, "0")
	}

	for i := 0; i < len(digits); i++ {
		dst := *pt.dst["small"]
		dst.X -= int32(i) * dst.W
		dst.Y -= int32(pt.panelOffset)
		render.Copy(image, pt.src["small"+digits[i]], &dst)
	}

	digits = digits[:0]
	for number := best; number != 0; number = number / 10 {
		digits = append(digits, strconv.FormatInt(number%10, 10))
	}
	if len(digits) == 0 {
		digits = append(digits, "0")
	}
	for i := 0; i < len(digits); i++ {
		dst := *pt.dst["small"]
		dst.X -= int32(i) * dst.W
		dst.Y -= int32(pt.panelOffset) - 42
		render.Copy(image, pt.src["small"+digits[i]], &dst)
	}

	if record {
		dst := *pt.dst["new"]
		dst.Y -= int32(pt.panelOffset)
		render.Copy(image, pt.src["new"], &dst)
	}

	image.SetAlphaMod(255)
}

func (pt *paint) medal(render *sdl.Renderer, image *sdl.Texture, score int64) {
	if pt.medalAlphaVal == 0 {
		return
	}

	image.SetAlphaMod(pt.medalAlphaVal)

	// 白金 < 40、黄金 < 30、白银 < 20、青铜 < 10
	var src *sdl.Rect
	switch {
	case score <= 10:
		src = pt.src["bronze"]
	case score > 10 && score <= 20:
		src = pt.src["silver"]
	case score > 20 && score <= 30:
		src = pt.src["gold"]
	default:
		src = pt.src["platina"]
	}

	if src != nil {
		dst := *pt.dst["medal"]
		dst.X += pt.medalReduce
		dst.Y += pt.medalReduce
		dst.W -= pt.medalReduce * 2
		dst.H -= pt.medalReduce * 2
		render.Copy(image, src, &dst)
	}

	image.SetAlphaMod(255)
}

func (pt *paint) splash(render *sdl.Renderer, image *sdl.Texture) {
	if pt.splashAlphaVal == 0 {
		return
	}
	image.SetAlphaMod(pt.splashAlphaVal)
	for _, scene := range []string{"title", "ready", "tutorial"} {
		render.Copy(image, pt.src[scene], pt.dst[scene])
	}
	image.SetAlphaMod(255)
}

func (pt *paint) score(render *sdl.Renderer, image *sdl.Texture, score int64) {
	defer image.SetAlphaMod(255)
	image.SetAlphaMod(pt.scoreAlphaVal)
	dst := pt.dst["big"]
	if score == 0 {
		render.Copy(image, pt.src["big0"], dst)
		return
	}

	var digits []string
	for number := score; number != 0; number = number / 10 {
		digits = append(digits, strconv.FormatInt(number%10, 10))
	}

	offset := (int32(windowWidth) - dst.W*int32(len(digits))) / 2
	for i := len(digits) - 1; i >= 0; i-- {
		dst.X = offset + int32(len(digits)-i-1)*dst.W
		render.Copy(image, pt.src["big"+digits[i]], dst)
	}
}

func (pt *paint) pauseButton(render *sdl.Renderer, image *sdl.Texture) {
	render.Copy(image, pt.src["pause"], pt.dst["pause"])
}

func (pt *paint) clickedPlayButton(event sdl.Event) bool {
	if pt.panelOffset < 30 {
		return false
	}

	switch e := event.(type) {
	case *sdl.MouseButtonEvent:
		dst := pt.dst["play"]
		point := &sdl.Point{X: e.X, Y: e.Y}
		rect := &sdl.Rect{X: dst.X, Y: dst.Y - pt.panelOffset, W: dst.W, H: dst.H}
		return point.InRect(rect)
	case *sdl.KeyboardEvent:
		return true
	default:
	}

	return false
}

func (pt *paint) runEffect() {
	if pt.splashAlphaVal >= 5 {
		pt.splashAlphaVal -= pt.fadeRate
	}

	if pt.splashAlphaVal == 0 && pt.scoreAlphaVal <= 250 {
		pt.scoreAlphaVal += pt.fadeRate
	}

}

func (pt *paint) overEffect() {
	if pt.panelAlphaVal < 250 {
		pt.panelAlphaVal += pt.fadeRate
	}

	if pt.panelOffset <= 80.0 {
		pt.panelOffset += 5
	}

	if pt.panelAlphaVal > 80.0 && pt.medalAlphaVal < 250 {
		pt.medalAlphaVal += pt.fadeRate
	}

	if pt.panelAlphaVal > 80.0 && pt.medalAlphaVal < 250 && pt.medalReduce < 52 {
		pt.medalReduce += 1
	}
}

func random(choice []string) string {
	return choice[rand.Intn(len(choice))]
}
