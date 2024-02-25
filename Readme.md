# Game of life

## Rules

- Underpopulation: An alive cell with 0-1 live neighbors will die
- Right conditions: An alive cell with 2-3 live neighbors will be kept alive
- Overpopulation: An alive cell with more than 3 live neighbots will die
- Regeneration: A dead cell with exactly 3 live neighbors will live

# App
## Run

Run the game of life directly with go

`go run main.go`

## Tests

Execute all the test cases

`go test`

# Remaining tasks

- [X] run it forever and update results in console
- [ ] Review pending features from the origianl documentation
- [ ] refactor
    - [ ] Modularize the code (Board Functionality != Game of Life Rules) 
        - Is there a generic implementation to inject the Cell
        - How to restrict the types to be initialized just with specific functions
        - How to package the code?
- [ ] Board: connect edges so every cell will have 8 neighbors
- [ ] Refactor 2:
    - [ ] Memory Optimization? 
        - Less memory consumption vs Inmutability
    - [ ] Process optimization
        - What if the cell can reffer its neighboards to determine its next state? New Matrix Structure
- [ ] Create CLI tool
- [ ] CLI Visor (part with the cell, part with changes)
- [ ] Read the initial state from a file (CLI)
- [ ] Create Docker file
- [ ] Create WebHook
- [ ] Create Ract Webhook component

# References

[Programming Projects for Advanced Beginners #2: Game of Life](https://robertheaton.com/2018/07/20/project-2-game-of-life/)