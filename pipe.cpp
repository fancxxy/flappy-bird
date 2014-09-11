#include "pipe.h"
#include <ctime>

Pipe::Pipe(int xFrame) : xFrame(xFrame) {
    up = {112, 646, 52, 320};
    down = {168, 646, 52, 320};
    xVel = 5;

    srand(time(0));
    upHeight = 0;
    downHeight = 0;
}

void Pipe::showPipe(ScreenSurface* screen, ImageSurface* image) {
    movePipe();

    //显示上面的管子
    screen->applySurface(*image, up.w, upHeight, up.x, up.y + up.h - upHeight, xFrame, 0);
    //显示下面的管子
    screen->applySurface(*image, down.w, downHeight, down.x, down.y, xFrame, upHeight + 100);
}

void Pipe::movePipe() {
    xFrame -= xVel;
    if (xFrame < -235) {
        xFrame = 705;

        upHeight = getUpHeight();
        downHeight = 300 - upHeight;
    }
}

int Pipe::getUpHeight() {
    return  50 + rand() % (270 - 50); //管子长度控制在(50, 270)之内
}

int Pipe::getxFrame() {
    return xFrame;
}

PipeBoundary Pipe::getPipeBoundary() {
    PipeBoundary pipeBoundary;
    pipeBoundary.upLeft = xFrame;
    pipeBoundary.upRight = xFrame + 52;
    pipeBoundary.upBottom = upHeight;
    pipeBoundary.upTop = 0;

    pipeBoundary.downLeft = xFrame;
    pipeBoundary.downRight = xFrame + 52;
    pipeBoundary.downTop = 400 - downHeight;
    pipeBoundary.downBottom = 400;

    return pipeBoundary;
}

void Pipe::stopPipe() {
    xVel = 0;
}
