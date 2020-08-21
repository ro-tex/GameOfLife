# GameOfLife

This is one of those "2am, lying awake" ideas.

The idea:
- write a Go library that calculates a generation of the [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)
  - rectangular world, grows in all directions simultaneously, has pre-defined max size
  - Seed(), NextGen(), SkipNGens(n uint)
- compile to WASM
- (optional) build with Tiny Go
- publish to NPM
- (optional) make available to Deno
- build a single page JS application that opens a canvas and uses the library to draw generations of the game
  - set the size of each cell (1-10 pixels)
  - set the number of generations per second (sane default)
  - load a starting set by uploading a file (simple text grid of 1s and 0s will do, nothing fancy)
  - pause/resume, speed up and slow down buttons (transparency will be important)
  - save and download?
  - grey background, orange life
- host on github pages and Skynet

Drawbacks:
- communicating between JS and WASM on each generation will be a severe performance hit. It might be faster to just implement in JS/TS but that kind of kills most of the idea (which is to play with Go and WASM).

Stretch goals:
 - check this: https://www.reddit.com/r/golang/comments/ial219/tool_by_go_team_to_get_go_doc_badge_for_your/
