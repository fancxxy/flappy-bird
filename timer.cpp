#include "timer.h"

Timer::Timer() {
    startTicks = 0;
    pauseTicks = 0;
    started = false;
    paused = false;
}

void Timer::start() {
    started = true;
    paused = false;
    startTicks = SDL_GetTicks();
}

void Timer::stop() {
    started = false;
    paused = false;
}

void Timer::pause() {
    if (started && !paused) {
        paused = true;
        pauseTicks = SDL_GetTicks() - startTicks;
    }
}

void Timer::unpause() {
    if (paused) {
        paused = false;
        startTicks = SDL_GetTicks() - pauseTicks;
        pauseTicks = 0;
    }
}

int Timer::getTicks() {
    if (started) {
        if (paused)
            return pauseTicks;
        else
            return SDL_GetTicks() - startTicks;
    }
    return 0;
}

bool Timer::isStarted() {
    return started;
}

bool Timer::isPaused() {
    return paused;
}
