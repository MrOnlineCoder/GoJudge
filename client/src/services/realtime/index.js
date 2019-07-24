let socket = null;

function init(callbacks) {
	var socket = new WebSocket(`ws://localhost:1337/api/realtime/ws`);

	socket.onopen = function () {
		
	};

	socket.onmessage = function (e) {
		let msg = JSON.parse(e.data);

		console.log('[Realtime] Received event '+msg.type);

		if (msg.type === 'submission_update') {
			callbacks.onSubmissionUpdate(msg.data.submission_id, msg.data.verdict, msg.data.passed_tests);
			return;
		}
	};
}

export default {
	init
}