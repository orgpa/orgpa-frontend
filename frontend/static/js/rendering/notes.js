class NotesRendering {

	AllNotesHomePage(allNotes) {

		// Display the AllNotes div if there is any notes to display
		if (allNotes.success == "false") {
			// TODO : handle error and display a pop up 
			document.getElementById("AllNotes").style.display = "none";
			return;
		} else if (allNotes.number_of_record == 0) {
			console.log("no data to show");
			return;
		} else {
			document.getElementById("AllNotes").style.display = "block";
		}
		var table = document.getElementById("AllNotesTableBody");
		var html = "";

		if (table == undefined)
			return;

		// Table's body
		for (var i = 0; i < allNotes.data.length; i++) {
			html += "<tr>";
			html += "<td>" + allNotes.data[i].title + "</td>";
			html += "<td>" + allNotes.data[i].content + "</td>";
			html += "<td>" + "<a href='/note/"+ allNotes.data[i].id +"'><button class='btn red'>see</button></a>" + "</td>";
			html += "</tr>";
		}

		// Rendering
		table.innerHTML = html;
	}
}