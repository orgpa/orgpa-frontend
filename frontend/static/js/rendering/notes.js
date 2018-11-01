class NotesRendering {

	AllNotesHomePage(allNotes) {
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