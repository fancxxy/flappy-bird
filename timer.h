#ifndef _TIMER_H_
#define _TIMER_H_

#include <SDL/SDL.h>

class Timer {
    private:
        int startTicks;
        int pauseTicks;
        bool started;
        bool paused;
    public:
        Timer();
        void start();
        void stop();
        void pause();
        void unpause();

        int getTicks();

        bool isStarted();
        bool isPaused();
};

#endif
