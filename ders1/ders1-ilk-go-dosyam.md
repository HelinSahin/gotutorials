# Ders 1

## go dili nedir

Go, Google'da 2007 yılından itibaren geliştirilmeye başlayan açık kaynak programlama dilidir. İlk web sitesi golang.org alan adına sahip olduğundan golang ismiyle anılsa da doğru adı Go'dur. Daha çok sistem programlama için tasarlanmış olup, derlenmiş ve statik tipli bir dildir

## Ilk go dosyam

### Go dosyasi:
Go dile ile programlama yapabilmek icin go dosyalarina ihtiyacimiz var.

go dosyalari ise `.go` uzantisi ile olusturulur

ornegin;
```cmd
dosya.go
```

#### Devam etmeden once ders1 icindeki main.go dosyasini inceleyin

### package:
Go dilinde, kodu organize etmek ve tekrar kullanılabilirlik sağlamak için paketler kullanılır. Bir Go dosyasının başında, o dosyanın hangi pakete ait olduğunu belirtmek zorundayiz. 

main paketi ayni zamanda uygulamanin giris noktasidir. Bir go compaileri uygulamayi derlemeye main paketinden baslar. O yuzden tek bir paketten olusan uygulamalar icin main paketi kullanilir.


go dosyasinda package tanimi dosyanin en basina yapilir
```go
package main
```
yukaridaki kod parcasi main paketini tanimlamistir. main.go dosyasini inceleyerek bu tanimi bulabilirsin

not: package konusu cok daha detayli sekilde bakacaz suanlik bu kadarini bilmen yeterli. 

----



### Ana Fonksiyon (Main Function):

Bir Go programında, yürütme işlemi main adlı özel bir fonksiyonla başlar. Bu fonksiyon, uygulamanın ana giriş noktasıdır.

Yani go ile bir program olusturdugumuzda, programin yurutulmesi main paketindeki main fonksiyonundan baslar. Eger programda main fonksiyonu bulunmuyorsa program execute  edilemez (yurutulemez)


```go
func main() {
   // program calistirildiginda calisacak kodlar buraya gelecek
}
```

Yukarıdaki kod parçası, main fonksiyonunu tanımlar. Program buradan başlayacak ve bu fonksiyon içindeki kodlar çalıştırılacak.


### İçe Aktarma (Import) İfadeleri:

Başka bir paketin fonksiyonlarını veya özelliklerini kullanmak için import ifadesi kullanılır. Örneğin, standart kütüphaneden fmt paketini içe aktararak, konsola çıktı yazdırabiliriz.

```go
import "fmt"
```

Yukarıdaki kod parçası, fmt paketini içe aktardigimizi (import) ettigimizi ifade ediyor.


not: package yapisi ile birlikte standart kutuphaneyide isleyecez suan buna takilmana gerek yok sadece fmt (format) paketinin icinde Println gibi consola cikti almamizi saglayan bir fonksiyon oldugunu bilmen yeterli

---

ornek olarak ders1 icindeki main.go dosyasini inceleyebilirsin


### sonuc

1. Go dosyalarinin `.go` uzantisi ile olusturuldugunu ogrendik
2. Her go dosyasinin hangi pakete ait oldugunu `package` ifadesi ile belirtmek zorunda oldugumuzu ogrendik
3. Her go uygulamasinin `main` paketinden derlenmeya basladigini ogrendik
4. Her go uygulamasinin calismaya baslarken giris noktasinin `main()` fonksiyonu oldugunu ogrendik.
5. Uygulamamiza diger paketleri `import` ifadesi ile katabilecegimizi ogrendik


### alistirma

calistirildiginda terminale `Tatlı dil yılanı deliğinden çıkarır. ` ciktisini versin

ip ucu#: go dilinde terminale cikti alabilmek icin `fmt` paketi altindaki `Println()` fonksiyonunu kullanabilirsiniz

ip ucu2#: go uygulamasini bulundugu dizin icinde terminalden `go run dosya_adi.go` seklinde calistirabilirsiniz 