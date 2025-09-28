package utils

import "strings"

var blockedCommands = []string{
    "rm -rf /", 
    "mkfs", 
    ":(){:|:&};:", // known destructive commands
    "dd if=",
    "shutdown",
    "reboot",
}

func IsBlocked(gptCommand string) bool {
    for i := 0; i < len(blockedCommands); i++ {
        blockedCommand := blockedCommands[i]
        if (strings.Contains(gptCommand, blockedCommand)){
            return true
        }
    }
    return false
}
