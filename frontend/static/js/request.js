class DatabaseAPI {

	constructor(url) {
		this.url = url
	}

	RequestList() {
		console.log("click");
		$.ajax({
			method: "GET",
			url: "/api/notes",
			data: {
				"token": "ceci est un test"
			},
			success: function(result) {
				console.log("Request Valid")
				console.log(result)
				return result
			},
			error: function(result) {
				console.log("Could not request API at /api/url")
				console.log(result);
			}
		});
	}
}