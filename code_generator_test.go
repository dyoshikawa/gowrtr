package gowrtr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateCode(t *testing.T) {
	expected := `// THIS CODE WAS AUTO GENERATED

package mypkg

import (
	"fmt"
)

type MyInterface interface {
	MyFunc(foo string) (string, error)
}

type MyStruct struct {
	Foo string
	Bar int64
}

func (m *MyStruct) MyFunc(foo string) (string, error) {
	{
		str := 		func(bar string) string {
			return bar
		}(foo)

		if str == "" {
			for i := 0; i < 3; i++ {
				fmt.Printf("%d\n", i)
			}
		}
		return str, nil
	}
}
`
	myFuncSignature := NewFuncSignatureGenerator("MyFunc").
		AddFuncParameters(
			NewFuncParameter("foo", "string"),
		).
		AddReturnTypes("string", "error")

	generator := NewCodeGenerator(
		NewCommentGenerator(" THIS CODE WAS AUTO GENERATED"),
		NewNewlineGenerator(),
		NewPackageGenerator("mypkg"),
		NewNewlineGenerator(),
		NewImportGenerator("fmt"),
		NewNewlineGenerator(),
		NewInterfaceGenerator("MyInterface").
			AddFuncSignature(myFuncSignature),
		NewNewlineGenerator(),
		NewStructGenerator("MyStruct").
			AddField("Foo", "string").
			AddField("Bar", "int64"),
		NewNewlineGenerator(),
		NewFuncGenerator(
			NewFuncReceiverGenerator("m", "*MyStruct"),
			NewFuncSignatureGenerator("MyFunc").
				AddFuncParameters(
					NewFuncParameter("foo", "string"),
				).
				AddReturnTypes("string", "error"),
		).AddStatements(
			NewCodeBlockGenerator(
				NewRawStatementGenerator("str := "),
				NewInlineFuncGenerator(
					false,
					NewInlineFuncSignatureGenerator().
						AddFuncParameters(NewFuncParameter("bar", "string")).
						AddReturnTypes("string"),
					NewReturnStatementGenerator("bar"),
				).AddFuncInvocation(NewFuncInvocationGenerator("foo")),
				NewNewlineGenerator(),
				NewIfGenerator(`str == ""`).
					AddStatements(
						NewForGenerator(`i := 0; i < 3; i++`).AddStatements(
							NewRawStatementGenerator(`fmt.Printf("%d\n", i)`, true),
						),
					),
				NewReturnStatementGenerator("str", "nil"),
			),
		),
	)

	generated, err := generator.Generate(0)

	assert.NoError(t, err)
	assert.Equal(t, expected, generated)
}

