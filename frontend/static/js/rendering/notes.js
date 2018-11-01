class NotesRendering {

	AllNotesHomePage(allNotes) {
		var table = document.getElementById("AllNotesTable");
		var html = "";

		if (table == undefined)
			return;

		// Table's header
		html += "<table>";
		html += "<tr>";
		html += "<td>Title</td>";
		html += "<td>Content</td>";
		html += "</tr>";

		// Table's body
		for (var i = 0; i < allNotes.length; i++) {
			html += "<tr>";
			html += "<td>" + allNotes[i].Title + "</td>";
			html += "<td>" + allNotes[i].Content + "</td>";
			html += "</tr>";
		}
		
		// End of the table
		html += "</table>";

		// Rendering
		table.innerHTML = html;
	}
}