## cred migrate

Migrate the credential store to a new GPG key

### Synopsis

The migrate command allows you to re-encrypt your credential store with a new GPG key.
This operation will create a backup of your current store and re-encrypt all files with the new key.

	Example Usage: cred migrate <new-gpg-key-id>

```
cred migrate [flags]
```

### Options

```
  -h, --help   help for migrate
```

### SEE ALSO

* [cred](cred.md)	 - A password and environment variables manager

