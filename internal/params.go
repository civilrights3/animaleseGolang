package animalizer

import (
	_ "embed"
	"errors"
	"flag"
)

type Parameters struct {
	OpenText, SourcePath, DestPath string
	AudioPitch                     float64
	OverwriteSave                  bool
}

func ProcessParams() (Parameters, error) {
	// -t raw text
	// -f file to read
	// -p pitch (0.2 - 2)
	// -o output file path
	// -Y overwrite output file if exists
	openText := flag.String("t", "", "raw text to be converted to animalese (cannot use t and f at the same time)")
	filePath := flag.String("f", "", "file to read input text from (cannot use f and t at the same time)")
	audioPitch := flag.Float64("p", 1, "pitch of the animalese voice (0.2 - 2)")
	destPath := flag.String("o", "anim.wav", "path to output the file")
	overwriteSave := flag.Bool("Y", false, "overwrite output file if it already exists (returns error otherwise)")
	flag.Parse()

	p := Parameters{
		OpenText:      *openText,
		SourcePath:    *filePath,
		DestPath:      *destPath,
		AudioPitch:    *audioPitch,
		OverwriteSave: *overwriteSave,
	}

	return p, p.validate()
}

func (p Parameters) validate() error {
	text := p.OpenText
	file := p.SourcePath
	switch {
	case text == "" && file == "":
		return errors.New("one of -t or -f flags must be provided")
	case text != "" && file != "":
		return errors.New("only one of -t or -f flags can be provided, not both")
	case p.AudioPitch > 2.0 || p.AudioPitch < 0.2:
		return errors.New("audio pitch flag -p be between the values 0.2 and 2.0 inclusively")
	}
	return nil
}
