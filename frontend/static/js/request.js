class DatabaseAPI {

	constructor(url) {
		this.url = url
	}

	RequestAllNotes(callback) {
		$.ajax({
			method: "GET",
			url: "/api/notes",
			success: function(data) {
				callback(data);
			},
			error: function(data) {
				console.error("Could not request API at /api/url")
				console.error(data);
			}
		});
	}
}