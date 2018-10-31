class DatabaseAPI {

	constructor(url) {
		this.url = url
	}

	RequestList() {
		console.log("click");
		$.ajax({
			method: "GET",
			url: "/api/notes",
			success: function(result) {
				console.log("Request Valid")
				console.log(result[0])
				document.getElementById("response-request").innerHTML = 
					"ID: " + result[0].ID +
					" Title: " + result[0].Title +
					" Content: " + result[0].Content +
					" Last Edited: " + result[0].LastEdit;
				return result
			},
			error: function(result) {
				console.log("Could not request API at /api/url")
				console.log(result);
			}
		});
	}
}