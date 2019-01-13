package generator

import (
	"fmt"
	"log"
)

func ExampleRootGenerator_Generate() {
	myFuncSignature := NewFuncSignatureGenerator("MyFunc").
		AddFuncParameters(
			NewFuncParameter("foo", "string"),
		).
		AddReturnTypes("string", "error")

	generator := NewRootGenerator(
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
	).AddStatements(
		NewFuncGenerator(
			NewFuncReceiverGenerator("m", "*MyStruct"),
			NewFuncSignatureGenerator("MyFunc").
				AddFuncParameters(
					NewFuncParameter("foo", "string"),
				).
				AddReturnTypes("string", "error"),
		).AddStatements(
			NewCodeBlockGenerator(
				NewRawStatementGenerator("str := ", false),
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
				NewNewlineGenerator(),
				NewSwitchGenerator("str").
					AddCaseStatements(
						NewCaseGenerator(
							`""`,
							NewCommentGenerator(" empty string"),
						),
						NewCaseGenerator(
							`"foo"`,
							NewCommentGenerator(" foo string"),
						),
					).
					SetDefaultStatement(
						NewDefaultCaseGenerator(NewCommentGenerator(" default")),
					),
				NewNewlineGenerator(),
				NewReturnStatementGenerator("str", "nil"),
			),
		),
	)

	generated, err := generator.
		EnableGofmt("-s").
		EnableGoimports().
		Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
