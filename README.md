# Go Snake Go!

Nothing fancy, just a basic snake game written in Go; a VERY basic version of this.

I mostly made this for two reasons:

1) To learn Go! It is a fantastic language, and I have wanted to learn it more!
2) Generating the snake, and food were both interesting problems to me.

## Explination

Most snake games you see use arrays, and a lot of looping. People also seem to generate the food by randomly selecting a tile, and then checking if the snake is on it, and if so spawns food on the tile, or generates another one until an empty spot is found. This method works, but I am not a fan of it, as it is theoretically possible to never find an empty space like this! Especially once the snake begins to fill up the screen, generation can take longer and longer (technically).

My version differs in the fact that, instead of using arrays, uses linked lists. Linked lists also have the benefit of faster removal and additions than arrays. This is beneficial because all the pieces of the snake connect, and it is super easy to keep track of the head, and also grow the snake. Prepending to the front of array involves shifting all other elements to the right, with a linked list it is much easier.

For food generation this is nice because I can keep a list of empty spaces, and choose a random tile from this list, to get a tile. You are always guaranteed to get an open space this way. Then whenever a space opens up, we can add it back to the list, and when one gets taken away, we can simply remove it from the list (again in a faster way than arrays can).

## Missing features

There are many things that could be added to this, to imporve it as a game:

1) Score
2) Easier reseting
3) Difficulties
4) etc, etc...

However this has served well enough as a learning project, and whats left is more trivial/tedious to add than is worth (especially for a snake clone 😄).