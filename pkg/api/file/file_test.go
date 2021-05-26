package file_test

import (
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/konstantinfarrell/go-example"
)

func TestCreate(t *testing.T) {
	type args struct {
		c 	echo.Context
		req	gox.File
	}
	cases := []struct {
		name		string
		args		args	
		wantErr 	bool
		wantData	gox.File
	}

}