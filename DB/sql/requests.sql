-- Active: 1664892261064@@127.0.0.1@3306@adverts
sqlite3 upgrade.db 

-- создаём таблицу пользователей
CREATE TABLE
    users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name varchar(255),
        telegram_id INT,
        first_name varchar(255),
        last_name varchar(255),
        chat_id INT
    );

.quit 

-- если надо удалить кого-то
-- sqlite3 upgrade.db 
-- DELETE FROM users WHERE id == 1;
