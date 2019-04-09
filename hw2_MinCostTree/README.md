## Graph Theory Homework 2

* Using GO language to implement this homework
    * [How to install GO compiler](https://golang.org/doc/install)
* Folder 
    * bin  : The exe file for Linux
    * src  : The source code 
    * test : The testing data

## How to use 

* Using binary file to run
    * You have to run on environment of Linux.
    ```
    $ cd bin/
    $ ./main "input file path"
    ```
* Using GO compiler to run (Recommend for other environment)
    * You have to intall GO compiler by fallowing the above link, and using the GO to run.
    ```
    $ cd src/
    $ go run main.go "input file path"
    ```
    * You can create the exe file for your environment
    ```
    $ cd src/
    $ go build main.go
    ```

## My environment

```
$ cat /etc/os-release
NAME="Ubuntu"
VERSION="18.04.2 LTS (Bionic Beaver)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 18.04.2 LTS"
VERSION_ID="18.04"
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
VERSION_CODENAME=bionic
UBUNTU_CODENAME=bionic

```

## Using testing input file
* These are some testing file in `/test`, you can chose what you want.
* Using exe file on Linux
    ```
    $ cd bin/
    $ ./main ../test/input.txt
    ```
* Using Go run
    ```
    $ cd src/
    $ go run main.go ../test/input.txt
    ```
