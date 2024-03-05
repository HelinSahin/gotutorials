
# Ders 4

## Go da 'Array' ve 'Slice'lar


### Array: 

Sabit boyutlu ve aynı türdeki verileri içeren veri yapılarıdır.

```go
var numbers [5]int // 5 elemanlı bir tamsayı array'i

```

* Boyutu tanımlandığında sabittir.

* Bellekte ardışıl olarak yer kaplar.

* Genellikle belirli bir boyutta ve tipde veri koleksiyonu için kullanılır.

### len() fonksiyonu:

Bir dizenin kactane elemandan olustugunu doner.

```go

var numbers [5]int

fmt.Print(len(numbers)) // 5
```


###  Slice'lar:
Esnek boyutlu ve dinamik olarak büyüyen veri yapılarıdır. Array'lerin üzerine inşa edilmiştir. Go'da en cok kullanilacak dize tipi slicelardir.

Slice'lar esnek boyutlarıyla veri manipülasyonunda daha kullanışlıdır.

### append() fonksiyonu:

Bir slice yeni bir eleman eklemek icin kullanilir, orjinal dizeyi degistirmez. ekleme isleminden sonra yeni bir dize doner.


```go

numbers = append(numbers, 1)
numbers = append(numbers, 2)
```


```go
var numbers []int // boş bir slice
```

* Array'lerin bir görünümüdür ve boyutları yoktur.
* Verileri kesirli (subarray) bir kısmı alabilirler.
* make() fonksiyonu veya slice literali kullanılarak oluşturulurlar.
* Dinamik boyutları sayesinde veri eklemek, çıkarmak ve değiştirmek için kullanılırlar.


### Array ve Slice Kullanımı

Array:
```go
var numbers [5]int
numbers[0] = 1
numbers[1] = 2
```



slice:


not#1: Arraylar sabit boyutludur, slicelar ise bilinmeyen(esnek) boyutludur

not#2: Array tanimlarken boyut bilgisi vermek gereklidir, Slice tanimi yaparken boyut bilgisi belirtilmez





ornek:

```go
package main

import "fmt"

func main() {
    // Array tanımı
    var numbers [5]int
    numbers[0] = 10
    numbers[1] = 20
    fmt.Println("Array:", numbers)

    // Slice tanımı ve kullanımı
    var slice []int
    slice = append(slice, 30)
    slice = append(slice, 40)
    fmt.Println("Slice:", slice)
}
```