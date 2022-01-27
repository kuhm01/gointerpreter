package evaluator

import (
	"bufio"
	"fmt"
	"monkey/object"
	"os"
	"strconv"
)

/*
All variable that written by book is
"function name": &object.Builtin{
	Fn: ~~~
	~~~
}
*/

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
` //Monkey Face is Copy from the code provided book

var functionName = []string{"len", "first", "last", "rest", "push", "puts",
	"pop", "integer", "float", "range", "Itoa", "typeOf", "gets"}

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to len not supported, got %s", args[0].Type())
			}
		},
	},

	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to first must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)

			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},

	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to last must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},

	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to rest must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if length > 0 {
				newElements := make([]object.Object, length-1) //By Book. make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},

	"push": {
		Fn: func(args ...object.Object) object.Object {
			switch len(args) {
			case 2:
				if args[0].Type() != object.ARRAY_OBJ {
					return newError("argument to push must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*object.Array)
				length := len(arr.Elements)

				newElements := make([]object.Object, length+1) //By Book. make([]object.Object, length+1, length+1)
				copy(newElements, arr.Elements)
				newElements[length] = args[1]

				return &object.Array{Elements: newElements}

			case 3:
				if args[0].Type() != object.ARRAY_OBJ || args[2].Type() != object.INTEGER_OBJ {
					return newError("argument to push must be ARRAY and INTEGER, got %s %s", args[0].Type(), args[2].Type())
				}

				arr := args[0].(*object.Array).Elements
				length := len(arr)
				index := args[2].(*object.Integer).Value

				if index < 0 || int(index) > length-1 {
					return newError("Wrong index. that over array scale, got %d", index)
				}

				temp := make([]object.Object, length)
				copy(temp, arr)

				newElements := make([]object.Object, length+1)

				temp1 := temp[:index]
				temp2 := temp[index:]

				for i, v := range temp1 {
					newElements[i] = v
				}

				newElements[index] = args[1]
				tempindex := index + 1

				for _, v := range temp2 {
					newElements[tempindex] = v
					tempindex++
				}

				return &object.Array{Elements: newElements}

			default:
				return newError("wrong number of arguments. got=%d, want=1, 2", len(args))
			}
		},
	},

	"puts": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},

	//실습자 정의 함수
	//신뢰도 : 몰?루
	"pop": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to pop must be ARRAY, got %s", args[0].Type())
			}

			if args[1].Type() != object.INTEGER_OBJ {
				return newError("argument to pop must be INTEGER, got %s", args[1].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if length == 0 {
				return newError("array is empty, length %d, got %s", length, arr.Elements)
			}

			userIndex := args[1].(*object.Integer).Value

			if 0 > userIndex || userIndex > int64(length-1) {
				return newError("argument is beyond the scope of an array, got %d", userIndex)
			}

			newElements := make([]object.Object, length-1)
			i := 0
			j := 0
			for i < length-1 {
				if j == int(userIndex) {
					j++
					newElements[i] = arr.Elements[j]
					i++
					j++
				} else {
					newElements[i] = arr.Elements[j]
					i++
					j++
				}
			}

			return &object.Array{Elements: newElements}
		},
	},

	"monkey": {
		Fn: func(args ...object.Object) object.Object {
			return &object.String{Value: MONKEY_FACE}
		},
	},

	"integer": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.FLOAT_OBJ {
				return newError("argument to must be FLOAT, got %s", args[0].Type())
			}

			value := args[0].(*object.Float).Value

			return &object.Integer{Value: int64(value)}
		},
	},

	"float": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.INTEGER_OBJ {
				return newError("argument to must be INTEGER, got %s", args[0].Type())
			}

			value := args[0].(*object.Integer).Value

			return &object.Float{Value: float64(value)}
		},
	},

	"range": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.INTEGER_OBJ {
				return newError("argument to must be INTEGER, got %s", args[0].Type())
			}

			value := args[0].(*object.Integer).Value

			arr := make([]object.Object, value)

			for i := 0; i < int(value); i++ {
				arr[i] = &object.Integer{Value: int64(i)}
			}

			return &object.Array{Elements: arr}
		},
	},

	"Itoa": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.INTEGER_OBJ {
				return newError("argument to must be INTEGER, got %s", args[0].Type())
			}

			str := strconv.Itoa(int(args[0].(*object.Integer).Value))

			return &object.String{Value: str}
		},
	},

	"typeOf": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			return &object.String{Value: string(args[0].Type())}
		},
	},

	"gets": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return newError("gets function should be no parameters. got=%d", len(args))
			}

			fmt.Fprintf(os.Stdout, "<< ")

			scanner := bufio.NewScanner(os.Stdin)
			scanned := scanner.Scan()
			if !scanned {
				return newError("Critical Scan Error")
			}

			line := scanner.Text()

			return &object.String{Value: line}
		},
	},
}
