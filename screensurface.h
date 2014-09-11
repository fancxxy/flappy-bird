#ifndef _SCREENSURFACE_H_
#define _SCREENSURFACE_H_

#include <SDL/SDL.h>
#include <SDL/SDL_image.h>
#include "surface.h"
#include <string>

class ScreenSurface {
    private:
        static bool screenCreated;
        int bpp;
        Uint32 flags;
        SDL_Surface* screen;
        int width;
        int height;
    public:
        //ScreenSurface();
        ScreenSurface(const int width, const int height,
                const int bpp = 32, const Uint32 flags = SDL_SWSURFACE);
        ~ScreenSurface();
        bool flip() const;
        void setTitle(const std::string &title) const ;
        void applySurface(Surface& surface, const int width, const int height, 
                const int srcX = 0, const int srcY = 0,
                const int dstX = 0, const int dstY = 0) const;
    private:
        void init();
};

#endif
