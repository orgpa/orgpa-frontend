class NotesRendering {

	AllNotesHomePage(allNotes) {

		// Display the AllNotes div if there is any notes to display
		if (allNotes.length == 0) {
			document.getElementById("AllNotes").style.display = "none";
			return;
		} else {
			document.getElementById("AllNotes").style.display = "block";
		}
		var table = document.getElementById("AllNotesTableBody");
		var html = "";

		if (table == undefined)
			return;

		// Table's body
		for (var i = 0; i < allNotes.length; i++) {
			html += "<tr>";
			html += "<td>" + allNotes[i].Title + "</td>";
			html += "<td>" + allNotes[i].Content + "</td>";
			html += "<td>" + "<a href='/note/"+ allNotes[i].ID +"'><button class='btn red'>see</button></a>" + "</td>";
			html += "</tr>";
		}

		// Rendering
		table.innerHTML = html;
	}
}