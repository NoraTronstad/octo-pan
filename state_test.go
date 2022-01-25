import "testing"import "testing"  
func TestViewState(t *testing.T) {func TestViewState(t *testing.T) {  
    wanted := "[kylling rev korn mann ---\ \__/ _________________/---]"    wanted := "[kylling rev korn mann ---\ \__/ _________________/---]"  
    state := ViewState();    state := ViewState();  
    if state != wanted {    if state != wanted {  
        t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)         t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)   
    }    }
}}  
