package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Will support multiple file types, so needs to be more abstract than
type AudioFileMetadata struct {
	Path   string                  `json:"path"`
	Name   string                  `json:"name"`
	Type   string                  `json:"type"`
	Detail AudioFileMetadataDetail `json:"detail"`
}

type AudioFileMetadataDetail struct {
	SampleRate uint32 `json:"sampleRate"` // Hz
	BitRate    uint8  `json:"bitRate"`    // bits per sample
	Size       uint32 `json:"size"`       // bytes
}

func GetAudioFileMetadata(filePath string) (AudioFileMetadata, error) {
	file := AudioFileMetadata{}
	file.Path = filePath

	fileName, err := GetFileName(filePath)
	if err != nil {
		return file, err
	}
	file.Name = fileName

	fileType, err := GetFileType(filePath)
	if err != nil {
		return file, err
	}
	file.Type = fileType

	// Trying to do strategy pattern but this is unsightly
	detail, err := AudioFileMetadataDetail{}, nil
	switch file.Type {
	case "FLAC":
		detail, err = GetAudioFileDetailFlac(filePath)
	default:
		detail, err = AudioFileMetadataDetail{}, fmt.Errorf("Unimplemented file type")
	}
	if err != nil {
		return file, err
	}
	file.Detail = detail

	return file, nil
}

func GetFileName(filePath string) (string, error) {
	filePathSplit := strings.Split(filePath, "/")
	fileName := filePathSplit[len(filePathSplit)-1]
	if len(fileName) == 0 {
		return "", fmt.Errorf("Invalid file name")
	}

	return fileName, nil
}

func GetFileType(filepath string) (string, error) {
	filePathSplit := strings.Split(filepath, ".")
	if len(filePathSplit) < 2 {
		return "", fmt.Errorf("Could not identify the file type")
	}
	fileType := filePathSplit[len(filePathSplit)-1]

	return strings.ToUpper(fileType), nil
}

func (af AudioFileMetadata) Print() {
	b, err := json.MarshalIndent(af, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
}
