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
      --allow-digit     should allow digits in the password (default true)
      --allow-lower     should allow lower-case characters in the password (default true)
      --allow-special   should allow special characters in the password (default true)
      --allow-upper     should allow upper-case characters in the password (default true)
  -h, --help            help for generate
  -l, --length int      length of the generated password (default 12)
```

### SEE ALSO

* [cred pass](cred_pass.md)	 - A command to manage passwords

