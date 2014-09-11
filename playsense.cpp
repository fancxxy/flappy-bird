#include "playsense.h"

extern int FRAMES_PER_SECOND;

PlaySense::PlaySense() {
    background = new Background();
    bird = new Bird(); number = new Number();
    pipe[0] = new Pipe(100);
    pipe[1] = new Pipe(100 + 235);
    pipe[2] = new Pipe(100 + 235 * 2);
    pipe[3] = new Pipe(100 + 235 * 3);
    gameover = false;
    gamequit = false;
    gamestop = false;
    goal = -3;
    fps = new Timer();
}

PlaySense::~PlaySense() {
    delete background;
    delete bird;
    delete number;
    delete[] *pipe;
    delete fps;
} void PlaySense::showPlaySense(ScreenSurface* screen, ImageSurface* image) {
    while ( gamequit == false && gameover == false) {
        fps->start();

        background->showBackground(screen, image);

        pipe[0]->showPipe(screen, image);
        pipe[1]->showPipe(screen, image);
        pipe[2]->showPipe(screen, image);
        pipe[3]->showPipe(screen, image);

        Uint8* keystates = SDL_GetKeyState(NULL);

        if (gamestop == false)
            bird->handleEvents(keystates);
        bird->showBird(screen, image);

        number->showNumber(screen, image, goal);

        if ( (pipe[0]->getxFrame() == 175 || pipe[1]->getxFrame() == 175 || 
              pipe[2]->getxFrame() == 175 || pipe[3]->getxFrame() == 175) && gamestop == false )
            goal++;

        if (fps->getTicks() < 1000 / FRAMES_PER_SECOND) {
            SDL_Delay(1000 / FRAMES_PER_SECOND - fps->getTicks());
        }

        screen->flip();

        if (SDL_PollEvent(&event) && event.type == SDL_QUIT)
            gamequit = true;

        if ( collision(bird, pipe[0]) || collision(bird, pipe[1]) || 
             collision(bird, pipe[2]) || collision(bird, pipe[3]) ||
             bird->getBirdBoundary().bottom >= 400 ) {
            //gameover = true;
            pipe[0]->stopPipe();
            pipe[1]->stopPipe();
            pipe[2]->stopPipe();
            pipe[3]->stopPipe();
            background->stopGround();
            gamestop = true;
        }

        if (bird->getBirdBoundary().bottom >= 400) {
            SDL_Delay(1000);
            gameover = true;
        }
    }
}

bool PlaySense::gameOver() {
    return gameover;
}

bool PlaySense::gameQuit() {
    return gamequit;
}

//碰撞
bool PlaySense::collision(Bird* bird, Pipe* pipe) { 
   int birdLeft = bird->getBirdBoundary().left;
   int birdRight = bird->getBirdBoundary().right;
   int birdTop = bird->getBirdBoundary().top;
   int birdBottom = bird->getBirdBoundary().bottom;

   int pipeUpLeft = pipe->getPipeBoundary().upLeft;
   int pipeUpRight = pipe->getPipeBoundary().upRight;
   int pipeUpTop = pipe->getPipeBoundary().upTop;
   int pipeUpBottom = pipe->getPipeBoundary().upBottom;
   int pipeDownLeft = pipe->getPipeBoundary().downLeft;
   int pipeDownRight = pipe->getPipeBoundary().downRight;
   int pipeDownTop = pipe->getPipeBoundary().downTop;
   int pipeDownBottom = pipe->getPipeBoundary().downBottom;

   if ( (birdRight >= pipeUpLeft && birdTop <= pipeUpBottom && birdLeft <= pipeUpRight) ||
        (birdRight >= pipeDownLeft && birdBottom >= pipeDownTop && birdLeft <= pipeDownRight) )
       return true;
   else return false;
}
