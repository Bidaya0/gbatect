# gbatect

gbatect is a tool help users adopt [batect](https://github.com/batect/batect) , a Build And Testing Environments as Code Tool , use exists [docker-compose](https://docs.docker.com/compose/) configurations.

gbatect take the `docker-compose.yml` and translates it to `batect.yml`.

# Installation

get the latest version
```
curl -s https://api.github.com/repos/bidaya0/gbatect/releases/latest | grep 'tag_name' | cut -d\" -f4
```

download and install 
```
sh -c "curl -L https://github.com/bidaya0/gbatect/releases/download/${GBATECT_VERSION}/gbatect-${GBATECT_VERSION}-${System}-amd64.tar.gz  | tar xvz -C > /usr/local/bin/"
chmod +x /usr/local/bin/gbatect
```

validation install success
```
> gbatect
gbatect is a tool help users move exists docker-compose to batect.
	gbatect take the docker-compose.yml and translates it to batect.yml.

Usage:
  gbatect [command]

Available Commands:
  convert     convert docker-compose file to batect format
  help        Help about any command

Flags:
  -h, --help   help for gbatect

Use "gbatect [command] --help" for more information about a command.
```


# Roadmap


## N+1 

- [ ] bidirectional conversion between docker-compose and batect
- [ ] basic support Kubernetes 
- [ ] convert `docker run` command to batect container

## v0.0.2
- [ ] update exists batect.yml
	- [ ] show diff between exists batect.yml and updated one
	- [ ] update and keep origin format
	- [ ] choose update parts of batect.yml
- [ ] serve as a web service
	- [ ] offer online convert 

## 0.0.1

- [x] Convert docker-compose.yml to batect.yml
	- [x] Support all batect field types
	- [x] basic output batect
- [x] basic cli
	- [x] basic help document
- [x] project quality flow
	- [x] coverage test
	- [x] code style check
	- [x] code scaner
