# SSH login password-less

For example, using vagrant keys
```
fanhonglingdeMacBook-Pro:docker-compose fanhongling$ curl -jkSL https://raw.githubusercontent.com/mitchellh/vagrant/master/keys/vagrant -o vagrant
fanhonglingdeMacBook-Pro:docker-compose fanhongling$ chmod 600 vagrant 
fanhonglingdeMacBook-Pro:docker-compose fanhongling$ curl -jkSL https://raw.githubusercontent.com/mitchellh/vagrant/master/keys/vagrant.pub -o vagrant.pub
fanhonglingdeMacBook-Pro:docker-compose fanhongling$ chmod 600 vagrant.pub
```

Create or Add into Remote Linux Server
```
[vagrant@localhost ~]$ cat .ssh/authorized_keys 
ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEA6NF8iallvQVp22WDkTkyrtvp9eWW6A8YVr+kz4TjGYe7gHzIw+niNltGEFHzD8+v1I2YJ6oXevct1YeS0o9HZyN1Q9qgCgzUFtdOKLv6IedplqoPkcmF0aYet2PkEDo3MlTBckFXPITAMzF8dJSIFo9D8HfdOV0IAdx4O7PtixWKn5y2hMNG0zQPyUecp4pzC6kivAIhyfHilFR61RGL+GPXQ2MWZWFYbAGjyiYJnAmCP3NOTd0jMZEnDkbUvxhMmBYSdETk1rRgm+R4LOzFUGaHqHDLKLX+FIPKcF96hrucXzcWyLbIbEgE98OHlnVYCzRdK8jlqm8tehUc9c9WhQ== vagrant insecure public key
```

Make sure to secure file
```
[vagrant@localhost ~]$ chmod 600 .ssh/authorized_keys
```

Login
```
fanhonglingdeMacBook-Pro:sample-microservices-springboot fanhongling$ ssh -i ~/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/docker-compose/vagrant vagrant@172.17.4.50
Last login: Thu Jul 20 04:55:36 2017 from 172.17.4.1
```

Want own keys, refer to `ssh-keygen` - https://www.openssh.com/

For example
```
[vagrant@localhost ~]$ ssh-keygen
```

The keys will generated into __./ssh__ directory, default named id_rsa & id_rsa.pub