# Ders 5

## Standart kutuphane, strings ve sort paketi


Go'daki standart kitaplık, ortak programlama görevleri için çeşitli işlevler sunan zengin bir paket seti sağlar. Standart kitaplıktaki iki önemli paket dizeler ve sıralamadır. Strings paketi, stringleri işlemek için işlevler sağlarken, sort paketi farklı veri türleri için sıralama algoritmaları sunar.

### I. Standard kitaplik
Go'daki standart kitaplık, Go programlama diline dahil olan paketlerin bir koleksiyonudur.
Bu paketler çok çeşitli işlevler sunarak Go'da uygulama geliştirmeyi kolaylaştırır.

Bir paketin işlevlerini kullanmak için paketin `import` anahtar sözcüğü kullanılarak içe aktarılması gerekir.
Örnek: `import "fmt"`, biçimlendirilmiş giriş/çıkış için fmt paketini içe aktarır.


Standart kitaplık, dosya G/Ç, ağ iletişimi, eşzamanlılık, şifreleme ve daha fazlasına yönelik paketler içerir.
Örnekler: fmt, os, net, senkronizasyon, kripto vb.


### strings Paketi

Strings paketi, stringleri işlemek ve onlarla çalışmak için işlevler sağlar.

* strings.Contains(s, substr), s dizesinin bir alt dize substr içerip içermediğini kontrol eder.
* strings.HasPrefix(s, prefix), bir dizenin belirli bir önekle başlayıp başlamadığını kontrol eder.
* strings.Join(strs, sep), strs dizelerinin bir dilimini sep ile ayrılmış tek bir dizede birleştirir.
* strings.Split(s, sep), bir sep ayırıcı kullanarak bir s dizesini alt dizelere böler.

### sort Paketi
sort paketi, farklı veri türleri için sıralama algoritmaları sağlar.

#### dizeleri siralama
sort.Ints(slice) bir tamsayı dilimini artan düzende sıralar.
sort.Strings(slice), bir dize dilimini alfabetik olarak sıralar.