#ifndef _PIPE_H_
#define _PIPE_H_

#include "imagesurface.h"
#include "screensurface.h"

struct PipeBoundary {
    int upLeft;
    int upRight;
    int upBottom;
    int upTop;

    int downLeft;
    int downRight;
    int downTop;
    int downBottom;
};

class Pipe {
    private:
        SDL_Rect up, down;
        int xVel;       //管子移动的速度
        int xFrame;     //管子水平方向的位置
        int upHeight;   //上面管子的长度
        int downHeight; //下面管子的长度
    public:
        Pipe(int xFrame);   //设置初始位置
        void movePipe();
        void showPipe(ScreenSurface* screen, ImageSurface* image);
        int getUpHeight();
        int getxFrame();    //得分
        PipeBoundary getPipeBoundary();
        void stopPipe();
};

#endif
