## Kubeswitches

Kubeswitches is a tool to help you switch between different Kubernetes configurations.

### Requirements

- `KUBECONFIGS` environment variable set to the path of the kubeconfig file you want to switch between
- `go` language installed for build the binary

#### Setup kube contexts folder

```bash
mkdir -p ~/.kube/contexts # example contexts folder who contains all your kube contexts 
echo "export KUBECONFIGS=~/.kube/contexts" >> ~/.zshrc
source ~/.zshrc
```
### Installation

```bash
go build -o kubeswitches main.go && sudo mv kubeswitches /usr/local/bin/
```

#### Verify installation

```bash
kubeswitches -h
```

### Usage

## for help
```bash
kubeswitches -h
```

## set a config
```bash
kubeswitches set <configfilename>
```

## list configs
```bash
kubeswitches list
```

## get current config
```bash
kubeswitches current
```
### Add Alias

```bash
echo "alias ks='kubeswitches'" >> ~/.zshrc
source ~/.zshrc
```

