/*
Copyright © 2023 IITS-Consulting

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals // commands as globals is the cleanest and easiest way to create (sub)commands
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: fmt.Sprintf(`To load completions:

Bash:

    $ source <(%[1]s completion bash)

    # To load completions for each session, execute once:
    # Linux:
    $ %[1]s completion bash > /etc/bash_completion.d/%[1]s
    # macOS:
    $ %[1]s completion bash > $(brew --prefix)/etc/bash_completion.d/%[1]s

Zsh:

    # If shell completion is not already enabled in your environment,
    # you will need to enable it.  You can execute the following once:

    $ echo "autoload -U compinit; compinit" >> ~/.zshrc

    # To load completions for each session, execute once:
    $ %[1]s completion zsh > "${fpath[1]}/_%[1]s"

    # You will need to start a new shell for this setup to take effect.

fish:

    $ %[1]s completion fish | source

    # To load completions for each session, execute once:
    $ %[1]s completion fish > ~/.config/fish/completions/%[1]s.fish

PowerShell:

    PS> %[1]s completion powershell | Out-String | Invoke-Expression

    # To load completions for every new session, run:
    PS> %[1]s completion powershell > %[1]s.ps1
    # and source this file from your PowerShell profile.
`, "charter"),
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		switch args[0] {
		case "bash":
			err = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			err = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			err = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			err = cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
		return err
	},
}
