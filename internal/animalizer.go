package animalizer

import (
	"civilrights3.com/animalese/internal/filesystem"
	"civilrights3.com/animalese/internal/mp3"
	"strings"
)

func Convert(params Parameters) error {
	// get the data
	data, err := getData(params)
	if err != nil {
		return err
	}

	// process the data
	cleanData := cleanText(data)

	// convert to animalese
	// wrap in mp3 headers
	audioData, err := mp3.CreateSoundData(cleanData, params.AudioPitch)
	if err != nil {
		return err
	}

	// write to file
	err = filesystem.SaveFile(params.DestPath, audioData, params.OverwriteSave)
	if err != nil {
		return err
	}

	return nil
}

func getData(params Parameters) ([]byte, error) {
	if d := params.OpenText; d != "" {
		return []byte(d), nil
	}

	path := params.SourcePath
	b, err := filesystem.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// cleanText Will take text input and replace special characters with space, upper case all letters
func cleanText(data []byte) []byte {
	dataAsStr := string(data)

	dataAsStr = strings.ToUpper(dataAsStr)
	dataAsStr = strings.Map(func(r rune) rune {
		if r < 'A' || r > 'Z' {
			return 32 // 32 is the rune for space
		}
		return r
	}, dataAsStr)

	return []byte(dataAsStr)
}
