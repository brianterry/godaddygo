package godaddygo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
)

type records struct {
	c *Config
}

func newRecords(c *Config) records {
	return records{c}
}

func (r records) List(ctx context.Context) ([]Record, error) {
	url := "/domains/" + r.c.domainName + "/records"
	result, err := makeDo(ctx, r.c, MethodGet, url, nil, 200)
	if err != nil {
		return nil, exception.listingRecords(err, r.c.domainName)
	}

	return readRecordListResponse(result)
}

func (r records) FindByType(ctx context.Context, t string) ([]Record, error) {
	url := "/domains/" + r.c.domainName + "/records/" + t
	result, err := makeDo(ctx, r.c, MethodGet, url, nil, 200)
	if err != nil {
		return nil, exception.findingRecordsByType(err, r.c.domainName, t)
	}

	return readRecordListResponse(result)
}

func (r records) FindByTypeAndName(ctx context.Context, t string, n string) ([]Record, error) {
	url := "/domains/" + r.c.domainName + "/records/" + t + "/" + n
	result, err := makeDo(ctx, r.c, MethodGet, url, nil, 200)
	if err != nil {
		return nil, exception.findingRecordsByTypeAndName(err, r.c.domainName, t, n)
	}

	return readRecordListResponse(result)
}

func (r records) Update(ctx context.Context, rec Record) error {
	url := "/domains/" + r.c.domainName + "/records/" + rec.Name
	body, err := buildUpdateRecordRequest([]Record{rec}) // Must be []Record{} !!!
	if err != nil {
		return exception.updatingRecord(err, r.c.domainName, rec.Name)
	}

	if _, err = makeDo(ctx, r.c, MethodGet, url, body, 200); err != nil {
		return exception.updatingRecord(err, r.c.domainName, rec.Name)
	}

	return nil
}

func (r records) Delete(ctx context.Context, rec Record) error {
	/* return r.c.Delete("/domains/" + r.domain + "/records/" + rec.Name) */
	return nil
}

func readRecordListResponse(r []byte) ([]Record, error) {
	var zone []Record
	if err := json.Unmarshal(r, &zone); err != nil {
		return []Record{}, exception.invalidJSONResponse(err)
	}

	return zone, nil
}

// buildUpdateRecordRequest gives us our dns record as io.Reader
func buildUpdateRecordRequest(rec []Record) (io.Reader, error) {
	b, e := json.Marshal(rec)
	if e != nil {
		return nil, fmt.Errorf("ErrorCannotMarshalRecords : %w", e)
	}

	return bytes.NewBuffer(b), nil
}

func readRecordResponse(result io.ReadCloser) (Record, error) {
	//TODO..
	defer result.Close()
	return Record{}, nil
}
