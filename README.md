## forego

<a href="https://travis-ci.org/heroku/forego">
  <img align="right" src="https://travis-ci.org/heroku/forego.svg?branch=master">
</a>

Foreman in Go.

### Installation

##### OS X (Homebrew)

    brew install forego

##### Compile from Source

    $ go get -u github.com/heroku/forego

### Usage

    $ cat Procfile
    web: bin/web start -p $PORT
    worker: bin/worker queue=FOO

    $ forego start
    web    | listening on port 5000
    worker | listening to queue FOO

### License

Apache 2.0 &copy; 2015 David Dollar
