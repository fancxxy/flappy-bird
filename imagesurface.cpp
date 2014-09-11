#include "imagesurface.h"

ImageSurface::ImageSurface(const std::string &img) {
    image = loadImage(img);
}

ImageSurface::~ImageSurface() {
    freeSurface();
}

SDL_Surface* ImageSurface::loadImage(const std::string &img) {
    freeSurface();
    //SDL_Surface* tmp = IMG_Load(image.c_str());
    //if (tmp) {
    //    surface = SDL_DisplayFormat(tmp);
    //    SDL_FreeSurface(tmp);
    //}
    surface = IMG_Load(img.c_str());
    return surface;
}
