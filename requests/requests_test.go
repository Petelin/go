package requests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/thedevsaddam/gojsonq"

	"github.com/stretchr/testify/assert"
)

func TestM(t *testing.T) {
	r, _ := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/users", nil)
	jq, err := Do(r)
	assert.NoError(t, err, "should no error")
	fmt.Println(jq("[0].name"))
}

func TestJ(t *testing.T) {
	//const json = `{"city":"dhaka","type":"weekly","temperatures":[30,39.9,35.4,33.5,31.6,33.2,30.7]}`
	const json = `[30,39.9,35.4,33.5,31.6,33.2,30.7]`
	avg := gojsonq.New().JSONString(json).Find("[2]")
	fmt.Println(avg) // 33.471428571428575
}
