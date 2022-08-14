package mp3

import (
	_ "embed"
	"encoding/binary"
	"math"
	"unicode/utf8"
)

const (
	sampleFreq        = 44100
	libraryLetterSecs = 0.15
	outputLetterSecs  = 0.075
)

var (
	//go:embed res/animalese.wav
	letterLibrary []uint8

	librarySamplesPerLetter = int32(math.Floor(libraryLetterSecs * sampleFreq))
	outputSamplesPerLetter  = int(math.Floor(outputLetterSecs * sampleFreq))
)

type header struct {
	ChunkId       []byte `json:"chunkId"`
	ChunkSize     uint32 `json:"chunkSize"`
	Format        []byte `json:"format"`
	SubChunk1Id   []byte `json:"subChunk1Id"`
	SubChunk1Size uint32 `json:"subChunk1Size"`
	AudioFormat   uint16 `json:"audioFormat"`
	NumChannels   uint16 `json:"numChannels"`
	SampleRate    uint32 `json:"sampleRate"`
	ByteRate      uint32 `json:"byteRate"`
	BlockAlign    uint16 `json:"blockAlign"`
	BitsPerSample uint16 `json:"bitsPerSample"`
	SubChunk2Id   []byte `json:"subChunk2Id"`
	SubChunk2Size uint32 `json:"subChunk2Size"`
}

// CreateSoundData Will take audio data and wrap it, as appropriate to create a valid mp3 file
func CreateSoundData(data []byte, audioPitch float64) ([]byte, error) {
	animalizedData := englishToAnimalese(data, audioPitch)

	h := createHeader(animalizedData)

	return serialize(h, animalizedData), nil
}

// englishToAnimalese Will take a set of text and bit shift it accordingly to map to the audio aspects of the animalese audio bytes
func englishToAnimalese(data []byte, pitch float64) []byte {
	convertedData := make([]byte, len(data)*outputSamplesPerLetter)

	for strInd := 0; strInd < len(data); strInd++ {
		c := data[strInd]
		r, _ := utf8.DecodeRune([]byte{c})
		if c >= 'A' && c <= 'Z' {
			libraryLetterStart := librarySamplesPerLetter * (r - ('A'))

			for i := 0; i < outputSamplesPerLetter; i++ {
				pitchOffset := int32(math.Floor(float64(i) * pitch))
				convertedData[strInd*outputSamplesPerLetter+i] =
					letterLibrary[44+libraryLetterStart+pitchOffset]
			}
		} else { // non-pronounceable character or space
			for i := 0; i < outputSamplesPerLetter; i++ {
				convertedData[strInd*outputSamplesPerLetter+i] = 127
			}
		}
	}

	return convertedData
}

func createHeader(data []byte) *header {
	h := newHeader()

	h.SampleRate = sampleFreq
	h.BlockAlign = (h.NumChannels * h.BitsPerSample) >> 3
	//h.ByteRate = uint32(h.BlockAlign) * h.SampleRate // this line in the code this is based does not do what it should
	h.SubChunk2Size = uint32(len(data) * int(h.BitsPerSample>>3))
	h.ChunkSize = 36 + h.SubChunk2Size

	return h
}

func newHeader() *header {
	return &header{ // OFFS SIZE NOTES
		ChunkId:       []byte{0x52, 0x49, 0x46, 0x46}, // 0    4    "RIFF" = 0x52494646
		ChunkSize:     0,                              // 4    4    36+SubChunk2Size = 4+(8+SubChunk1Size)+(8+SubChunk2Size)
		Format:        []byte{0x57, 0x41, 0x56, 0x45}, // 8    4    "WAVE" = 0x57415645
		SubChunk1Id:   []byte{0x66, 0x6d, 0x74, 0x20}, // 12   4    "fmt " = 0x666d7420
		SubChunk1Size: 16,                             // 16   4    16 for PCM
		AudioFormat:   1,                              // 20   2    PCM = 1
		NumChannels:   1,                              // 22   2    Mono = 1, Stereo = 2...
		SampleRate:    8000,                           // 24   4    8000, 44100...
		ByteRate:      0,                              // 28   4    SampleRate*NumChannels*BitsPerSample/8
		BlockAlign:    0,                              // 32   2    NumChannels*BitsPerSample/8
		BitsPerSample: 8,                              // 34   2    8 bits = 8, 16 bits = 16
		SubChunk2Id:   []byte{0x64, 0x61, 0x74, 0x61}, // 36   4    "data" = 0x64617461
		SubChunk2Size: 0,                              // 40   4    data size = NumSamples*NumChannels*BitsPerSample/8
	}
}

func serialize(h *header, data []byte) []byte {
	ad := make([]byte, 0, len(data))

	ad = append(ad, h.ChunkId...)
	ad = append(ad, u32ToArray(h.ChunkSize)...)
	ad = append(ad, h.Format...)
	ad = append(ad, h.SubChunk1Id...)
	ad = append(ad, u32ToArray(h.SubChunk1Size)...)
	ad = append(ad, u16ToArray(h.AudioFormat)...)
	ad = append(ad, u16ToArray(h.NumChannels)...)
	ad = append(ad, u32ToArray(h.SampleRate)...)
	ad = append(ad, u32ToArray(h.ByteRate)...)
	ad = append(ad, u16ToArray(h.BlockAlign)...)
	ad = append(ad, u16ToArray(h.BitsPerSample)...)
	ad = append(ad, h.SubChunk2Id...)
	ad = append(ad, u32ToArray(h.SubChunk2Size)...)
	ad = append(ad, data...)

	return ad
}

func u32ToArray(i uint32) []byte {
	a := make([]byte, 4)
	binary.LittleEndian.PutUint32(a, i)
	return a
}

func u16ToArray(i uint16) []byte {
	a := make([]byte, 2)
	binary.LittleEndian.PutUint16(a, i)
	return a
}
