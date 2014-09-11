#include "surface.h"

Surface::Surface() {
    surface = NULL;
}

Surface::~Surface() {
    freeSurface();
}

void Surface::freeSurface() {
    if (surface) 
        SDL_FreeSurface(surface);
    surface = NULL;
}

SDL_Surface* Surface::getSurface() {
    return surface;
}
