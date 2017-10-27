package datamodels

type Crserver struct {
	BaseModel
	Address   string
	Version   string
	BinPath   string
	Port      string
	Range     string
	Directory string
	Tags      []*Tag `pg:",polymorphic:Trackable"`
}

type Tag struct {
	TrackableID   int64
	TrackableType string
	Text          string
}
