# is105-ica04

Navnliste: Abdikani Gureye, Brede Knutsen Meli, Eirik Aanestad Fintland, Jan Kevin Henriksen, Mats Skjærvik, Mikael Kimerud, Morten Johan Mygland, Nils Fredrik Iselvmo Bjerk.

- I denne oppgaven har alle bidratt. Ikke alle jobber på hver sin maskin, vi jobber av og til bare på noen av pcene og det er da de som laster opp det som er gjort. Derfor kan det vise at det er bare 5/8 som har lastet opp noe. 

## Oppgave 1
## a)

(Bildet til venstre)
54 65 73 74 65 72 20 6C 69 6E 6A 65 73 6B 69 66 74 2E 0A 4F 67 20 65 6E 20 74 69 6C 20 2E 2E 2E 0A 4F 67 20 65 6E 20 74 69 6C 20 2E 2E 2E 0A 4F 67 20 65 6E 20 74 69 6C 20 2E 2E 2E 0A 
[T e s t e r   l i n j e s k i f t . 
 O g   e n   t i l   . . . 
 O g   e n   t i l   . . . 
 O g   e n   t i l   . . . 
 
( Bildet til høyre)
54 65 73 74 65 72 20 6C 69 6E 6A 65 73 6B 69 66 74 2E 0D 0A 4F 67 20 65 6E 20 74 69 6C 20 2E 2E 2E 0D 0A 4F 67 20 65 6E 20 74 69 6C 20 2E 2E 2E 0D 0A 4F 67 20 65 6E 20 74 69 6C 20 2E 2E 2E 0D 0A 
[T e s t   r   l i n j e s k i f t . 
 O g   e n   t i l   . . . 
 O g   e n   t i l   . . . 
 O g   e n   t i l   . . . 
 
 
Operativsystem har ingenting å si på hvordan en fil blir vist på skjermen. Det er applikasjonene. Hvis man lager noe i notepad, kan de endre line shift etc. Windows systemer kan legge på character/carriage return.
Hvis man skriver ut byte for byte kan man se hva som egentlig skjuler seg i teksten.


Som vi kan se i text1.txt er byten 0D lagt til 4 ganger, noe som forklarer hvorfor den ene filen er 4 bytes større enn den andre.. 0D er carriage return, 0A er linjeskift. Carriage return resetter enheten til å begynne på starten av ny linje. 


## b)
Koden kjøres med en innlagt fil (asdf.txt) i programmet.
```
go run main.go
```

Linjeskift m/ punktum, mac: Vi ser tegnet 0A  fordi tekstfiler opprettet i unix og nyere mac versjoner kun bruker line feed, mens de i de første mac versjonene kun brukte  carriage return (0D) i tillegg til 0A.

Bilde for linjeskift i filen på mac: 


Bilde for linjeskift windows: Dette er samme fil som i illustrasjonen over. Siden filen er opprettet i windows sin notepad, så vil tegnene 0D og 0A dukke opp når en trykker på enter tasten. 0A står for line feed, mens 0D står for carriage return. Vi har brukt en funksjon for å telle antall “\x0A” i filen, som vi får ut som tekst. Dette viser oss at det er 4 linjeskift i teksten.









Vi ble ikke helt ferdige med oppgave 4e, så den vedlagte koden er uferdig.
