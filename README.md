# Cкилл «Эхо»

Скилл доступен по Webhook URL: https://kna0ifamyg.execute-api.eu-central-1.amazonaws.com/dev/webhook

Повторяет сообщение пользователя: `response.text` и `response.tts` равны полю `request.original_utterance` запроса.

Подробнее о создании скилов для Маруси можно узнать тут: [https://vk.com/dev/marusia_skill_docs](https://vk.com/dev/marusia_skill_docs)

## Деплой на AWS Lambda

1. [Установить Serverless Framework](https://www.serverless.com/framework/docs/providers/aws/guide/installation/) и [настроить AWS Credentials](https://www.serverless.com/framework/docs/providers/aws/guide/credentials/)
2. Выполнить `make deploy` в директории проекта. Присвоенный Webhook URL будет выведен в терминал в разделе `endpoints`.
