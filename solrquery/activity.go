package solrquery

import (
	"fmt"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/vanng822/go-solr/solr"
)

const (
	ivHost       = "Host"
	ivPort       = "Port"
	ivUserName   = "UserName"
	ivPassword   = "Password"
	ivCollection = "Collection"

	ivQuery             = "Query"
	ivFilterQuery       = "FilterQuery"
	ivSort              = "Sort"
	ivStart             = "Start"
	ivRow               = "Row"
	ivFieldList         = "FieldList"
	ivDefaultSearchList = "DefaultSearchList"
	ivRawQueryParameter = "RawQueryParameter"

	ovOutput = "Output"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// do eval
	host := context.GetInput(ivHost).(string)
	port := context.GetInput(ivPort).(int)
	username := context.GetInput(ivUserName).(string)
	password := context.GetInput(ivPassword).(string)
	collection := context.GetInput(ivCollection).(string)

	query := context.GetInput(ivQuery).(string)
	filterquery := context.GetInput(ivFilterQuery).(string)
	sort := context.GetInput(ivSort).(string)
	fieldlist := context.GetInput(ivFieldList).(string)
	//defaultsearchlist := context.GetInput(ivDefaultSearchList).(string)
	rawqueryparameter := context.GetInput(ivRawQueryParameter).(string)
	start := context.GetInput(ivStart).(int)
	row := context.GetInput(ivRow).(int)

	LocalURL := fmt.Sprintf("http://%s:%d/solr", host, port)

	si, _ := solr.NewSolrInterface(LocalURL, collection)
	if username != "" && password != "" {
		si.SetBasicAuth(username, password)

	}

	// Query/Search Solr
	solrquery := solr.NewQuery()
	solrquery.Q(query)

	if filterquery != "" {
		solrquery.FilterQuery(filterquery)

	}
	if sort != "" {
		solrquery.Sort(sort)

	}

	if fieldlist != "" {
		solrquery.FieldList(fieldlist)

	}
	solrquery.Start(start)
	solrquery.Rows(row)

	if rawqueryparameter != "" {
		params := strings.Split(rawqueryparameter, "&")
		for i := 0; i < len(params); i++ {
			keyval := strings.Split(params[i], "=")
			solrquery.AddParam(keyval[0], keyval[1])
		}
	}

	s := si.Search(solrquery)
	r, _ := s.Result(nil)
	fmt.Println(r.Results.Docs)
	context.SetOutput("Output", r.Results.Docs)

	return true, nil
}
