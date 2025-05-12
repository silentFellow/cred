## cred ssh connect

Establish SSH connection using stored connection string and private key

### Synopsis

Looks for the 'connection.gpg' and 'private.gpg' files stored under the given <keyname>
in the cred store. Decrypts these files and uses the connection string and private key
to establish an SSH connection to the remote server.

Example:
  cred ssh connect my-server
This will look for:
  my-server/connection.gpg
  my-server/private.gpg
and use the decrypted contents to connect.

```
cred ssh connect [flags]
```

### Options

```
  -h, --help   help for connect
```

### SEE ALSO

* [cred ssh](cred_ssh.md)	 - Manage SSH keys and connections

