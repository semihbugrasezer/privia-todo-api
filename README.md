Privia Staj Backend TODO App

Proje Yapısı 

├── main.go
├── controllers
│   ├── todo_controller.go
│   ├── user_controller.go
├── models
│   ├── todo.go
│   ├── user.go
├── middlewares
│   └── auth_middleware.go
├── utils
│   ├── jwt.go
│   └── response.go
├── routes
│   └── routes.go
├── mock
│   └── data.go
└── go.mod

Programlama Dili: GoLang
Framework: Gin
JWT için: github.com/dgrijalva/jwt-go/v4
Mock veri saklama: Dahili veri yapılarını kullanarak
Yayınlama: Vercel

API'nin çeşitli uç noktaları bulunmaktadır:

Genel Rotalar

POST /api/v1/login : Kullanıcı giriş yapma.
To-Do Listeleri (Yetkilendirme Gerektirir)

GET /api/v1/todos : Tüm to-do listelerini getir.
POST /api/v1/todos : Yeni bir to-do listesi oluştur.
GET /api/v1/todos/:id : Belirli bir to-do listesini getir.
PUT /api/v1/todos/:id : Belirli bir to-do listesini güncelle.
DELETE /api/v1/todos/:id : Belirli bir to-do listesini sil.
To-Do Öğeleri (Yetkilendirme Gerektirir)

GET /api/v1/todos/:todoId/items : Belirli bir to-do listesindeki tüm öğeleri getir.
POST /api/v1/todos/:todoId/items : Belirli bir to-do listesine yeni bir öğe ekle.
GET /api/v1/todos/:todoId/items/:itemId : Belirli bir to-do listesindeki belirli bir öğeyi getir.
PUT /api/v1/todos/:todoId/items/:itemId : Belirli bir to-do listesindeki belirli bir öğeyi güncelle.
DELETE /api/v1/todos/:todoId/items/:itemId : Belirli bir to-do listesindeki belirli bir öğeyi sil.
