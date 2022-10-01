package main

import (
	"encoding/json"
	"fmt"

	"github.com/jmpsec/osctrl/queries"
)

// GetQueries to retrieve nodes from osctrl
func (api *OsctrlAPI) GetQueries(env string) ([]queries.DistributedQuery, error) {
	var qs []queries.DistributedQuery
	reqURL := fmt.Sprintf("%s%s%s/%s", api.Configuration.URL, APIPath, APIQueries, env)
	rawQs, err := api.GetGeneric(reqURL, nil)
	if err != nil {
		return qs, fmt.Errorf("error api request - %v - %s", err, string(rawQs))
	}
	if err := json.Unmarshal(rawQs, &qs); err != nil {
		return qs, fmt.Errorf("can not parse body - %v", err)
	}
	return qs, nil
}

// GetQuery to retrieve one node from osctrl
func (api *OsctrlAPI) GetQuery(env, name string) (queries.DistributedQuery, error) {
	var q queries.DistributedQuery
	reqURL := fmt.Sprintf("%s%s%s/%s/%s", api.Configuration.URL, APIPath, APIQueries, env, name)
	rawQ, err := api.GetGeneric(reqURL, nil)
	if err != nil {
		return q, fmt.Errorf("error api request - %v - %s", err, string(rawQ))
	}
	if err := json.Unmarshal(rawQ, &q); err != nil {
		return q, fmt.Errorf("can not parse body - %v", err)
	}
	return q, nil
}

// DeleteQuery to delete node from osctrl
func (api *OsctrlAPI) DeleteQuery(env, identifier string) error {
	return nil
}

// CompleteQuery to complete a query from osctrl
func (api *OsctrlAPI) CompleteQuery(env, identifier string) error {
	return nil
}