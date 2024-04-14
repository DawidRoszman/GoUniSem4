# Silna i słaba liczba

Program wczytuje imię i nazwisko, a następnie tworzy nick składający się z trzech pierwszych liter imienia i nazwiska np. Z "Dawid Roszman" tworzy się nick "DawRos". Ten nick następnie przechodzi przez parę funckcji które usuwają polskie litery i zamieniają wielkie na małe litery. Po zakończeniu tych funkcji nasz nick ma postać "dawros".

Następnie program zamienia literki nicku na odpowiadający im liczby ASCII. Z nicku "dawros" zostaje utworzona tablica [100 97 119 114 111 115].

Kolejnym krokiem jest znalezienie naszej silnej liczby. Program szuka silni która zawiera wszyskie wymienione powyżej liczby. Dla nicku "dawros" tą liczbą jest 261. Wyniki tak dużych silni nie mieściły się w zakresię int64, więc musiałem zastosować bibliotekę "math/big".

Żeby znaleźć naszą słabą liczbę program musi rekurencyjnie obliczyć wartość Ciągu Fibonacciego dla liczby 30 oraz wypisać liczbę wywołań dla wartości 1,2,3 i kolejnych. Wartość której wywyołania są najbliższe naszej silnej liczbie jest naszą słabą liczbą. Dla nicku "dawros" słabą liczbą jest 18, której ilość wywołań to 233.

Potem program oblicza ile czasu zajmuje obliczenie wartości ciągu Fibonacciego dla pewnych liczb. Oraz oszacowuje ile czasu program będzie obliczać bardzo duże liczby Fibonacciego.
Oszacowanie polega na obliczeniu średniej różnicy czasu pomiędzy wykonaniem F(n) i F(n-1), a następnie oszacowaniu czasu wykonania funkcji F(m+k) = F(m) \* avg^k, gdzie avg to średnia obliczona wcześniej, m to czas wykonania funkcji F dla pewnego m (w programie jest to 30).
