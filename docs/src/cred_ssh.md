## cred ssh

Manage SSH keys and connections

### Synopsis

The ssh command allows you to manage SSH keys and establish connections.
It provides functionalities to add keys, list keys, and connect to servers.

Examples:
- Add a new SSH key: ssh add <key-name>
- List all SSH keys: ssh ls
- Connect to a server: ssh connect <key-name>

```
cred ssh [flags]
```

### Options

```
  -h, --help   help for ssh
```

### SEE ALSO

* [cred](cred.md)	 - A password and environment variables manager
* [cred ssh connect](cred_ssh_connect.md)	 - Establish SSH connection using stored connection string and private key
* [cred ssh copy](cred_ssh_copy.md)	 - Copies the stored ssh to system clipboard
* [cred ssh cp](cred_ssh_cp.md)	 - copies files and directories
* [cred ssh download](cred_ssh_download.md)	 - Download SSH key files for the specified entry
* [cred ssh edit](cred_ssh_edit.md)	 - edit a new ssh entry
* [cred ssh generate](cred_ssh_generate.md)	 - Generate a new SSH key pair and save it securely
* [cred ssh insert](cred_ssh_insert.md)	 - Insert a new ssh entry
* [cred ssh ls](cred_ssh_ls.md)	 - List files and directories
* [cred ssh mkdir](cred_ssh_mkdir.md)	 - Create directories
* [cred ssh mv](cred_ssh_mv.md)	 - Move files and directories
* [cred ssh rm](cred_ssh_rm.md)	 - Remove files and directories
* [cred ssh show](cred_ssh_show.md)	 - Displays the stored ssh

