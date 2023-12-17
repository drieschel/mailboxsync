# mailboxsync
Easy-to-use tool to synchronize multiple mailboxes at the same time. It uses [imapsync](https://github.com/imapsync/imapsync) for the synchronization process.

## Usage (with Docker)
1. Create `config/mailboxes.json` and fill it with the required data (see section below)
2. Execute `make run`
3. Grab a coffee :coffee: and make yourself comfortable

> [!TIP]
> 
> You can tail the logs in `var/log` from different mailboxes during the sync

> [!TIP]
> 
> Use the variable `CONCURRENT_SYNCS` in case you want to adjust the amount of concurrent mailbox syncs. Default is 3.
> 
> `make run CONCURRENT_SYNCS=5`

## Mailbox JSON example
Required server and mailbox data has to be provided in a JSON file
```
[
  {
    "src": {
      "host": "src.imap-server.tld"
    },
    "dst": {
      "host": "dst.imap-server.tld"
    },
    "mailboxes": [
      {
        "user": "john.doe@domain.tld",
        "password": "verysecretpassword",
        "srcUser": "john.doe@domain.tld",
        "srcPassword": "verysecretpassword",
        "dstUser": "john.doe@domain.tld",
        "dstPassword": "verysecretpassword"
      },
      {
        "user": "jane.doe@domain.tld",
        "password": "passwordsecretvery",
        "srcUser": "jane.doe@domain.tld",
        "srcPassword": "passwordsecretvery",
        "dstUser": "jane.doe@domain.tld",
        "dstPassword": "passwordsecretvery",
        "active": false
      }
    ]
  }
]
```
Every element in the list is a `Sync` object. A `Sync` object consists of all required data for synchronizing mailboxes between `src` and `dst`.

## Types
Reference about the types which have to be defined in the JSON file. 
### ImapServer
Consists of all required data to establish a connection to an IMAP server

| Name | Type    | Description                     |                  |
|------|---------|---------------------------------|------------------|
| host | string  | Domain or IP of the IMAP server | **Required**     |
| port | integer | Port of the IMAP server         | **Default: 143** |

### Mailbox
Consists of all required credentials to connect to the mailboxes which have to be synchronized

| Name        | Type    | Description                                                                                                 |                   |
|-------------|---------|-------------------------------------------------------------------------------------------------------------|-------------------|
| user        | string  | Username for source and/or destination mailbox                                                              | **Required**      |
| password    | string  | Password for source and/or destination mailbox                                                              | **Required**      |
| srcUser     | string  | Will be used as username for the source mailbox in case it is set. Otherwise `user` will be taken.          |                   |
| srcPassword | string  | Will be used as password for the source mailbox in case it is set. Otherwise `password` will be taken.      |                   |
| dstUser     | string  | Will be used as username for the destination mailbox in case it is set. Otherwise `user` will be taken.     |                   |
| dstPassword | string  | Will be used as password for the destination mailbox in case it is set. Otherwise `password` will be taken. |                   |
| active      | boolean | De-/Activate mailbox for synchronization                                                                    | **Default: true** |

### Sync
Contains all required data to synchronize one or more mailboxes between two servers at the same time

| Name      | Type       | Description                                                             |              |
|-----------|------------|-------------------------------------------------------------------------|--------------|
| src       | ImapServer | Connection data for the source IMAP server                              | **Required** |
| dst       | ImapServer | Connection data for the destination IMAP server                         | **Required** |
| mailboxes | Mailbox[]  | List of mailboxes which have to be synchronized between `src` and `dst` | **Required** |

## Misc
In case you miss an option from [imapsync](https://github.com/imapsync/imapsync) or a specific feature, don't hesitate to create an [issue](https://github.com/drieschel/mailboxsync/issues) or a PR.  