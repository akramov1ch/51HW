# Uyga vazifa: Go yordamida gRPC Server Streaming bilan Ob-havo Yangilash Tizimi

## Maqsad

Bu vazifaning maqsadi client uchun ma'lum bir joy uchun real vaqt rejimida ob-havo yangilanishlarini olish imkoniyatini yaratishdir. Siz bu funktsiyani amalga oshirish uchun `Go` yordamida `gRPC` server streamingdan foydalanasiz.

## Vazifa Tafsilotlari

1. **Protokol Buffers (protobuf) faylini aniqlash**: `.proto` faylini `gRPC` xizmati va xabar turlarini belgilash uchun aniqlang.
2. **`.proto` faylidan Go kodini generatsiya qilish**: `protoc` kompilyatoridan foydalanib `Go` kodini generatsiya qiling.
3. **gRPC serverni amalga oshirish**: Ob-havo yangilanishlarini uzatadigan serverni amalga oshirish.
4. **gRPC clientini amalga oshirish**: Ob-havo yangilanishlariga obuna bo'lib, ularni chop etadigan clientni amalga oshirish.

## Bosqichma-bosqich yo'riqnoma

### 1. `.proto` Faylini aniqlash
### 2. `proto` Faylidan Go Kodini Generatsiya Qilish
### 3. `gRPC` Serverni Amalda Qo'llash. Ob-havo ma'lumotlari tasodifiy ravishda yoki `OpenWeatherMap API` orqali olinishi mumkin
### 4. `gRPC` Clientni Amalda Qo'llash


```plaintext
+-------------------+   WeatherRequest   +----------------------+
|                   | -----------------> |                      |
|    gRPC Client    |                    |      gRPC Server     |
|                   | <----------------- |                      |
+-------------------+  WeatherUpdate     +----------------------+
                                            |                  |
                                            | HTTP Request     |
                                            v                  v
                            +-------------------+         +--------------------+
                            |                   |         |                    |
                            |   Weather API     |         |   Random Weather   |
                            |    (Optional)     |         |      Generator     |
                            |                   |         |                    |
                            +-------------------+         +--------------------+
```


