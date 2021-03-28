# Codagotchi

This is a little game experiment where the universe is slowly dying, and your goal is to write code to save it.
The game is hosted here: https://codagotchi.edznux.fr
To participate in this experiment, submit a PR, when it's merged, it will be immediatly deployed (should be less than 3min)

## Rules and goal of the game

- Try to keep our little universe alive and running
- Every solution to save our little friends MUST have downsides that also kills them.

Example ideas:
- Hunger (takes damage if hungry, must find food sources, slows down if it eats to much)
- Sleep schedule (need to sleep, takes damage when sleep deprived)
- Takes damage if it doesn't move enough

PR can also be visual aspects only:
- World background (different dimensions possible ?)
- Day / night cycles

## Build & Run

### Requirements:

- go: 1.16+ (We are using the new `embeds` package)
- Ebiten dependencies: https://ebiten.org/documents/install.html

Build everything you need: 
```bash
./build.sh
```

Run the compiled binary:
```bash
./codagotchi
# or the web version (starts a webserver on port 8080)
./codagotchi web
```

Alternatively, you can use docker (recommendend only on servers, only supports web version):
```bash
docker build -t codagotchi:latest

docker run -d \
--net host \
-v $(pwd)/save.json:/app/save.json \
--name codagotchi \
codagotchi:latest
```

## License

Copyright (c) 2021 Edznux and Contributors

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.