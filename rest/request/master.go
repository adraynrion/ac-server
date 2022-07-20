package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type (
	MasterCreate struct {
		Name string
	}

	MasterRead struct {
		Name string
	}

	MasterUpdate struct {
		Name string
	}

	MasterDelete struct {
		Name string
	}
)

// NewMasterCreate request
func NewMasterCreate() *MasterCreate {
	return &MasterCreate{}
}

// Auditable returns all auditable/loggable parameters
func (r *MasterCreate) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"name": r.Name,
	}
}

// Auditable returns all auditable/loggable parameters
func (r *MasterCreate) GetName() string {
	return r.Name
}

func (r *MasterCreate) Fill(req *http.Request) (err error) {
	if strings.HasPrefix(strings.ToLower(req.Header.Get("content-type")), "application/json") {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w", err)
		}
	}

	// Caching 32MB to memory, the rest to disk
	if err = req.ParseMultipartForm(32 << 20); err != nil && err != http.ErrNotMultipart {
		return err
	} else if err == nil {
		// Multipart params
		if val, ok := req.MultipartForm.Value["name"]; ok && len(val) > 0 {
			r.Name, err = val[0], nil
		}
	}

	if err = req.ParseForm(); err != nil {
		return err
	}

	// POST Params
	if val, ok := req.Form["name"]; ok && len(val) > 0 {
		r.Name, err = val[0], nil
		if err != nil {
			return err
		}
	}

	return err
}

// NewMasterRead request
func NewMasterRead() *MasterRead {
	return &MasterRead{}
}

// Auditable returns all auditable/loggable parameters
func (r *MasterRead) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"name": r.Name,
	}
}

// Auditable returns all auditable/loggable parameters
func (r *MasterRead) GetName() string {
	return r.Name
}

func (r *MasterRead) Fill(req *http.Request) (err error) {
	// GET Params
	tmp := req.URL.Query()

	if val, ok := tmp["name"]; ok && len(val) > 0 {
		r.Name, err = val[0], nil
		if err != nil {
			return err
		}
	}

	return err
}

// NewMasterUpdate request
func NewMasterUpdate() *MasterUpdate {
	return &MasterUpdate{}
}

// Auditable returns all auditable/loggable parameters
func (r *MasterUpdate) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"name": r.Name,
	}
}

// Auditable returns all auditable/loggable parameters
func (r *MasterUpdate) GetName() string {
	return r.Name
}

func (r *MasterUpdate) Fill(req *http.Request) (err error) {
	if strings.HasPrefix(strings.ToLower(req.Header.Get("content-type")), "application/json") {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w", err)
		}
	}

	// Caching 32MB to memory, the rest to disk
	if err = req.ParseMultipartForm(32 << 20); err != nil && err != http.ErrNotMultipart {
		return err
	} else if err == nil {
		// Multipart params
		if val, ok := req.MultipartForm.Value["name"]; ok && len(val) > 0 {
			r.Name, err = val[0], nil
			if err != nil {
				return err
			}
		}
	}

	if err = req.ParseForm(); err != nil {
		return err
	}

	// POST params
	if val, ok := req.Form["name"]; ok && len(val) > 0 {
		r.Name, err = val[0], nil
		if err != nil {
			return err
		}
	}

	// GET params
	tmp := req.URL.Query()

	if val, ok := tmp["name"]; ok && len(val) > 0 {
		r.Name, err = val[0], nil
		if err != nil {
			return err
		}
	}

	return err
}

// NewMasterDelete request
func NewMasterDelete() *MasterDelete {
	return &MasterDelete{}
}

// Auditable returns all auditable/loggable parameters
func (r *MasterDelete) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"name": r.Name,
	}
}

// Auditable returns all auditable/loggable parameters
func (r *MasterDelete) GetName() string {
	return r.Name
}

func (r *MasterDelete) Fill(req *http.Request) (err error) {
	tmp := req.URL.Query()

	if val, ok := tmp["name"]; ok && len(val) > 0 {
		r.Name, err = val[0], nil
		if err != nil {
			return err
		}
	}

	return err
}
