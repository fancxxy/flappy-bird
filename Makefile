MAIN = main

CC = g++


OBJS += surface.o \
		imagesurface.o \
		screensurface.o \
		background.o \
		timer.o \
		bird.o \
		number.o \
		pipe.o \
		playsense.o \
		getreadysense.o \
		game.o \
		main.o 

FLAGS += -std=c++11
FLAGS += $(shell sdl-config --cflags) $(shell sdl-config --libs)
FLAGS += -lSDL_image -lSDL_gfx -g

$(MAIN) : $(OBJS)
	$(CC) $(FLAGS) $(OBJS) -o $@

$(OBJS) : %.o : %.cpp
	$(CC) -c $(FLAGS) $< -o $@

clean:
	-rm -rf $(MAIN) $(OBJS)
