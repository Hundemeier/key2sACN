# key2sACN

**Currently in a test stadium and therefore no binaries are provided!**

This program is a Linux tool for using keystrokes as sACN triggers. When you press a key on a keyboard
a DMX channel will send 100% and you can use this as a simple remote for eg. grandMA.

This program does only work on Linux, because only under Linux it is simple enough to differentiate
between multiple keyboards. So this tool supports multiple keyboards as different keys.

key2sACN can be configurated via Web-Interface in a normal browser. For more information, see the wiki page.

This program is written in go and some compiled binaries are available under [releases](https://github.com/Hundemeier/key2sACN/releases).

It uses a modified version of the [keylogger](https://github.com/MarinX/keylogger) by MarinX and this sACN library: [sACN](https://github.com/Hundemeier/go-sacn).