## Assignment 1 for Cloud Technologies at NTNU Gjøvik

# Hvordan vi jobber

Master branchen er låst fra å pushe direkte til denne. Jeg har satt det opp slik at vi oppretter separate brancher for hver oppgave, og merger denne inn via pull request - da er det også anledning for å diskutere og kommentere kodeendringer i pull requesten, før de merges inn i masterbranchen. Slik kan vi passe på at vi ikke roter ting til selv om vi jobber parallellt, samt at vi kan gjøre koden bedre gjennom å sjekke hverandre.

Når du skal jobbe med en oppgave eller problem, opprett en lokal branch med:

```git checkout -b branch-navn```

Git oppretter den lokale branchen utfra den branchen man er i (typisk master), og bytter automatisk over til den. Her kan du gjøre endringer, og committe som ellers med:

```git commit -am "oppsummering av endring"```

Commit ofte, så kan man lettere gå tilbake dersom noe har fucka seg. For å stille tilbake repoet til den nyligste commiten (sett at man ikke har committet tabben man har gjort) gjør man:

```git checkout -- .```

Legg merke til punktumet. Må gjøres i root-mappen.


## Legge til nye filer

Skal du legge inn nye filer i prosjektet, er det noen ting å passe på:

1. Legg til filen under `target_sources` i `CMakeLists.txt`, dersom den er en kildefil. Dersom det er et bibliotek som bare inkluderes, som med `ListTool2B`, legges filen i mappen `include`, og filnavnet med relativ sti (e.g. "include/FILNAVN") legges til under `target_sources`. Dersom den nye filen ikke er en kildefil, men f.eks. generert av visual studio, cmake eller annet, ikke tenk på dette steget, men gå videre.
2. Legg til filen i git ved `git add filnavn`
3. Husk å committe endringen, siden `git add` ikke gjør dette selv

Dersom du skulle få problemer med å kompilere programmet etter å ha lagt inn en ny fil, forsøk å gjøre "rebuild all" fra "CMake" menyen i visual studio.


## Opplasting av lokal branch

Det kan være greit å jobbe på samme branch, og da må man laste denne opp til bitbucket. Dette må man også gjøre uansett, når det er klart for å lage et pull-request til master branchen. Da gjør man:

```git push --set-upstream origin branch-navn```

Bruker du git bash, kan man stort sett påbegynne ord, og trykke TAB, og så completer shellet for deg. Etter dette opprettes den samme branchen i bitbucket, og man kan pulle og pushe til denne som ellers med git.

## Hvordan checkout'e en remote branch opprettet av noen andre
```
git fetch
git checkout branch-navn
```

## Pull-request

Når koden i en branch er klar for å inkluderes i master, oppretter man et pull-request, og da bør minst en annen av oss gå over koden, før den inkluderes eller endres videre. Man kan fortsette å pulle og pushe til branchen, og pull-requesten vil være oppdatert. Til slutt merger man pull requesten og sletter branchen (dette vil jeg gjøre ihvertfall i starten, så tar vi evt en seanse på det når jeg kommer på land igjen, slik at alle er gode på det).

Opprettelse av pull-requests kan gjøres via commandovinduet, men jeg er ikke komfortabel med det, og den grafiske løsningen på bitbucket er ganske egnet for det. Når du har lastet opp en lokal branch, finner du den i repoet, under branches, derfra kan man klikke på linken "Create" helt til høyre i tabellen (under "pull request" kollonnen).

Vinduet som kommer opp er ganske selvforklarende. Feltet "reviewers" er hvor man kan assigne en av de andre til å reviewe koden, men jeg tror ikke vi trenger å tenke på å gjøre det så formelt siden vi er bare tre. Default-tittelen er ofte grei, men etterhvert som kodebasen blir litt større kan det være lurt å legge litt arbeid i tittel og beskrivelse her.

### Å sette sitt "stamp of approval" på en pull request

Når du har gått over en annens pull request, og er fornøyd, klikker du knappen "approve" øverst til høyre for tittelen, for å vise at du har sett over og syns den er klar for å gå inn i master. (obs: jeg er usikker på hvordan denne tilgangen er, men tror alle skal kunne klikke på denne. Si fra om ikke)

## Ellers

Dersom det ellers oppstår problemer, eller kommandoer vi har bruk for, som ikke allerede er dekket her, endrer vi dette dokumentet.
