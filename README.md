# Doprax CLI
Doprax CLI is a unofficial CLI tool to interact with doprax.com dashboard. you can view and manage your projects within CLI.

## Commands
To view help run `doprax help`

### login
with login command you can login into doprax account it currently supports email/password style of login and OATH is not supported
Example
```
doprax login --api-key [api key] --username [username] --password [password]
```
Or you can supply values interactively with just calling `doprax login`

### Project

- Pull project source code
```
doprax project pull [project title]
```

- Deploy project main service
```
doprax project deploy restart [project title]
```

- List projects
```
doprax project list
```
