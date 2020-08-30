# comper-crawler-go

### How to install Go

To run this code, you'll must have Go in your machine. Follow [this](https://linuxize.com/post/how-to-install-go-on-ubuntu-20-04/) steps to install on Ubuntu.

### How to run

```sh
git clone https://github.com/igorgbianchi/comper-crawler-go
cd comper-crawler-go
go build
./comper-crawler-go
```

Result can be viewed on ```output.json```.

### Output example
```json
[
    {
        "name": "Cerveja Patagônia 473ml Amber Lager",
        "url": "https://www.comperdelivery.com.br/cerveja-patagonia-473ml-amber-lager-/p",
        "img_url": "https://comper.vteximg.com.br/arquivos/ids/165427-292-292/1970550_1.jpg?v=637224382837530000",
        "price": 7.9,
        "availability": true
    },
    {
        "name": "Vinho Chileno Heroes 750ml Reserva Especial Cabernet",
        "url": "https://www.comperdelivery.com.br/vinho-chileno-heroes-750ml-reserva-especial-cabern/p",
        "img_url": "https://comper.vteximg.com.br/arquivos/ids/171035-292-292/2262096.png?v=637274747952270000",
        "price": 53.99,
        "availability": true
    },
    {
        "name": "Vinho Português Gabarito 750ml Tinto",
        "url": "https://www.comperdelivery.com.br/vinho-portugues-gabarito-tinto-750ml/p",
        "img_url": "https://comper.vteximg.com.br/arquivos/ids/159934-292-292/1669982.jpg?v=637210528616100000",
        "price": 59.99,
        "availability": true
    }
]
```