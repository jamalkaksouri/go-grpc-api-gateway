## Installation

```bash
$ make proto
```

## Running the app

```bash
$ make server
```

## How to install make command in windows

```make``` is a GNU command so the only way you can get it on Windows is installing a Windows version like the one provided by [GNUWin32](http://gnuwin32.sourceforge.net/install.html). Anyway, there are several options for getting that:

The most simple choice is using [Chocolatey](https://chocolatey.org/install). First you need to install this package manager. Once installed you simlpy need to install ```make``` (you may need to run it in an elevated/admin command prompt) :

```bash
choco install make
```
Other recommended option is [installing a Windows Subsystem for Linux (WSL/WSL2)](https://docs.microsoft.com/en-us/windows/wsl/install-win10), so you'll have a Linux distribution of your choice embedded in Windows 10 where you'll be able to install make, gccand all the tools you need to build C programs.

For older Windows versions (MS Windows 2000 / XP / 2003 / Vista / 2008 / 7 with msvcrt.dll) you can use [GNUWin32](http://gnuwin32.sourceforge.net/install.html).

An outdated alternative was [MinGw](https://www.ics.uci.edu/~pattis/common/handouts/mingweclipse/mingw.html), but the project seems to be abandoned so it's better to go for one of the previous choices.
