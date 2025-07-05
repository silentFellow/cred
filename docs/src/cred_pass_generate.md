## cred pass generate

Generate a new password and store it securely

### Synopsis

The generate command creates a new password of specified length and stores it securely in the password store.
You can specify the length of the password using the -l flag. If the file already exists, you will be prompted to overwrite it.

Examples:
  pass generate mypassword -l 16
  pass generate anotherpassword -l 24

```
cred pass generate [flags]
```

### Options

```
      --allow-digit              should allow digits in the password (default true)
      --allow-lowercase          should allow lower-case characters in the password (default true)
      --allow-special            should allow special characters in the password (default true)
      --allow-uppercase          should allow upper-case characters in the password (default true)
      --allowed-special string   allowed special characters in the password (default "!@#$%^&*()-_=+[]{}|;:,.<>?/`~")
  -e, --editor                   open password in editor for editing extra details after insertion
  -h, --help                     help for generate
  -l, --length int               length of the generated password (default 12)
```

### SEE ALSO

* [cred pass](cred_pass.md)	 - A command to manage passwords

