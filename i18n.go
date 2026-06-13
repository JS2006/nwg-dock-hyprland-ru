package main

import (
        "os"
        "strings"
)

// Lang определяет текущий язык интерфейса
var Lang string

// translations содержит переводы для всех поддерживаемых языков
// Ключ - английская фраза, значение - map[lang]translation
var translations = map[string]map[string]string{
        // === Флаги командной строки ===
        "Alignment in full width/height: \"start\", \"center\" or \"end\"": {
                "ru": "Выравнивание по полной ширине/высоте: \"start\", \"center\" или \"end\"",
        },
        "auto-hiDe: show dock when hotspot hovered, close when left or a button clicked": {
                "ru": "Автоскрытие: показывать док при наведении на горячую зону, скрывать при уходе или клике",
        },
        "Styling: css file name": {
                "ru": "Стилизация: имя css-файла",
        },
        "turn on debug messages": {
                "ru": "Включить отладочные сообщения",
        },
        "display Version information": {
                "ru": "Показать информацию о версии",
        },
        "set eXclusive zone: move other windows aside; overrides the \"-l\" argument": {
                "ru": "Установить исключительную зону: сдвинуть другие окна; отменяет аргумент \"-l\"",
        },
        "take Full screen width/height": {
                "ru": "Занять полную ширину/высоту экрана",
        },
        "quote-delimited, space-separated class list to iGnore in the dock": {
                "ru": "Список классов через пробел для игнорирования в доке",
        },
        "Hotspot Delay [ms]; the smaller, the faster mouse pointer needs to enter hotspot for the dock to appear; set 0 to disable": {
                "ru": "Задержка горячей зоны [мс]; чем меньше, тем быстрее курсор должен войти в зону; установите 0 для отключения",
        },
        "Hotspot Layer \"overlay\" or \"top\"": {
                "ru": "Слой горячей зоны \"overlay\" или \"top\"",
        },
        "alternative name or path for the launcher ICOn": {
                "ru": "Альтернативное имя или путь к иконке лаунчера; для эмодзи используйте формат emoji:СИМВОЛ",
        },
        "Ignore the running applications on these Workspaces based on the workspace's name or id, e.g. \"special,10\"": {
                "ru": "Игнорировать запущенные приложения на этих рабочих столах по имени или id, напр. \"special,10\"",
        },
        "Icon size": {
                "ru": "Размер иконок",
        },
        "Command assigned to the launcher button": {
                "ru": "Команда, назначенная кнопке лаунчера",
        },
        "Launcher button position, 'start' or 'end'": {
                "ru": "Позиция кнопки лаунчера: 'start' или 'end'",
        },
        "Layer \"overlay\", \"top\" or \"bottom\"": {
                "ru": "Слой \"overlay\", \"top\" или \"bottom\"",
        },
        "Margin Bottom": {
                "ru": "Отступ снизу",
        },
        "Margin Left": {
                "ru": "Отступ слева",
        },
        "Margin Right": {
                "ru": "Отступ справа",
        },
        "Margin Top": {
                "ru": "Отступ сверху",
        },
        "don't show the launcher button": {
                "ru": "Не показывать кнопку лаунчера",
        },
        "number of Workspaces you use": {
                "ru": "Количество используемых рабочих столов",
        },
        "Position: \"bottom\", \"top\" \"left\" or \"right\"": {
                "ru": "Позиция: \"bottom\", \"top\", \"left\" или \"right\"",
        },
        "Leave the program resident, but w/o hotspot": {
                "ru": "Оставить программу резидентной, без горячей зоны",
        },
        "name of Output to display the dock on": {
                "ru": "Имя вывода для отображения дока",
        },
        "allow Multiple instances of the dock (skip lock file check)": {
                "ru": "Разрешить несколько экземпляров дока (пропустить проверку файла блокировки)",
        },

        // === Сообщения сигналов ===
        "toggle dock visibility (USR1 has been deprecated)": {
                "ru": "переключить видимость дока (USR1 устарел)",
        },
        "show the dock": {
                "ru": "показать док",
        },
        "hide the dock": {
                "ru": "скрыть док",
        },

        // === Контекстное меню ===
        "Unpin": {
                "ru": "Открепить",
        },
        "Pin": {
                "ru": "Закрепить",
        },
        "New window": {
                "ru": "Новое окно",
        },
        "Close all windows": {
                "ru": "Закрыть все окна",
        },
        "closewindow": {
                "ru": "Закрыть окно",
        },
        "togglefloating": {
                "ru": "Плавающее окно",
        },
        "fullscreen": {
                "ru": "Полноэкранный режим",
        },

        // === Лог-сообщения ===
        "autohiDe and Resident arguments are mutually exclusive, ignoring -d!": {
                "ru": "Аргументы автоскрытия и резидентного режима взаимоисключающие, -d игнорируется!",
        },
        "HYPRLAND_INSTANCE_SIGNATURE not found, terminating.": {
                "ru": "HYPRLAND_INSTANCE_SIGNATURE не найден, завершение.",
        },
        "Starting in autohiDe mode": {
                "ru": "Запуск в режиме автоскрытия",
        },
        "Starting in resident mode": {
                "ru": "Запуск в резидентном режиме",
        },
        "SIGTERM received, bye bye!": {
                "ru": "Получен SIGTERM, до свидания!",
        },
        "SIGUSR1 for toggling visibility is deprecated, use SIGRTMIN+1": {
                "ru": "SIGUSR1 для переключения видимости устарел, используйте SIGRTMIN+1",
        },
        "Running instance found, terminating...": {
                "ru": "Найден запущенный экземпляр, завершение...",
        },
        "Sending sigToggle to running instance and bye, bye!": {
                "ru": "Отправка sigToggle запущенному экземпляру, до свидания!",
        },
        "Neither 'nwg-drawer' nor 'nwggrid' command found, and no other launcher specified; hiding the launcher button.": {
                "ru": "Ни 'nwg-drawer', ни 'nwggrid' не найдены, и другой лаунчер не указан; кнопка лаунчера скрыта.",
        },
        "Unable to start program": {
                "ru": "Невозможно запустить программу",
        },
        "Unable to launch command!": {
                "ru": "Невозможно запустить команду!",
        },
        "Error getting data directory:": {
                "ru": "Ошибка получения каталога данных:",
        },
        "Couldn't determine cache directory location": {
                "ru": "Не удалось определить расположение каталога кэша",
        },
        "Error connecting to the socket:": {
                "ru": "Ошибка подключения к сокету:",
        },
        "Error reading from socket2:": {
                "ru": "Ошибка чтения из socket2:",
        },
        "Unknown signal": {
                "ru": "Неизвестный сигнал",
        },

        // === Формат версии ===
        "nwg-dock-hyprland version": {
                "ru": "Версия nwg-dock-hyprland",
        },

        // === Прочие строки ===
        "Usage of signals": {
                "ru": "Использование сигналов",
        },
}

// T переводит строку на текущий язык. Если перевод не найден, возвращает оригинал.
func T(s string) string {
        if Lang == "" || Lang == "en" {
                return s
        }
        if m, ok := translations[s]; ok {
                if t, ok := m[Lang]; ok {
                        return t
                }
        }
        return s
}

// detectLang определяет язык из переменной окружения LANG/LC_ALL
func detectLang() string {
        lang := os.Getenv("LANG")
        if lang == "" {
                lang = os.Getenv("LC_ALL")
        }
        if lang == "" {
                lang = os.Getenv("LC_MESSAGES")
        }
        lang = strings.ToLower(lang)

        // Извлекаем базовый язык (ru_RU.UTF-8 -> ru)
        if len(lang) >= 2 {
                baseLang := lang[:2]
                switch baseLang {
                case "ru":
                        return "ru"
                default:
                        return "en"
                }
        }
        return "en"
}
