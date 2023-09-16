package utils

import (
	"os"

	"github.com/go-flac/go-flac"
)

func GetAudioFileDetailFlac(filePath string) (AudioFileMetadataDetail, error) {
	f, err := flac.ParseFile(filePath)
	if err != nil {
		return AudioFileMetadataDetail{}, err
	}

	streamInfo, err := f.GetStreamInfo()
	if err != nil {
		return AudioFileMetadataDetail{}, err
	}

	osFileInfo, err := os.Stat(filePath)
	if err != nil {
		return AudioFileMetadataDetail{}, err
	}

	return AudioFileMetadataDetail{
		SampleRate: uint32(streamInfo.SampleRate),
		BitRate:    uint8(streamInfo.BitDepth),
		Size:       uint32(osFileInfo.Size()),
	}, nil
}
