# ДЗ №2

## Задание 1

Даны файлы [one.go](./one.go) и [one_test.go](./one_test.go). В первом файле находятся прототипы для двух функций: TwoSum и Equal.

Первая функция принимает массив целых чисел и цель, а возвращает массив, состоящий из двух чисел. Необходимо реализовать функцию так, чтобы она возвращала такие два индекса массива-аргумента, что элементы массива по этим индексам в сумме давали цель.

Например:

```go
nums := []int{2, 7, 11, 15}
target := 9

result := TwoSum(nums, target) // return []int{0, 1}
```

, потому что `num[0] + num[1] == 9`

Вторая функция сравнивает два переданных в неё массива и проверяет, равны ли они (порядок важен). Если равны, то возвращается `true`, иначе `false`

Проверить реализацию функций можно с помощью команды в консоли:

```shell
go test
```

Тесты реализованы в файле [one_test.go](./one_test.go) (если люботно, посмотри)

Функция Equal используется в тестах для двух заданий, поэтому будь внимательным при её реализации)

## Задание 2

Даны файлы [two.go](./two.go) и [two_test.go](./two_test.go). В первом файле находятся прототип функции Intersection. Эта функция должна принимать слайсы целочисленных значений, а возвращать множество их пересечений.

Например:

Например:

```go
a := []int{23, 3, 1, 2}
b := []int{6, 2, 4, 23}

result := Intersection(a, b) // return []int{2, 23}
```

Для проверки воспользуйся тестами, реализованными в файле [two_test.go](./two_test.go), с помощью команды:

```shell
go test
```

Вот так вот мы с тобой узнали о том, как легко можно тестировать код в Go)
