
package project

// Contains data about a project

// https://github.com/thiagoralves/OpenPLC_v3/blob/master/webserver/openplc.py#L71

type Project struct {
	File string
	Name string
	Nick string
	Desription string
}


func NewProject() Project {

	p := Project{}
	
	return p
}

	