package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func apiCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Start the API server",
		Run: func(cmd *cobra.Command, args []string) {
			r := gin.Default()
			r.GET("/ping", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
				})
			})
			r.Run()
		},
	}
}

func main() {
	rootCmd := &cobra.Command{Use: "radio-station"}
	rootCmd.AddCommand(apiCommand())
	rootCmd.Execute()
}
