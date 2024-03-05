# Ders 7

## Go dilinde donguler ve dongu yapilari


Go'da, koleksiyonlar üzerinde yineleme yapmak veya bir kod bloğunu tekrarlamak için kullanılabilen birkaç döngü yapısı vardır. Yaygın olarak kullanılan döngü yapıları for, for-in, while ve do-while'dir. 

### For dongusu

For döngüsü Go'daki en yaygın ve çok yönlü döngü yapısıdır. Belirli bir koşul karşılanana kadar bir kod bloğunu tekrar tekrar çalıştırmanıza olanak tanır. Go'daki for döngüsünün genel sözdizimi aşağıdaki gibidir:

```go

for baslangic_degeri; kosul; gecis {
    // calistiralacak kodlar
}
```

* `baslangic_degeri`: döngü başlamadan önce değişkenleri başlatan veya değerleri atayan **isteğe bağlı** bir ifadedir.
* `kosul`: her yinelemeden önce değerlendirilen bir `Boolean` ifadesidir. Koşul **doğruysa** döngü devam eder; aksi takdirde sona erer.

* `gecis`: her yinelemenin sonunda yürütülen **isteğe bağlı** bir ifadedir.

ornek: 

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

```

cikti: 
```bash
0
1
2
3
4
```

### for-in dongusu:
For-in döngüsü array, slice veya map gibi koleksiyonlar üzerinde yineleme yapmak için kullanılır. Bir indeks değişkenine ihtiyaç duymadan koleksiyonun öğeleri arasında otomatik olarak yinelenir.

not#: map ileriki derslerin konusu

```go
for index, eleman := range koleksiyon {
    // calisacak kod
}

```

* `index`: mevcut elamanin dizedeki pozisyonunu(index)'ini verir
* `eleman`: koleksiyonun mevcut öğesini temsil eder.
* `range`: koleksiyonun yinelecegini belirtir
* `koleksiyon`: yinelecek olan dizi yada map objesi

not#: eger index'e ihtiyac yoksa, indexi yok saymak için  `_`  kullanılabilir.


```go
numaralar := []int{1, 2, 3, 4, 5}
for _, num := range numaralar {
    fmt.Println(num)
}
```

terminal ciktisi 
```bash
1
2
3
4
5
```

### while dongusu

Go'nun yerleşik bir while döngüsü yapısı yoktur, ancak for döngüsünü kullanarak bunu simüle edebilirsiniz. While döngüsü, belirtilen koşul doğru olduğu sürece kod bloğunu çalıştırmaya devam eder. Go'daki while döngüsünün genel sözdizimi aşağıdaki gibidir:

```go
for kosul {
    // calistirilacak kodlar
}
```

ornek:
```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```
terminal ciktisi:
```bash
0
1
2
3
4
```


### do-while dongusu

Go ayrıca yerleşik bir do-while döngüsü yapısına sahip değildir, ancak for döngüsünü kullanarak bunu simüle edebilirsiniz. Do-while döngüsü kod bloğunu en az bir kez çalıştırır ve belirtilen koşul doğru olduğu sürece devam eder. Go'daki do-while döngüsünün genel sözdizimi aşağıdaki gibidir:

```go
for {
    // calisacak kodlar
    if !condition {
        break
    }
}
```


ornek: 


```go
i := 0
for {
    fmt.Println(i)
    i++
    if i >= 5 {
        break
    }
}
```

terminal ciktisi
```bash
0
1
2
3
4
```