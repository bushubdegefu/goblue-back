package manager

import (
	"github.com/spf13/cobra"
	"semay.com/bluerabbit"
)

var (
	backgroundCmd = &cobra.Command{
		Use:   "startconsumer",
		Short: "Start rabbit qeue task consumers Database",
		Long:  `Start RabbitMQ consumer`,
		Run: func(cmd *cobra.Command, args []string) {
			start_consumer()
		},
	}
)

func start_consumer() {

	bluerabbit.BlueConsumer()
}

func init() {
	goBlueCmd.AddCommand(backgroundCmd)

}
