# umka tutorial

Hello! Welcome to my umka tutorial. Even though umka already has a tour available in the readme, it is pretty short. This is largely inpired by [tour of go](https://tour.golang.org/welcome/1).

## step 1: hello world

Hello world in umka is not much different from other languages. Here it is:

```
fn main() {
	printf("Hello world!\n")
}
```

We can notice, that there isn't need to import any packages (like in go), since `printf` is a builtin function. Another noteworthy thing is, that we need to include `\n` since `printf` in umka behaves the same as in c.

## step 2: functions

If you look at step one, you can already see two occurences of a function.
First is `fn main`. That is a function declaration. The second is the print statement.
That is a function call. Functions can take arguments and return values. There isn't a reachable limit to the number of arguments/return values, but their count and type has to be specified.

```
fn add_and_sub(x, y: int): (int, int) {
	return x + y, x - y
}

fn main() {
	printf("%d, %d\n", add_and_sub(4, 6))
}
```

In this example, you can see a function with name `add_and_sub`, that takes int x and y as arguments and returns x + y and x - y. If multiple argument after each other have the same type, you can specify it only once. If a function returns multiple values, you have to enclose them in (). Otherwise there can be only a value.

## variables

Variables in umka are statically typed. There are two ways of defining them.

```
fn main() {
	var i: int = 2

	var j: int
	j = 2

	k := 2
}
```

All these three do the same thing, but the last one is preffered, since it is shorter.

You can also group variable together.

```
var i, j, k: int
```

## types

- ints
```
int8  int16  int32  int
uint8 uint16 uint32 uint
```
As you can see, int is 8 bytes by default in umka.

- reals
Reals can hold decimal numbers.
```
real32 real
```

- string
```
char str
```

`str` is a string as a whole. If you want to index string, you get chars.

- bool

Same as everywhere else

- interfaces, structs

Structs and interfaces have to be user defined. We will touch on them later.

## 
