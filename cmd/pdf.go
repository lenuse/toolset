/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lenuse/toolset/pkg/pdf"
	"github.com/spf13/cobra"
)

// pdfCmd represents the pdf command
var pdfCmd = &cobra.Command{
	Use:   "pdf",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args:      cobra.MinimumNArgs(2),
	ValidArgs: []string{"i", "c"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return pdf.AddBookmarks(pc)
	},
}

func init() {
	rootCmd.AddCommand(pdfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//pdfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	pdfCmd.Flags().StringVar(&pc.InfilePath, "i", "", "Help message for toggle")
	pdfCmd.Flags().StringVar(&pc.OutfilePath, "o", "", "Help message for toggle")
	pdfCmd.Flags().StringVar(&pc.BookmarksPath, "c", "", "Help message for toggle")

}

var pc pdf.PathConf
