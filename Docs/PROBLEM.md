# Labyrinth Game - gRPC Implementation

## Introduction

This project implements a single-server, single-client architecture for the Labyrinth game using gRPC. The goal of the game is to navigate through a grid-based maze (Labyrinth) to reach the bottom-right corner, while collecting as many coins as possible and avoiding walls.

### Tile Types:

- **Empty Tile**: No object on this tile.
- **Coin Tile**: Contains a coin. Once collected, the tile turns into an empty tile.
- **Wall Tile**: Contains a wall that prevents movement. Can be destroyed to turn into an empty tile.

The player starts at the top-left corner, aiming to reach the bottom-right corner while collecting coins and avoiding walls. The player has three health points and can perform up to three spells during the game.

## Server

The server is responsible for managing the game state, including:

- The layout of the Labyrinth.
- Player status (health, score, position, remaining spells).
- Labyrinth updates as the player moves.

### RPC Methods:

1. **GetLabyrinthInfo**:

   - Request: Empty
   - Response: Returns the width and height of the Labyrinth.

2. **GetPlayerStatus**:

   - Request: Empty
   - Response: Returns the player’s score, health points, current position, and remaining spells.

3. **RegisterMove**:

   - Request: The desired direction (up, down, left, right).
   - Response: Provides the status of the move (success, failure, victory, or death).

4. **Revelio (Spell)**:

   - Request: Target position and the desired tile type (coin/wall).
   - Response: A stream of all positions surrounding the target that match the desired tile type.

5. **Bombarda (Spell)**:
   - Request: A stream of target positions.
   - Response: Destroys the objects on the selected tiles (wall/coin), turning them into empty tiles.

The request and response objects must adhere to the descriptions above. Streaming should be used where indicated.

## Client

The client will interact with the server using the provided gRPC methods. It should:

- Display the player's progress (score, health, position).
- Allow the player to move and perform spells.
- Handle all possible game scenarios (victory, death, etc.).

### Minimum Requirement:

- A terminal-based interface to call each RPC method and interact with the game.

## Bonus Task:

- Create a fully interactive terminal interface to play the Labyrinth game.

---
