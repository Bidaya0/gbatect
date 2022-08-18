# gbatect

gbatect is a tool help users move exists docker-compose to batect .
gbatect take the `docker-compose.yml` and translates it to `batect.yml`.


# Roadmap


## N+1 

- [ ] bidirectional conversion between docker-compose and batect
- [ ] basic support Kubernetes 
- [ ] convert `docker run` command to batect container
- [ ] update exists batect.yml

## 0.0.1

- [x] Convert docker-compose.yml to batect.yml
	- [x] Support all batect field types
	- [x] basic output batect
- [ ] basic cli
	- [ ] basic help document
- [ ] project quality flow
	- [ ] coverage test
	- [ ] code style check
	- [x] code scaner
