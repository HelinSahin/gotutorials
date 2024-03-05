# Ders3

## Go dilinde string formatlama ve yazdirma

Go dilinde string bicimlendirme ve yazdirmak icin `fmt` paketi kullanilir.


### Print() fonksiyonu:

Go dilinde string yazdirmak icim kullanilan en basit fonksiyondur

```go

fmt.Print("merhaba, ")
fmt.Print("dunya!")
```

terminal ciktisi su sekilde olacaktir
```bash
merhaba, dunya!
```

goruldugu uzere cursor bir alt satira gecmedi, eger cursor alt satira gecmesini istiyorsak `\n` alt satir ozel karakterini kullanabiliriz


```go
fmt.Print("selam, \n")
fmt.Print("dunya!")
```
terminal ciktisi su sekilde olacaktir
```bash
selam,
dunya!
```


### Println() fonksiyonu:

Go dilinde string yazdirmak icin kullanilan fonksiyonlardan biri, `Print()`'e benzer ancak ciktidan sonra alt satira gecer

```go

fmt.Println("merhaba, ")
fmt.Println("dunya!")
```

terminal ciktisi su sekilde olacaktir
```bash
merhaba, 
dunya!
```


not: `Print()` ve `Println()` fonksiyonlari `Print(1,2,3,...)` seklinde sinirsiz arguman alabilir. Aldigi butun argumanlari birlestirip tek bir yazi dizesine cevirecektir.   

### Printf() fonksiyonu:

Bicimlendirilmis yazi dizesi yazdirmamizi saglar, Bu fonksiyon C dilindeki printf fonksiyonuna benzer.


```go
package main

import "fmt"

func main() {
    isim := "Ahmet"
    yas := 30
    fmt.Printf("İsim: %s, Yaş: %d\n", isim, yas)
}
```

### Biçimlendirme Dizgileri:

Printf fonksiyonu tek basina bir anlam ifade etmiyor, o yuzden bicimlendirme dizgileri ile birlikte kullaniyoruz. Formatlanacak metin icinde ilgili yerlere istedigimiz bilgileri yerlestirir. Go daki bicimlendirme dizgileri sunlardir

`%s` yazi dizisi

`%d` tamsayi

`%f` Ondalık sayı (float)

`%t` Boolean (true/false)

`%v` Değişkenin doğal türüne göre biçimlendirme

`%%` Yüzde işareti

`%T` Degiskenin tipini yazdirir

`%p` Pointerlar


```go

package main

import "fmt"

func main() {
    miktar := 10
    fiyat := 19.99
    urun := "Kedi mamasi"
    fmt.Printf("Product: %s, Quantity: %d, Price: %.2f\n", urun, miktar, fiyat)
}
```

cikti

```bash
Product: Golang Course, Quantity: 10, Price: 19.99
```



## Sprintf() fonksiyonu:

`Printf()` fonksiyonuna benzer ancak sonucu yazdirmak yerine, bicimlendirilmis metini geri cevirir


```go
package main

import "fmt"

func main() {
    isim := "Helin"
    yas := 28
    bicimlendirilmisMetin := fmt.Sprintf("Name: %s, Age: %d", isim, yas)
    fmt.Println(bicimlendirilmisMetin)
}
```