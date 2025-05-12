## cred ssh download

Download SSH key files for the specified entry

### Synopsis

The download command retrieves SSH key files (public, private, and connection info)
for a given SSH entry and saves them into a 'downloads' directory.

Example:
  cred ssh download my-key-name
This will create:
  downloads/my-key-name/public.key
  downloads/my-key-name/private.key
  downloads/my-key-name/connection.txt

```
cred ssh download [flags]
```

### Options

```
  -h, --help   help for download
```

### SEE ALSO

* [cred ssh](cred_ssh.md)	 - Manage SSH keys and connections

