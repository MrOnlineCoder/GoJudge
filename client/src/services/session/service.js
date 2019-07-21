import axios from 'axios'

let subscribeCb = null;

let cachedUser = {};
let loggedIn = false;

function subscribe(cb) {
	subscribeCb = cb;
}

function getCachedUser() {
	return cachedUser;
}

function logout() {
	loggedIn = false;
	cachedUser = {};
	if (subscribeCb) subscribeCb();
}

function whoami() {
	axios.get('/api/auth/me').then(response => {
		if (!response.data.success) {
			loggedIn = false;
			console.error('[Session] whoami failed.');
			return;
		}

		if (!response.data.user) {
			loggedIn = false;
			return;
		}

		cachedUser = response.data.user;

		loggedIn = true;

		if (subscribeCb) subscribeCb();
	}).catch(error => {
		loggedIn = false;
		console.error('[Session] Init request failed');
	});
}

function isLoggedIn() {
	return loggedIn;
}

export default {
	getCachedUser,
	subscribe,
	whoami,
	logout,
	isLoggedIn
}