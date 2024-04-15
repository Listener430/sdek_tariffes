# As requested here is an updated custom read.me for a GCP project
# вызов СДЕК API для функции расчета тарифов

Входящие данные просто подставляются в main:

Результат выводится в консоль в формате(пример результата для входящих данных):
```json

{
    "tariff_codes": [
        {
            "tariff_code": 62,
            "tariff_name": "Магистральный экспресс склад-склад",
            "delivery_mode": 4,
            "delivery_sum": 750.0,
            "period_min": 8,
            "period_max": 12,
            "calendar_min": 8,
            "calendar_max": 12
        },
        {
            "tariff_code": 139,
            "tariff_name": "Посылка дверь-дверь",
            "delivery_mode": 1,
            "delivery_sum": 490.0,
            "period_min": 8,
            "period_max": 12,
            "calendar_min": 8,
            "calendar_max": 12
        },
        {
            "tariff_code": 138,
            "tariff_name": "Посылка дверь-склад",
            "delivery_mode": 2,
            "delivery_sum": 350.0,
            "period_min": 8,
            "period_max": 12,
            "calendar_min": 8,
            "calendar_max": 12
        },
        {
            "tariff_code": 366,
            "tariff_name": "Посылка дверь-постамат",
            "delivery_mode": 6,
            "delivery_sum": 350.0,
            "period_min": 8,
            "period_max": 12,
            "calendar_min": 8,
            "calendar_max": 12
        },
        {
            "tariff_code": 137,
            "tariff_name": "Посылка склад-дверь",
            "delivery_mode": 3,
            "delivery_sum": 350.0,
            "period_min": 8,
            "period_max": 12,
            "calendar_min": 8,
            "calendar_max": 12
        },
        {
            "tariff_code": 136,
            "tariff_name": "Посылка склад-склад",
            "delivery_mode": 4,
            "delivery_sum": 210.0,
            "period_min": 8,
            "period_max": 12,
            "calendar_min": 8,
            "calendar_max": 12
        },
        {
            "tariff_code": 368,
            "tariff_name": "Посылка склад-постамат",
            "delivery_mode": 7,
            "delivery_sum": 210.0,
            "period_min": 8,
            "period_max": 12,
            "calendar_min": 8,
            "calendar_max": 12
        },
        {
            "tariff_code": 480,
            "tariff_name": "Экспресс дверь-дверь",
            "tariff_description": "Экспресс-доставка",
            "delivery_mode": 1,
            "delivery_sum": 540.0,
            "period_min": 2,
            "period_max": 4,
            "calendar_min": 2,
            "calendar_max": 4
        },
        {
            "tariff_code": 481,
            "tariff_name": "Экспресс дверь-склад",
            "tariff_description": "Экспресс-доставка",
            "delivery_mode": 2,
            "delivery_sum": 460.0,
            "period_min": 2,
            "period_max": 4,
            "calendar_min": 2,
            "calendar_max": 4
        },
        {
            "tariff_code": 485,
            "tariff_name": "Экспресс дверь-постамат",
            "tariff_description": "Экспресс-доставка",
            "delivery_mode": 6,
            "delivery_sum": 460.0,
            "period_min": 2,
            "period_max": 4,
            "calendar_min": 2,
            "calendar_max": 4
        },
        {
            "tariff_code": 482,
            "tariff_name": "Экспресс склад-дверь",
            "tariff_description": "Экспресс-доставка",
            "delivery_mode": 3,
            "delivery_sum": 460.0,
            "period_min": 2,
            "period_max": 4,
            "calendar_min": 2,
            "calendar_max": 4
        },
        {
            "tariff_code": 483,
            "tariff_name": "Экспресс склад-склад",
            "tariff_description": "Экспресс-доставка",
            "delivery_mode": 4,
            "delivery_sum": 380.0,
            "period_min": 2,
            "period_max": 4,
            "calendar_min": 2,
            "calendar_max": 4
        },
        {
            "tariff_code": 486,
            "tariff_name": "Экспресс склад-постамат",
            "tariff_description": "Экспресс-доставка",
            "delivery_mode": 7,
            "delivery_sum": 380.0,
            "period_min": 2,
            "period_max": 4,
            "calendar_min": 2,
            "calendar_max": 4
        }
    ]
}

```
