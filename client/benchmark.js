const axios = require('axios');

function send() {
	axios.request({
		method: 'POST',
		url: 'http://localhost:1337/api/contest/submit',
		data: {
			problem_index: 0,
			language: "cpp",
			sourcecode: "#include <iostream>\nusing namespace std;\nint main() {int a; int b; cin >> a >> b; cout << a + b; return 0;}"
		},
		headers: {
			'Cookie': 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.549ajXQemOEpF725HotszuhqNOz34Aa2DFDvpkkGeoU;'
		}
	}).then(response => {
		console.log(response.data);
	}).catch(error => {
		console.log(error);
	});
}

for (let i = 0; i < 20; i++) send();