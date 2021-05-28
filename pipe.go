package main

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type pipe struct {
	src, dst *sdl.Rect
}

// 管理水管的生成和消失
type pipes struct {
	src       map[string]*sdl.Rect
	picked    string
	choice    []string
	heightMax int32
	heightMin int32
	gap       int32
	moveSpeed int32
	distance  int32
	list      [][2]*pipe
}

func newPipes(position map[string]*sdl.Rect) *pipes {
	ps := &pipes{
		src: make(map[string]*sdl.Rect),
	}

	ps.src["green"] = position["pipe_up"]
	ps.src["red"] = position["pipe2_up"]

	ps.choice = []string{"green", "red"}
	ps.picked = random(ps.choice)
	ps.heightMax = 220
	ps.heightMin = 70
	ps.gap = 110
	ps.distance = 216
	ps.moveSpeed = 2
	ps.init()

	return ps
}

func (ps *pipes) init() {
	ps.picked = random(ps.choice)

	ps.list = ps.list[:0]
	for i := 0; i < 5; i++ {
		up, down := ps.produce(ps.src[ps.picked], 100+int32(i)*ps.distance)
		ps.list = append(ps.list, [2]*pipe{up, down})
	}
}

func (ps *pipes) move() {
	i := 0
	for _, pair := range ps.list {
		pair[0].dst.X -= ps.moveSpeed
		pair[1].dst.X -= ps.moveSpeed
		if pair[0].dst.X >= -pair[0].dst.W {
			ps.list[i] = pair
			i++
		}
	}

	// 第一个水管超出屏幕后删掉重新创建水管
	if i != len(ps.list) {
		up, down := ps.produce(ps.src[ps.picked], ps.distance-ps.list[0][0].dst.W)
		ps.list[len(ps.list)-1] = [2]*pipe{up, down}
	}
}

func (ps *pipes) paint(render *sdl.Renderer, image *sdl.Texture) {
	for i := range ps.list {
		up, down := ps.list[i][0], ps.list[i][1]
		render.Copy(image, up.src, up.dst)
		render.CopyEx(image, down.src, down.dst, 180.0, nil, sdl.FLIP_NONE)
	}
}

func (ps *pipes) produce(src *sdl.Rect, offset int32) (*pipe, *pipe) {
	h := rand.Int31n(ps.heightMax-ps.heightMin) + ps.heightMin

	up := &pipe{
		src: &sdl.Rect{
			X: src.X,
			Y: src.Y,
			W: src.W,
			H: h,
		},
		dst: &sdl.Rect{
			X: windowWidth + offset,
			Y: 400 - h,
			W: src.W,
			H: h,
		},
	}

	down := &pipe{
		src: &sdl.Rect{
			X: src.X,
			Y: src.Y,
			W: src.W,
			H: 400 - ps.gap - h,
		},
		dst: &sdl.Rect{
			X: windowWidth + offset,
			Y: 0,
			W: src.W,
			H: 400 - ps.gap - h,
		},
	}

	return up, down
}
