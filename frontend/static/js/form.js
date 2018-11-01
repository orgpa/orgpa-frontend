/**
 * Display the btnID if origin is different than actual
 * Or make the btnID disapear if origin and actual are equal
 */
function ShowValidateButton(origin, actual, btnID) {
	if (origin != actual) {
		document.getElementById(btnID).style.display = "block";
	} else {
		document.getElementById(btnID).style.display = "none";
	}
}

function CheckNewNoteForm() {
	// Check form validity
}