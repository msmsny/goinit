package {{.SubName}}

import (
	"{{.Repo}}/pkg/cmd/{{.Name}}/options"
	"github.com/spf13/cobra"
)

func New{{.UpperSubName}}Command(opt *options.Option) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "sub command",
		Short: "sub command",
		Long:  "sub command",
		RunE: func(cmd *cobra.Command, args []string) error {
			// implement sub command
			return nil
		},
		SilenceErrors: false,
		SilenceUsage:  false,
	}

	// describe flags
	flags := cmds.Flags()
	flags.StringVar(&opt.Opt2, "opt2", "default-value", "opt2 description")

	return cmds
}
