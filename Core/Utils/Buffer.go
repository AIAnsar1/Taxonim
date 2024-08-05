package Utils

import (
	"os"
	"strings"
)

/*
|========================================================
|	StringInSlice() -> boolean
|========================================================
|
| Проверяет, содержится ли строка A в срезе строк List.
| Возвращает true, если строка найдена в срезе, иначе false.
|
*/
func StringInSlice(A string, List []string) bool {
	for _, B := range List {
		if B == A {
			return true
		}
	}
	return false
}

/*
|========================================================
|	IsSystemInTestMode() -> boolean
|========================================================
|
| Проверяет, запущена ли система в тестовом режиме
| анализируя аргументы командной строки.
| Возвращает true, если найдется аргумент,
| начинающийся с префикса "-test.", иначе false.
|
*/
func IsSystemInTestMode() bool {
	for _, Arg := range os.Args {
		if strings.HasPrefix(Arg, "-test.") {
			return true
		}
	}
	return false
}
