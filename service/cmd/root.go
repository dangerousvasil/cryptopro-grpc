package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	cobra.OnInitialize(initConfig)
}

var rootCmd = &cobra.Command{
	Use:   "cpcrypto-service",
	Short: "A crypto pro service",
	Long:  `Программное обеспечение, необходимое для работы с электронной подписью`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	binFile string
)

func initConfig() {
	log.SetFlags(0)
	// узнаём бинфайл процесса
	var err error = nil
	binFile, err = os.Readlink("/proc/self/exe")
	if err != nil || binFile == "" {
		log.Fatalln(err)
	}

}
