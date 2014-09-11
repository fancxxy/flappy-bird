#ifndef _BACKGROUND_H_
#define _BACKGROUND_H_

#include "screensurface.h"
#include "imagesurface.h"

class Background {
    private:
        SDL_Rect background;    //背景
        SDL_Rect ground;    //地面
        int xVel;   //地面移动的速度
        int xFrame; //地面在水平方向的位置
    public:
        Background();
        void showBackground(ScreenSurface* screen, ImageSurface* image);
        void moveGround();
        void stopGround();
};

#endif
