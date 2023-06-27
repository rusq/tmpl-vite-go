# Template for Vite-Go

Heavily based on sample program for [vite-go][1] library, the difference is
that the vite-go integration is only used for the development environment.
For production it relies on httpex to serve the dist static files, as I had
trouble running vite-go in production mode (in particular, it has some
problems finding the files and serving assets).

It relies on build tags.  To build development version using go command,
use the following command:

```sh
go build -tags=devel
```

Production is built by default.

On the first `make` execution it will initialise the vite-vue environment.
It is also possible to initialise the latest Vue environment by running:
```sh
npm create vue@latest
```
then make will not attempt to initialise the environment, but use the existing.

## TODO
- [ ] there is a problem with the vite-go dev server, it does not serve the
      index.html file, so the page is blank.  It is possible to work around
      this by using the httpex server, but it would be nice to fix this.
- [ ] Makefile is a mess.  It should be possible to simplify it.

## Usage

Default port for vite is 5173, and for the go application: 4000.

```sh
make dev
```

starts the dev environment.  To start dev manually:
```sh
cd frontend
npm run dev &
cd ..
go run -tags=devel .
```

```sh
make stop_dev
```
stops the dev environment.

```sh
make build
```

builds the production version

[1]: https://github.com/torenware/vite-go
