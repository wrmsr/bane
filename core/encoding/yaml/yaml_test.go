package yaml

import (
	"fmt"
	"testing"

	"gopkg.in/yaml.v3"

	tu "github.com/wrmsr/bane/core/dev/testing"
)

func TestNodes(t *testing.T) {
	type Person struct {
		Name    yaml.Node
		Address yaml.Node
	}

	data := `
name: John Doe
#comment here
address: 
    street: 123 E 3rd St
    city: Denver
    state: CO
    zip: 81526
`

	var person Person
	tu.AssertNoErr(t, yaml.Unmarshal([]byte(data), &person))
	fmt.Printf("%v", person)

	address := &yaml.Node{}
	tu.AssertNoErr(t, person.Address.Decode(address))
	fmt.Printf("%v", address.HeadComment)

	var root yaml.Node
	tu.AssertNoErr(t, yaml.Unmarshal([]byte(data), &root))
	fmt.Printf("%v", root)

	buf, err := yaml.Marshal(root)
	tu.AssertNoErr(t, err)
	fmt.Printf("%s", string(buf))
}
