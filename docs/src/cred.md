## cred

A password and environment variables manager

### Synopsis

Cred is a powerful CLI tool built in Go for managing passwords and environment variables.
It uses GPG encryption to securely store and manage sensitive information.

Examples and usage:
- Initialize with a GPG key: cred init <gpg-key-id>
- Store a new credentials: cred {pass/env} insert <file-name>
- Retrieve a credentials: cred {pass/env} show <file-name>
- Retrieve a credentials: cred {pass/env} copy <file-name>
- List all stored credentials: cred {pass/env} list

### Options

```
      --generate-docs   Creates markdown documentation for the CLI
  -h, --help            help for cred
```

### SEE ALSO

* [cred completion](cred_completion.md)	 - Generate the autocompletion script for the specified shell
* [cred env](cred_env.md)	 - A command to manage env-variables
* [cred git](cred_git.md)	 - Manage cred-store git repository and operations
* [cred init](cred_init.md)	 - Initialize the credential store
* [cred migrate](cred_migrate.md)	 - Migrate the credential store to a new GPG key
* [cred pass](cred_pass.md)	 - A command to manage passwords
* [cred quick-setup](cred_quick-setup.md)	 - Creates a GPG ID and initiates store usage credentials
* [cred ssh](cred_ssh.md)	 - Manage SSH keys and connections

