package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowTitle  = "flappy bird"
	windowWidth  = 864
	windowHeight = 512
	imagePath    = "./resources/images/atlas.png"
	positionPath = "./resources/images/atlas.txt"
	soundPath    = "./resources/sounds"
	frameRate    = 60
)

type Game struct {
	window   *sdl.Window
	render   *sdl.Renderer
	image    *sdl.Texture
	events   chan sdl.Event
	position map[string]*sdl.Rect
	sounds   map[string]*mix.Music
	paint    *paint
	running  bool
	sound    chan string
	score    int64
	best     int64
	quit     bool
	record   bool // 新记录
}

func NewGame() (game *Game, err error) {
	game = new(Game)

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}

	if game.window, err = sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight, sdl.WINDOW_SHOWN); err != nil {
		return nil, err
	}

	if game.render, err = sdl.CreateRenderer(game.window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		return nil, err
	}

	if game.image, err = img.LoadTexture(game.render, imagePath); err != nil {
		return nil, err
	}

	if game.position, err = loadPosition(positionPath); err != nil {
		return nil, err
	}

	if err := mix.Init(mix.INIT_OGG); err != nil {
		return nil, err
	}

	if game.sounds, err = loadSounds(soundPath); err != nil {
		return nil, err
	}

	game.paint = newPaint(game.position)
	game.events = make(chan sdl.Event)
	game.sound = make(chan string)
	game.best = readBest()

	return
}

func (game *Game) Run() {
	var wg sync.WaitGroup
	wg.Add(3)
	game.playSound(&wg)
	game.mainLoop(&wg)
	game.waitEvent(&wg)
	wg.Wait()
	game.destroy()
}

func (game *Game) playSound(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for !game.quit {
			name := <-game.sound
			game.sounds[name].Play(1)
			if name == "point" || name == "hit" {
				sdl.Delay(800)
			} else if name == "wing" {
				sdl.Delay(50)
			}
		}
	}()
}

func (game *Game) mainLoop(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for !game.quit {
			// 三个场景无限循环
			game.scene(game.ready())
			game.scene(game.run())
			game.scene(game.over())
		}
	}()
}

func (game *Game) waitEvent(wg *sync.WaitGroup) {
	defer wg.Done()
	for !game.quit {
		game.events <- sdl.WaitEvent()
	}
}

func (game *Game) scene(init func(), paint func(), handle func(e sdl.Event)) {
	if game.quit {
		return
	}

	init()
	game.running = true
	for game.running {
		select {
		case event := <-game.events:
			switch e := event.(type) {
			case *sdl.QuitEvent:
				close(game.sound)
				close(game.events)
				game.quit = true
				return
			case *sdl.MouseButtonEvent:
				if e.GetType() == sdl.MOUSEBUTTONDOWN {
					handle(e)
				}
			case *sdl.KeyboardEvent:
				if e.GetType() == sdl.KEYDOWN {
					if e.Keysym.Scancode == sdl.SCANCODE_RETURN {
						game.pause()
					} else if e.Keysym.Scancode == sdl.SCANCODE_SPACE {
						handle(e)
					}
				}
			}
		default:
		}

		paint()
		game.render.Present()
		sdl.Delay(1000 / frameRate)
	}
}

func (game *Game) pause() {
	game.paint.pauseButton(game.render, game.image)
	game.render.Present()
	for {
		event := <-game.events
		switch e := event.(type) {
		case *sdl.QuitEvent:
			close(game.sound)
			close(game.events)
			game.quit = true
			return
		case *sdl.KeyboardEvent:
			if e.GetType() == sdl.KEYDOWN && e.Keysym.Scancode == sdl.SCANCODE_RETURN {
				return
			}
		}
	}
}

func (game *Game) destroy() {
	for _, sound := range game.sounds {
		sound.Free()
	}
	mix.CloseAudio()
	mix.Quit()

	game.image.Destroy()
	game.render.Destroy()
	game.window.Destroy()
	sdl.Quit()
	writeBest(game.best)
}

func (game *Game) ready() (func(), func(), func(event sdl.Event)) {
	init := func() {
		game.score = 0
		game.record = false
		game.paint.init()
	}
	paint := func() {
		game.paint.ready(game.render, game.image)
	}
	handle := func(event sdl.Event) {
		game.running = false
	}
	return init, paint, handle
}

func (game *Game) run() (func(), func(), func(event sdl.Event)) {
	init := func() {}
	paint := func() {
		game.paint.run(game.render, game.image, game.score)
		collision, score := game.paint.bird.collisionAndScore(game.paint.pipes)
		if collision {
			game.paint.bird.dead()
			game.running = false
		}
		if score {
			go game.mustPlay("point")
			game.score += 1
		}
	}
	handle := func(event sdl.Event) {
		game.play("wing")
		game.paint.bird.jump()
	}
	return init, paint, handle
}

func (game *Game) over() (func(), func(), func(event sdl.Event)) {
	init := func() {
		go func() {
			game.mustPlay("hit")
			game.mustPlay("die")
			// game.mustPlay("swooshing")
		}()

		if game.score > game.best {
			game.record = true
			game.best = game.score
		}
	}
	paint := func() {
		game.paint.over(game.render, game.image, game.score, game.best, game.record)
	}
	handle := func(event sdl.Event) {
		if game.paint.clickedPlayButton(event) {
			game.running = false
		}
	}
	return init, paint, handle
}

func (game *Game) play(sound string) {
	select {
	case game.sound <- sound:
	default:
	}
}

func (game *Game) mustPlay(sound string) {
	game.sound <- sound
}

func loadSounds(path string) (map[string]*mix.Music, error) {
	ret := make(map[string]*mix.Music)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if err := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		return nil, err
	}

	for _, file := range files {
		sound, err := mix.LoadMUS(filepath.Join(path, file.Name()))
		if err != nil {
			return nil, err
		}

		name := strings.TrimSuffix(strings.TrimPrefix(file.Name(), "sfx_"), ".ogg")
		ret[name] = sound
	}

	return ret, nil
}

func loadPosition(filename string) (map[string]*sdl.Rect, error) {
	ret := make(map[string]*sdl.Rect)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		slice := strings.Split(scanner.Text(), " ")
		name := slice[0]
		w, _ := strconv.ParseInt(slice[1], 10, 32)
		h, _ := strconv.ParseInt(slice[2], 10, 32)
		x, _ := strconv.ParseInt(slice[3], 10, 32)
		y, _ := strconv.ParseInt(slice[4], 10, 32)
		ret[name] = &sdl.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}

func readBest() int64 {
	bs, err := ioutil.ReadFile("./data.db")
	if err != nil {
		return 0
	}
	best, err := strconv.ParseInt(string(bs), 10, 64)
	if err != nil {
		return 0
	}
	return best
}

func writeBest(best int64) {
	ioutil.WriteFile("./data.db", []byte(strconv.FormatInt(best, 10)), 0644)
}
