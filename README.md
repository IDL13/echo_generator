# Утилита для генерации скелета Echo кода

### Проблема:

Постоянное создание и написаение одинакового кода для создания скелета микросервиса.

### Задача:

  1) Создать CLI приложение для генерации Echo сервисов с дальнейшим диплоем в продакшен;
  2) Поддержка документации;
  3) Большое количество функционала:
  4) Подгрузка нужных зависимостей;
  5) Совместимость с Git и дальнейший диплой на GitHub;

**Сценарии использования:**

- Запуск кода при помощи Makefile:
  - flag -new "name" позволяет создать полный echo код для папки с названием "name"
  - flag -git "link" позволяет создать зависимости с учетом git акаунта

**Пример использования:**
  - Для Windows:
      `make build_windows name_your_file`
  - Для Linux:
      `make build_linux name_your_file`
  