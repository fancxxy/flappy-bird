#include "background.h"

Background::Background() {
    background = {0, 0, 288, 512};
    ground = {584, 0, 336, 112};
    xVel = 5;   
    xFrame = 0; 
}

void Background::showBackground(ScreenSurface* screen, ImageSurface* image) {
    moveGround();

    //显示背景
    screen->applySurface(*image, background.w, background.h, background.x, background.y, 0, 0);
    screen->applySurface(*image, background.w, background.h, background.x, background.y, background.w, 0);
    screen->applySurface(*image, background.w, background.h, background.x, background.y, background.w * 2, 0);

    //显示地面
    screen->applySurface(*image, ground.w, ground.h, ground.x, ground.y, xFrame, 400);
    screen->applySurface(*image, ground.w, ground.h, ground.x, ground.y, xFrame + ground.w, 400);
    screen->applySurface(*image, ground.w, ground.h, ground.x, ground.y, xFrame + ground.w * 2, 400);
}

void Background::moveGround() {
    xFrame -= xVel;
    if (xFrame == -120)
        xFrame = 0;
}

void Background::stopGround() {
    xVel = 0;
}
