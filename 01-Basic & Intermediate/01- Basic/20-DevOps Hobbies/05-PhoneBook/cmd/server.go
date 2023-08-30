package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"AcronPhoneBook/Internal/config"
)

type Server struct{}

func (cmd Server) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string)  {
		cmd.run(&config.Config{}, trap)
	}
	return &cobra.Command{
		Use: "server",
		Short: "run AcronPhoneBook server",
		Run: run,
	}
}

func (cmd *Server) run(cfg *config.Config, trap chan os.Signal) {

}
