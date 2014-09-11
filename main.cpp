#include "game.h"

int FRAMES_PER_SECOND = 30;   //限制fps

int main(int argc, char* argv[]) {
    Game* game = new Game();
    game->gameRun();

    return 0;
}
