#ifndef _NUMBER_H_
#define _NUMBER_H_

#include "imagesurface.h"
#include "screensurface.h"

class Number {
    public:
        SDL_Rect num[10];
    public:
        Number();
        SDL_Rect getNumber(int n);
        void showNumber(ScreenSurface* screen, ImageSurface* image, int n);
};

#endif
