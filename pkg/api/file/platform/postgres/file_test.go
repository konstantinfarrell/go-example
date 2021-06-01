package postgres_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/konstantinfarrell/go-example"
	"github.com/konstantinfarrell/go-example/mocks"
	pg"github.com/konstantinfarrell/go-example/pkg/util/postgres"
	"github.com/konstantinfarrell/go-example/pkg/api/file/platform/postgres"
)

func TestReadFile(t *testing.T) {
	assert := assert.New(t)
	mdb := new(mocks.Databaser)
	mdb.On("Query", mock.Anything, mock.Anything).Return(nil, nil)
	d := &pg.Database{ mdb }
	f := new(gox.File)
	pf := new(postgres.File)
	_, err := pf.ReadFile(d, f)
	assert.Nil(err)
}

func TestReadAllFiles(t *testing.T) {
	assert := assert.New(t)
	mdb := new(mocks.Databaser)
	mdb.On("Query", mock.Anything, mock.Anything).Return(nil, nil)
	d := &pg.Database{ mdb }
	pf := new(postgres.File)
	_, err := pf.ReadAllFiles(d)
	assert.Nil(err)
}