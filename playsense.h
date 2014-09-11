#ifndef _PLAYSENSE_H_
#define _PLAYSENSE_H_

#include "imagesurface.h"
#include "screensurface.h"
#include "pipe.h"
#include "bird.h"
#include "background.h"
#include "number.h"
#include "timer.h"

class PlaySense {
    private:
        Background* background;
        Bird* bird;
        Pipe* pipe[4];
        Number* number;
        bool gameover;
        bool gamequit;
        bool gamestop;
        int goal;
        Timer* fps;
        SDL_Event event;
    public:
        PlaySense();
        ~PlaySense();
        void showPlaySense(ScreenSurface* screen, ImageSurface* image);
        bool gameOver();
        bool gameQuit();
        bool collision(Bird* bird, Pipe* pipe);
};

#endif
