
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


**Metode for å gå fra hexadesimale tall til binære tall**


**Metode for å gå fra binære tall til desimaltall**


**Metode for å gå fra desimaltall til binære tall**



### Oppgave B. Beskriv kort metoden for å gå fra hexadesimale tall til desimaltall og motsatt.

**Metode for å gå fra hexadesimale tall til desimaltall**


**Metode for å gå fra desimaltall til hexadesimale tall**


## 2. Forstå algoritmer og utføre "benchmark"-tester på koden.

### Oppgave A. Skriv en modifisert bubble-sort funksjon benchmarkBSortModified basert på eksempel-funksjon Bubble_sort i filen sorting.go 


### Oppgave B. Skriv "benchmark"-tester for benchmarkBSortModified funksjonen basert på eksempel-funksjon benchmarkBSort i filen sorting_test.go


### Oppgave C. Utfør alle benchmark- testene med kommando “go test -bench=.” og presenter resultatene grafisk. Hva kan du si om big-O for alle 3 algoritmene som du har testet?


## 3. Forstå prosessadministrasjon på en platform. 

### Skriv et program som består av en evig løkke. 
Se loop.go i mappen "Oblig1". Den evige løkken skriver ut "venter på signal..." helt til den mottar et signal. 

### Hvor mye minne og CPU bruker programmet når det kjører. 
På MAC - Som bildet viser øker både CPU og minne betraktelig når programmet kjøres i terminal. 

-- (legg inn bilder + forsøk på flere enheter) --

### Generer ulike avslutningssignaler til prosessen og dokumenter hvilke avslutningskommandoer programmet håndterer og som trigger avslutningsmeldingen.

INT signal ("interrupt", SIGINT) håndteres av programmet fordi vi har spesifikt programmert det til å gjøre det. Denne trigger også avslutningsmeldingen. 

TSTP signal ("terminal stop", SIGTSTP) håndteres av programmet. Selv om dette ikke er programmert inn får signalet terminalen til å avslutte. Den genererer derimot ikke avlsutningsmeldingen i etterkant. 

QUIT signal (SIGQUIT) håndteres ikke av programmet, og skriver heller ikke ut avslutningsmelding. 


## 4. Typografiske symboler.

### Oppgave A. Bruk filen ascii.go i Oblig1 mappen og lag en funksjon som itererer (går i en løkke over)  over tegn med byte-verdier fra 0x80 til 0xFF, dvs. det utvidede ASCII settet. 


### Oppgave B. Lag funksjonen ExtendedASCIIText () i samme filen iso.go, som skriver ut: " € ÷ ¾ dollar "


### Oppgave C. Implementer en test for funksjonen ExtendedASCIIText(String) i egen fil iso_test.go, som tester om input-verdier (av type string) inneholder kun tegn fra en Extended ASCII.


