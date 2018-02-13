
# Obligatorisk oppgave 1: 


## 1. Fyll ut manglende tall.
 
<table style="width: 687px;" border="1">
<tbody>
<tr>
<td style="width: 265px;">

**Binære tall (mest signifikant bit til venstre)**

</td>
<td style="width: 204px;">

**Hexadesimaltall**

</td>
<td style="width: 195px;">

**Desimaltall**

</td>
</tr>
<tr>
<td style="width: 265px; text-align: center;">

<p>1101</p>

</td>
<td style="width: 204px; text-align: center;">

<p>0xD</p>

</td>
<td style="width: 195px; text-align: center;">

<p>13</p>

</td>
</tr>
<tr>
<td style="width: 265px; text-align: center;">

<p>110111101010</p>

</td>
<td style="width: 204px; text-align: center;">

<p>0xDEA</p>

</td>
<td style="width: 195px; text-align: center;">

<p>3562</p>

</td>
</tr>
<tr>
<td style="width: 265px; text-align: center;">

<p>1010111100110100</p>

</td>
<td style="width: 204px; text-align: center;">

<p>0xAF34</p>

</td>
<td style="width: 195px; text-align: center;">

<p>44852</p>

</td>
</tr>
<tr>
<td style="width: 265px; text-align: center;">

<p>1111111111111111</p>

</td>
<td style="width: 204px; text-align: center;">

<p>0xFFFF</p>

</td>
<td style="width: 195px; text-align: center;">

<p>65535</p>

</td>
</tr>

<tr>
<td style="width: 265px; text-align: center;">

<p>10001011110001010</p>

</td>
<td style="width: 204px; text-align: center;">

<p>0x1178A</p>

</td>
<td style="width: 195px; text-align: center;">

<p>71562</p>

</td>
</tr>
</tbody>
</table>
<h2><span style="font-size: 18pt;"></span></h2>



### Oppgave A. Beskriv kort metode for å gå fra binære tall til hexadesimale tall og motsatt. Beskriv kort metoden for å gå fra binære tall til desimaltall og motsatt

**Metode for å gå fra binære tall til hexadesimale tall**

I denne metoden begynner vi bakerst med å dele det binære tallet i grupper på fire siffer hver. Deretter gjør vi hver gruppe med fire siffer i de binære tallene om til hexadesimale tall. Til slutt legger vi tallene og bokstavene vi får sammen i tilsvarende rekkefølge. 

For eksempel: 
Ta utgangspunkt i det binære tallet 101111101101. 
```
1011 = B
1110 = E
1101 = D

Det hexadesimale tallet blir dermed ordet BED. 
```

**Metode for å gå fra hexadesimale tall til binære tall**

Først bruker vi hvert enkelt siffer i det hexadesimale tallet til å finne det tilsvarende binære tallet. Til slutt setter vi disse tallene sammen til et helt tall, ved å følge rekkefølgen tallene stod i som hexadesimale tall. Dersom det binære tallet vi får blir langt, deler vi det opp i grupper.

For eksempel: 
La oss ta utgangspunkt i det hexadesimale tallet BE4D.
```
B (11) = 1011
E (14) = 1110
4 (4)  = 0100
D (13) = 1101

Slår vi disse sammen får vi det binære tallet 1011 1110 0100 1101. 
```

**Metode for å gå fra binære tall til desimaltall**

I denne metoden tar vi i bruk følgende formel: *verdi = siffer x grunntall^posisjon*. Når vi har regnt ut hvert siffer summerer vi tallene vi får. Binærsystemet har grunntallet 2, og består kun av siffrene 0 og 1. 


For eksempel: 
Ta utgangspunkt i det binære tallet 100101. 
```
Posisjon   Siffer
0           1         1 * 2^0 = 1.  
1           0         0 * 2^1 = 0. 
2           1         1 * 2^2 = 4. 
3           0         0 * 2^3 = 0.
4           0         0 * 2^4 = 0. 
5           1         1 * 2^5 = 32. 

32 + 4 +1 = 37. 
```


**Metode for å gå fra desimaltall til binære tall**

Først deler man tallet på 2. Får man en rest skriver man 1, og dersom man ikke får rest skriver man 0. Så kjører man denne prosessen helt til svaret blir 0. Det ferdige binære tallet leser vi deretter i omvendt rekkefølge, altså nedenfra og opp. 

For eksempel: 
Ta utgangspunkt i tallet 133. 
```
133/2 = 1. 
66/2 = 0 
33/2 = 1 
16/2 = 0 
8/2 = 0 
4/2 = 0 
2/2 = 0 
1/2 = 1

Det binære tallet av 133 er derfor 10000101.
```

### Oppgave B. Beskriv kort metoden for å gå fra hexadesimale tall til desimaltall og motsatt.

**Metode for å gå fra hexadesimale tall til desimaltall**

I denne metoden tar vi, som i metoden fra binære tall til desimaltall, i bruk følgende formel: *verdi = siffer x grunntall^posisjon*. Når vi har regnt ut hvert siffer summerer vi tallene vi får. Det hexadesimale tallsystemet har grunntallet 16, og består av verdiene 0 til 9 og a til f. 

