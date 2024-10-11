# 1. Labyrinth

In this problem, you are required to implement a **single-server single-client architecture** for the **Labyrinth game**.

A Labyrinth is simply an **m × n grid** containing different types of tiles. In our version, there are **three types of tiles** possible:

- **Empty tile**: As the name suggests, there is nothing on this tile.
- **Coin tile**: This tile contains a coin. After the coin is collected, this tile will turn into an empty tile.
- **Wall tile**: This tile contains a wall, and it is not possible to visit this tile. This tile will be converted to an empty tile if the wall is destroyed.

The player will start at the **top left corner** of the Labyrinth, and the goal is to reach the **bottom right corner** while collecting as many coins as possible. In one move, the player can either move one tile or perform a spell. The player can move in **four directions**: up, down, left, or right. One of three things can happen:

- The player moved into an **empty tile**.
- The player moved into a tile with a **coin**. The coin will be collected, and the tile will become empty.
- The player moved into a **wall**. In this case, a collision will happen, and the player will lose a health point.

The player may also perform one of **two types of spells**:

- **Revelio**: This spell will be performed on a target tile with a specific tile type in mind. All the surrounding tiles (consider the **3 × 3 square** around the target) with the desired tile type (coin/wall) will be revealed.
- **Bombarda**: This spell will destroy all objects on **three tiles** of the player's choice, converting these tiles to empty tiles.

Initially, the player has **three health points** and can perform a total of **three spells**.

## 1.1 Server

To initialize the game, the server will read a Labyrinth from a file and initialize all player attributes. From then on, the server needs to keep track of all the player attributes and Labyrinth updates: the player score, player health, remaining spells, and the current position. A total of **five RPCs** need to be implemented:

- **GetLabyrinthInfo**: Take an empty request and provide the width and height of the Labyrinth.
- **GetPlayerStatus**: Take an empty request and provide the player's score, health points, current position in the Labyrinth, and the number of remaining spells.
- **RegisterMove**: Take a request containing the desired direction to move and move the player in that direction. Provide the status of the move informing the client whether the move 1. was a success, 2. was a failure, 3. led to victory, or 4. resulted in player death.
- **Revelio**: Take a request containing the target position of the spell and the desired tile type. Provide a stream of positions of all the tiles surrounding the target tile having the same tile type as the one in the request. Note that the target tile is also included under the spell effect.
- **Bombarda**: Take a stream of target positions as input. Destroy anything that may be present on the target tiles (wall or coin), converting them into empty tiles.

The request and response objects for each method must contain the attributes mentioned in its description. You cannot change the method names, and streaming must be used wherever it is mentioned. However, the exact implementation of these objects is up to you. You may include additional fields if you feel it is necessary.

## 1.2 Client

On the client side, you need to demonstrate and prove that each of the RPC methods implemented on the server works as expected. Cover all cases that can be encountered in the game. The minimum requirement is to provide an interface to call each RPC method through the terminal.

## 1.3 Bonus

This section can recover the points you lost in this problem only. Create a **terminal interface** through which the Labyrinth game can be played.
