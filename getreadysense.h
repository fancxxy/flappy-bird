#ifndef _GETREADYSENSE_H_
#define _GETREADYSENSE_H_

#include "screensurface.h"
#include "imagesurface.h"
#include "bird.h"
#include "number.h"
#include "background.h"
#include "timer.h"

class GetReadySense {
    private:
        SDL_Rect getReady;
        SDL_Rect tap;
        Bird* bird;
        Number* number;
        Background* background;
        bool gamequit;
        bool gamestart;
        Timer* fps;
        SDL_Event event;
    public:
        GetReadySense();
        ~GetReadySense();
        void showGetReadySense(ScreenSurface* screen, ImageSurface* image);
        bool gameStart();
        bool gameQuit();
};

#endif