For eksempel: 
La oss ta utgangspunkt i det engelske ordet BED. 
```
Posisjon   Siffer
0           D (13)         13 * 16^0 = 13
1           E (14)         14 * 16^1 = 224.  
2           B (11)         11 * 16^2 = 2816.  


13 + 224 + 2816 = 3053. 
```


**Metode for å gå fra desimaltall til hexadesimale tall**

Først deler vi tallet på 16. Dersom vi får rest skriver vi ned hva vi får igjen, og dersom vi ikke får rest skriver vi 0. Så starter vi fra begynnelsen igjen helt til svaret blir 0. Til slutt leser vi det heksadesimale tallet vi får nedenfra og opp. Vi kan se sammenhenger mellom denne metoden og metoden for å gå fra desimaltall til binære tall. 

For eksempel: 
Ta utgangspunkt i tallet 2392. 
```
2392: 16 = 149,5
149 * 16 = 2384
2392 - 2384 =    8
8 i heksadesimal er 8

149:16 = 9,3
9*16 = 144
149 - 144 = 5
5 er 5 i heksadesimal

Resten er 9, tallet i heksadesimal blir altså 958.
```


## 2. Forstå algoritmer og utføre "benchmark"-tester på koden.

### Oppgave A. Skriv en modifisert bubble-sort funksjon benchmarkBSortModified basert på eksempel-funksjon Bubble_sort i filen sorting.go 



### Oppgave B. Skriv "benchmark"-tester for benchmarkBSortModified funksjonen basert på eksempel-funksjon benchmarkBSort i filen sorting_test.go



### Oppgave C. Utfør alle benchmark- testene med kommando “go test -bench=.” og presenter resultatene grafisk. Hva kan du si om big-O for alle 3 algoritmene som du har testet?




## 3. Forstå prosessadministrasjon på en platform. 

### Skriv et program som består av en evig løkke. 
Se loop.go i mappen "Oblig1". Den evige løkken skriver ut "venter på signal..." helt til den mottar et signal. 


### Hvor mye minne og CPU bruker programmet når det kjører. 
På MAC - Som bildet viser øker både CPU og minne betraktelig når programmet kjøres i terminal. 

-- (legg inn bilder) --

På Windows - Som bildet viser øker både CPU og minne litt når programmet kjøres i terminal.

|![Windows](https://github.com/Prestholm/TeamStovsuger/blob/master/Oblig1/105/terminalloop.png)|

### Generer ulike avslutningssignaler til prosessen og dokumenter hvilke avslutningskommandoer programmet håndterer og som trigger avslutningsmeldingen.

INT signal ("interrupt", SIGINT) håndteres av programmet fordi vi har spesifikt programmert det til å gjøre det. Denne trigger også avslutningsmeldingen. 

TSTP signal ("terminal stop", SIGTSTP) håndteres av programmet. Selv om dette ikke er programmert inn får signalet terminalen til å avslutte. Den genererer derimot ikke avlsutningsmeldingen i etterkant. 

QUIT signal (SIGQUIT) håndteres ikke av programmet, og skriver heller ikke ut avslutningsmelding. 



## 4. Typografiske symboler.

### Oppgave A. Bruk filen ascii.go i Oblig1 mappen og lag en funksjon som itererer (går i en løkke over)  over tegn med byte-verdier fra 0x80 til 0xFF, dvs. det utvidede ASCII settet. 

I mappen Oppgave 4 ser man filen ascii.go i mappen ascii som kan kjøres via asciimain.go. Dette programmet skal kunne itererer over tegn med byte-verdier fra 0x80 til 0xFF. Utskriften fra funksjonen har de fleste symbolene, noen tegn er utilgjenelige i goland, derfor viser de fram som en boks istedet for et symbol. De binære tallene virker korrekte i forhold til hexadesimalene. 


### Oppgave B. Lag funksjonen ExtendedASCIIText () i samme filen iso.go, som skriver ut: " € ÷ ¾ dollar "

I mappen Oppgave 4 i samme fil som i 4a) finner du ascii.go filen med ExtendedASCIItext(). Den skal skrive ut "€ ÷ ¾ dollar " fra byte-verdier lagt inn. For å få skrivet ut  € ÷ ¾ dollar, bruker vi x80\xf7\xbe\x64\x6f\x6c\x6c\x61\x72. Den er skrevet og testet først i Windows goland. Noen av symbolene vil ikke vises i kommandovinduet. Resultatene er likt selv i goland, powershell og terminalen. Terminalene har ikke tegnsettet til for alle disse symbolene i sin rekkevidde, derfor vil noen av symbolene være markert som spørsmålstegn i boks, eller bare en boks. Selv tegnesettet i windows-terminalen gir ikke symbolene selv om standarden sier at symbolene fins i settet.

### Oppgave C. Implementer en test for funksjonen ExtendedASCIIText(String) i egen fil iso_test.go, som tester om input-verdier (av type string) inneholder kun tegn fra en Extended ASCII.


I mappen oppgave 4/ascii finner du iso_test.go som tar ut verdiene fra ExtendedASCIItext() og tester om de er i en verdi i det extended ascii. Alle verdiene er ikke i det extended ascii. € ÷ ¾ er alle i det extended ascii, men " dollar" er tegn som ikke er i det extended ascii.
