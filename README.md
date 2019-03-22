
# CPanel / WHM Exim monitoring

This is the missing piece of software that cPanel have yet to add.
Created to detect malware / virus infected / hacked email accounts that attempt to use a legitimate email and send out spams.
It costed me multiple time, but biggest was USD 90+ on mailgun with 197k of spam emails being billed upon me.
I couldn't wait any longer. Please join me to maintain this :)

# How it works?

It scans /var/log/exim_mainlog on interval basis, and logs all email activity being received from authenticated user.
Take note that it only support dovecot_plain at the moment. If you need other type of login, i may add it soon.

# How to operate?

Before starting, please go through this steps to limit the number of emails allowed per hour per domain.

[https://documentation.cpanel.net/display/CKB/How+to+Prevent+Spam+with+Mail+Limiting+Features](https://documentation.cpanel.net/display/CKB/How+to+Prevent+Spam+with+Mail+Limiting+Features)

* login to whm
* development > Manage Api Token
* create an api token, grant "Everything", and copy token (for API_TOKEN)
* add this to /etc/rc.local (on cpanel as root): replace xxxxx, your-email with your configurations

```
cd /root/eximmon && 
API_TOKEN=xxxxx NOTIFY_EMAIL=your-email EXIM_LOG=/var/log/exim_mainlog ./eximmon start > out.log &
```

* to see logs
```
tail -f /root/eximmon/out.log
```

* for helps

```
./eximmon help
```

## Environment variables
* MAX_PER_MIN=4
* MAX_PER_HOUR=100
* NOTIFY_EMAIL=email
* EXIM_LOG=/var/log/exim_mainlog
* WHM_API_HOST=node.servername.com

## Available command

* start - continue from last position or start from yesterday, and repeats from last position
* run - continue from last position or start from beginning for one time
* rerun - continue from a specific date
* skip - skip all existing data and repeats for new logs
* reset - reset all data, huh, what?
* suspend - suspend outgoing email
* unsuspend - unsuspend outgoing email
* info - get information of a domain
* test-notify - test send notification mail


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
GOOS=linux go build -o bin/eximmon main.go

```

# Changes

* 2019 Mar 23 - Fixed dovecot_login, added rerun
* 2018 Sep 25 - First public respository up


# TODO

* count numbers of external relayed recipients per email
* support for other type of dovecot_login methods
* auto delete old data directory by date on every 100 runs
* query email to show list of dates with hour and mins count

# Cleanup

Add this in cronjob, will delete any directory older than 30days.
Remember to change /path-to to your eximmon parent directory of "data"

```
0 1 * * * find /path-to/data/*/ -type f -name '*' -mtime +30 -exec rm -Rf {} \;
```

# Help wanted!

* feel free to report bugs here
* if you need additional features, feel free to add into issue list
* if you need support, contact me @ james@mercstudio.com
