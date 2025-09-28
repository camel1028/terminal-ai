package cmd

import (
	"os"
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"github.com/camel1028/terminal-ai/openai"
    "github.com/camel1028/terminal-ai/openai/utils"
    "github.com/fatih/color"
)


func ParseGPTOutput(raw string) (string, string) {
    parts := strings.SplitN(raw, "Explanation:", 2)
    command := strings.TrimPrefix(parts[0], "Command:")
    explanation := ""
    if len(parts) > 1 {
        explanation = parts[1]
    }
    return strings.TrimSpace(command), strings.TrimSpace(explanation)
}


var rootCmd = &cobra.Command{
    Use:   "goshell",
    Short: "Natural language to shell command assistant",
    Run: func(cmd *cobra.Command, args []string) {
        input := strings.Join(args, " ")

        // terminal colors
        green := color.New(color.FgGreen).SprintFunc()
        yellow := color.New(color.FgYellow).SprintFunc()
        red := color.New(color.FgRed).SprintFunc()

        fmt.Println(utils.IsBlocked)

        if input == "" {
            fmt.Println("Please provide a command.")
            return
        }

        fmt.Println("Asking GPT for:", yellow(input))

        result, err := openai.AskGPT(input)
        if err != nil {
            fmt.Println(red("GPT Error:", err))
            return
        }
        if utils.IsBlocked(input) {
            fmt.Println(red("dangerous command. Execution blocked"))
            return
        }

        command, explanation := ParseGPTOutput(result)

        green(fmt.Println("GPT Response:\n"))
        fmt.Println("Command: ", green(command))
        fmt.Println("Explanation: ", green(explanation))
    },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.terminal-ai.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


