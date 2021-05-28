package main

import "github.com/veandco/go-sdl2/sdl"

const (
	gravity = 0.2
)

type bird struct {
	src        map[string][]*sdl.Rect
	x, y       float64
	wingVal    float64
	wingRate   float64
	fallSpeed  float64
	rotateVal  float64
	rotateRate float64
	picked     string
	choice     []string
}

func newBird(position map[string]*sdl.Rect) *bird {
	bd := &bird{
		src: make(map[string][]*sdl.Rect),
	}
	bd.src["yellow"] = []*sdl.Rect{
		position["bird0_0"],
		position["bird0_1"],
		position["bird0_2"],
		position["bird0_1"],
	}
	bd.src["blue"] = []*sdl.Rect{
		position["bird1_0"],
		position["bird1_1"],
		position["bird1_2"],
		position["bird1_1"],
	}
	bd.src["red"] = []*sdl.Rect{
		position["bird2_0"],
		position["bird2_1"],
		position["bird2_2"],
		position["bird2_1"],
	}

	bd.choice = []string{"yellow", "blue", "red"}
	bd.picked = random(bd.choice)
	// 初始坐标
	bd.x, bd.y = 180, 240
	// 煽动翅膀频率
	bd.wingVal = 0.0
	bd.wingRate = 0.2
	// 下落速度
	bd.fallSpeed = 0.4
	//  旋转角度
	bd.rotateVal = 0.0
	bd.rotateRate = 2.5

	return bd
}

func (bd *bird) init() {
	bd.picked = random(bd.choice)
	bd.x, bd.y = 180, 240
	bd.fallSpeed = 0.4
	bd.rotateVal = 0.0
}

func (bd *bird) collisionAndScore(pipes *pipes) (collision bool, score bool) {
	// 低于地面
	if bd.y >= windowHeight-155 {
		collision = true
		return
	}

	// 找出在碰撞检测范围之内的管道
	var up, down *sdl.Rect
	for _, pair := range pipes.list {
		pipe := pair[0].dst
		if pipe.X >= int32(bd.x)-pipe.W && pipe.X <= int32(bd.x)+bd.src[bd.picked][0].W {
			up, down = pair[0].dst, pair[1].dst
			break
		}
	}

	if up == nil || down == nil {
		return
	}

	// 穿过管道1/2宽度就认为得分
	if up.X == int32(bd.x)-up.W/2 {
		score = true
	}

	bird := &sdl.Rect{X: int32(bd.x) + 12, Y: int32(bd.y) + 12, W: 24, H: 24}
	if bird.HasIntersection(up) || bird.HasIntersection(down) {
		collision = true
	}

	return
}

func (bd *bird) flow() {
	bd.y += bd.fallSpeed
	if bd.y > 250.0 || bd.y < 230.0 {
		bd.fallSpeed *= -1
	}
}

func (bd *bird) dead() {
	bd.fallSpeed = 0.0
}

func (bd *bird) jump() {
	// 每次jump给的速度
	bd.fallSpeed = -4.0
	// 旋转最小角度
	bd.rotateVal = -45.0
}

func (bd *bird) fall() {
	// 10.5是最大下落速度
	if bd.fallSpeed < 10.5 {
		bd.fallSpeed += gravity
	}
	if bd.y < windowHeight-150 {
		bd.y += bd.fallSpeed
	}

	// 超过上边界
	if bd.y < -10.0 {
		bd.y = -10.0
		bd.fallSpeed = 0.0
	}

	// 60.0是旋转最大角度
	if bd.rotateVal < 60.0 {
		bd.rotateVal += bd.rotateRate
	}
}

func (bd *bird) wing() {
	bd.wingVal += bd.wingRate
	if bd.wingVal >= 4.0 {
		bd.wingVal = 0.0
	}
}

func (bd *bird) paint(render *sdl.Renderer, image *sdl.Texture) {
	picked := bd.src[bd.picked]
	src := picked[int(bd.wingVal)%len(picked)]
	dst := &sdl.Rect{X: int32(bd.x), Y: int32(bd.y), W: src.W, H: src.H}
	render.CopyEx(image, src, dst, bd.rotateVal, nil, sdl.FLIP_NONE)
}
