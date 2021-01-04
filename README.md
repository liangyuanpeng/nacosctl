# Nacosctl
A CLI for [Nacos](https://github.com/alibaba/nacos) written in Go

# Install 
```
git clone https://github.com/liangyuanpeng/nacosctl.git \
cd nacosctl 
```  
```
go run main.go   
```   
and you will get the `Hi here is Nacos cli.`

### Build from source code

```bash
make build
chmod +x nacosctl && mv nacosctl /usr/local/bin/nacosctl
```
# TODO List
- [ ] Nacos v1 Config (http for publish/get/del/listen)
- [ ] Nacos v2 Config (grpc for publish/get/del/listen)  
- [ ] Manager Multiple Cluster for Config

# Contribute
Contributions are welcomed and greatly appreciated. For more information about how to submit a patch and the contribution workflow, see [CONTRIBUTING.md](CONTRIBUTING.md).

# License
Licensed under the Apache License Version 2.0: http://www.apache.org/licenses/LICENSE-2.0
