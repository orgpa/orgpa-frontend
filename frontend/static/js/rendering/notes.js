class NotesRendering {

	AllHomePage(allNotes) {

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
			html += "<td><h6 style='margin:0px;'>" + allNotes.data[i].title + "</h6></td>";
			html += "<td>" + "<a href='/note/"+ allNotes.data[i].id +"'><button class='btn red' style='padding-left: 25px;padding-right: 25px;'>see</button></a>" + "</td>";
			html += "</tr>";
		}

		// Rendering
		table.innerHTML = html;
	}
}