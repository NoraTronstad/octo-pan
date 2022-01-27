import "testing"

func TestViewState(t *testing.T) { 
    wanted := "[kylling rev korn mann ---\ \__/ _________________/---]" 
    state := ViewState();   
    if state != wanted {  
        t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)   
    }
    
func PutInBoat() String{
        return "rev"
    }
    
func TestPutInBoat(t *testing.T) {
        wanted := "rev"
        var state = PutInBoat()
        if state != wanted {
            t.Errorf("Feil, fikk #{state}, ønsket #{wanted}")
        }
    }


