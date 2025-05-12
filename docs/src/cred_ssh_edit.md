## cred ssh edit

edit a new ssh entry

### Synopsis

The edit command allows you to add a new ssh entry to the ssh store.
You will be prompted to enter and confirm the ssh, which will be stored securely.
If the entry already exists, you will be asked whether you want to overwrite it.

Examples:
ssh edit <key-name> --public-key <key-path> --private-key <key-path> --connection <connection-string>

```
cred ssh edit [flags]
```

### Options

```
      --connection-string string   connection string
  -h, --help                       help for edit
      --private-key string         private key file path
      --public-key string          public key file path
```

### SEE ALSO

* [cred ssh](cred_ssh.md)	 - Manage SSH keys and connections

