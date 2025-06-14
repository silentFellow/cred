## cred env set

Sets environment variables from the cred-store into a .env file

### Synopsis

Retrieve environment variables for a specified file from the cred-store and write them to a local .env file.

You provide the file name as an argument. The command fetches the environment variables associated with that file from the cred-store and writes them into a new or existing .env file in your current directory.

	Example usage: cred env set <filepath>

```
cred env set [flags]
```

### Options

```
  -h, --help   help for set
```

### SEE ALSO

* [cred env](cred_env.md)	 - A command to manage env-variables

