package main

func registerReferenceFlags() {
	rootCommand.PersistentFlags().StringVarP(
		&referencePostgresHost,
		"reference-host", "A",
		getStrEnvDefault("POSTGRES_REFERENCE_HOST", "localhost"),
		"defines the reference PostgreSQL database host address (excluding port), defaults to `localhost` but can be derived from the POSTGRES_REFERENCE_HOST environment variable if it is specified and the flag is not present",
	)
	rootCommand.PersistentFlags().IntVarP(
		&referencePostgresPort,
		"reference-port", "P",
		getIntEnvDefault("POSTGRES_REFERENCE_PORT", 5432),
		"defines the reference PostgreSQL database host port, defaults to `5432` but can be derived from the POSTGRES_REFERENCE_PORT environment variable if it is specified and the flag is not present",
	)
	rootCommand.PersistentFlags().StringVarP(
		&referencePostgresUser,
		"reference-user", "U",
		getStrEnvDefault("POSTGRES_REFERENCE_USER", "postgres"),
		"defines the reference PostgreSQL database username, defaults to `postgres` but can be derived from the POSTGRES_REFERENCE_USER environment variable if it is specified and the flag is not present",
	)
	rootCommand.PersistentFlags().StringVarP(
		&referencePostgresPassword,
		"reference-password", "W",
		getStrEnvDefault("POSTGRES_REFERENCE_PASSWORD", ""),
		"defines the reference PostgreSQL database password, defaults to nothing (blank) but can be derived from the POSTGRES_REFERENCE_PASSWORD environment variable if it is specified and the flag is not present",
	)
	rootCommand.PersistentFlags().StringVarP(
		&referencePostgresDatabase,
		"reference-database", "D",
		getStrEnvDefault("POSTGRES_REFERENCE_DATABASE", "postgres"),
		"defines the reference PostgreSQL database, defaults to `postgres` but can be derived from the POSTGRES_REFERENCE_DATABASE environment variable if it is specified and the flag is not present",
	)
}

func registerDesiredFlags() {
	rootCommand.PersistentFlags().StringVarP(
		&desiredPostgresHost,
		"desired-host", "a",
		getStrEnvDefault("POSTGRES_DESIRED_HOST", "localhost"),
		"defines the desired PostgreSQL database host address (excluding port), defaults to `localhost` but can be derived from the POSTGRES_DESIRED_HOST environment variable if it is specified and the flag is not present",
	)
	rootCommand.PersistentFlags().IntVarP(
		&desiredPostgresPort,
		"desired-port", "p",
		getIntEnvDefault("POSTGRES_DESIRED_PORT", 5432),
		"defines the desired PostgreSQL database host port, defaults to `5432` but can be derived from the POSTGRES_DESIRED_PORT environment variable if it is specified and the flag is not present",
	)
	rootCommand.PersistentFlags().StringVarP(
		&desiredPostgresUser,
		"desired-user", "u",
		getStrEnvDefault("POSTGRES_DESIRED_USER", "postgres"),
		"defines the desired PostgreSQL database username, defaults to `postgres` but can be derived from the POSTGRES_DESIRED_USER environment variable if it is specified and the flag is not present",
	)
	rootCommand.PersistentFlags().StringVarP(
		&desiredPostgresPassword,
		"desired-password", "w",
		getStrEnvDefault("POSTGRES_DESIRED_PASSWORD", ""),
		"defines the desired PostgreSQL database password, defaults to nothing (blank) but can be derived from the POSTGRES_DESIRED_PASSWORD environment variable if it is specified and the flag is not present",
	)
	rootCommand.PersistentFlags().StringVarP(
		&desiredPostgresDatabase,
		"desired-database", "d",
		getStrEnvDefault("POSTGRES_DESIRED_DATABASE", "postgres"),
		"defines the desired PostgreSQL database, defaults to `postgres` but can be derived from the POSTGRES_DESIRED_DATABASE environment variable if it is specified and the flag is not present",
	)
}

func registerOtherFlags() {
	rootCommand.PersistentFlags().StringVarP(
		&outputMode,
		"output", "o",
		getStrEnvDefault("MIGRATE_OUTPUT_MODE", "text"),
		"defines how the migration should be output, valid values are: text, file; defaults to `text` but can be derived from the MIGRATE_OUTPUT_MODE environment variable if it is specified and the flag is not present",
	)
	rootCommand.PersistentFlags().StringVarP(
		&inputSchema,
		"input", "i",
		"",
		"provides the desired schema, this is applied to the desired database. This can be a filepath or can be - for stdin",
	)
	rootCommand.PersistentFlags().StringVarP(
		&outputDirectory,
		"folder", "f",
		getStrEnvDefault("MIGRATE_OUTPUT_FOLDER", "./"),
		"defines where generated schema migrations should be written, only valid if output is file; defaults to `./` but can be derived from the MIGRATE_OUTPUT_DIRECTORY environment variable if it is specified and the flag is not present",
	)
}

