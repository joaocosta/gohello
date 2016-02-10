# gohello
Limited and Simple echo server implemented in go, as part of a personal learning effort.

Based on https://howistart.org/posts/go/1

# Usage
## Start server
    git clone https://github.com/joaocosta/gohello.git
    cd gohello
    go build
    ./gohello

## Send client requests
    curl -X POST -d '{"message": "Hello World"}' http://localhost:49001/
