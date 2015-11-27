package client

import (
	"bytes"
)

//Cloner describes an object that can be deep cloned
type Cloner interface {
	//Clone returns the deep cloned object, and if for
	//some reason, the object cannot be cloned, the
	//appropriate error will be returned
	Clone() (interface{}, error) //clone()
}

//ChangeCounter describes an object whose changes
//are counted monotonically. For example, every time a field
//of an object is changed, an internal write safe index would
//be updated accordingly
type ChangeCounter interface { //util.ChangeCountable
	//CountChanges returns the internal counter describing
	//how many times the object's fields have been changed
	CountChanges() int64 //getChangeCount
}

//Signer describes an object that has been signed
//with a signature
type Signer interface { //Signature
	//GetSignature returns the signature blob of a
	//given data packet
	GetSignature() (*bytes.Buffer, error) //getSignature
	//SetSignature configures the signer with the given
	//signature blob with the configuration on whether
	//to deep clone the blob or not
	SetSignature(blob *bytes.Buffer, clone bool) //setSignature
}
