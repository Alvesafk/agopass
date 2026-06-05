# Agopass, CLI password manager made in Golang.

Agopass is a CLI password/secret manager made in Go, it's very simple, few commands is all you need.

Want to install?

**Requirements**:
- Go
- A clipboard manager:
    - xsel (X11)
    - xclip (X11)
    - wl-clipboard (Wayland)
> If you're in MacOS agopass will use your system clipboard, so don't worry about it.

Compatibility:
- Unix-like systems (Linux, MacOS)
> I don't think it's compatible with windows and don't really care enough to test it.

Compile from source:
```
git clone https://github.com/Alvesafk/agopass --depth=1
cd agopass
go build -o bin/agopass main.go
```
Then just put on your bin directory for you to be able to access whenever.

Go install:

```
go install github.com/Alvesafk/agopass@latest
```

It will install for you, will be on your `$GOPATH/bin`, you should export your Go bin dir with your normal bin folder, than you can use whenever and wherever.

## How to use it?

It's very simple, like i said few commands are necessary.

First you have to initialize it, i will assume your agopass binary is on your $PATH.

Use `agopass init` to create a **Master Key**, this key is the only password that you have to remember now (make sure to remember!), it's used to authenticate you on some commands that modify your DB or show your *secrets*, it's also used as *salt* when encrypting your secrets, you can't use `agopass init` more than one time, so make sure you don't forget it.
> If you forgot your password the only thing that you can really do it's removing the DB and starting over, the Master Key can't be retrieved, it's on purpose and it's inegotiable.

After you created your **Master Key** you can use the other commands, in this moment the commands are:
- `add`
    - It's how you insert your secrets, just do `agopass add` and you will be prompted to add a name and your secret key, names are case sensitive for all the commands that use them, so make sure you're writing correctly.
- `list`
    - You will not be prompted to use your **Master Key** with this command, it shows the registered Secrets on the DB, it does not show the real secret key nor the hash.
- `get`
    - This is how you retrieve your secret keys from the DB, just do `agopass get <Secret>` and it will be coppied to your clipboard.
- `delete`
    - Delete a secret registered on your DB based on it's name, let's say i want to delete my Github password from agopass, i will do `agopass delete Github`, will prompt me asking if i really want to do it. Remember that names are case sensitive.
- `update`
    - Update a secret registered on the DB, the name and the key can be changed, `agopass update github` will prompt me to update my Github key, it has a way to check typing so in most cases the program will find the key that you're looking for. 

This is the project for now, some features that i want to work next:
- [x] Update command to update or change the registered key.
- [ ] A way to export and import the DB, without just copying the DB binary.

Thanks for your attention, Alvesafk.
