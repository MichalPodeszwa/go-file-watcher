package main

import "flag"

type commandLineArgs struct {
	InputDir   string
	OutputDir  string
	OutputType string
}

func handleCliArguments() commandLineArgs {
	parsedCliArgs := commandLineArgs{}
	flag.StringVar(&parsedCliArgs.InputDir, "input", "input", "Defines what directory it should watch for the new files")
	flag.StringVar(&parsedCliArgs.OutputDir, "output", "output", "Defines what directory it should put the new files")
	flag.StringVar(&parsedCliArgs.OutputType, "type", "json", "Defines type of output files")

	flag.Parse()
	return parsedCliArgs
}

func (args commandLineArgs) GetDirectories() []string {
	return []string{args.InputDir, args.OutputDir}
}
