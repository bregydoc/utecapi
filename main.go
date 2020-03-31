package main

import (
	"context"
	"fmt"
	"github.com/MontFerret/ferret/pkg/drivers"
	"github.com/MontFerret/ferret/pkg/runtime"
	"io/ioutil"

	"github.com/MontFerret/ferret/pkg/compiler"
	"github.com/MontFerret/ferret/pkg/drivers/cdp"
	"github.com/pkg/errors"
)

const enrollmentQueryPath = "./queries/enrollment.fql"

func main() {
	enrollmentQuery, err := ioutil.ReadFile(enrollmentQueryPath)
	if err != nil {
		panic(errors.Wrap(err, "error at read enrollment query"))
	}

	comp := compiler.New()

	program, err := comp.Compile(string(enrollmentQuery))
	if err != nil {
		panic(errors.Wrap(err, "error at compile enrollment program"))
	}

	ctx := context.Background()

	ctx = drivers.WithContext(ctx, cdp.NewDriver(), drivers.AsDefault())

	out, err := program.Run(ctx, runtime.WithParams(map[string]interface{}{
		"email": "bregy.malpartida@utec.edu.pe",
		"password": "alanturing1802",
	}))
	if err != nil {
		panic(errors.Wrap(err , "error at run enrollment program"))
	}


	fmt.Println(string(out))
}