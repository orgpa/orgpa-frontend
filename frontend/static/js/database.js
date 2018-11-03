class DatabaseAPI {

	constructor(url) {
		this.url = url
	}

	RequestAllNotes(callback) {
		$.ajax({
			method: "GET",
			url: "/api/notes",
			success: function(data) {
				callback(data)
			},
			error: function(data) {
				console.error("Could not request API")
				console.error(data)
			}
		});
	}

	NewNote(idTitle, idContent) {
		var titleInput = document.getElementById(idTitle)
		var contentInput = document.getElementById(idContent)
		if (!titleInput || !contentInput) {
			console.warn("Invalid HTML.")
			return false
		}
		var title = titleInput.value
		var content = contentInput.value
		if (title == "" || content == "") {
			console.warn("Invalid form.")
			return false
		}

		$.ajax({
			method: "POST",
			url: "/api/notes",
			data: {
				"title": title,
				"content": content
			},
			success: function(data) {
				titleInput.value = ""
				contentInput.value = ""

			},
			error: function(data) {
				console.error("Could not request API")
				console.error(data)
			}
		})
	}

	DeleteNote(idNote, redirectURL) {
		$.ajax({
			method: "DELETE",
			url: "/api/notes/"+idNote,
			success: function(data) {
				console.log(data)
				window.location = redirectURL
			},
			error: function(data) {
				console.error("Could not request API")
				console.error(data)
			}
		})
	}

	ModifyNote(idNote, idModifiedContent) {
		var content = document.getElementById(idModifiedContent)
		if (!content)
			return false
		$.ajax({
			method: "PATCH",
			url: "/api/notes",
			data: {
				"id": idNote,
				"content": content.value
			},
			success: function(data) {
				window.location = "/note/" + idNote
			},
			error: function(data) {
				console.error("Could not request API")
				console.error(data)
			}
		})
	}
}