#include "number.h"

Number::Number() {
    num[0] = {992, 120, 24, 36};
    num[1] = {272, 910, 24, 36};
    num[2] = {584, 320, 24, 36};
    num[3] = {612, 320, 24, 36};
    num[4] = {640, 320, 24, 36};
    num[5] = {668, 320, 24, 36};
    num[6] = {584, 368, 24, 36};
    num[7] = {612, 368, 24, 36};
    num[8] = {640, 368, 24, 36};
    num[9] = {668, 368, 24, 36};
}

SDL_Rect Number::getNumber(int n) {
    return num[n];
}

void Number::showNumber(ScreenSurface* screen, ImageSurface* image, int n) {
    if (n <= 0) {
        SDL_Rect tmp = getNumber(0);
        screen->applySurface(*image, tmp.w, tmp.h, tmp.x, tmp.y, 339, 50);
    } else if (n < 10 && n > 0) {
        SDL_Rect tmp = getNumber(n);
        screen->applySurface(*image, tmp.w, tmp.h, tmp.x, tmp.y, 339, 50);
    } else if (n >= 10 && n < 100) {
        int one = n % 10;
        int ten = n / 10;
        SDL_Rect tmpOne = getNumber(one);
        SDL_Rect tmpTen = getNumber(ten);
        screen->applySurface(*image, tmpTen.w, tmpTen.h, tmpTen.x, tmpTen.y, 329, 50);
        screen->applySurface(*image, tmpOne.w, tmpOne.h, tmpOne.x, tmpOne.y, 349, 50);
    } else if (n >= 100 && n < 1000) {
        int one = n % 10;
        int ten = n % 100 / 10;
        int hundred = n / 100;
        SDL_Rect tmpOne = getNumber(one);
        SDL_Rect tmpTen = getNumber(ten);
        SDL_Rect tmpHundred = getNumber(hundred);
        screen->applySurface(*image, tmpHundred.w, tmpHundred.h, tmpHundred.x, tmpHundred.y, 319, 50);
        screen->applySurface(*image, tmpTen.w, tmpTen.h, tmpTen.x, tmpTen.y, 339, 50);
        screen->applySurface(*image, tmpOne.w, tmpOne.h, tmpOne.x, tmpOne.y, 359, 50);
    } else {
        SDL_Rect tmp= getNumber(9);
        screen->applySurface(*image, tmp.w, tmp.h, tmp.x, tmp.y, 319, 50);
        screen->applySurface(*image, tmp.w, tmp.h, tmp.x, tmp.y, 339, 50);
        screen->applySurface(*image, tmp.w, tmp.h, tmp.x, tmp.y, 359, 50);
    }
}
