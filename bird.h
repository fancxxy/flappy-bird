#ifndef _BIRD_H_
#define _BIRD_H_

#include "imagesurface.h"
#include "screensurface.h"
#include "timer.h"

struct BirdBoundary {
    int top;
    int bottom;
    int left;
    int right;
};

class Bird {
    private:
        SDL_Rect bird[3];
        int yVel;   //y方向速度
        int yFrame; //y方向的坐标位置
        int yGravity;   //加速度
        int frame;  //帧切换
        Timer* time; //计时，每隔一段时间速度增加一次
        int direction;  //方向
    public:
        Bird();
        Bird(int direction);
        ~Bird();
        void showBird(ScreenSurface* screen, ImageSurface* image);
        void moveBird();

        void showBirdFly(ScreenSurface* screen, ImageSurface* image);
        void moveBirdFly();

        void handleEvents(Uint8* keystates);
        BirdBoundary getBirdBoundary();
};

#endif
