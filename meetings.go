package meetings

import "gopkg.in/mgo.v2/bson"

type meetings struct{

  ID        bson.ObjectId `bson:"_id" json:"id"`
	Title        string        `bson:"title" json:"title"`
	Participants  string        `bson:"participant" json:"participant"`
	StartTime    bson.Date       `bson:"start" json:"start"`
  EndTime    bson.Date       `bson:"end" json:"end"`
  CRTimeStamp  bson.TimeStamp  `bson:"ts" json:"ts"`



 }
