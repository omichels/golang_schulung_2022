# Agenda

## Go - Allgemein

- Was ist Go? Oder die Kunst des Weglassens.

## I - Basics (Teil 1)

- Pakete / Module / Programmstruktur, Kommentare
- Funktionen und Aufrufe: `func`, Call by value
    - init-Funktionen
- Kontrollstrukturen und Syntax:
    - "Oberon mit C-Syntax" [Evolution of Go (Robert Griesemer)](https://talks.golang.org/2015/gophercon-goevolution.slide)
    - `if`, `for`, `switch`, `goto`
- Operatoren
- Datenstrukturen:
    - Variablen `var`:
      - einfache: `int`, `float32`, `rune`, ...
          - [Numeric types](https://golang.org/ref/spec#Numeric_types)
      - zusammengesetzte:
          - Structs und Arrays
          - struct-Tags
      - gebundene:
          - Slices
              - [SliceTricks](https://github.com/golang/go/wiki/SliceTricks)
          - Maps
              - [Presentation: Inside the Map Implementation (Keith Randall)](https://docs.google.com/presentation/d/1CxamWsvHReswNZc7N2HMV7WPFqS8pvlPVZcDegdC_T4/edit#slide=id.p)
              - [Video: Inside the Map Implementation (Keith Randall)](https://www.youtube.com/watch?v=Tl7mi9QmLns)
        - Funktionen II / Closures / [Variadische Funktionen](https://de.wikipedia.org/wiki/Variadische_Funktion)
            - Bsp. für variadische Funktion [func Printf aus dem fmt Package](https://golang.org/pkg/fmt/#Printf)
      - Kontrollstrukturen II: `range`
      - Pointer
- Builtin-Funktionen: `append`, `copy`, `len`, `cap`, ...
- `defer`, `panic`, `recover`
    - [Beispielaufgabe](https://gitlab.com/sascha.l.teichmann/go-examples/-/blob/master/methodexp/main.go)
- `return`: Tuple und Fehlerbehandlung
    - [Errors are values (Rob Pike)](https://blog.golang.org/errors-are-values)
- `string`
    - Hinweise auf UTF8-Encoding
    - siehe auch [package unicode/utf8](https://golang.org/pkg/unicode/utf8/)

- Konstanten `const`: typsiert, untypisiert
    - [Bsp. time.go in package time](https://golang.org/pkg/time/#Duration)
    - [Iota](https://golang.org/ref/spec#Iota)


## II - Typsystem

- Typen
- Methoden, Method-Sets
- Reciever, Pointer-Reciever
- Interfaces, implizites Erfüllen
    - [implizites Erfüllen am Beispiel type Stringer im package fmt](https://golang.org/pkg/fmt/#Stringer)
- Embedding (Exkurs Objektorientierung mit anderen Mitteln)
    - Komposit-Gedanke
    - Vererbung mittels Embedding
- Kontrollstrukturen III:
    - Typ-Zusicherungen mittels Type-Assertions
    - type switch
      - [Beispiel aus dem Package fmt](https://golang.org/src/fmt/print.go#L623)
- Funktionen III:
    - Method-Values
      - [Bsp. intests/controller.go](https://bitbucket.org/s_l_teichmann/intests/src/9e39252ea071feaa101c570f66aee6502caa65bf/controller.go?at=default&fileviewer=file-view-default#controller.go-42)
- Method-Expressions
    - [Bsp. go-examples/methodexp/main.go](https://gitlab.com/sascha.l.teichmann/go-examples/-/blob/master/methodexp/main.go)
- Generics
    - Type Parameters für Funktionen und Typen
    - Type Sets durch Interfaces
    - Type Interference


## III - Concurrency (Nebenläufigkeit)

- Einleitung
    - [Communicating Sequential Processes (CSP) von Tony Hoare](https://de.wikipedia.org/wiki/Communicating_Sequential_Processes)
    - [Concurrency is not parallelism (Rob Pike)](https://blog.golang.org/concurrency-is-not-parallelism)
    - [C10k problem](https://en.wikipedia.org/wiki/C10k_problem)
- Go-Routines `go f()`
- Channels `chan`
    - [Effective Go / channels](https://golang.org/doc/effective_go.html#channels)
- Kontrollstrukturen IV:
    - `range`, `<-`
    - `select`
- Orchestrierend vs. serialisierend (`sync`-Paket)
    - Orchestrierend
        - Channels `chan`
    - serialisierend
        - waitGroup
        - Mutex
        - atomic

## IV - Standard-Bibliothek (optional)

- `io`, `fmt`, `os`, ...

## V: Bauen (optional)

- Build Tags
- CGo: Interaktion mit C-Bibliotheken
- Cross-Compiling
- Vendoring

## VI: Meta-Programming (optional)

- Reflection: Types, Values (Hinweis: Konzept der Pointer in Go muss verstanden sein!!!)
- Generierung von Code: `go generate`
- interne Data-Repräsentierung
- `unsafe`
