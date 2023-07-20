# Setting up Golang locally

[Download Go](https://go.dev/dl/)

<b> Go to the featured section from the above link and download the latest version of the go package based on the OS </b>

## Installation instructions
### Linux

<b>1. Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:</b>
```
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.6.linux-amd64.tar.gz
```

<b>Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.</b>

<b>2. Add /usr/local/go/bin to the PATH environment variable.</b>
```
nano ~/.bashrc
```
<b>Then in the nano editor enter the below command to add go to the path</b>
```
export PATH=$PATH:/usr/local/go/bin
```
```
source ~/.bashrc
```
<b>3. Verify the installation by entering the command in the command line terminal</b>
```
go version
```


