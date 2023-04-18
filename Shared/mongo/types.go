package mongo


type BlockMeta struct {
    BlockId string
    Type string
    string lang = 3;
    string updateAt = 4;
    repeated Author authors = 5;
    repeated string categories = 6;
    repeated string tags = 7;
}
