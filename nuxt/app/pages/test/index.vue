<script>
export default {
	data() {
		return {
			socket: new WebSocket("ws://localhost:8080/ws"),
			pc: null,
			peerId: null,
			usersIds: ["1", "2"],
			offer: null,
			answer: null,
			candidateQueue: [],
		};
	},
	methods: {
		async startCall(to) {
			const stream = await navigator.mediaDevices.getUserMedia({
				audio: {
					echoCancellation: true, // подавление эха
					noiseSuppression: true, // шумоподавление
					autoGainControl: true, // автоподстройка громкости
				},
			});
			this.pc = new RTCPeerConnection({
				iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
			});
			stream.getTracks().forEach((track) => this.pc.addTrack(track, stream));
			const offer = await this.pc.createOffer({ offerToReceiveAudio: true });
			await this.pc.setLocalDescription(offer);
			this.socket.send(
				JSON.stringify({
					type: "offer",
					from: this.peerId,
					to: to,
					payload: offer,
				})
			);

			this.pc.ontrack = (event) => {
				this.$refs.remoteAudio.srcObject = event.streams[0];

				// this.pc
				// 	.getSenders()
				// 	.filter((s) => s.track.kind === "audio")
				// 	.forEach((sender) => {
				// 		let params = sender.getParameters();
				// 		if (!params.encodings) params.encodings = [{}];
				// 		params.encodings[0].fec = true;
				// 		params.encodings[0].priority = "high";
				// 		sender.setParameters(params);
				// 	});

				const receiver = this.pc
					.getReceivers()
					.find((r) => r.track.kind === "audio");
				receiver.playoutDelayHint = 0.2; // 200 мс задержки ради плавности
			};
			this.pc.onicecandidate = (event) => {
				if (event.candidate) {
					this.socket.send(
						JSON.stringify({
							type: "candidate",
							from: this.peerId,
							to: to,
							payload: event.candidate,
						})
					);
				}
			};
		},
		async acceptCall() {
			const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
			this.pc = new RTCPeerConnection({
				iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
			});
			stream.getTracks().forEach((track) => this.pc.addTrack(track, stream));
			const offer = this.offer.payload;
			await this.pc.setRemoteDescription(offer);
			const answer = await this.pc.createAnswer();
			await this.pc.setLocalDescription(answer);

			this.socket.send(
				JSON.stringify({
					type: "answer",
					from: this.peerId,
					to: this.offer.from,
					payload: answer,
				})
			);

			this.pc.ontrack = (event) => {
				this.$refs.remoteAudio.srcObject = event.streams[0];
			};
			this.pc.onicecandidate = (event) => {
				if (event.candidate) {
					this.socket.send(
						JSON.stringify({
							type: "candidate",
							from: this.peerId,
							to: this.offer.from,
							payload: event.candidate,
						})
					);
				}
			};
		},
	},
	mounted() {
		const params = new URLSearchParams(window.location.search);
		this.peerId = params.get("id");
		const index = this.usersIds.indexOf(this.peerId);
		if (index !== -1) {
			this.usersIds.splice(index, 1);
		}
		this.socket.onopen = () => {
			// отправляем свой peerId
			console.log("отправляем свой peerId");
			this.socket.send(
				JSON.stringify({
					type: "peerId",
					from: this.peerId,
				})
			);
		};
		this.socket.onmessage = (event) => {
			const msg = JSON.parse(event.data);
			if (msg.type == "offer") {
				this.offer = msg;
			}
			if (msg.type == "answer") {
				this.answer = msg;
				this.pc.setRemoteDescription(this.answer.payload);
			}
			if (msg.type == "candidate") {
				if (!this.pc) {
					this.candidateQueue.push(msg.payload);
				} else {
					this.candidateQueue.forEach((candidate, i) => {
						this.pc.addIceCandidate(new RTCIceCandidate(candidate));
						this.candidateQueue.splice(i, 1);
					});
					this.pc.addIceCandidate(new RTCIceCandidate(msg.payload));
				}
			}
		};
		this.socket.onerror = (error) => {
			console.error("WebSocket error:", error);
		};
		this.socket.onclose = () => {
			console.log("WebSocket connection closed");
		};
	},
};
</script>

<template>
	<div>
		<h1>Минимальная аудиосвязь на WebRTC</h1>
		<ul>
			<li v-for="userId of usersIds" :key="userId">
				<div>друг: {{ userId }}</div>
				<button @click="acceptCall(userId)" v-if="offer">Принять</button>
				<button @click="startCall(userId)" v-else>Начать звонок</button>
			</li>
		</ul>
		<audio ref="remoteAudio" autoplay></audio>
	</div>
</template>
