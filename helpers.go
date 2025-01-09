package gscXml

import "encoding/xml"

func toEncodeXmlName(name XmlName) xml.Name {
	return xml.Name{
		Local: name.Local,
		Space: name.Space,
	}
}

func toGscXmdName(name xml.Name) XmlName {
	return XmlName{
		Space: name.Space,
		Local: name.Local,
	}
}

func Name(local string, space string) XmlName {
	return XmlName{
		Space: space,
		Local: local,
	}
}

func GetNode(m map[XmlName][]*XmlNode, name XmlName) *XmlNode {
	return m[name][0]
}

func GetNodes(m map[XmlName][]*XmlNode, name XmlName) []*XmlNode {
	return m[name]
}
