#ifndef _IMAGESURFACE_H_
#define _IMAGESURFACE_H_

#include "surface.h"
#include <SDL/SDL_image.h>
#include <string>

class ImageSurface : public Surface {
    public:
        SDL_Surface* image;
        ImageSurface(const std::string &img);
        ~ImageSurface();
        SDL_Surface* loadImage(const std::string &img);
};

#endif
