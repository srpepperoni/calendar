# Calendar

Simple calendar generator in Golang

Nostalgia for the past, maybe my first "project" in high school

My first promgram language was turbo Pascal at High School, and my first medium project (huge for me at that moment) was a calendar generator. So I thouht to create similar project with golang for the feels making some similarities with the project.

## Console mode

This code is very straighforward. It shows a simple three options menu like following one:

```bash
Calendar
========
Choose your option:
1) Print Year
2) Print Month
3) Print Day of year
```

And after that, pick the options by standard input.
It will make some validation for different parameters and return result in console.

```bash
--------------------
       JANUARY
--------------------
 S  M  T  W  T  F  S
--------------------
                   1
 2  3  4  5  6  7  8
 9 10 11 12 13 14 15
16 17 18 19 20 21 22
23 24 25 26 27 28 29
30 31
```

## Compile and exec

```bash
go build .\cmd\main\. # From root project directory
.\main.exe
```