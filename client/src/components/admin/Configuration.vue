<template>
	<div>
		<ErrorBlock :err="error"/>
		<b-alert variant="success" :show="saved">
			<font-awesome-icon icon="check"/>
			New configuration has been saved!
		</b-alert>

		<b-card title="Server">
			<b-form>
				<b-form-group 
					label="Port"
					description="Web server port.">
					<b-form-input type="number" v-model.number="config.server.port"/>
				</b-form-group>	
				<b-form-group 
					label="Max workers"
					description="Maximum amount of judge workers that can be started. Use 0 to set it to numbers of CPU cores avaliable on your system. Default: 0 (CPUs num)">
					<b-form-input type="number" v-model.number="config.server.max_workers"/>
				</b-form-group>	
			</b-form>
		</b-card>	
		<br>
		<b-card title="Limits">
			<b-form>
				<b-form-group 
					label="Submission source code limit"
					description="Maximum size of submission source code, in bytes, including whitespace and blank lines. Default: 128 KB (131072 bytes)">
					<b-form-input type="number" v-model.number="config.limits.sourcecode"/>
				</b-form-group>	
			</b-form>
		</b-card>	
		<br>
		<b-button variant="success" @click="saveConfig()">
			<font-awesome-icon icon="save"/>
			Save configuration
		</b-button>
		<br>	
		<br>	
	</div>
</template>

<script>
import axios from 'axios'

export default {
	data() {
		return {
			config: {
				server: {
					port: 1337,
					max_workers: 0
				},
				limits: {
					sourcecode: 131072
				}
			},
			saved: false,
			error: null
		}
	},
	methods: {
		fetchConfig() {
			axios.get('/api/admin/config').then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.config = response.data.config;
			}).catch(error => {
				this.error = error;
			});
		},
		saveConfig() {
			this.saved = false;
			this.error = null;
			axios.post('/api/admin/config/save', {
				config: this.config
			}).then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.saved = true;
			}).catch(error => {
				this.error = error;
			});
		}
	},
	created() {
		this.fetchConfig();
	}
}
</script>