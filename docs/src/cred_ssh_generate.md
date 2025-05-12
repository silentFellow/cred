## cred ssh generate

Generate a new SSH key pair and save it securely

### Synopsis

The generate command allows you to create a new SSH key pair and store it
in the credential store. You can optionally add a connection string during generation.

Example:
  cred ssh generate <key-name> --connection <connection-string>

```
cred ssh generate [flags]
```

### Options

```
      --connection-string string   connection string
  -h, --help                       help for generate
```

### SEE ALSO

* [cred ssh](cred_ssh.md)	 - Manage SSH keys and connections

