# Dojo DevMT S01E02 (Versão em Go Lang) [![Build Status](https://travis-ci.org/alvarowolfx/dojo-s01e02-escrevendo-no-celular.go.svg?branch=master)](https://travis-ci.org/alvarowolfx/dojo-s01e02-escrevendo-no-celular.go)
#### Versão original desenvolvida em Clojure. [Disponivel aqui](https://github.com/devmatogrosso/dojo-s01e02-escrevendo-no-celular)

## Data e Local
Data: 04/07/2015
Local: Nuvem Tecnologia

## Problema
Um dos serviços mais utilizados pelos usuários de aparelhos celulares são os SMS (Short Message Service), que permite o envio de mensagens curtas (até 255 caracteres em redes GSM e 160 caracteres em redes CDMA).

Para digitar uma mensagem em um aparelho que não possui um teclado QWERTY embutido é necessário fazer algumas combinações das 10 teclas numéricas do aparelho para conseguir digitar. Cada número é associado a um conjunto de letras como a seguir:

	Letras  ->  Número
	ABC    ->  2
	DEF    ->  3
	GHI    ->  4
	JKL    ->  5
	MNO    ->  6
	PQRS    ->  7
	TUV    ->  8
	WXYZ   ->  9
	Espaço -> 0

Desenvolva um programa que, dada uma mensagem de texto limitada a 255 caracteres, retorne a seqüência de números que precisa ser digitada. Uma pausa, para ser possível obter duas letras referenciadas pelo mesmo número, deve ser indicada como _.

Por exemplo, para digitar "SEMPRE ACESSO O DOJOPUZZLES", você precisa digitar:

	77773367_7773302_222337777_777766606660366656667889999_9999555337777

[Link do problema](http://dojopuzzles.com/problemas/exibe/escrevendo-no-celular/)

### Alguns pré-requisitos

Este projeto asume que você está trabalhando com um -workspace-- padrão para Go como descrito em [Golang Code](http://golang.org/doc/code.html). Para este projeto foi utilizado Go 1.3.3, porém acredito que qualquer versão acima de 1.1 vá funcionar.

### Resolver dependências do projeto 

```shell
go get github.com/smartystreets/goconvey/convey
```

### Como rodar

```shell
goconvey
```
e acessar [locahost:8080](http://localhost:8080)
