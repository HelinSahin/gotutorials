# Ders 2

## Go Dilinde Değişkenler, Değişken Tanımlama, Stringler ve Sayılar


### Degisken tanimlama

Go statik tipli derlenebilir bir dil oldugu icin, degiskenler tutacagi degerin tipine gore tanimlanir, ornegin sayisal bir deger tutan degiskene, metinsel bir deger atayamayiz.

Go da degisken tanimlamak icin `var` anahtar kelimesini kullaniriz. 

2 farkli degisken atama yontemi vardir; Dolayli veya dogrudan tip belirterek.


#### Dogrudan tip belirterek: 
Degiskenin tipini dogrudan `var degisken_adi degisken_tipi` seklinde belirtiriz

ornek;
```go
var degisken_adi_bir int
```
ustteki ornekte degisken_adi_bir adinda tamsayi tipinde bir degisken tanimlamis olduk


```go

var sayi int
var isim string
```
burdaki ornekte ise sayi adinda tamsayi tipi bir degisken, isim adinda string tipinde bir degisken tanimlamis olduk

#### Değişkenlerin İlk Değerleri:
Go'da bir değişken tanımlandığında, ona otomatik olarak bir başlangıç değeri atanır. Bu başlangıç değeri, değişkenin veri tipine bağlıdır. Örneğin, tamsayılar için başlangıç değeri 0'dır ve metin dizileri için başlangıç değeri boş bir dizedir ("").
#### Degiskene deger atama:

Go dilinde bir degiskene deger atamak icin `=` operatoru kullanilir.


```go
sayi = 24
isim = "Helin"
```
yukaridaki ornek Yukarıdaki örnek, sayi değişkenine 24'u ve isim değişkenine "Helin" metnini atar.



not#: string (metin dizisi) ifadeler atanirken `""` degeri cift tirnaklar arasina yazariz 


#### Dolayli yoldan tip belirterek degisken tanimlama

Degiskenin tipini belirtmeden dogrudan degerini atayarak gerceklestirilir. Degisken tipini atanan degere uyan veri tipinden alir

```go

var sayi = 10 
```
ustteki ornek bize tipini belirtmedigimiz halde, tamsayi (int) tipinde sayi isimli bir degisken tanimlar 


#### kisa degisken tanimlama

var anahtar kelimesi kullanilmadan `:=` operatoru ile atanarak yapilir, degisken tipini dolayli yolla alir. kisa yollu tanimlama global scope icinde yapilamaz (scope kavrami bir kac sonraki dersin konusu)

```go
sayi := 10

```

burda sayi adinda int tipinde bir degisken tanimlamis olduk


```go
var sayi1 = 10
sayi2 := 10
```

bu iki ifade aslinda ayni seydir diyebiliriz.


### strings

Stringler, metin verilerini temsil eder. Bir string, çift veya tek tırnak içinde belirtilir. Örneğin:

```go
ad := "Helin"
mesaj := "Miau!"
```

### Sayilar

o dilinde farklı tiplerde sayılar bulunur. Bunlar arasında tamsayılar (int), kayan noktalı sayılar (float), karmaşık sayılar (complex), kesirli sayılar (rational) vb. bulunur.

```go
tamsayi := 10 // int
kesirli := 3.14 // float
```

## Ozet

1. go ile dolayli ve dogrudan tip atamasi yapabilecegimizi ogrendik
2. go ile bir degisken olusturuldugunda, sadece o tipte veriler tutabilecegini ogrendik
3. go ile bir degisken olusturuldugunda, o degiskene eger bir deger atanmamis ise, go compailerinin degisken tipine gore varsayilan bir deger atadigini ogrendik
4. metinsel ve sayisal veri tiplerini gorduk


## alistirma

1. bir go dosyasi olusturun

2. string tipinde degisken olusturun

3. string tipinde degiskene helin ifadesi atayin

4. dolayli yoldan tipini belirttiginiz bir degisken olusturun

5. kisa yoldan atama yaptiginiz bir degisken olusturunuz

6. tamsayi ve kesirli sayi tutan degiskenler tanimlayiniz

7. olusturulan degiskenleri konsol ekranina yazdiriniz