## cred completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	cred completion fish | source

To load completions for every new session, execute once:

	cred completion fish > ~/.config/fish/completions/cred.fish

You will need to start a new shell for this setup to take effect.


```
cred completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [cred completion](cred_completion.md)	 - Generate the autocompletion script for the specified shell

