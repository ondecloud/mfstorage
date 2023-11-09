package mfstorage

import "fmt"

// AnyDatastoreConfig returns a DatastoreConfig from a spec based on
// the "type" parameter
func AnyDatastoreConfig(params map[string]interface{}) (DatastoreConfig, error) {
	which, ok := params["type"].(string)
	if !ok {
		return nil, fmt.Errorf("'type' field missing or not a string")
	}
	dss.Lock()
	fun, ok := dss.cfgMaps[which]
	dss.Unlock()
	if !ok {
		return nil, fmt.Errorf("unknown datastore type: %s", which)
	}
	return fun(params)
}
