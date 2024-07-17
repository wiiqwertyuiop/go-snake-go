# Go Snake Go!

<img src="https://github.com/user-attachments/assets/0e8981a0-5f76-4fd7-ab9c-62c3c29208e7" width="300">

Nothing fancy, just a basic snake game written in Go; a VERY basic version of this.

I mostly made this for two reasons:

1) To learn Go! It is a fantastic language, and I have wanted to learn it more!
2) Generating the snake, and food were both interesting problems to me.

## Explanation

When you think about how to generate food on the map, you realize you need to spawn the food, WITHOUT overlapping on the snake. A lot of people seem to solve this by randomly selecting a tile, and then checking if a snake tile is on it; and if so spawns food on the tile, or selects another random space until an empty spot is found.

This method works, but I am not a fan of it. It is theoretically possible to never find an empty space like this! Especially once the snake begins to fill up the screen, generation can take longer and longer (technically). 

 My version differs in the fact that, instead of using arrays (like you will see in most snake games), we can use linked lists. ~~Linked lists have the known benefit of faster removal and item additions than arrays.~~ This is not true! However it was still fun doing it this way!

 For food generation this is nice because I can keep a list of empty spaces, and choose a random tile from this list, to get a valid tile. Whenever a space gets taken away, we can remove it from this list, and when one opens up, we can add it back to the list! You are always guaranteed to get a free open space in this way. This is also beneficial for the snake, because all of the pieces connect with each other.

## Missing features

There are many things that could be added to this, to improve it as a game:

1) Score display
2) Easier resetting
3) Difficulties
4) Sound
5) Few more optimizations
6) etc, etc...

However this has served well enough as a learning project, and whats left is more trivial/tedious to add than is worth (especially for a snake clone 😄).
