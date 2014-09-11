#include "getreadysense.h"

extern int FRAMES_PER_SECOND;

GetReadySense::GetReadySense() {
    getReady = {590, 118, 184, 50};
    tap = {584, 182, 114, 98};
    bird = new Bird(0);
    background = new Background();
    number = new Number();
    gamequit = false;
    gamestart = false;
    fps = new Timer();
}

GetReadySense::~GetReadySense() {
    delete bird;
    delete background;
    delete number;
    delete fps;
}

void GetReadySense::showGetReadySense(ScreenSurface* screen, ImageSurface* image) {
    while (gamequit == false && gamestart == false) {
        fps->start();

        Uint8* keystates = SDL_GetKeyState(NULL);
        if (keystates[SDLK_SPACE])
            gamestart = true;

        background->showBackground(screen, image);

        number->showNumber(screen, image, 0);

        bird->showBirdFly(screen, image);

        screen->applySurface(*image, 184, 50, 590, 118, 258, 125); 
        screen->applySurface(*image, 114, 98, 584, 182, 293, 200); 

        if (fps->getTicks() < 1000 / FRAMES_PER_SECOND) {
            SDL_Delay(1000 / FRAMES_PER_SECOND - fps->getTicks());
        }

        screen->flip();

        if (SDL_PollEvent(&event) && event.type == SDL_QUIT)
            gamequit = true;
    }
}

bool GetReadySense::gameStart() {
    return gamestart;
}

bool GetReadySense::gameQuit() {
    return gamequit;
}
