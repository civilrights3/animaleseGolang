package animalizer_test

import (
	animalizer "civilrights3.com/animalese/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestSanity(t *testing.T) {
	assert.True(t, true)
}

func TestA(t *testing.T) {
	expectedFile := "../tests/a.wav"
	expectedPayload, err := os.ReadFile(expectedFile)
	require.NoError(t, err)

	targetFile := "../tests/output/anim-test-a.wav"

	params := animalizer.Parameters{
		OpenText:      "a",
		SourcePath:    "",
		DestPath:      targetFile,
		AudioPitch:    1,
		OverwriteSave: true,
	}

	err = animalizer.Convert(params)
	assert.NoError(t, err)
	actual, err := os.ReadFile(targetFile)
	assert.NoError(t, err)

	for i, _ := range actual {
		assert.Equal(t, expectedPayload[i], actual[i], "Payloads differ at byte %d", i)
	}
	assert.Equal(t, len(expectedPayload), len(actual))
	assert.Equal(t, expectedPayload, actual)
}
