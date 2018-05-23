# Prefab Mattermost Installing
## you need to have jumpscale installed follow this docs https://github.com/Jumpscale/bash

## please choose a machine that have at least 2 GB of RAM
to install prefab on a remote machine 
then type `js9`

```
prefab = j.tools.prefab.getFromSSH("<machine_ip>", '<ssh_port>')
prefab.apps.mattermost.install('<mysql_password>') # this will creates a mysql database and add a user to it with this password 
```

# User docs 
https://docs.mattermost.com/guides/user.html