func TestShouldGenerateCodeWithGofmt(t *testing.T) {
	expected := `// THIS CODE WAS AUTO GENERATED

package mypkg

import (
	"fmt"
)

type MyInterface interface {
	MyFunc(foo string) (string, error)
}

type MyStruct struct {
	Foo string
	Bar int64
}

func (m *MyStruct) MyFunc(foo string) (string, error) {
	{
		str := func(bar string) string {
			return bar
		}(foo)

		if str == "" {
			for i := 0; i < 3; i++ {
				fmt.Printf("%d\n", i)
			}
		}
		return str, nil
	}
}
`
	myFuncSignature := NewFuncSignatureGenerator("MyFunc").
		AddFuncParameters(
			NewFuncParameter("foo", "string"),
		).
		AddReturnTypes("string", "error")

	generator := NewCodeGenerator(
		NewCommentGenerator(" THIS CODE WAS AUTO GENERATED"),
		NewNewlineGenerator(),
		NewPackageGenerator("mypkg"),
		NewNewlineGenerator(),
		NewImportGenerator("fmt"),
		NewNewlineGenerator(),
		NewInterfaceGenerator("MyInterface").
			AddFuncSignature(myFuncSignature),
		NewNewlineGenerator(),
		NewStructGenerator("MyStruct").
			AddField("Foo", "string").
			AddField("Bar", "int64"),
		NewNewlineGenerator(),
		NewFuncGenerator(
			NewFuncReceiverGenerator("m", "*MyStruct"),
			NewFuncSignatureGenerator("MyFunc").
				AddFuncParameters(
					NewFuncParameter("foo", "string"),
				).
				AddReturnTypes("string", "error"),
		).AddStatements(
			NewCodeBlockGenerator(
				NewRawStatementGenerator("str := "),
				NewInlineFuncGenerator(
					false,
					NewInlineFuncSignatureGenerator().
						AddFuncParameters(NewFuncParameter("bar", "string")).
						AddReturnTypes("string"),
					NewReturnStatementGenerator("bar"),
				).AddFuncInvocation(NewFuncInvocationGenerator("foo")),
				NewNewlineGenerator(),
				NewIfGenerator(`str == ""`).
					AddStatements(
						NewForGenerator(`i := 0; i < 3; i++`).AddStatements(
							NewRawStatementGenerator(`fmt.Printf("%d\n", i)`, true),
						),
					),
				NewReturnStatementGenerator("str", "nil"),
			),
		),
	).EnableGofmt("-s")

	generated, err := generator.Generate(0)

	assert.NoError(t, err)
	assert.Equal(t, expected, generated)
}

func TestShouldGenerateCodeWithGoimport(t *testing.T) {
	expected := `// THIS CODE WAS AUTO GENERATED

package mypkg

import "fmt"

type MyInterface interface {
	MyFunc(foo string) (string, error)
}

type MyStruct struct {
	Foo string
	Bar int64
}

func (m *MyStruct) MyFunc(foo string) (string, error) {
	{
		str := func(bar string) string {
			return bar
		}(foo)

		if str == "" {
			for i := 0; i < 3; i++ {
				fmt.Printf("%d\n", i)
			}
		}
		return str, nil
	}
}
`
	myFuncSignature := NewFuncSignatureGenerator("MyFunc").
		AddFuncParameters(
			NewFuncParameter("foo", "string"),
		).
		AddReturnTypes("string", "error")

	generator := NewCodeGenerator(
		NewCommentGenerator(" THIS CODE WAS AUTO GENERATED"),
		NewNewlineGenerator(),
		NewPackageGenerator("mypkg"),
		NewInterfaceGenerator("MyInterface").
			AddFuncSignature(myFuncSignature),
		NewNewlineGenerator(),
		NewStructGenerator("MyStruct").
			AddField("Foo", "string").
			AddField("Bar", "int64"),
		NewNewlineGenerator(),
		NewFuncGenerator(
			NewFuncReceiverGenerator("m", "*MyStruct"),
			NewFuncSignatureGenerator("MyFunc").
				AddFuncParameters(
					NewFuncParameter("foo", "string"),
				).
				AddReturnTypes("string", "error"),
		).AddStatements(
			NewCodeBlockGenerator(
				NewRawStatementGenerator("str := "),
				NewInlineFuncGenerator(
					false,
					NewInlineFuncSignatureGenerator().
						AddFuncParameters(NewFuncParameter("bar", "string")).
						AddReturnTypes("string"),
					NewReturnStatementGenerator("bar"),
				).AddFuncInvocation(NewFuncInvocationGenerator("foo")),
				NewNewlineGenerator(),
				NewIfGenerator(`str == ""`).
					AddStatements(
						NewForGenerator(`i := 0; i < 3; i++`).AddStatements(
							NewRawStatementGenerator(`fmt.Printf("%d\n", i)`, true),
						),
					),
				NewReturnStatementGenerator("str", "nil"),
			),
		),
	).EnableGoimports()

	generated, err := generator.Generate(0)

	assert.NoError(t, err)
	assert.Equal(t, expected, generated)
}
