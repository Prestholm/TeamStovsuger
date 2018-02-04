# TeamStovsuger

Medlemmer: 
- Henning Van Ta.
- Victor Prestholm.
- Arian Nuland. 


Oppgavesett 1 - Oppgave 3:
----------------------------

**1. Finn ut hva  objektfiler heter for de mest brukte platformene (Unix/Linux, MS Windows, Mac OS X). Hvorfor, etter din mening, har disse platformene så forskjellige objektfil-formater?**

  - Unix/Linux = ELF (Executable Linkable Format).
  - MS Windows = PE (Portable Executable). 
  - Mac OS X = Mach-O.

    Disse platformene har så forskjellige objektfil-formater fordi hver av operativsystemene fungerer på fundamentalt forskjellige måter. 
    Derfor må de fleste filer være spesifikt forskjellige fra hverandre for å være brukbare for hvert enkelt system.


**2. Gå gjennom https://tour.golang.org/basics/1 - Hvilke forskjeller ser du i forhold til programmeringsspråket Java?** 
   
   - Go har ikke private og public. 
   - I Go bruker man pakker istedenfor klasser.
   - I Go importerer vi i parantes.
   - Exporterte navn starter med stor bokstav , f.eks. "Pi" er exportert fra "math" pakken. Hvis man bruker "pi" bruker man ikke
     exporterte navn, men heller fra den interne pakken.
   - Metoder er mer eller mindre det samme. Men det er litt friere med hvordan man definerer int parametere f. eks. func metode(x, y int) 
     mens i java ville man heller ha skrevet (int x, int y). Metodene kan også returnere flere resultat i Go. I Go kan man returnere 
     "nakne" resultater, men dette vil bare fungere i små metoder, siden det kan hemme hvordan man leser verdiene som er utgitt i lengre 
     metoder.
   - For å forklare variabler må man bruke "var" før vi erklarer hva variabel vi bruker. F.eks. "var x, y int" eller "var hei bool". 
     Inne i f.eks en Main-metode kan man bruke ":=" for å forkorte bruken av "var" - "i := 3" eller "hei := false". Erklærte variabler som 
     ikke er blitt gitt en spesifikk verdi er null eller false. 
  
