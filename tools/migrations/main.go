package main

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

var (
	rootCommand = &cobra.Command{
		Use:   "migrate",
		Short: "migrate is a tool to generate schema migrations",
		Long: "migrate is a tool to generate schema migrations for the harder than it needs to be application\n" +
			"it will compare the reference database to the desired database and generate schema changes needed to\n" +
			"make the reference database match the desired schema",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			generateDiff()
		},
	}

	referencePostgresHost     string
	referencePostgresPort     int
	referencePostgresUser     string
	referencePostgresPassword string
	referencePostgresDatabase string

	desiredPostgresHost     string
	desiredPostgresPort     int
	desiredPostgresUser     string
	desiredPostgresPassword string
	desiredPostgresDatabase string

	outputMode string

	inputSchema     string
	outputDirectory string
	author          string
)

func init() {
	registerReferenceFlags()
	registerDesiredFlags()
	registerOtherFlags()
}

func getStrEnvDefault(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}

	return value
}

func getIntEnvDefault(name string, defaultValue int) int {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		log.Fatalf(
			"invalid %s environment variable, expected valid integer found; `%s` - %+v",
			name, value, errors.Wrap(err, "failed to parse value"),
		)
		return -1
	}

	return int(intValue)
}

func main() {
	rootCommand.Execute()
}
