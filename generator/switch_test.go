package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"
	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateSwitch(t *testing.T) {
	generator := NewSwitch("foo")

	{
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		expected := `switch foo {
}
`
		assert.Equal(t, expected, gen)
	}

	generator = generator.AddCaseStatements(
		NewCase("1", NewComment(" one")),
		nil,
		NewCase("2", NewComment(" two")),
	).SetDefaultStatement(
		NewDefaultCase(NewComment(" default")),
	)

	{
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		expected := `switch foo {
case 1:
	// one
case 2:
	// two
default:
	// default
}
`
		assert.Equal(t, expected, gen)
	}

	{
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		expected := `		switch foo {
		case 1:
			// one
		case 2:
			// two
		default:
			// default
		}
`
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateSwitchRaisesErrorWhenCaseRaisesError(t *testing.T) {
	generator := NewSwitch("foo").AddCaseStatements(
		NewCase(""),
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.CaseConditionIsEmptyError().Error())
}

func TestShouldGenerateSwitchRaisesErrorWhenDefaultRaisesError(t *testing.T) {
	generator := NewSwitch("foo").SetDefaultStatement(
		NewDefaultCase(
			NewFunc(nil, NewFuncSignature("")),
		),
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}