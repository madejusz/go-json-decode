package main

import (
    "flag"
    "fmt"
    "os"
    "encoding/json"
)

func main() {

    jsonEnv  := flag.String("env", "", "Environment variable with json data inside (required).")
    shell    := flag.String("shell", "/bin/bash", "Shell type {bash|csh};")
    debug    := flag.Bool("debug", false, "Turn on debug mode")
    flag.Parse()

    if *debug {
        fmt.Printf("# Run flags: shell=%s, debug=%t, validate=%t, env=%s\n", *shell, *debug, *jsonEnv)
    }

    if *jsonEnv == "" {
        fmt.Printf("# AWS unpack JSON Secret into variable 1.0.0\n")
        fmt.Printf("# Copyright (C) 2020 Jakub K. Boguslaw\n\n")
        flag.PrintDefaults()

        fmt.Printf("# Example: `eval %s -env ENV_WITH_JSON -shell /bin/bash`\n", os.Args[0])
        os.Exit(1)
    }

    jsonEnvVal := os.Getenv(*jsonEnv)

    var result map[string]interface{}
    err := json.Unmarshal([]byte(jsonEnvVal), &result)

    if err != nil {
            fmt.Printf("# ERROR: %s (%s)\n", err, *jsonEnv)
    }
    for k, v := range result {
        if *shell == "/bin/csh" {
            fmt.Printf("setenv %s %s\n", k, v)
        } else {
            fmt.Printf("export %s=%s\n", k, v)
        }
    }

    if *debug {
        fmt.Printf("# Ending....\n")
    }
}
