# kclip
Command line interface to windows clipboard over [KiTTY](http://www.9bis.net/kitty/#!index.md) remote-control printing

## About
This tool behaves like the `cat` command, it just tries to pad your content with the appropriate characters for your ssh client to send to a printer. KiTTY allows you to set your printer to the windows clip board so this in effect outputs whatever is fed to it to your windows clipboard.

## Setup
To configure KiTTY for clipboard printing first pull up configuration or reconfiguration. Then under the `Remote-controlled printing` section select `Windows clipboard` as your printer [see screenshot](doc/termcfg.png). Then install this tool on the remote system you want to copy to your clipboard from via `go install github.com/jacobalberty/kclip@latest`. Now use it as you would `cat` and instead of the files being output to your terminal they should output to your windows clipboard.
