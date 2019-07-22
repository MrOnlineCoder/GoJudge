import axios from 'axios'

let active = false;
let contest = {};

function loadStatus() {
	return new Promise((resolve, reject) => {
		axios.get('/api/contest/status').then(response => {
			if (!response.data.success) {
				console.error(`[Contest] Status error: ${response.data.message}`);
				reject(response.data.message);
				return;
			}

			active = response.data.active;

			if (active) {
				contest = response.data.contest;
			}

			resolve(response.data);
		}).catch(error => {
			console.error(`[Contest] Status request failed: ${error}`);
		});
	});
}

function isActive() {
	return active;
}

function getContest() {
	return contest;
}

export default {
	loadStatus,
	isActive,
	getContest
}