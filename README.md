## Kubeswitches

Kubeswitches is a tool to help you switch between different Kubernetes configurations.

- [Requirements](#requirements)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)

### Requirements

- `KUBECONFIGS` environment variable set to the path of the kubeconfig file you want to switch between
- `go` language installed for build the binary

####  Prerequisites

- Setup kube contexts folder:

```bash
mkdir -p ~/.kube/contexts # example contexts folder who contains all your kube contexts 
echo "export KUBECONFIGS=/Users/eatmore/.kube/contexts" >> ~/.zshrc
source ~/.zshrc
```
### Installation

#### Build from code
```bash
go build -o kubeswitches main.go && sudo mv kubeswitches /usr/local/bin/
```

#### Install a binary from releases

##### for Macos
```bash
curl -L https://github.com/eatmore01/kubeswitches/releases/download/v1.0.0/kubeswitches_darwin_amd64 -o kubeswitches
chmod +x kubeswitches
sudo mv kubeswitches /usr/local/bin/
```

##### for Linux
```bash
curl -L https://github.com/eatmore01/kubeswitches/releases/download/v1.0.0/kubeswitches_linux_amd64 -o kubeswitches
chmod +x kubeswitches
sudo mv kubeswitches /usr/local/bin/
```
##### for Windows download manualy the binary from [releases](https://github.com/eatmore01/kubeswitches/releases)


#### Verify installation

```bash
kubeswitches -h
```

## Usage

### for help
```bash
kubeswitches -h
```

### set a config
```bash
kubeswitches set <configfilename>
```

### list configs
```bash
kubeswitches list
```

### get current config
```bash
kubeswitches current
```
### Add Alias

```bash
echo "alias ks='kubeswitches'" >> ~/.zshrc
source ~/.zshrc
```

