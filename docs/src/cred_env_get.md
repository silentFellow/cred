## cred env get

Retrieve and store environment variables

### Synopsis

Retrieve environment variables from a specified file or default .env files and store them securely.

You can specify the file containing the environment variables using the -f flag. If the -f flag is not provided, the command will look for the following files in order: .env, .env.local, .env.development, .env.production, .env.test.

	Example usage: cred env get [flags: -f (filepath)]

```
cred env get [flags]
```

### Options

```
  -f, --file string   path to the file containing env
  -h, --help          help for get
```

### SEE ALSO

* [cred env](cred_env.md)	 - A command to manage env-variables

