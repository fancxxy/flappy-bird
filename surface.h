#ifndef _SURFACE_H_
#define _SURFACE_H_

#include <SDL/SDL.h>

class Surface {
    protected:
        SDL_Surface* surface;
    public:
        Surface();
        ~Surface();
        SDL_Surface* getSurface();
    protected:
        void freeSurface();
};

#endif
