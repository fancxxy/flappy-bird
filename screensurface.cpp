#include "screensurface.h"

bool ScreenSurface::screenCreated = false;

void ScreenSurface::init() {
    if (screenCreated)
        throw "Can't create more than one screen.";
    if (SDL_Init(SDL_INIT_EVERYTHING) == -1)
        throw SDL_GetError();

    screen = SDL_SetVideoMode(width, height, bpp, flags);
    if (screen)
        screenCreated = true;
    else
        throw "Create screen surface failed.";
}

ScreenSurface::ScreenSurface(const int width, const int height,
        const int bpp, const Uint32 flags) : 
    width(width), height(height), bpp(bpp), flags(flags) {
    init();
}

ScreenSurface::~ScreenSurface() {
    SDL_Quit();
}

bool ScreenSurface::flip() const {
    return SDL_Flip(screen) != -1;
}

void ScreenSurface::applySurface(Surface& surface, const int width, const int height, 
        const int srcX, const int srcY,
        const int dstX, const int dstY) const { 
    SDL_Rect offset, clip;

    offset.x = dstX;
    offset.y = dstY;
    clip.x = srcX;
    clip.y = srcY;
    clip.w = width;
    clip.h = height;

    SDL_BlitSurface(surface.getSurface(), &clip, screen, &offset);
}

void ScreenSurface::setTitle(const std::string &title) const {
    SDL_WM_SetCaption(title.c_str(), NULL);
}
