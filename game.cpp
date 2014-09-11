#include "game.h"

Game::Game() : gamequit(false), gameover(true), gamestart(false) {
    screen = new ScreenSurface(705, 512);
    screen->setTitle("Flappy Bird");
    image = new ImageSurface("./resource/images/atlas.png");
    getreadysense = new GetReadySense;
}

Game::~Game() {
    if (gamestart == false && gameover == true)
        delete getreadysense;
    else if (gamestart == true && gameover == false)
        delete playsense;
    delete screen;
    delete image;
}

void Game::gameRun() {
    while (gamequit == false) {

        if (gamestart == false && gameover == true) {
            GetReadySense* getreadysense = new GetReadySense;
            getreadysense->showGetReadySense(screen, image);

            gamestart = getreadysense->gameStart();
            gamequit = getreadysense->gameQuit();

            if (gamestart == true) {
                gameover = false;
                delete getreadysense;
                playsense = new PlaySense;
            }
        } else if (gamestart == true && gameover == false) {
            playsense->showPlaySense(screen, image);

            gameover = playsense->gameOver();
            gamequit = playsense->gameQuit();

            if (gameover == true) {
                gamestart = false;
                delete playsense;
                getreadysense = new GetReadySense;
            }
        }
    }
}
