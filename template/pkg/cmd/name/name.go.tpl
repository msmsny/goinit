package {{.Name}}

import (
	"{{.Repo}}/pkg/cmd/{{.Name}}/options"
	"{{.Repo}}/pkg/cmd/{{.Name}}/{{.SubName}}"
	"github.com/spf13/cobra"
)

func Execute() error {
	return New{{.UpperName}}Command().Execute()
}

func New{{.UpperName}}Command() *cobra.Command {
	opt := options.NewOption()
	cmds := &cobra.Command{
		Use:           "main command",
		Short:         "main command short description",
		Long:          "main command long description",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	// describe persistent flags
	pflags := cmds.PersistentFlags()
	pflags.StringVar(&opt.Opt1, "opt1", "default-value", "opt1 description")

	// initialize
	cobra.OnInitialize(func() {
		// ...
	})

	// add sub command
	cmds.AddCommand({{.SubName}}.New{{.UpperSubName}}Command(opt))

	return cmds
}
