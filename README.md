# kclip
Command line interface to windows clipboard over [KiTTY](http://www.9bis.net/kitty/#!index.md) remote-control printing

## About
This tool behaves like the `cat` command, it just tries to pad your content with the appropriate characters for your ssh client to send to a printer. KiTTY allows you to set your printer to the windows clip board so this in effect outputs whatever is fed to it to your windows clipboard.

## Setup
To configure KiTTY for clipboard printing first pull up configuration or reconfiguration. Then under the `Remote-controlled printing` section select `Windows clipboard` as your printer [see screenshot](doc/termcfg.png). Then install this tool on the remote system you want to copy to your clipboard from via `go install github.com/jacobalberty/kclip@latest`. Now use it as you would `cat` and instead of the files being output to your terminal they should output to your windows clipboard.

## OSC 52
If you are using terminal with OSC 52 support this will attempt to automatically detect that and use it instead. Clients like [nassh](https://chromium.googlesource.com/apps/libapps/+/master/nassh) support this.

## TODO (in no particular order)

* Add a list of terminals that OSC 52 supports
* Improve the robustness of OSC 52 detection
* Add documentation for a workaround for PuTTY without the printer clipboard patch
* Add chunking to OSC 52 and detect when we need to use it
* Add a CLI option to force OSC 52 or printer mode
