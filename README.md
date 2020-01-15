# EmbargoUpdater
Embargo updater tool for in-network people (sadly not yet available for external use)

## Usage
The source code is provided for those interested, but most of you should be fine with just using an executable (found on /bin)

If you just want the embargo exec, just download a zip or clone, and grab it straight out of /bin (exe for windows, linux/mac can use the empty extension with ./)

If you want to run the code, you need the go tools installed: https://golang.org/doc/install, then:
1. Clone the repo
2. Cd into the /src/embargo_updater/
3. Either run the code directly with `go run *.go` or build with `go build *.go` (if running, provide creds as stated below)
4. If built, retrieve your exec and run with (unix) `./embargo_updater <username> <password>` where, username and password are, of course, your valid xpo user creds

## Notes
This is just a tool to help out, there might/will be bugs, let me (jose.martinez19@xpo.com) know if you need any help.
