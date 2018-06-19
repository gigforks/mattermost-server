# Prefab Mattermost Installing
- you need to have jumpscale installed follow this docs https://github.com/Jumpscale/bash to start installation from it on a remote machine

- please choose a machine that have at least 2 GB of RAM
- to install mattermost on a remote machine type `js9`

```
prefab = j.tools.prefab.getFromSSH("<machine_ip>", '<ssh_port>')
prefab.apps.mattermost.install('<mysql_password>') # this will creates a mysql database and add a user to it with this password 
```

# Full docs for user and Administration is here 
https://docs.mattermost.com/overview/index.html


# To integrate with itsyou.online
- from `system console` > `Authentication` > `Itsyouonline`
- set enable to True
- fill in the applicationId, secret, and itsyouonline url
- click save
