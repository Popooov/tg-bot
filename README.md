# tg-bot

Привіт! Це Telegram-бот на мові програмування Golang з використанням фреймворків `github.com/spf13/cobra` та `gopkg.in/telebot.v3`.

[Бот у Telegram](https://t.me/popov2_bot)

## Інструкції для встановлення

1. **Клонуйте репозиторій:**
    ```bash
    git clone https://github.com/Popooov/tg-bot.git
    cd tg-bot
    ```

2. **Встановіть залежності:**
    ```bash
    go get -u github.com/spf13/cobra
    go get -u gopkg.in/telebot.v3
    ```

3. **вкажіть токен свого Telegram-бота:**
    ```go
    package cmd

    var (
        TbotToken = os.Getenv("TELE_TOKEN")
    )
    ```

4. **Запустіть бота:**
    ```bash
    ./tg-bot start
    ```

## Приклади використання команд

1. **/start**
   - Початок роботи з ботом.

2. **/help**
   - Вивести список доступних команд.

3. **/version**
   - Вивести поточну версію бота.

4. **/time**
   - Вивести поточний час.
---
