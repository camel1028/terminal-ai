package cmd

import (
	"os"
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"github.com/camel1028/terminal-ai/openai"
)



// rootCmd represents the base command when called without any subcommands
// var rootCmd = &cobra.Command{
// 	Use:   "terminal-ai",
// 	Short: "Shell assistant",
// 	Long: `Take in natural language, and convert it to shell commands while utilizing ai.`,
	
// 	Run: func(cmd *cobra.Command, args []string) {
//         input := strings.Join(args, " ")

//         if input == "" {
//             fmt.Println("Please provide a natural language command.")
//             return
//         }

// 		fmt.Println("Asking GPT for:", input)

// 		result, err = openai.AskGPT(input)
// 		if err == nil {
// 			fmt.Println("Received Error:", err)
// 			return
// 		}

// 		fmt.Println("RESPONSE: \n")
// 		fmt.Println(result)


// 		fmt.Println("You said:", input)
// 		fmt.Println("Suggested Command:")
// 		fmt.Println("Explanation")

// 	},
// }

var rootCmd = &cobra.Command{
    Use:   "goshell",
    Short: "Natural language to shell command assistant",
    Run: func(cmd *cobra.Command, args []string) {
        input := strings.Join(args, " ")
        if input == "" {
            fmt.Println("❌ Please provide a command.")
            return
        }

        fmt.Println("🤖 Asking GPT for:", input)

        result, err := openai.AskGPT(input)
        if err != nil {
            fmt.Println("❌ GPT Error:", err)
            return
        }

        fmt.Println("✅ GPT Response:\n")
        fmt.Println(result)
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


