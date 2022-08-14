# animaleseGolang
A generator to create an animalese (animal crossing) sound effect and save it to a file

## Usage

Command line flags:
- `-t` string value of the text to be converted. for strings with spaces, enclose in quotation marks "" (use only -t OR -f)
- `-f` name and path of the file which contains the text to be converted
- `-p` the pitch of the animalese'd voice (0.2 - 2) the lower the number the deeper the voice - default: 1
- `-o` name and path of the file desired for output - default: `anim.wav`
- `-Y` force overwriting the output file if it already exists - default: false


### Example
For creating a file from text input:

`animalese.exe -t="this is some sample text" -p="1.8" -o="samples/test_file.wav" -Y=true`

For creating a file from a source text file:

`animalese.exe -f="bad_idea/the_script_of_the_bee_movie.txt" -p="0.6" -o="samples/bee_movie.wav" -Y=true`

## Building

The code requires Go v1.19 or higher. It can be built with `go build -o=bin/animalese ./cmd/main.go`. Specify your OS with the `GOOS` environment variable
