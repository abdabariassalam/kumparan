# Test Kumparan

Features:

- mempunyai 2 api dengan bahasa Golang

## Setup explained

### Tooling

-   docker-compose:

    -   `docker-compose up --build` - running dengan perintah build atau rebuild applikasi.
    -   `docker-compose up -d` - running app di background dengan settingan terakhir kali build.
    -   `docker-compose down` - stop dan remove resource.


### Included sample packages

-   -   **app**
    -   [Gin](https://github.com/gin-gonic/gin) untuk route aplikasi.
    -   Listens on http://localhost:8080 .

## dev

kita bisa menjalan langsung semua app sekaligus dengan menggunakan syntax berikut:
```
docker-compose up --build
```

### Get
untuk mencoba api get articles bisa menggunkan **postman**

atau menggunakan syntax curl dibawah ini

```
curl --location --request GET 'localhost:8080/articles?search=test&author=test1'
```

search dan author adalah variable optional bisa di pake atau tidak
kalau untuk get all article hapus saja search dan author nya seperti berikut:

```
curl --location --request GET 'localhost:8080/articles'
```

### Post
untuk mencoba api create articles bisa menggunkan **postman**

atau menggunakan syntax curl dibawah ini

curl --location --request POST 'localhost:8080/articles' \
--header 'Content-Type: application/json' \
--data-raw '{
    "author": **author**,
    "title": **title**,
    "body": **body**
}'

semua parameter harus ada
dan ditambahkan penjagaan tidak boleh author dan title memiliki isi yang sama berturut-turut menghindari data duplicate