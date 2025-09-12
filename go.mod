module beckx.online/yaatt

go 1.24.6

require (
	beckx.online/butils/fileutils v0.0.0-00010101000000-000000000000
	github.com/go-flac/flacvorbis/v2 v2.0.2
	github.com/go-flac/go-flac/v2 v2.0.4
	github.com/rs/zerolog v1.34.0
	github.com/spf13/cobra v1.10.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	golang.org/x/sys v0.12.0 // indirect
)

replace beckx.online/butils/fileutils => ../butils/fileutils/
