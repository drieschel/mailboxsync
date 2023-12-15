Mailboxsync is an easy-to-use tool to synchronize multiple mailboxes between different servers at the same time. It uses [imapsync](https://github.com/imapsync/imapsync) for the synchronization process.

Required server and mailbox data has to be provided in a json file
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

## Usage (with Docker)
1. Create `config/config.json` and fill it with the required data
2. Run `make run`
3. Make yourself comfortable and grab a coffee :coffee:

> [!TIP]
> You can tail the logs in `var/log` from different mailboxes during the sync

## Types

### ImapServer
Consists of all required data to establish a connection to an IMAP server

#### Properties
**`host [string]`**
Domain or IP of the IMAP server (**required**)

**`port [integer]`**
Port of the IMAP server (**default: 143**)

### Mailbox
Consists of all required credentials to connect to the mailboxes which need to be synchronized

#### Properties
**`user [string]`**
Username for source and/or destination mailbox (**required**)

**`password [string]`**
Password for source and/or destination mailbox (**required**)

**`srcUser [string]`**
It will be used as username for the source mailbox in case it is set. Otherwise `user` will be taken.

**`srcPassword [string]`**
It will be used as password for the source mailbox in case it is set. Otherwise `password` will be taken.

**`dstUser [string]`**
It will be used as username for the destination mailbox in case it is set. Otherwise `user` will be taken.

**`dstPassword [string]`**
It will be used as password for the destination mailbox in case it is set. Otherwise `password` will be taken.

**`active [boolean]`**
In case active is false, mailbox will not get synchronized (**default: true**)

### Sync
Contains all required data to synchronize one or more mailboxes at the same time

#### Properties
**`src [ImapServer]`**
Connection data for the source IMAP server (**required**)

**`dst [ImapServer]`**
Connection data for the destination IMAP server (**required**)

**`mailboxes [Mailbox[]]`**
List of mailboxes which have to be synchronized between `src` and `dst` (**required**)