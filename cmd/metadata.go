/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/opera22/audiofile/utils"
	"github.com/spf13/cobra"
)

// metadataCmd represents the metadata command
var metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Get metadata for an audio file. Only accepts .flac files.",
	Long:  `Get metadata for an audio file. Only accepts .flac files.`,
	Run: func(cmd *cobra.Command, args []string) {
		filepath := args[0]
		fmt.Println("metadata called")
		fmt.Println("You entered a filepath of", filepath)
		audioFileMetadata, err := utils.GetAudioFileMetadata(filepath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		audioFileMetadata.Print()
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(metadataCmd)
}
