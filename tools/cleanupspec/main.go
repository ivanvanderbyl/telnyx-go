package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/inflect"
)

func main() {
	ctx := context.Background()
	loader := &openapi3.Loader{
		Context:               ctx,
		IsExternalRefsAllowed: true,
	}

	doc, err := loader.LoadFromFile("./openapi.yaml")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	err = doc.Validate(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	for typeName, v := range doc.Components.Schemas {
		objectType, _ := jq(v, "type")
		fmt.Println(typeName, objectType)
	}

	for _, pathItem := range doc.Paths.InMatchingOrder() {
		fmt.Println(pathItem)

		// for _, method := range pathItem.Operations() {
		// 	for _, param := range method.Parameters {
		// 		processParameters(param)
		// 	}
		// }

		// if pathItem.Get != nil {
		// 	for _, param := range pathItem.Get.Parameters {
		// 		processParameters(param)
		// 	}
		// }
		// if pathItem.Post != nil {
		// 	for _, param := range pathItem.Post.Parameters {
		// 		processParameters(param)
		// 	}
		// }
		// if pathItem.Patch != nil {
		// 	for _, param := range pathItem.Patch.Parameters {
		// 		processParameters(param)
		// 	}
		// }
		// if pathItem.Delete != nil {
		// 	for _, param := range pathItem.Delete.Parameters {
		// 		processParameters(param)
		// 	}
		// }
		// if pathItem.Put != nil {
		// 	for _, param := range pathItem.Put.Parameters {
		// 		processParameters(param)
		// 	}
		// }
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(doc)
	if err != nil {
		log.Fatalln(err)
	}
}

func processParameters(param *openapi3.ParameterRef) {
	switch param.Value.Name {
	case "id":
		param.Value.Extensions["x-go-name"] = "ID"
		param.Value.Required = true

	case "property_id":
		if param.Value.Schema.Value.Type == "array" {
			param.Value.Extensions["x-go-name"] = "PropertyIDs"
		} else {
			param.Value.Extensions["x-go-name"] = "PropertyID"
		}
		param.Value.Required = true
	default:
		param.Value.Extensions["x-go-name"] = pascal(param.Value.Name)
		if param.Value.Schema.Value.Type == "array" {
			param.Value.Required = true
		}
	}
}

type SchemaLike interface {
	JSONLookup(token string) (interface{}, error)
}

func jq(v SchemaLike, path string) (any, bool) {
	if v == nil {
		return nil, false
	}
	props, err := v.JSONLookup(path)
	if err != nil {
		return nil, false
	}
	return props, true
}

var (
	acronyms = make(map[string]struct{})
	rules    = ruleset()
)

func pascalWords(words []string) string {
	for i, w := range words {
		upper := strings.ToUpper(w)
		if _, ok := acronyms[upper]; ok {
			words[i] = upper
		} else {
			words[i] = rules.Capitalize(w)
		}
	}
	return strings.Join(words, "")
}

// pascal converts the given name into a PascalCase.
//
//	user_info 	=> UserInfo
//	full_name 	=> FullName
//	user_id   	=> UserID
//	full-admin	=> FullAdmin
func pascal(s string) string {
	words := strings.FieldsFunc(s, isSeparator)
	return pascalWords(words)
}

func isSeparator(r rune) bool {
	return r == '_' || r == '-' || unicode.IsSpace(r)
}

func ruleset() *inflect.Ruleset {
	rules := inflect.NewDefaultRuleset()
	// Add common initialisms from golint and more.
	for _, w := range []string{
		"DS", "IATA", "HREF",
		"ACL", "API", "ASCII", "AWS", "CPU", "CSS", "DNS", "EOF", "GB", "GUID",
		"HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "KB", "LHS", "MAC", "MB",
		"QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SQL", "SSH", "SSO", "TCP",
		"TLS", "TTL", "UDP", "UI", "UID", "URI", "URL", "UTF8", "UUID", "VM",
		"XML", "XMPP", "XSRF", "XSS",
	} {
		acronyms[w] = struct{}{}
		rules.AddAcronym(w)
	}
	return rules
}
