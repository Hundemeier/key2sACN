# key2sACN

If you are looking for a simpler solution that also works on Windows, see [key2sACN-simple](https://github.com/Hundemeier/key2sACN-simple).

This program is a Linux tool for using keystrokes as sACN triggers. When you press a key on a keyboard
a DMX channel will send 100% and you can use this as a simple remote for eg. grandMA.

This program does only work on Linux, because only under Linux it is simple enough to differentiate
between multiple keyboards. So this tool supports multiple keyboards as different keys.

key2sACN can be configurated via Web-Interface in a normal browser. For more information, see the 
[wiki](https://github.com/Hundemeier/key2sACN/wiki) page.

This program is written in go and some compiled binaries are available under [releases](https://github.com/Hundemeier/key2sACN/releases). Simply download the executable of your choice and execute it.

It uses a modified version of the [keylogger](https://github.com/MarinX/keylogger) by MarinX and this sACN library: [sACN](https://github.com/Hundemeier/go-sacn).