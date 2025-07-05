## Features
- Map parsing from file
- Adventurer movement and orientation
- Obstacles: mountains/other adventurers
- Treasure collection
- Simulation result output

## Project Structure
- `main.go`: Entry point, runs the simulation
- `game_service.go`: Core game logic, including map parsing and player movement
- `domain/`: Core domain entities
  - `constants/`: Game and error constants
  - `enum/`: Enums for player moves, orientation...
  - `types/`: Data structures for adventurers, board, cells, mountains, treasures
- `tests/`: Unit tests and test maps
- `utils/`: Utility functions (file I/O, move logic)


### Run

```sh
make FILEPATH=<path_to_initial_map>  
```
Or manually:
```sh
go run main.go <path_to_initial_map>
```

## Testing
Run all tests:
```sh
make test
```

## Linting
Run linter:
```sh
make lint
```
