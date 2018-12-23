# simulation-warehouse-simple
This website for simulation warehouse (still api)

Im using Golang (go1.9) Framework is Beego v1.7.1 and im using Database is SQLite3

# Setup 

1. Install Golang, you can visit [This Link](https://golang.org/doc/install)
2. Install Beego, you can visit [This Link](https://beego.me/docs/install/)
3. Open your terminal then execute command `git clone https://github.com/FadhilAhsan/simulation-warehouse-simple.git` in your `$GOPATH/src`
4. Go to folder simulation-warehouse-simple
5. Open your terminal then execute command`./"executeable-project.sh"` for run this project


# Note 
if failed follow step `5.`, please follow this step,

1. install gcc  Linux\Ubuntu 
	* `sudo apt install gcc`
	* `sudo apt install build-essential`
	* `gcc --version` for check gcc version
2. Open your terminal then execute command `go get github.com/astaxie/beego/orm` for install lib orm for beego
3. Open your terminal then execute command `go get github.com/mattn/go-sqlite3` for install driver sqlite3
