package transport

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	"github.com/konstantinfarrell/go-example"
	"github.com/konstantinfarrell/go-example/pkg/api/file"

)

// HTTP represents the file http service. 
type HTTP struct {
	svc file.Service // pkg/api/file/service.go
}

// Modify the pointer to the echo service to include the endpoints for file related operations
func NewHTTP(svc file.Service, r *echo.Group){
	h := HTTP{svc}
	ur := r.Group("/files")

	// Define routes for REST API
	// TODO: swagger
	ur.POST("", h.create)
	ur.GET("", h.readAll)
	ur.GET("/:id", h.read)
	ur.DELETE("/:id", h.delete)
}

// Model for file create request
type createReq struct {
	Filename			string	`json:"filename" validate:"required"`
	Path				string	`json:"path" validate:"required"`
	Permissions			string	`json:"permissions" validate:"required"`
	Created				string	`json:"created" validate:"required"`
	Modified			string	`json:"modified" validate:"required"`
	Data				string  `json:"data" validate:"required"`
}

func (h HTTP) create(c echo.Context) error {
	r := new(createReq)

	// Bind takes a context and attempts to bind it to a struct.
	// No validation appears to happen here
	if err := c.Bind(r); err != nil {
		return err
	}
	created := time.Now().UTC().Format(time.RFC3339)

	fileId := uuid.New().String()

	fc := make(chan gox.FileChannel)
	defer close(fc)
	
	go h.svc.Create(fc, c, &gox.File{
		Filename:		r.Filename,
		Path:			r.Path,
		Permissions:	r.Permissions,
		Created:		r.Created,
		Modified:		r.Modified,
		FileId:			fileId,
		Data:			[]byte(r.Data),
		Received:		created,
	})

	result := <- fc
	file := result.File
	err := result.Err

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &file)
}

//TODO: pagination
func (h HTTP) readAll(c echo.Context) error {
	fc := make(chan gox.FileChannel)
	defer close(fc)
	go h.svc.ReadAll(fc, c)


	result := <- fc
	file := result.File
	err := result.Err
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &file)
}

func (h HTTP) read(c echo.Context) error {
	fc := make(chan gox.FileChannel)
	defer close(fc)
	id := c.Param("id")

	go h.svc.Read(fc, c, &gox.File{
		FileId: id,
	})
	result := <- fc
	err := result.Err

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &result)	
}

func (h HTTP) delete(c echo.Context) error {
	fc := make(chan gox.FileChannel)
	defer close(fc)
	id := c.Param("id")

	go h.svc.Delete(fc, c, &gox.File{FileId: id})
	result := <- fc
	err := result.Err
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}