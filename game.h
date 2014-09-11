#ifndef _GAME_H_
#define _GAME_H_

#include <SDL/SDL.h>
#include "imagesurface.h"
#include "screensurface.h"
#include "playsense.h"
#include "getreadysense.h"

#include <string>

class Game {
    private:
        bool gamequit;
        bool gameover;
        bool gamestart;

        SDL_Event event;
        GetReadySense* getreadysense;
        PlaySense* playsense;

        ScreenSurface* screen;
        ImageSurface* image;
    public:
        Game();
        ~Game();
        void gameRun();
};

#endif
