# Используем официальный образ PostgreSQL
FROM postgres:15

# Устанавливаем переменные окружения для базы
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=mydatabase

# Копируем скрипты инициализации (если есть)
# COPY ./init-db.sh /docker-entrypoint-initdb.d/

# Экспонируем стандартный порт PostgreSQL
EXPOSE 5432