# Example Configuration File for Your Go CLI

## Location: `cred-store/config`

## Syntax: `key = value`

### Note:

1. Each configuration entry must be on a new line.
2. Each line is split into key and value by "=".
3. Whitespace around keys and values is allowed and will be trimmed automatically.
4. Duplicate keys are allowed; the last occurrence will overwrite previous values.
5. Lines starting with `#` are treated as comments and ignored by the parser.
6. Invalid config will be ignored by the parser.

### Example Configuration:

```ini
# Automatically push all changes to Git (true/false)
auto_git = false

# Suppress stderr output (true/false)
suppress_stderr = false

# Editor to use for inserting and editing files
# Provide the full path or name of the editor (e.g., vim, nano, code).
editor = vim
```
