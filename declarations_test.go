package gopkg_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/neurotempest/gopkg"
	"github.com/neurotempest/gopkg/tmpl"
)

func TestDeclFunc_RequiredImports(t *testing.T) {

	testCases := []struct{
		Name string
		F gopkg.DeclFunc
		Expected map[string]bool
	}{
		{
			Name: "empty func",
		},
		{
			Name: "composite args and return",
			F: gopkg.DeclFunc{
				Args: []gopkg.DeclVar{
					{
						Type: gopkg.TypeUnknownNamed{
							Import: "import/1",
						},
					},
					{
						Type: gopkg.TypeUnknownNamed{
							Import: "import/2",
						},
					},
				},
				ReturnArgs: tmpl.UnnamedReturnArgs(
					gopkg.TypeUnknownNamed{
						Import: "import/3",
					},
					gopkg.TypeUnknownNamed{
						Import: "import/1",
					},
				),
			},
			Expected: map[string]bool{
				"import/1": true,
				"import/2": true,
				"import/3": true,
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			actual := test.F.RequiredImports()

			if len(test.Expected) == 0 && len(actual) == 0{
				// Allow actual to be either nil or an empty map
				return
			}

			require.Equal(t, test.Expected, test.F.RequiredImports())
		})
	}
}
