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

function whoami(cb) {
	axios.get('/api/auth/me').then(response => {
		if (!response.data.success) {
			loggedIn = false;
			console.error('[Session] whoami failed.');
			if(cb) cb();
			return;
		}

		if (!response.data.user) {
			loggedIn = false;
			if(cb) cb();
			return;
		}

		cachedUser = response.data.user;

		loggedIn = true;

		if (subscribeCb) subscribeCb();

		if(cb) cb();
	}).catch(error => {
		loggedIn = false;
		console.error('[Session] Init request failed');
		if(cb) cb();
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