This repository is a sample of launching API on docker containers.

## How to build and launch

You can build and launch API by

```
$ docker-compose up
$ docker-compose up
```

or

```
$ docker-compose up postgres
$ docker-compose up apig-sample
```

You have to run ``` $ docker-compose up ``` twice or up each service separately because ``` docker-compose ``` will not wait until dependant services are "ready". 

You can call API by these URLs

```
http://<docker-ip>:8080/users 
http://<docker-ip>:8080/email
```