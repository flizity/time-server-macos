# 🕒 NTP Time Sync - Гибкое приложение для синхронизации времени

![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)
![Fyne Version](https://img.shields.io/badge/Fyne-v2.4%2B-9cf)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

Кроссплатформенное приложение для синхронизации системного времени с NTP-серверами, написанное на Go с использованием GUI-фреймворка Fyne.

## 🌟 Особенности

- 🖥️ **Кроссплатформенность** (Windows, macOS, Linux)
- ⏱️ Точное сравнение локального времени с NTP
- 🔄 Автоматический перебор нескольких серверов при ошибках
- 🔔 Уведомления о завершении синхронизации
- ⚠️ Предупреждения о большом расхождении времени
- 🛠️ Рекомендации по ручной синхронизации для macOS
- 🎨 Простой и интуитивно понятный интерфейс

## 📦 Установка

### Требования
- Установленный Go (версия 1.21 или выше)
- Fyne CLI (для сборки): `go install fyne.io/fyne/v2/cmd/fyne@latest`

### Сборка из исходников
```bash
git clone https://github.com/flizity/time-server-macos.git
cd time-server-macos
go build -o ntp-sync ./cmd/main.go

🚀 Использование
Запустите приложение

Введите предпочитаемый NTP-сервер (по умолчанию time.apple.com)

Нажмите кнопку "Sync Time"

Просмотрите результаты:

Точное время NTP-сервера

Ваше локальное время

Разницу между ними

🖼️ Скриншоты интерфейса
![image](https://github.com/user-attachments/assets/48c7eff8-d8e2-413c-af75-5cb6c8bb64fe)

⚙️ Поддерживаемые NTP-серверы
Приложение автоматически проверяет несколько серверов:

time.apple.com

time.google.com

ntp1.stratum2.ru

pool.ntp.org

Вы можете указать любой другой NTP-сервер вручную.
![image](https://github.com/user-attachments/assets/0a93e6f2-9521-4774-8988-937defbfae6e)

🛠️ Особенности для macOS
При обнаружении расхождения более 5 секунд приложение:

Покажет предупреждение

Предоставит инструкции для ручной синхронизации:

bash
sudo sntp -sS time.apple.com
