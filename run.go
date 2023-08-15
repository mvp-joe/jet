package jet

import (
	"encoding/json"
	"github.com/dop251/goja"
	"gopkg.in/errgo.v2/errors"
)

func RunJs(script string, input string) ([]byte, error) {

	// run script to create main function
	vm := goja.New()
	_, err := vm.RunString(script)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	// get main function
	main, ok := goja.AssertFunction(vm.Get("main"))
	if !ok {
		return nil, errors.New("main is not a function")
	}

	// unmarshal input
	data := map[string]any{}
	err = json.Unmarshal([]byte(input), &data)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	// call main function
	res, err := main(goja.Undefined(), vm.ToValue(data))
	if err != nil {
		return nil, errors.Wrap(err)
	}

	// marshal result
	jsonResult, err := json.Marshal(res.Export())
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return jsonResult, nil
}
