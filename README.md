
# CPanel / WHM Exim monitoring

This is the missing piece of software that cPanel have yet to add.
Created to detect malware / virus infected / password hacked email accounts that attempt to use a legitimate email and send out spams.
It once costed me USD 90+ on mailgun with 197k of spam emails being billed upon me. I couldn't wait any longer. Please join me to maintain this :)

# development setup

* copy exim_mainlog to local machine

```
scp root@yourserver.com:/home/user/exim_mainlog ./
```

* clone repository
```
cd $GOPATH/src
git clone git@github.com:u007/eximmon.git

```

* install required dependencies
```
go get
```

* first time run

```
go run main.go start
```

* to build for production

```
GOOS=linux go build -o eximmon main.go

```

# production setup

* login to whm
* development > Manage Api Token
* create an api token, grant "Everything", and copy token (for API_TOKEN)
* eximmon start # first time run, ctrl+c after seing "ended: ..."
* add this to /etc/rc.local (on cpanel as root)
```
cd /root/eximmon && 
API_TOKEN=xxxxx NOTIFY_EMAIL=your-email EXIM_LOG=/var/log/exim_mainlog ./eximmon start > out.log &
```
* to see logs
```
tail -f /root/eximmon/out.log
```

## suspend an email

```
API_TOKEN=... eximmon suspend emailhere
```

## unsuspend an email

```
API_TOKEN=... eximmon unsuspend emailhere
```

## reset all past data

```
API_TOKEN=... eximmon reset
```


# TODO

* auto delete old data directory by date on every 100 runs
* query email to show list of dates with hour and mins count


# Help wanted!

* feel free to report bugs here
* if you need additional features, feel free to add into issue list
* if you need support, contact me @ james@mercstudio.com
