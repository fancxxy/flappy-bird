package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type land struct {
	src       *sdl.Rect
	dst       [4]*sdl.Rect
	moveSpeed int32
}

func newLand(position map[string]*sdl.Rect) *land {
	ld := new(land)
	ld.moveSpeed = 2
	ld.src = position["land"]
	for i := 0; i < 4; i++ {
		ld.dst[i] = &sdl.Rect{
			X: ld.src.W * int32(i),
			Y: 512 - ld.src.H,
			W: ld.src.W,
			H: ld.src.H,
		}
	}
	return ld
}

func (ld *land) move() {
	for i := 0; i < 4; i++ {
		dst := ld.dst[i]
		dst.X -= ld.moveSpeed
		if dst.X <= -dst.W {
			dst.X = windowWidth
		}
	}
}

func (ld *land) paint(render *sdl.Renderer, image *sdl.Texture) {
	for i := 0; i < 4; i++ {
		render.Copy(image, ld.src, ld.dst[i])
	}
}
