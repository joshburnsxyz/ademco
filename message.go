package ademco

import (
	"strings"
	"errors"
)

type Message struct {
	ClientCode       string
	IsRestore        bool
	EventCode        int
	EventDescription string
	PartitionCode    string
	ZoneId           string
}

var (
	clientCode        string
	qualifierCode     string
	eventCode         string
	partitionCode     string
	zoneId            string
	qualiferEventType bool
)

func (m *Message) Marshal(str string) (*Message, error) {
	s := strings.Split(str, " ")
	clientCode = s[0]
	qualifierCode = s[1]
	eventCode = s[2]
	partitionCode = s[3]
	zoneId = s[4]
	
	if qualifierCode == "E" {
		qualiferEventType = false
	} else if qualifierCode == "R" {
		qualiferEventType = true
	} else {
		return &Message{},errors.New("Invalid, qualifier")
	}
	
	msg = Message{
		ClientCode:       clientCode,
		IsRestore:        qualiferEventType,
		EventCode:        eventCode,
		EventDescription: GetMessageDescription(eventCode,qualiferEventType),
		PartitionCode:    partitionCode,
		ZoneId:           zoneId
	}
	return &msg,nil
}


func (m *Message) Unmarshal(*str string) (&str, error) {
	// TODO: Actual format template
	str = fmt.Sprintf("%v")
	return str
}


func NewMessage(str string) (*Message,error) {
	msg := Message{}
	msg = *msg.Marshal(&str)
	return &msg,nil
}
