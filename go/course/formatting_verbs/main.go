package main

import "fmt"

/*
В Go пакете fmt есть много так называемых "глаголов форматирования" (formatting verbs), которые используются с
функциями типа Printf или Sprintf, чтобы указать, как должен быть отформатирован тот или иной тип данных.

Вот основные из них, сгруппированные по типам данных:
	%v (value): 	Выводит значение в формате по умолчанию. Это универсальный глагол, если вы не уверены, какой конкретный тип данных перед вами.
	%+v: 			То же, что и %v, но для структур включает имена полей.
	%#v: 			Выводит значение в синтаксисе Go (как если бы вы его объявили в коде).
	%T: 			Выводит тип значения.
	%%: 			Выводит сам символ процента (%). Не принимает аргументов.
	%t: 			Выводит true или false.
*/

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"Alice", 30}
	pi := 3.14159
	num := 15

	fmt.Printf("Default value: %v\n", p)
	fmt.Printf("Value with field names: %+v\n", p)
	fmt.Printf("Go syntax representation: %#v\n", p)
	fmt.Printf("Type of p: %T\n", p)
	fmt.Printf("Boolean: %t\n", true)
	fmt.Printf("Decimal: %d\n", num)
	fmt.Printf("Binary: %b\n", num)
	fmt.Printf("Hex (lowercase): %x\n", num) //	%x: Шестнадцатеричный дамп байтов строки (строчные буквы)
	fmt.Printf("Hex (uppercase): %X\n", num) //	%X: Шестнадцатеричный дамп байтов строки (заглавные буквы)
	fmt.Printf("Float (default): %f\n", pi)
	fmt.Printf("Float (2 decimal places): %.2f\n", pi)
	fmt.Printf("String: %s\n", "hello")
	fmt.Printf("Quoted string: %q\n", "hello world") //	%q: Строка в кавычках, как в Go, с экранированием специальных символов (например, "hello\nworld")
	fmt.Printf("Pointer: %p\n", &p)
	fmt.Printf("Literal percent sign: %%\n")
}
