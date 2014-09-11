#include "bird.h"
#include <SDL/SDL_rotozoom.h>

Bird::Bird()  {
    bird[0] = {6, 982, 34, 23};
    bird[1] = {62, 982, 34, 23};
    bird[2] = {118, 982, 34, 23};

    yVel = -8;
    yFrame = 0;
    yGravity = 2;
    frame = 0;
    time = new Timer;
    time->start();
}

Bird::Bird(int direction) : direction(direction) {
    bird[0] = {6, 982, 34, 23};
    bird[1] = {62, 982, 34, 23};
    bird[2] = {118, 982, 34, 23};
    yFrame = 200;
    frame = 0;
    time = new Timer;
}

Bird::~Bird() {
    delete time;
}


void Bird::showBird(ScreenSurface* screen, ImageSurface* image) {
    moveBird();
    screen->applySurface(*image, bird[frame/2].w, bird[frame/2].h, 
            bird[frame/2].x, bird[frame/2].y, 141,  200 + yFrame);
}

void Bird::moveBird() {
    frame++;
    if (frame > 5)
        frame = 0;

    if (time->getTicks() > 200) {
        yVel += yGravity;   //y方向速度
        if (yVel >= 10)
            yVel = 10;
        time->start();
    }

    yFrame += yVel; //y方向位置
}

void Bird::handleEvents(Uint8* keystates) {
    if (keystates[SDLK_SPACE]) {
        yFrame -= 6;
        yVel = -2;
    }
}

void Bird::showBirdFly(ScreenSurface* screen, ImageSurface* image) {
    moveBirdFly();
    screen->applySurface(*image, bird[frame/2].w, bird[frame/2].h, 
            bird[frame/2].x, bird[frame/2].y, 141,  yFrame);
}

void Bird::moveBirdFly() {
    frame++;
    if (frame > 5)
        frame = 0;

    if (direction == 0)
        yFrame++;
    else if (direction == 1)
        yFrame--;
    
    if (yFrame == 205)
        direction = 1;
    else if (yFrame == 195)
        direction = 0;

}

BirdBoundary Bird::getBirdBoundary() {
    BirdBoundary birdboundary;
    birdboundary.top = 200 + yFrame;
    birdboundary.bottom  = 223 + yFrame;
    birdboundary.left = 141;
    birdboundary.right = 175;
    return birdboundary;
}
