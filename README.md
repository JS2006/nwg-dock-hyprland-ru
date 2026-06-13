# nwg-dock-hyprland-ru

Это русифицированная версия [nwg-dock-hyprland](https://github.com/nwg-piotr/nwg-dock-hyprland) — настраиваемого (через аргументы командной строки и CSS) дока, написанного на Go, предназначенного исключительно для композитора [Hyprland](https://github.com/hyprwm/Hyprland). Док содержит закреплённые кнопки, кнопки клиентов и кнопку лаунчера. Последняя по умолчанию запускает [nwg-drawer](https://github.com/nwg-piotr/nwg-drawer).

This is a Russian-localized version of [nwg-dock-hyprland](https://github.com/nwg-piotr/nwg-dock-hyprland) — a configurable dock (via command line arguments and CSS), written in Go, aimed exclusively at the [Hyprland](https://github.com/hyprwm/Hyprland) Wayland compositor. It features pinned buttons, client buttons and the launcher button. The latter by default starts [nwg-drawer](https://github.com/nwg-piotr/nwg-drawer).

## Отличия от оригинального nwg-dock-hyprland / Differences from original:

- **Русификация**: автоматическое определение языка системы (LANG/LC_ALL). При русской локали все сообщения, пункты контекстного меню и описания флагов отображаются на русском языке. При английской — на английском.
- **Эмодзи-иконка лаунчера**: возможность использовать эмодзи вместо SVG-иконки для кнопки лаунчера. Укажите `-ico emoji:СИМВОЛ`, например `-ico "emoji:🚀"`. Стиль эмодзи-кнопки настраивается через CSS-селектор `#launcher-emoji`.

[![Packaging status](https://repology.org/badge/vertical-allrepos/nwg-dock-hyprland.svg)](https://repology.org/project/nwg-dock-hyprland/versions)

## Installation

### Requirements

- `go`>=1.20 (just to build)
- `gtk3`
- `gtk-layer-shell`
- [nwg-drawer](https://github.com/nwg-piotr/nwg-drawer): optionally. You may use another launcher (see help),
or none at all. The launcher button won't show up, if so.

### Steps

1. Clone the repository, cd into it.
2. Install golang libraries with `make get`. First time it may take ages, be patient.
3. `make build`
4. `sudo make install`

## Running

Either start the dock permanently in `hyprland.conf`:

```text
exec-once = nwg-dock-hyprland [arguments]
```

or assign the command to some key binding. Running the command again kills the existing program instance, so that
you could use the same key to open and close the dock.

## Running the dock residently

If you run the program with the `-d` or `-r` argument (preferably in autostart), it will be running residently.

```text
exec_always nwg-dock-hyprland -d
```

or

```text
exec_always nwg-dock-hyprland -r
```

### `-d` for autohiDe

Move the mouse pointer to expected dock location for the dock to show up. It will be hidden a second after you leave the
window. Invisible hot spots will be created on all your outputs, unless you specify one with the `-o` argument.

### `-r` for just Resident

No hotspot will be created. To show/hide the dock, bind the `exec nwg-dock-hyprland` command to some key or button.
How about the `Menu` key, which is usually useless?

Re-execution of the same command hides the dock. If a resident instance found, the `nwg-dock-hyprland` command w/o
arguments sends `SIGUSR1` to it. Actually `pkill -USR1 nwg-dock-hyprland` could be used instead. This also works in autohiDe
mode.

Re-execution of the command with the `-d` or `-r` argument won't kill the running instance. If the dock is
running residently, another instance will just exit with 0 code. In case you'd like to terminate it anyway, you need 
to `pkill -f nwg-dock-hyprland`.

*NOTE: you need to kill the running instance before reloading Hyprland, if you've just changed the arguments you
auto-start the dock with.*

```txt
$ nwg-dock-hyprland -h
Usage of nwg-dock-hyprland:
  -a string
        Alignment in full width/height: "start", "center" or "end" (default "center")
  -c string
        Command assigned to the launcher button (default "nwg-drawer")
  -d    auto-hiDe: show dock when hotspot hovered, close when left or a button clicked
  -debug
        turn on debug messages
  -f    take Full screen width/height
  -g string
        quote-delimited, space-separated class list to iGnore in the dock
  -hd int
        Hotspot Delay [ms]; the smaller, the faster mouse pointer needs to enter hotspot for the dock to appear; set 0 to disable (default 20)
  -hl string
        Hotspot Layer "overlay" or "top" (default "overlay")
  -i int
        Icon size (default 48)
  -ico string
        alternative name or path for the launcher ICOn
  -iw string
        Ignore the running applications on these Workspaces based on the workspace's name or id, e.g. "special,10"
  -l string
        Layer "overlay", "top" or "bottom" (default "overlay")
  -lp string
        Launcher button position, 'start' or 'end' (default "end")
  -m    allow Multiple instances of the dock (skip lock file check)
  -mb int
        Margin Bottom
  -ml int
        Margin Left
  -mr int
        Margin Right
  -mt int
        Margin Top
  -nolauncher
        don't show the launcher button
  -o string
        name of Output to display the dock on
  -p string
        Position: "bottom", "top" "left" or "right" (default "bottom")
  -r    Leave the program resident, but w/o hotspot
  -s string
        Styling: css file name (default "style.css")
  -v    display Version information
  -w int
        number of Workspaces you use (default 10)
  -x    set eXclusive zone: move other windows aside; overrides the "-l" argument

Usage of signals:
 SIGRTMIN+1 (signal 35): toggle dock visibility (USR1 has been deprecated)
 SIGRTMIN+2 (signal 36): show the dock
 SIGRTMIN+3 (signal 37): hide the dock
```

![screenshot-2.png](https://raw.githubusercontent.com/nwg-piotr/nwg-shell-resources/master/images/nwg-dock/dock-2.png)

## Styling

Edit `~/.config/nwg-dock-hyprland/style.css` to your taste.

### Эмодзи-иконка лаунчера / Emoji Launcher Icon

To use an emoji instead of the default grid icon for the launcher button, use the `-ico` flag with the `emoji:` prefix:

```bash
nwg-dock-hyprland -ico "emoji:🚀"
nwg-dock-hyprland -ico "emoji:⚡"
nwg-dock-hyprland -ico "emoji:🏠"
nwg-dock-hyprland -ico "emoji:📱"
```

The emoji button can be styled in CSS using the `#launcher-emoji` selector:

```css
#launcher-emoji {
    font-size: 28px;
    padding: 4px;
    min-width: 40px;
    min-height: 40px;
    color: #ffffff;
}

#launcher-emoji:hover {
    background-color: rgba(255, 255, 255, 0.2);
    border-radius: 6px;
}
```

### Русификация / Russian Localization

The dock automatically detects the system language from the `LANG`, `LC_ALL`, or `LC_MESSAGES` environment variable. When Russian locale (`ru_*`) is detected, all interface elements are displayed in Russian:
- Command-line flag descriptions
- Context menu items (Закрепить/Открепить, Новое окно, Закрыть все окна, etc.)
- Log messages and error notifications
- Signal usage descriptions

To force English output:
```bash
LANG=en_US.UTF-8 nwg-dock-hyprland
```

## Troubleshooting

### An application icon is not displayed

The only thing the dock knows about the app is it's class name.

```text
$ hyprctl clients
(...)
Window 55a62254b8c0 -> piotr@msi:~:
        mapped: 1
        hidden: 0
        at: 1204,270
        size: 2552,1402
        workspace: 6 (6)
        floating: 0
        monitor: 2
        class: foot
        title: piotr@msi:~
        initialClass: foot
        initialTitle: foot
        pid: 58348
        xwayland: 0
        pinned: 0
        fullscreen: 0
        fullscreenmode: 0
        fakefullscreen: 0
        grouped: 0
        swallowing: 0
```

Now it'll look for an icon named 'foot'. If that fails, it'll look for a .desktop file named foot.desktop, which should contain the icon name or path. If this fails as well, no icon will be displayed. I've added workarounds for some most common exceptions, but it's impossible to predict every single application misbehaviour. This is either programmers fault (improper class name), or bad packaging (.desktop file name different from the application class name).

If some app has no icon in the dock:

1. check the app class name (`hyprctl clients`);
2. find the app's .desktop file;
3. copy it to ~/.local/share/applications/` and rename to <class_name>.desktop.

If the .desktop file contains proper icon definition (`Icon=`), it should work now.

## Credits

This program uses some great libraries:

- [gotk4](https://github.com/diamondburned/gotk4) by [diamondburned](https://github.com/diamondburned) released under [GNU Affero General Public License v3.0](https://github.com/diamondburned/gotk4/blob/4/LICENSE.md)
- [go-singleinstance](github.com/allan-simon/go-singleinstance) Copyright (c) 2015 Allan Simon
